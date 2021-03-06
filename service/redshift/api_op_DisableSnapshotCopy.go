// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the automatic copying of snapshots from one region to another region
// for a specified cluster. If your cluster and its snapshots are encrypted using a
// customer master key (CMK) from AWS KMS, use DeleteSnapshotCopyGrant () to delete
// the grant that grants Amazon Redshift permission to the CMK in the destination
// region.
func (c *Client) DisableSnapshotCopy(ctx context.Context, params *DisableSnapshotCopyInput, optFns ...func(*Options)) (*DisableSnapshotCopyOutput, error) {
	stack := middleware.NewStack("DisableSnapshotCopy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDisableSnapshotCopyMiddlewares(stack)
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
	addOpDisableSnapshotCopyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableSnapshotCopy(options.Region), middleware.Before)
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
			OperationName: "DisableSnapshotCopy",
			Err:           err,
		}
	}
	out := result.(*DisableSnapshotCopyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DisableSnapshotCopyInput struct {

	// The unique identifier of the source cluster that you want to disable copying of
	// snapshots to a destination region. Constraints: Must be the valid name of an
	// existing cluster that has cross-region snapshot copy enabled.
	//
	// This member is required.
	ClusterIdentifier *string
}

type DisableSnapshotCopyOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDisableSnapshotCopyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDisableSnapshotCopy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableSnapshotCopy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableSnapshotCopy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DisableSnapshotCopy",
	}
}
