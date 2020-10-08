// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a test.
func (c *Client) GetTest(ctx context.Context, params *GetTestInput, optFns ...func(*Options)) (*GetTestOutput, error) {
	if params == nil {
		params = &GetTestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTest", params, optFns, addOperationGetTestMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the get test operation.
type GetTestInput struct {

	// The test's ARN.
	//
	// This member is required.
	Arn *string
}

// Represents the result of a get test request.
type GetTestOutput struct {

	// A test condition that is evaluated.
	Test *types.Test

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTestMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTest{}, middleware.After)
	if err != nil {
		return err
	}
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
	addOpGetTestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTest(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetTest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetTest",
	}
}
