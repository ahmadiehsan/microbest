#!/bin/bash

kibana_address="http://127.0.0.1:${KIBANA_PORT}/${KIBANA_PUBLIC_BASE_PATH}"

# Wait for Kibana to be ready
until curl -s -X GET "${kibana_address}/api/status" >/dev/null; do
  echo "[INFO ][import.sh] Waiting for Kibana to be ready..."
  sleep 5
done

# Import
import_file() {
  local file="$1"
  local response
  response=$(curl -s -w "%{http_code}" -X POST "${kibana_address}/api/saved_objects/_import" -H "kbn-xsrf: true" --form file=@"$file")
  echo "$response"
}

import_objects() {
  local dir="$1"
  local type="$2"

  echo "[INFO ][import.sh] Importing ${type}..."

  for file in "/etc/provisioning/${dir}/"*.ndjson; do
    echo "[INFO ][import.sh] Processing $file..."
    local response
    response=$(import_file "$file")
    http_code=${response: -3}
    body=${response:0:-3}

    if [ "$http_code" = "415" ]; then
      echo "[INFO ][import.sh] Got 415 status code, retrying..."

      for retry in {1..10}; do
        echo "[INFO ][import.sh] Retry attempt $retry..."
        response=$(import_file "$file")
        http_code=${response: -3}
        body=${response:0:-3}

        if [ "$http_code" = "200" ]; then
          break
        fi

        sleep 5
      done
    fi

    if [ "$http_code" != "200" ]; then
      echo "[ERROR][import.sh] importing $file | Status code: $http_code"
      echo "[ERROR][import.sh] Response: $body"
    fi
  done
}

import_objects "data_views" "Data Views"
import_objects "dashboards" "Dashboards"

echo "[INFO ][import.sh] Import completed"
