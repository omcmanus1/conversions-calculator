# Go Backend

A basic Go server with endpoints to calculate conversions from US measurements to metric format (recipes, heights etc.). Runs using a local server in development mode, and cloud functions in production.

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
