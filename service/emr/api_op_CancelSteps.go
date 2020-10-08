// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels a pending step or steps in a running cluster. Available only in Amazon
// EMR versions 4.8.0 and later, excluding version 5.0.0. A maximum of 256 steps
// are allowed in each CancelSteps request. CancelSteps is idempotent but
// asynchronous; it does not guarantee a step will be canceled, even if the request
// is successfully submitted. You can only cancel steps that are in a PENDING
// state.
func (c *Client) CancelSteps(ctx context.Context, params *CancelStepsInput, optFns ...func(*Options)) (*CancelStepsOutput, error) {
	if params == nil {
		params = &CancelStepsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelSteps", params, optFns, addOperationCancelStepsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelStepsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input argument to the CancelSteps () operation.
type CancelStepsInput struct {

	// The ClusterID for which specified steps will be canceled. Use RunJobFlow () and
	// ListClusters () to get ClusterIDs.
	//
	// This member is required.
	ClusterId *string

	// The list of StepIDs to cancel. Use ListSteps () to get steps and their states
	// for the specified cluster.
	//
	// This member is required.
	StepIds []*string

	// The option to choose for cancelling RUNNING steps. By default, the value is
	// SEND_INTERRUPT.
	StepCancellationOption types.StepCancellationOption
}

// The output for the CancelSteps () operation.
type CancelStepsOutput struct {

	// A list of CancelStepsInfo (), which shows the status of specified cancel
	// requests for each StepID specified.
	CancelStepsInfoList []*types.CancelStepsInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelStepsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelSteps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelSteps{}, middleware.After)
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
	addOpCancelStepsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelSteps(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCancelSteps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "CancelSteps",
	}
}
