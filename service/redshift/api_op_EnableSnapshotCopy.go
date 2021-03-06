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

// Enables the automatic copy of snapshots from one region to another region for a
// specified cluster.
func (c *Client) EnableSnapshotCopy(ctx context.Context, params *EnableSnapshotCopyInput, optFns ...func(*Options)) (*EnableSnapshotCopyOutput, error) {
	stack := middleware.NewStack("EnableSnapshotCopy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpEnableSnapshotCopyMiddlewares(stack)
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
	addOpEnableSnapshotCopyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableSnapshotCopy(options.Region), middleware.Before)
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
			OperationName: "EnableSnapshotCopy",
			Err:           err,
		}
	}
	out := result.(*EnableSnapshotCopyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type EnableSnapshotCopyInput struct {

	// The unique identifier of the source cluster to copy snapshots from. Constraints:
	// Must be the valid name of an existing cluster that does not already have
	// cross-region snapshot copy enabled.
	//
	// This member is required.
	ClusterIdentifier *string

	// The destination AWS Region that you want to copy snapshots to. Constraints: Must
	// be the name of a valid AWS Region. For more information, see Regions and
	// Endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region) in
	// the Amazon Web Services General Reference.
	//
	// This member is required.
	DestinationRegion *string

	// The number of days to retain newly copied snapshots in the destination AWS
	// Region after they are copied from the source AWS Region. If the value is -1, the
	// manual snapshot is retained indefinitely. The value must be either -1 or an
	// integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod *int32

	// The number of days to retain automated snapshots in the destination region after
	// they are copied from the source region. Default: 7. Constraints: Must be at
	// least 1 and no more than 35.
	RetentionPeriod *int32

	// The name of the snapshot copy grant to use when snapshots of an AWS
	// KMS-encrypted cluster are copied to the destination region.
	SnapshotCopyGrantName *string
}

type EnableSnapshotCopyOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpEnableSnapshotCopyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpEnableSnapshotCopy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableSnapshotCopy{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableSnapshotCopy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "EnableSnapshotCopy",
	}
}
