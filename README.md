# Recipe Converter App
A simple app to calculate common US > Metric / Metric > US ingredient list conversions. 

- Server / Backend written in Go
  - Mainly aiming to get used to language syntax & concepts
- Client / Frontend written in NextJS
  - Learning Typescript, NextJS and Tailwind CSS

## HOW TO RUN LOCALLY

### Server
- See [Makefile](Makefile) for all options (run in `/server/cmd` directory)
- `make run` will build and run the server for test usage if on mac
  - Or use `make run_linux` / `make_run_windows` as applicable
- `make run_dev` will run the project in development mode: 
  - Sets `ENVIRONMENT` variable to "local", which causes server to be launched on localhost
  - (hot reloading with [Air](https://github.com/cosmtrek/air))
- `make test` will run the test suite

- Server is currently set up to run on port 8080, but feel free to change setup in [main file](main.go)

### Client
- Create an `.env.local` file in `/client` root, with the following variable: 
  - `NEXT_PUBLIC_DEV_API_URL=http://localhost:<CHOSEN_PORT>/api`
- If you want to create your own cloud functions, generate URLs with suffixes corresponding to the route paths defined in `/server/routes.go`, then set the below NextJS environment variable as the URL path up to (excluding) the endpoint suffix
  - `NEXT_PUBLIC_PROD_API_URL=<YOUR_CLOUD_FUNCTION_URL>`