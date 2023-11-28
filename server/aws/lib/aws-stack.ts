import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";
import { RestApi, LambdaIntegration } from "aws-cdk-lib/aws-apigateway";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class AwsStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const myFunction = new lambda.Function(this, "MyLambda", {
      code: lambda.Code.fromAsset("../bin/build"),
      handler: "main",
      runtime: lambda.Runtime.PROVIDED_AL2023,
    });

    const gateway = new RestApi(this, "myGateway", {
      defaultCorsPreflightOptions: {
        allowOrigins: ["*"],
        allowMethods: ["GET", "POST"],
      },
    });

    const integration = new LambdaIntegration(myFunction);
    const resource = gateway.root.addResource("converter");
    resource.addMethod("GET", integration);
    resource.addMethod("POST", integration);
  }
}
