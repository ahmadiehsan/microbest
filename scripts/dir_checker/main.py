import logging
from pathlib import Path
from typing import NoReturn

from scripts.dir_checker._dto import DirSpecsDto
from scripts.dir_checker._empty_validator import EmptyValidator
from scripts.dir_checker._init_py_validator import InitPyValidator
from scripts.dir_checker._logger import setup_logger

_logger = logging.getLogger(__name__)


class DirChecker:
    def __init__(self) -> None:
        self._empty_validator = EmptyValidator()
        self._init_py_validator = InitPyValidator()

    def run(self) -> NoReturn:
        setup_logger()
        repo_abs_path = Path.cwd()
        errors = self._validate_dirs(repo_abs_path)

        if errors:
            for error in errors:
                _logger.error(error)
            raise SystemExit(1)

        _logger.info("all checks passed")
        raise SystemExit(0)

    def _validate_dirs(self, repo_abs_path: Path) -> list[str]:
        errors: list[str] = []

        for dir_abs_path in repo_abs_path.rglob("*"):
            dir_rel_path = dir_abs_path.relative_to(repo_abs_path)

            if dir_abs_path.is_dir() and not self._is_hidden(dir_rel_path) and not self._is_in_black_list(dir_rel_path):
                errors.extend(self._validate_dir(dir_abs_path, dir_rel_path, repo_abs_path))

        return errors

    def _is_hidden(self, dir_rel_path: Path) -> bool:
        return dir_rel_path.name.startswith(".") or any(p.name.startswith(".") for p in dir_rel_path.parents)

    def _is_in_black_list(self, dir_rel_path: Path) -> bool:
        black_list = ["__pycache__", "venv", "env"]
        return dir_rel_path.name in black_list or any(p.name in black_list for p in dir_rel_path.parents)

    def _validate_dir(self, dir_abs_path: Path, dir_rel_path: Path, repo_abs_path: Path) -> list[str]:
        dir_specs = DirSpecsDto(repo_abs_path=repo_abs_path, abs_path=dir_abs_path, rel_path=dir_rel_path, errors=[])
        self._run_validators(dir_specs)
        return dir_specs.errors

    def _run_validators(self, dir_specs: DirSpecsDto) -> None:
        self._empty_validator.validate(dir_specs)
        self._init_py_validator.validate(dir_specs)


if __name__ == "__main__":
    DirChecker().run()
