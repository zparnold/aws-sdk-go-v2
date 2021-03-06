// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets an Amazon Kinesis shard iterator. A shard iterator expires 5 minutes after
// it is returned to the requester. A shard iterator specifies the shard position
// from which to start reading data records sequentially. The position is specified
// using the sequence number of a data record in a shard. A sequence number is the
// identifier associated with every record ingested in the stream, and is assigned
// when a record is put into the stream. Each stream has one or more shards. You
// must specify the shard iterator type. For example, you can set the
// ShardIteratorType parameter to read exactly from the position denoted by a
// specific sequence number by using the AT_SEQUENCE_NUMBER shard iterator type.
// Alternatively, the parameter can read right after the sequence number by using
// the AFTER_SEQUENCE_NUMBER shard iterator type, using sequence numbers returned
// by earlier calls to PutRecord (), PutRecords (), GetRecords (), or
// DescribeStream (). In the request, you can specify the shard iterator type
// AT_TIMESTAMP to read records from an arbitrary point in time, TRIM_HORIZON to
// cause ShardIterator to point to the last untrimmed record in the shard in the
// system (the oldest data record in the shard), or LATEST so that you always read
// the most recent data in the shard. When you read repeatedly from a stream, use a
// GetShardIterator () request to get the first shard iterator for use in your
// first GetRecords () request and for subsequent reads use the shard iterator
// returned by the GetRecords () request in NextShardIterator. A new shard iterator
// is returned by every GetRecords () request in NextShardIterator, which you use
// in the ShardIterator parameter of the next GetRecords () request. If a
// GetShardIterator () request is made too often, you receive a
// ProvisionedThroughputExceededException. For more information about throughput
// limits, see GetRecords (), and Streams Limits
// (https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide. If the shard is closed,
// GetShardIterator () returns a valid iterator for the last sequence number of the
// shard. A shard can be closed as a result of using SplitShard () or MergeShards
// (). GetShardIterator () has a limit of five transactions per second per account
// per open shard.
func (c *Client) GetShardIterator(ctx context.Context, params *GetShardIteratorInput, optFns ...func(*Options)) (*GetShardIteratorOutput, error) {
	stack := middleware.NewStack("GetShardIterator", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetShardIteratorMiddlewares(stack)
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
	addOpGetShardIteratorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetShardIterator(options.Region), middleware.Before)
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
			OperationName: "GetShardIterator",
			Err:           err,
		}
	}
	out := result.(*GetShardIteratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for GetShardIterator.
type GetShardIteratorInput struct {

	// The shard ID of the Kinesis Data Streams shard to get the iterator for.
	//
	// This member is required.
	ShardId *string

	// Determines how the shard iterator is used to start reading data records from the
	// shard. The following are the valid Amazon Kinesis shard iterator types:
	//
	//     *
	// AT_SEQUENCE_NUMBER - Start reading from the position denoted by a specific
	// sequence number, provided in the value StartingSequenceNumber.
	//
	//     *
	// AFTER_SEQUENCE_NUMBER - Start reading right after the position denoted by a
	// specific sequence number, provided in the value StartingSequenceNumber.
	//
	//     *
	// AT_TIMESTAMP - Start reading from the position denoted by a specific time stamp,
	// provided in the value Timestamp.
	//
	//     * TRIM_HORIZON - Start reading at the last
	// untrimmed record in the shard in the system, which is the oldest data record in
	// the shard.
	//
	//     * LATEST - Start reading just after the most recent record in
	// the shard, so that you always read the most recent data in the shard.
	//
	// This member is required.
	ShardIteratorType types.ShardIteratorType

	// The name of the Amazon Kinesis data stream.
	//
	// This member is required.
	StreamName *string

	// The sequence number of the data record in the shard from which to start reading.
	// Used with shard iterator type AT_SEQUENCE_NUMBER and AFTER_SEQUENCE_NUMBER.
	StartingSequenceNumber *string

	// The time stamp of the data record from which to start reading. Used with shard
	// iterator type AT_TIMESTAMP. A time stamp is the Unix epoch date with precision
	// in milliseconds. For example, 2016-04-04T19:58:46.480-00:00 or 1459799926.480.
	// If a record with this exact time stamp does not exist, the iterator returned is
	// for the next (later) record. If the time stamp is older than the current trim
	// horizon, the iterator returned is for the oldest untrimmed data record
	// (TRIM_HORIZON).
	Timestamp *time.Time
}

// Represents the output for GetShardIterator.
type GetShardIteratorOutput struct {

	// The position in the shard from which to start reading data records sequentially.
	// A shard iterator specifies this position using the sequence number of a data
	// record in a shard.
	ShardIterator *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetShardIteratorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetShardIterator{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetShardIterator{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetShardIterator(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "GetShardIterator",
	}
}
