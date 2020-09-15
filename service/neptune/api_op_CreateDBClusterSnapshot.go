// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a snapshot of a DB cluster.
func (c *Client) CreateDBClusterSnapshot(ctx context.Context, params *CreateDBClusterSnapshotInput, optFns ...func(*Options)) (*CreateDBClusterSnapshotOutput, error) {
	stack := middleware.NewStack("CreateDBClusterSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBClusterSnapshotMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateDBClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBClusterSnapshot(options.Region), middleware.Before)

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
			OperationName: "CreateDBClusterSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateDBClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBClusterSnapshotInput struct {
	// The tags to be assigned to the DB cluster snapshot.
	Tags []*types.Tag
	// The identifier of the DB cluster to create a snapshot for. This parameter is not
	// case-sensitive. Constraints:
	//
	//     * Must match the identifier of an existing
	// DBCluster.
	//
	// Example: my-cluster1
	DBClusterIdentifier *string
	// The identifier of the DB cluster snapshot. This parameter is stored as a
	// lowercase string. Constraints:
	//
	//     * Must contain from 1 to 63 letters,
	// numbers, or hyphens.
	//
	//     * First character must be a letter.
	//
	//     * Cannot end
	// with a hyphen or contain two consecutive hyphens.
	//
	// Example:
	// my-cluster1-snapshot1
	DBClusterSnapshotIdentifier *string
}

type CreateDBClusterSnapshotOutput struct {
	// Contains the details for an Amazon Neptune DB cluster snapshot This data type is
	// used as a response element in the DescribeDBClusterSnapshots () action.
	DBClusterSnapshot *types.DBClusterSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBClusterSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBClusterSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBClusterSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBClusterSnapshot",
	}
}
