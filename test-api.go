package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/awslambdago"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type CdkLambdaStackProps struct {
	awscdk.StackProps
}

func NewCdkLambdaStack(scope constructs.Construct, id string, props *CdkLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// Create a new api HTTP api on gateway v2.
	api := awsapigatewayv2.NewHttpApi(stack, jsii.String("cdk-lambda-api"), &awsapigatewayv2.HttpApiProps{
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowOrigins: &[]*string{jsii.String("*")}, //
			AllowMethods: &[]awsapigatewayv2.CorsHttpMethod{awsapigatewayv2.CorsHttpMethod_ANY},
		},
	})

	// Create a new lambda function.
	helloFunc := awslambdago.NewGoFunction(stack, jsii.String("hello-func"), &awslambdago.GoFunctionProps{
		MemorySize: jsii.Number(128),
		Entry:      jsii.String("./app"),
	})

	// Add a lambda proxy integration.
	integ := awsapigatewayv2integrations.NewLambdaProxyIntegration(
		&awsapigatewayv2integrations.LambdaProxyIntegrationProps{
			Handler: helloFunc,
		})

	// Add a route to api.
	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Integration: integ,
		Path:        jsii.String("/hello"),
	})

	return stack

}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkLambdaStack(app, "CdkLambdaStack", &CdkLambdaStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
