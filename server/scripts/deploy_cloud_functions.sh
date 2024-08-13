#! /bin/bash

# This script deploys all cloud functions named with the
# corresponding route suffix as defined in `routes.go`

# Ensure the script fails on any error
set -e

# Source environment variables
if [ ! -f ../.env.local ]; then
  echo "Error: ../.env.local file not found!"
  exit 1
fi
# shellcheck source=../.env.local
. ../.env.local

# Check CORS_ORIGIN is set
if [ -z "$CORS_ORIGIN" ]; then
  echo "Error: CORS_ORIGIN is not set!"
  exit 1
fi

# Ensure routes.go exists
if [ ! -f ../routes.go ]; then
  echo "Error: routes.go file not found!"
  exit 1
fi

# Initialize an empty list of PIDs
pids=""

# read all `r.Post()` routes from `SetupRoutes()`
while IFS= read -r line; do
  # parse name from first argument
  name=$(echo "$line" | awk -F'api/ *' '{print $2}' | awk -F'["]' '{print $1}')
  # parse entry-point from second argument
  entry_point=$(echo "$line" | awk -F', *' '{print $2}' | awk -F'[)]' '{print $1}')

  if [ -z "$name" ] || [ -z "$entry_point" ]; then
    echo "Warning: Skipping malformed line: $line"
    continue
  fi

  echo "deploying '$name' with entrypoint '$entry_point'..."

  command="gcloud functions deploy $name --region=europe-west2 --trigger-http --runtime=go121 --gen2 --source=../. --entry-point=$entry_point --allow-unauthenticated --set-env-vars CORS_ORIGIN=$CORS_ORIGIN"

  # Execute the command in the background and capture the PID
  eval "$command" >"../scripts/logs/deploy_cloud_functions/${name}.log" 2>&1 &
  pid=$!
  echo "Started deployment for '$name'."
  echo "Output will be logged in ../scripts/logs/deploy_cloud_functions/${name}.log"

  # Track the background jobs
  pids="$pids $pid"
done < <(grep 'r\.Post' ../routes.go)

# Wait for all background jobs to complete and check their exit status
for pid in $pids; do
  wait "$pid"
  status=$?
  if [ $status -eq 0 ]; then
    echo "Deployment for process $pid was successful."
  else
    echo "Error deploying function for process $pid. See the corresponding log for details."
  fi
done

echo "All deployments completed."
