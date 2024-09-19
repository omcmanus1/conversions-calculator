# Conversions Calculator App

A simple app to calculate common US > Metric / Metric > US ingredient list conversions. Hosted at: https://conversions-calculator.vercel.app/

- Server / Backend written in Go
  - Mainly aiming to get used to language syntax & concepts
  - Serverless conversion requests using cloud functions
- Client / Frontend written in NextJS
  - Learning Typescript, NextJS and Tailwind CSS
  - Hosted on Vercel

## HOW TO RUN LOCALLY

### Server

- See [Makefile](Makefile) for all options (run in `/server/cmd` directory)
- `make run` will build and run the server for test usage if on mac
  - Or use `make run_linux` / `make_run_windows` as applicable
- `make run_dev` will run the project in development mode:
  - Sets `ENVIRONMENT` variable to "local", which causes server to be launched on localhost
  - (hot reloading with [Air](https://github.com/cosmtrek/air))
- `make test` will run the test suite
- `make_deploy` will deploy all cloud functions, and output logs to `scripts/logs/deploy_cloud_functions`
- `make clean_logs` will remove the log files

- Server is currently set up to run on port 8080, but feel free to change setup in [main.go file](server/cmd/main.go)

### Client

- Create an `.env.local` file in `/client` root, with the following variable:
  - `NEXT_PUBLIC_DEV_API_URL=http://localhost:<CHOSEN_PORT>/api`
  - Make sure the port matches up with the port in `/server/cmd/main.go`
- Run `npm run dev`
