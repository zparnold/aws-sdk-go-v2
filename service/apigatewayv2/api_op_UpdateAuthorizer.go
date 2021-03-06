// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an Authorizer.
func (c *Client) UpdateAuthorizer(ctx context.Context, params *UpdateAuthorizerInput, optFns ...func(*Options)) (*UpdateAuthorizerOutput, error) {
	stack := middleware.NewStack("UpdateAuthorizer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateAuthorizerMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateAuthorizerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAuthorizer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "UpdateAuthorizer",
			Err:           err,
		}
	}
	out := result.(*UpdateAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates an Authorizer.
type UpdateAuthorizerInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The authorizer identifier.
	//
	// This member is required.
	AuthorizerId *string

	// Specifies the required credentials as an IAM role for API Gateway to invoke the
	// authorizer. To specify an IAM role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To use resource-based permissions on the Lambda
	// function, specify null.
	AuthorizerCredentialsArn *string

	// Authorizer caching is not currently supported. Don't specify this value for
	// authorizers.
	AuthorizerResultTtlInSeconds *int32

	// The authorizer type. For WebSocket APIs, specify REQUEST for a Lambda function
	// using incoming request parameters. For HTTP APIs, specify JWT to use JSON Web
	// Tokens.
	AuthorizerType types.AuthorizerType

	// The authorizer's Uniform Resource Identifier (URI). For REQUEST authorizers,
	// this must be a well-formed Lambda function URI, for example,
	// arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form:
	// arn:aws:apigateway:{region}:lambda:path/{service_api} , where {region} is the
	// same as the region hosting the Lambda function, path indicates that the
	// remaining substring in the URI should be treated as the path to the resource,
	// including the initial /. For Lambda functions, this is usually of the form
	// /2015-03-31/functions/[FunctionARN]/invocations. Supported only for REQUEST
	// authorizers.
	AuthorizerUri *string

	// The identity source for which authorization is requested. For a REQUEST
	// authorizer, this is optional. The value is a set of one or more mapping
	// expressions of the specified request parameters. Currently, the identity source
	// can be headers, query string parameters, stage variables, and context
	// parameters. For example, if an Auth header and a Name query string parameter are
	// defined as identity sources, this value is route.request.header.Auth,
	// route.request.querystring.Name. These parameters will be used to perform runtime
	// validation for Lambda-based authorizers by verifying all of the identity-related
	// request parameters are present in the request, not null, and non-empty. Only
	// when this is true does the authorizer invoke the authorizer Lambda function.
	// Otherwise, it returns a 401 Unauthorized response without calling the Lambda
	// function. For JWT, a single entry that specifies where to extract the JSON Web
	// Token (JWT) from inbound requests. Currently only header-based and query
	// parameter-based selections are supported, for example
	// "$request.header.Authorization".
	IdentitySource []*string

	// This parameter is not used.
	IdentityValidationExpression *string

	// Represents the configuration of a JWT authorizer. Required for the JWT
	// authorizer type. Supported only for HTTP APIs.
	JwtConfiguration *types.JWTConfiguration

	// The name of the authorizer.
	Name *string
}

type UpdateAuthorizerOutput struct {

	// Specifies the required credentials as an IAM role for API Gateway to invoke the
	// authorizer. To specify an IAM role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To use resource-based permissions on the Lambda
	// function, specify null. Supported only for REQUEST authorizers.
	AuthorizerCredentialsArn *string

	// The authorizer identifier.
	AuthorizerId *string

	// Authorizer caching is not currently supported. Don't specify this value for
	// authorizers.
	AuthorizerResultTtlInSeconds *int32

	// The authorizer type. For WebSocket APIs, specify REQUEST for a Lambda function
	// using incoming request parameters. For HTTP APIs, specify JWT to use JSON Web
	// Tokens.
	AuthorizerType types.AuthorizerType

	// The authorizer's Uniform Resource Identifier (URI). ForREQUEST authorizers, this
	// must be a well-formed Lambda function URI, for example,
	// arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form:
	// arn:aws:apigateway:{region}:lambda:path/{service_api} , where {region} is the
	// same as the region hosting the Lambda function, path indicates that the
	// remaining substring in the URI should be treated as the path to the resource,
	// including the initial /. For Lambda functions, this is usually of the form
	// /2015-03-31/functions/[FunctionARN]/invocations. Supported only for REQUEST
	// authorizers.
	AuthorizerUri *string

	// The identity source for which authorization is requested. For a REQUEST
	// authorizer, this is optional. The value is a set of one or more mapping
	// expressions of the specified request parameters. Currently, the identity source
	// can be headers, query string parameters, stage variables, and context
	// parameters. For example, if an Auth header and a Name query string parameter are
	// defined as identity sources, this value is route.request.header.Auth,
	// route.request.querystring.Name. These parameters will be used to perform runtime
	// validation for Lambda-based authorizers by verifying all of the identity-related
	// request parameters are present in the request, not null, and non-empty. Only
	// when this is true does the authorizer invoke the authorizer Lambda function.
	// Otherwise, it returns a 401 Unauthorized response without calling the Lambda
	// function. For JWT, a single entry that specifies where to extract the JSON Web
	// Token (JWT) from inbound requests. Currently only header-based and query
	// parameter-based selections are supported, for example
	// "$request.header.Authorization".
	IdentitySource []*string

	// The validation expression does not apply to the REQUEST authorizer.
	IdentityValidationExpression *string

	// Represents the configuration of a JWT authorizer. Required for the JWT
	// authorizer type. Supported only for HTTP APIs.
	JwtConfiguration *types.JWTConfiguration

	// The name of the authorizer.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateAuthorizerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAuthorizer{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAuthorizer{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAuthorizer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateAuthorizer",
	}
}
