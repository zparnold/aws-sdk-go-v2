// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a batch of variables.
func (c *Client) BatchCreateVariable(ctx context.Context, params *BatchCreateVariableInput, optFns ...func(*Options)) (*BatchCreateVariableOutput, error) {
	if params == nil {
		params = &BatchCreateVariableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCreateVariable", params, optFns, addOperationBatchCreateVariableMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCreateVariableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateVariableInput struct {

	// The list of variables for the batch create variable request.
	//
	// This member is required.
	VariableEntries []*types.VariableEntry

	// A collection of key and value pairs.
	Tags []*types.Tag
}

type BatchCreateVariableOutput struct {

	// Provides the errors for the BatchCreateVariable request.
	Errors []*types.BatchCreateVariableError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchCreateVariableMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchCreateVariable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchCreateVariable{}, middleware.After)
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
	addOpBatchCreateVariableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateVariable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchCreateVariable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "BatchCreateVariable",
	}
}
