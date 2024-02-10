# Go Backend

A basic Go server with endpoints to translate US recipe measurements into UK formats.

## HOW TO RUN LOCALLY

- See [Makefile](Makefile) for all options
- `make run` will build and run the server for test usage if on mac
  - Or use `make run_linux` / `make_run_windows` as applicable
- `make run_dev` will run the project in development mode (hot reloading with [Air](https://github.com/cosmtrek/air))
- `make test` will run the test suite

- Server is currently set up to run on port 8080, but feel free to change setup in [main file](main.go)

## TODO

- ~~Update RecipeList function to accept array of Setup structs for full ingredient list~~
- Testing
  - ~~Add first test~~
  - ~~Check / fix output rounding~~
- Create server endpoints for potential frontend usage
  - ~~Initial GET requests~~
  - ~~Return JSON~~
  - ~~Add standalone POST requests per input~~
  - ~~Implement JSON response for each request~~
  - ~~Add CORS (if necessary)?~~
  - ~~Check error handling~~
- Write basic endpoint docs
- Refactor
  - ~~Abstract JSON logic from routes into util functions~~
  - ~~Refactor if chain - switch statements for readability~~
- Simplify Logic
  - Automatically return most appropriate units based on amount (e.g. 1.5l vs 1500ml) if no OutputUnit is provided?
- New Features
  - ~~Add height conversion endpoints~~
