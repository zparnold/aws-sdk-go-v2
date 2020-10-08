// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a running model. The operation might take a while to complete. To check
// the current status, call DescribeProjectVersions ().
func (c *Client) StopProjectVersion(ctx context.Context, params *StopProjectVersionInput, optFns ...func(*Options)) (*StopProjectVersionOutput, error) {
	if params == nil {
		params = &StopProjectVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopProjectVersion", params, optFns, addOperationStopProjectVersionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopProjectVersionInput struct {

	// The Amazon Resource Name (ARN) of the model version that you want to delete.
	// This operation requires permissions to perform the
	// rekognition:StopProjectVersion action.
	//
	// This member is required.
	ProjectVersionArn *string
}

type StopProjectVersionOutput struct {

	// The current status of the stop operation.
	Status types.ProjectVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopProjectVersionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopProjectVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopProjectVersion{}, middleware.After)
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
	addOpStopProjectVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopProjectVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopProjectVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StopProjectVersion",
	}
}
