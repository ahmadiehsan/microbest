set quiet
set no-exit-message

[private]
_project_name := "microbest"

[private]
_project_env_key := uppercase(_project_name) + "_ENV"

[private]
_project_env := env(_project_env_key, "")

[private]
_compose_cmd := "docker compose -f compose.yaml -f compose." + _project_env + ".yaml --env-file settings/compose/.env -p " + _project_name

[private]
_service_1_root := "./apps/service_1"

[private]
_service_2_root := "./apps/service_2"

[default]
help:
    just --list {{ module_path() }}

_env_is *envs:
    if [ -z "{{ _project_env }}" ]; then \
        echo "Error: {{ _project_env_key }} is not set. Run: export {{ _project_env_key }}=dev"; exit 1; \
    fi; \
    case " {{ envs }} " in \
        *" {{ _project_env }} "*) ;; \
        *) echo "Error: Target only for '{{ envs }}'. Current {{ _project_env_key }} is '{{ _project_env }}'"; exit 1;; \
    esac

_env_is_dev: (_env_is "dev")
_env_is_prod: (_env_is "prod")
_env_is_prod_or_dev: (_env_is "prod" "dev")
