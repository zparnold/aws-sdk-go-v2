// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a set of cluster snapshots.
func (c *Client) BatchDeleteClusterSnapshots(ctx context.Context, params *BatchDeleteClusterSnapshotsInput, optFns ...func(*Options)) (*BatchDeleteClusterSnapshotsOutput, error) {
	if params == nil {
		params = &BatchDeleteClusterSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteClusterSnapshots", params, optFns, addOperationBatchDeleteClusterSnapshotsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteClusterSnapshotsInput struct {

	// A list of identifiers for the snapshots that you want to delete.
	//
	// This member is required.
	Identifiers []*types.DeleteClusterSnapshotMessage
}

type BatchDeleteClusterSnapshotsOutput struct {

	// A list of any errors returned.
	Errors []*types.SnapshotErrorMessage

	// A list of the snapshot identifiers that were deleted.
	Resources []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeleteClusterSnapshotsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBatchDeleteClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchDeleteClusterSnapshots{}, middleware.After)
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
	addOpBatchDeleteClusterSnapshotsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteClusterSnapshots(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchDeleteClusterSnapshots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "BatchDeleteClusterSnapshots",
	}
}
