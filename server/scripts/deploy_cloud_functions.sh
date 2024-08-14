#! /bin/bash

# This script deploys all cloud functions named with the
# corresponding route suffix as defined in `routes.go`

LIGHT_CYAN="\033[1;36m"
GREEN="\033[0;32m"
RED="\033[0;31m"
NO_COLOUR="\033[0m"

# Source environment variables
if [ ! -f ../.env.local ]; then
  echo "Error: env.local file not found!"
  exit 1
fi
# Execute file to set env variable
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

LIGHT_CYAN="\033[1;36m"
RED="\033[0;31m"
NO_COLOUR="\033[0m"

echo -e "${LIGHT_CYAN}>---------------------------------------------------------------------< ${NO_COLOUR}"

# Initialize arrays to store PIDs and corresponding function names
pids=()
names=()

# Read all `r.Post()` routes
while IFS= read -r line; do
  # Parse name from first argument
  name=$(echo "$line" | awk -F'api/ *' '{print $2}' | awk -F'["]' '{print $1}')
  # Parse entry-point from second argument
  entry_point=$(echo "$line" | awk -F', *' '{print $2}' | awk -F'[)]' '{print $1}')

  if [ -z "$name" ] || [ -z "$entry_point" ]; then
    echo -e "${RED}Warning: Skipping malformed line: $line ${NO_COLOUR}"
    continue
  fi

  echo "deploying '$name' with entrypoint '$entry_point'..."

  command="gcloud functions deploy $name --region=europe-west2 --trigger-http --runtime=go121 --gen2 --source=../. --entry-point=$entry_point --allow-unauthenticated --set-env-vars CORS_ORIGIN=$CORS_ORIGIN"

  # Execute the command in the background and capture the Process ID
  eval "$command" >"../scripts/logs/deploy_cloud_functions/${name}.log" 2>&1 &
  pid=$!
  echo "Output will be logged in 'deploy_cloud_functions/${name}.log'"
  echo -e "${LIGHT_CYAN}>---------------------------------------------------------------------< ${NO_COLOUR}"

  # Track background jobs
  # Store the PID and function name in arrays
  pids+=("$pid")
  names+=("$name")
done < <(grep 'r\.Post' ../routes.go)

# Wait for all background jobs to complete and check their exit status
for i in "${!pids[@]}"; do
  pid="${pids[$i]}"
  name="${names[$i]}"
  log_file="../scripts/logs/deploy_cloud_functions/${name}.log"

  wait "$pid"
  status=$?

  if [ $status -eq 0 ]; then
    echo "Deployment for '$name' was successful."
  else
    echo -e "${RED}!!!!!!!!!!!!!!!!!!"
    echo -e "!!!!!!!!!!!!!!!!!!${NO_COLOUR}"
    echo "Error deploying '$name'. See log from output file for details:"
    echo -e "${RED}!!!!!!!!!!!!!!!!!!"
    echo -e "!!!!!!!!!!!!!!!!!!${NO_COLOUR}"
    cat "$log_file"
    exit 1
  fi
done

echo "All deployments completed."
