// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates auto scaling settings on your global tables at once. This operation only
// applies to Version 2019.11.21
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
// of global tables.
func (c *Client) UpdateTableReplicaAutoScaling(ctx context.Context, params *UpdateTableReplicaAutoScalingInput, optFns ...func(*Options)) (*UpdateTableReplicaAutoScalingOutput, error) {
	if params == nil {
		params = &UpdateTableReplicaAutoScalingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTableReplicaAutoScaling", params, optFns, addOperationUpdateTableReplicaAutoScalingMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTableReplicaAutoScalingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTableReplicaAutoScalingInput struct {

	// The name of the global table to be updated.
	//
	// This member is required.
	TableName *string

	// Represents the auto scaling settings of the global secondary indexes of the
	// replica to be updated.
	GlobalSecondaryIndexUpdates []*types.GlobalSecondaryIndexAutoScalingUpdate

	// Represents the auto scaling settings to be modified for a global table or global
	// secondary index.
	ProvisionedWriteCapacityAutoScalingUpdate *types.AutoScalingSettingsUpdate

	// Represents the auto scaling settings of replicas of the table that will be
	// modified.
	ReplicaUpdates []*types.ReplicaAutoScalingUpdate
}

type UpdateTableReplicaAutoScalingOutput struct {

	// Returns information about the auto scaling settings of a table with replicas.
	TableAutoScalingDescription *types.TableAutoScalingDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTableReplicaAutoScalingMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTableReplicaAutoScaling{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTableReplicaAutoScaling{}, middleware.After)
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
	addOpUpdateTableReplicaAutoScalingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTableReplicaAutoScaling(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opUpdateTableReplicaAutoScaling(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "UpdateTableReplicaAutoScaling",
	}
}
