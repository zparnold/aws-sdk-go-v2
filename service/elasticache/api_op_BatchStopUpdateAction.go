// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stop the service update. For more information on service updates and stopping
// them, see Stopping Service Updates
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/stopping-self-service-updates.html).
func (c *Client) BatchStopUpdateAction(ctx context.Context, params *BatchStopUpdateActionInput, optFns ...func(*Options)) (*BatchStopUpdateActionOutput, error) {
	if params == nil {
		params = &BatchStopUpdateActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchStopUpdateAction", params, optFns, addOperationBatchStopUpdateActionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchStopUpdateActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchStopUpdateActionInput struct {

	// The unique ID of the service update
	//
	// This member is required.
	ServiceUpdateName *string

	// The cache cluster IDs
	CacheClusterIds []*string

	// The replication group IDs
	ReplicationGroupIds []*string
}

type BatchStopUpdateActionOutput struct {

	// Update actions that have been processed successfully
	ProcessedUpdateActions []*types.ProcessedUpdateAction

	// Update actions that haven't been processed successfully
	UnprocessedUpdateActions []*types.UnprocessedUpdateAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchStopUpdateActionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBatchStopUpdateAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchStopUpdateAction{}, middleware.After)
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
	addOpBatchStopUpdateActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchStopUpdateAction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchStopUpdateAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "BatchStopUpdateAction",
	}
}
