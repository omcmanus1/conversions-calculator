# Recipe Measurements Converter - Go Backend

A basic Golang app to translate US recipe measurements into UK formats. Mainly aiming to get used to Golang syntax & concepts

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

- Build frontend calculator?
