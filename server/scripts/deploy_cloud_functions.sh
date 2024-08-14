#! /bin/bash

# This script deploys all cloud functions named with the
# corresponding route suffix as defined in `routes.go`

RED='\033[0;31m'
NO_COLOUR='\033[0m'

# Source environment variables
if [ ! -f ../.env.local ]; then
  echo "Error: env.local file not found!"
  exit 1
fi
# execute file to set env variable
. ../.env.local

# Check CORS_ORIGIN is set
if [ -z "$CORS_ORIGIN" ]; then
  echo "Error: CORS_ORIGIN is not set!"
  exit 1
fi

# Ensure target files exist
if [ ! -f ../routes.go ]; then
  echo "Error: routes.go file not found!"
  exit 1
fi
if [ ! -f ../functions.go ]; then
  echo "Error: functions.go file not found!"
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

  echo ">---------------------------------------------------------------------<"
  echo "deploying '$name' with entrypoint '$entry_point'..."

  command="gcloud functions deploy $name --region=europe-west222 --trigger-http --runtime=go121 --gen2 --source=../. --entry-point=$entry_point --allow-unauthenticated --set-env-vars CORS_ORIGIN=$CORS_ORIGIN"

  # Execute the command in the background and capture the Process ID
  eval "$command" >"../scripts/logs/deploy_cloud_functions/${name}.log" 2>&1 &
  pid=$!
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
    echo -e "${RED}!!!!!!!!!!!!!!!!!!"
    echo -e "!!!!!!!!!!!!!!!!!!${NO_COLOUR}"
    echo "Error deploying function for process $pid. See the corresponding log for details."
    echo -e "${RED}!!!!!!!!!!!!!!!!!!"
    echo -e "!!!!!!!!!!!!!!!!!!${NO_COLOUR}"
    set -e
  fi
done

echo "All deployments completed."
