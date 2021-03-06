// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the task logs for the specified task.
func (c *Client) DescribeReplicationInstanceTaskLogs(ctx context.Context, params *DescribeReplicationInstanceTaskLogsInput, optFns ...func(*Options)) (*DescribeReplicationInstanceTaskLogsOutput, error) {
	stack := middleware.NewStack("DescribeReplicationInstanceTaskLogs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeReplicationInstanceTaskLogsMiddlewares(stack)
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
	addOpDescribeReplicationInstanceTaskLogsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationInstanceTaskLogs(options.Region), middleware.Before)
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
			OperationName: "DescribeReplicationInstanceTaskLogs",
			Err:           err,
		}
	}
	out := result.(*DescribeReplicationInstanceTaskLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicationInstanceTaskLogsInput struct {

	// The Amazon Resource Name (ARN) of the replication instance.
	//
	// This member is required.
	ReplicationInstanceArn *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeReplicationInstanceTaskLogsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn *string

	// An array of replication task log metadata. Each member of the array contains the
	// replication task name, ARN, and task log size (in bytes).
	ReplicationInstanceTaskLogs []*types.ReplicationInstanceTaskLog

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeReplicationInstanceTaskLogsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationInstanceTaskLogs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationInstanceTaskLogs{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeReplicationInstanceTaskLogs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeReplicationInstanceTaskLogs",
	}
}
