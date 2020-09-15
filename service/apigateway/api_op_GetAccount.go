// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the current Account () resource.
func (c *Client) GetAccount(ctx context.Context, params *GetAccountInput, optFns ...func(*Options)) (*GetAccountOutput, error) {
	stack := middleware.NewStack("GetAccount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetAccountMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccount(options.Region), middleware.Before)

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
			OperationName: "GetAccount",
			Err:           err,
		}
	}
	out := result.(*GetAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to get information about the current Account () resource.
type GetAccountInput struct {
	Name             *string
	Title            *string
	TemplateSkipList []*string
	Template         *bool
}

// Represents an AWS account that is associated with API Gateway. To view the
// account info, call GET on this resource.
// Error Codes
//
// The following exception
// may be thrown when the request fails.
//
//     * UnauthorizedException
//
//     *
// NotFoundException
//
//     * TooManyRequestsException
//
// For detailed error code
// information, including the corresponding HTTP Status Codes, see API Gateway
// Error Codes
// (https://docs.aws.amazon.com/apigateway/api-reference/handling-errors/#api-error-codes)
// Example:
// Get the information about an account.
//
// Request
//
//     GET /account HTTP/1.1
// Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
// X-Amz-Date: 20160531T184618Z Authorization: AWS4-HMAC-SHA256
// Credential={access_key_ID}/us-east-1/apigateway/aws4_request,
// SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
//
// Response
//
// The
// successful response returns a 200 OK status code and a payload similar to the
// following: { "_links": { "curies": { "href":
// "https://docs.aws.amazon.com/apigateway/latest/developerguide/account-apigateway-{rel}.html",
// "name": "account", "templated": true }, "self": { "href": "/account" },
// "account:update": { "href": "/account" } }, "cloudwatchRoleArn":
// "arn:aws:iam::123456789012:role/apigAwsProxyRole", "throttleSettings": {
// "rateLimit": 500, "burstLimit": 1000 } }  In addition to making the REST API
// call directly, you can use the AWS CLI and an AWS SDK to access this resource.
// API Gateway Limits
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-limits.html)Developer
// Guide
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html), AWS
// CLI
// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-account.html)
type GetAccountOutput struct {
	// A list of features supported for the account. When usage plans are enabled, the
	// features list will include an entry of "UsagePlans".
	Features []*string
	// The version of the API keys used for the account.
	ApiKeyVersion *string
	// Specifies the API request limits configured for the current Account ().
	ThrottleSettings *types.ThrottleSettings
	// The ARN of an Amazon CloudWatch role for the current Account ().
	CloudwatchRoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetAccountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetAccount{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAccount{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetAccount",
	}
}
