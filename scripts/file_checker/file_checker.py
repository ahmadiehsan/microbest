import ast
import logging
import tomllib
from pathlib import Path
from typing import NoReturn

from scripts.file_checker._dto import FileSpecsDto
from scripts.file_checker._func_validator import FuncValidator
from scripts.file_checker._import_validator import ImportValidator
from scripts.file_checker._msg_validator import MsgValidator


class FileChecker:
    def __init__(self) -> None:
        self.ignore_rules = self._load_ignore_rules()
        self.msg_validator = MsgValidator()
        self.import_validator = ImportValidator()
        self.func_validator = FuncValidator()

    def run(self) -> NoReturn:
        repo_abs_path = Path.cwd()
        errors = self._validate_files(repo_abs_path)

        if errors:
            for error in errors:
                logging.error(error)
            raise SystemExit(1)

        logging.info("all checks passed")
        raise SystemExit(0)

    def _validate_files(self, repo_abs_path: Path) -> list[str]:
        errors: list[str] = []

        for file_abs_path in repo_abs_path.rglob("*.py"):
            if file_abs_path.is_file():
                errors.extend(self._validate_file(file_abs_path, repo_abs_path))

        return errors

    def _validate_file(self, file_abs_path: Path, repo_abs_path: Path) -> list[str]:
        with file_abs_path.open() as f:
            tree = ast.parse(f.read(), filename=str(file_abs_path))

        self._set_parents(tree)
        file_specs = FileSpecsDto(
            repo_abs_path=repo_abs_path,
            abs_path=file_abs_path,
            rel_path=file_abs_path.relative_to(repo_abs_path),
            errors=[],
        )
        self._run_validators(tree, file_specs)
        return file_specs.errors

    def _run_validators(self, tree: ast.AST, file_specs: FileSpecsDto) -> None:
        ignored_validators = self._get_ignored_validators(file_specs.rel_path)

        if "import_validator" not in ignored_validators:
            self.import_validator.validate(tree, file_specs)

        if "msg_validator" not in ignored_validators:
            self.msg_validator.validate(tree, file_specs)

        if "func_validator" not in ignored_validators:
            self.func_validator.validate(tree, file_specs)

    def _get_ignored_validators(self, file_rel_path: Path) -> list[str]:
        """Return validators to ignore for a given file."""
        for base_rel_path, ignored_validators in self.ignore_rules.items():
            if str(file_rel_path).startswith(base_rel_path):
                return ignored_validators

        return []

    def _set_parents(self, node: ast.AST, parent: ast.AST | None = None) -> None:
        """Recursively set parent attributes for AST nodes."""
        node.parent = parent  # type: ignore[attr-defined]
        for child in ast.iter_child_nodes(node):
            self._set_parents(child, node)

    @staticmethod
    def _load_ignore_rules() -> dict[str, list[str]]:
        pyproject_path = Path("pyproject.toml")

        if not pyproject_path.exists():
            return {}

        with pyproject_path.open("rb") as f:
            config = tomllib.load(f)

        return config.get("tool", {}).get("file_checker", {}).get("per-file-ignores", {})


if __name__ == "__main__":
    FileChecker().run()
