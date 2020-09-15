// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a particular global replication group. If no
// identifier is specified, returns information about all Global Datastores.
func (c *Client) DescribeGlobalReplicationGroups(ctx context.Context, params *DescribeGlobalReplicationGroupsInput, optFns ...func(*Options)) (*DescribeGlobalReplicationGroupsOutput, error) {
	stack := middleware.NewStack("DescribeGlobalReplicationGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeGlobalReplicationGroupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalReplicationGroups(options.Region), middleware.Before)

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
			OperationName: "DescribeGlobalReplicationGroups",
			Err:           err,
		}
	}
	out := result.(*DescribeGlobalReplicationGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalReplicationGroupsInput struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	MaxRecords *int32
	// Returns the list of members that comprise the Global Datastore.
	ShowMemberInfo *bool
	// The name of the Global Datastore
	GlobalReplicationGroupId *string
	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
}

type DescribeGlobalReplicationGroupsOutput struct {
	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords. >
	Marker *string
	// Indicates the slot configuration and global identifier for each slice group.
	GlobalReplicationGroups []*types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeGlobalReplicationGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeGlobalReplicationGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeGlobalReplicationGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGlobalReplicationGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeGlobalReplicationGroups",
	}
}
