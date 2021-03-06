// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the function or function version, with a link to
// download the deployment package that's valid for 10 minutes. If you specify a
// function version, only details that are specific to that version are returned.
func (c *Client) GetFunction(ctx context.Context, params *GetFunctionInput, optFns ...func(*Options)) (*GetFunctionOutput, error) {
	stack := middleware.NewStack("GetFunction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetFunctionMiddlewares(stack)
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
	addOpGetFunctionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunction(options.Region), middleware.Before)
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
			OperationName: "GetFunction",
			Err:           err,
		}
	}
	out := result.(*GetFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionInput struct {

	// The name of the Lambda function, version, or alias. Name formats
	//
	//     * Function
	// name - my-function (name-only), my-function:v1 (with alias).
	//
	//     * Function ARN
	// - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN
	// - 123456789012:function:my-function.
	//
	// You can append a version number or alias
	// to any of the formats. The length constraint applies only to the full ARN. If
	// you specify only the function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify a version or alias to get details about a published version of the
	// function.
	Qualifier *string
}

type GetFunctionOutput struct {

	// The deployment package of the function or version.
	Code *types.FunctionCodeLocation

	// The function's reserved concurrency
	// (https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html).
	Concurrency *types.Concurrency

	// The configuration of the function or version.
	Configuration *types.FunctionConfiguration

	// The function's tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html).
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetFunctionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetFunction{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunction{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFunction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetFunction",
	}
}
