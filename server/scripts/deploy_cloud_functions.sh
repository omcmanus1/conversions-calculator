#! /bin/bash

# This script deploys all cloud functions named with the
# corresponding route suffix as defined in `routes.go`

LIGHT_CYAN="\033[1;36m"
GREEN="\033[0;32m"
RED="\033[0;31m"
NO_COLOUR="\033[0m"

# Source environment variables
if [[ ! -f ../.env.local ]]; then
  echo "Error: env.local file not found!"
  exit 1
fi
. ../.env.local

# Check CORS_ORIGIN is set
if [[ -z "$CORS_ORIGIN" ]]; then
  echo "Error: CORS_ORIGIN is not set!"
  exit 1
fi

# Ensure target files exist
for file in ../routes.go ../functions.go; do
  if [[ ! -f $file ]]; then
    echo "Error: $file file not found!"
    exit 1
  fi
done

# Read all existing functions
existing_funcs=()
while IFS= read -r line; do
  existing_funcs+=("$line")
done < <(grep 'functions\.HTTP' ../functions.go)

add_func_if_missing() {
  # Search existing functions for match with entry point
  local found=false
  for func in "${existing_funcs[@]}"; do
    if [[ "$func" == *"$entry_point"* ]]; then
      found=true
      break
    fi
  done
  # Add function to 'functions.go' file and proceed with deployment
  if [[ "$found" == false ]]; then
    echo -e "${RED}!!!!!!!!!!!!!!!!!!${NO_COLOUR}"
    echo "Entry point $entry_point missing in functions list"
    echo -e "${RED}!!!!!!!!!!!!!!!!!!${NO_COLOUR}"
    echo -e "${GREEN}Adding to 'functions.go'... ${NO_COLOUR}"
    sed -i '' "/func init() {/a\\
  functions.HTTP(\"$entry_point\", requestHandler($entry_point))
    " ../functions.go
    echo -e "${GREEN}Successfully added $entry_point to functions.go${NO_COLOUR}"
  fi
}

echo -e "${LIGHT_CYAN}>---------------------------------------------------------------------< ${NO_COLOUR}"

# Initialize arrays to store PIDs and corresponding function names
pids=()
names=()

# Ensure log directory exists
log_dir="../scripts/logs/deploy_cloud_functions"
mkdir -p "$log_dir"

# Read all `r.Post()` routes
while IFS= read -r line; do
  # Parse name and entry-point
  name=$(echo "$line" | awk -F'api/ *' '{print $2}' | awk -F'["]' '{print $1}')
  entry_point=$(echo "$line" | awk -F', *' '{print $2}' | awk -F'[)]' '{print $1}')

  if [[ -z "$name" || -z "$entry_point" ]]; then
    echo -e "${RED}Warning: Skipping malformed line: $line ${NO_COLOUR}"
    continue
  fi

  add_func_if_missing

  echo "Deploying '$name' with entrypoint '$entry_point'..."

  command="gcloud functions deploy $name --region=europe-west2 --trigger-http --runtime=go121 --gen2 --source=../. --entry-point=$entry_point --allow-unauthenticated --set-env-vars CORS_ORIGIN=$CORS_ORIGIN"

  # Execute the command in the background and capture the Process ID
  eval "$command" >"$log_dir/${name}.log" 2>&1 &
  pid=$!
  echo "Output will be logged in '$log_dir/${name}.log'"
  echo -e "${LIGHT_CYAN}>---------------------------------------------------------------------< ${NO_COLOUR}"

  # Track background jobs
  pids+=("$pid")
  names+=("$name")
done < <(grep 'r\.Post' ../routes.go)

# Wait for all background jobs to complete and check their exit status
for i in "${!pids[@]}"; do
  pid="${pids[$i]}"
  name="${names[$i]}"
  log_file="$log_dir/${name}.log"

  wait "$pid"
  status=$?

  if [[ $status -eq 0 ]]; then
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
