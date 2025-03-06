import logging
from pathlib import Path
from typing import NoReturn

from scripts.dir_checker._dto import DirSpecsDto
from scripts.dir_checker._empty_validator import EmptyValidator


class DirChecker:
    def __init__(self) -> None:
        self.empty_validator = EmptyValidator()

    def run(self) -> NoReturn:
        repo_abs_path = Path.cwd()
        errors = self._validate_dirs(repo_abs_path)

        if errors:
            for error in errors:
                logging.error(error)
            raise SystemExit(1)

        logging.info("all checks passed")
        raise SystemExit(0)

    def _validate_dirs(self, repo_abs_path: Path) -> list[str]:
        errors: list[str] = []

        for dir_abs_path in repo_abs_path.rglob("*"):
            if dir_abs_path.is_dir():
                errors.extend(self._validate_dir(dir_abs_path, repo_abs_path))

        return errors

    def _validate_dir(self, dir_abs_path: Path, repo_abs_path: Path) -> list[str]:
        dir_specs = DirSpecsDto(
            repo_abs_path=repo_abs_path,
            abs_path=dir_abs_path,
            rel_path=dir_abs_path.relative_to(repo_abs_path),
            errors=[],
        )
        self._run_validators(dir_specs)
        return dir_specs.errors

    def _run_validators(self, dir_specs: DirSpecsDto) -> None:
        self.empty_validator.validate(dir_specs)


if __name__ == "__main__":
    DirChecker().run()
