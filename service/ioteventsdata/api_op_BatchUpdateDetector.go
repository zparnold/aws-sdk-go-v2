// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the state, variable values, and timer settings of one or more detectors
// (instances) of a specified detector model.
func (c *Client) BatchUpdateDetector(ctx context.Context, params *BatchUpdateDetectorInput, optFns ...func(*Options)) (*BatchUpdateDetectorOutput, error) {
	if params == nil {
		params = &BatchUpdateDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateDetector", params, optFns, addOperationBatchUpdateDetectorMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateDetectorInput struct {

	// The list of detectors (instances) to update, along with the values to update.
	//
	// This member is required.
	Detectors []*types.UpdateDetectorRequest
}

type BatchUpdateDetectorOutput struct {

	// A list of those detector updates that resulted in errors. (If an error is listed
	// here, the specific update did not occur.)
	BatchUpdateDetectorErrorEntries []*types.BatchUpdateDetectorErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchUpdateDetectorMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateDetector{}, middleware.After)
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
	addOpBatchUpdateDetectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateDetector(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchUpdateDetector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ioteventsdata",
		OperationName: "BatchUpdateDetector",
	}
}
