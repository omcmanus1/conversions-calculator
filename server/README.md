# Go Backend
A basic Go server with endpoints to translate US recipe measurements into UK formats. 

## TODO
- ~~Update Flow function to accept array of Setup structs for full ingredient list~~
- Testing
  - ~~Add first test~~
  - Learn more / complete
- Create server endpoints for potential frontend usage
  - ~~Initial GET requests~~
  - ~~Return JSON~~
  - ~~Add standalone POST requests per input~~
  - ~~Implement JSON response for each request~~
  - Check error handling
- Refactor
  - ~~Abstract JSON logic from routes into util functions~~
  - Create controllers for each endpoint (input / output validation checks)
- Simplify Logic
  - Automatically return most appropriate units based on amount (e.g. 1.5l vs 1500ml) if no OutputUnit is provided?
- New Features
  - Add list of recipe links with example calcs
- New Pages
  - Heights/weights


