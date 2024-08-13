# Go Backend

A basic Go server with endpoints to calculate conversions from US measurements to metric format (recipes, heights etc.). Runs using a local server in development mode, and cloud functions in production.

## Deploying Cloud Functions

- Ensure `CORS_ORIGIN` environment variable is set to the domain where the app is hosted (either via script or in standalone gcloud CLI command)
- Sample CLI command (`entry-point` must match first argument in `functions.HTTP()` call):

```
gcloud functions deploy height-metric --region=europe-west2 --trigger-http --runtime=go121 --gen2 --source=. --entry-point=PostHeightMetric --allow-unauthenticated --set-env-vars CORS_ORIGIN=https://recipe-converter-ebon.vercel.app
```
