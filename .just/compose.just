import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d

start: _env_is_dev
    {{ _compose_cmd }} up

destroy: _env_is_dev
    {{ _compose_cmd }} down -v

down: _env_is_prod_or_dev
    {{ _compose_cmd }} down

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs -f

ls: _env_is_prod_or_dev
    {{ _compose_cmd }} ps --format "table {{{{.ID}}\t{{{{.Name}}\t{{{{.Status}}"
