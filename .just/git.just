import "_share.just"
import "_working_dir.just"

init_hooks: _env_is_dev
    uvx pre-commit install -t pre-commit -t pre-push -t commit-msg

run_hooks: _env_is_dev
    uvx pre-commit run

run_hooks_for_all: _env_is_dev
    uvx pre-commit run --all-files
