import ast
from pathlib import Path
from typing import cast

from scripts.file_checker._dto import FileSpecsDto


class FuncValidator:
    def validate(self, tree: ast.AST, file_specs: FileSpecsDto) -> None:
        for node in ast.walk(tree):
            if not self._is_func(node):
                continue

            func_node = cast("ast.FunctionDef | ast.AsyncFunctionDef", node)

            if self._is_public(func_node) and self._is_file_level(node) and self._is_public_module(file_specs.rel_path):
                error = (
                    f"{file_specs.rel_path}:{func_node.lineno}: "
                    f"top-level public function '{func_node.name}' "
                    f"is not allowed in a public module"
                )
                file_specs.errors.append(error)

    @staticmethod
    def _is_func(node: ast.AST) -> bool:
        return isinstance(node, ast.FunctionDef | ast.AsyncFunctionDef)

    @staticmethod
    def _is_public(func: ast.FunctionDef | ast.AsyncFunctionDef) -> bool:
        return not func.name.startswith("_")

    @staticmethod
    def _is_file_level(node: ast.AST) -> bool:
        return isinstance(node.parent, ast.Module)  # type: ignore[attr-defined]

    @staticmethod
    def _is_public_module(file_rel_path: Path) -> bool:
        return not any(p.startswith("_") for p in file_rel_path.parts)
