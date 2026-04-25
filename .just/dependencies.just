import "_share.just"
import "_working_dir.just"

install: _env_is_dev
    echo ">>>>> Service 1"
    cd {{ _service_1_root }} && go mod download -x && go install tool

    echo ">>>>> Service 2"
    uv sync --project {{ _service_2_root }}

upgrade: _env_is_dev
    echo ">>>>> Service 1"
    cd {{ _service_1_root }} && go get -u ./...

    echo ">>>>> Service 2"
    uv sync --project {{ _service_2_root }} --upgrade

lock: _env_is_dev
    echo ">>>>> Service 1"
    cd {{ _service_1_root }} && go mod tidy

    echo ">>>>> Service 2"
    uv lock --project {{ _service_2_root }}
