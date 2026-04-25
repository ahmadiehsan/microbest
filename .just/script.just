import "_share.just"
import "_working_dir.just"

dir_checker: _env_is_prod_or_dev
    uv run --only-dev dir_checker

python_checker: _env_is_dev
    echo ">>>>> Service 2"
    cd {{ _service_2_root }} && uv run --only-dev python_checker

compose_checker: _env_is_dev
    uv run --only-dev compose_checker --env-file ./settings/compose/.env

just_indexer: _env_is_dev
    uv run --only-dev just_indexer .just
