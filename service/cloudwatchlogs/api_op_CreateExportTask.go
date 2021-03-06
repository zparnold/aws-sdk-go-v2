// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an export task, which allows you to efficiently export data from a log
// group to an Amazon S3 bucket. This is an asynchronous call. If all the required
// information is provided, this operation initiates an export task and responds
// with the ID of the task. After the task has started, you can use
// DescribeExportTasks
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeExportTasks.html)
// to get the status of the export task. Each account can only have one active
// (RUNNING or PENDING) export task at a time. To cancel an export task, use
// CancelExportTask
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CancelExportTask.html).
// You can export logs from multiple log groups or multiple time ranges to the same
// S3 bucket. To separate out log data for each export task, you can specify a
// prefix to be used as the Amazon S3 key prefix for all exported objects.
// Exporting to S3 buckets that are encrypted with AES-256 is supported. Exporting
// to S3 buckets encrypted with SSE-KMS is not supported.
func (c *Client) CreateExportTask(ctx context.Context, params *CreateExportTaskInput, optFns ...func(*Options)) (*CreateExportTaskOutput, error) {
	stack := middleware.NewStack("CreateExportTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateExportTaskMiddlewares(stack)
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
	addOpCreateExportTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExportTask(options.Region), middleware.Before)
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
			OperationName: "CreateExportTask",
			Err:           err,
		}
	}
	out := result.(*CreateExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExportTaskInput struct {

	// The name of S3 bucket for the exported log data. The bucket must be in the same
	// AWS region.
	//
	// This member is required.
	Destination *string

	// The start time of the range for the request, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC. Events with a timestamp earlier
	// than this time are not exported.
	//
	// This member is required.
	From *int64

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The end time of the range for the request, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC. Events with a timestamp later than
	// this time are not exported.
	//
	// This member is required.
	To *int64

	// The prefix used as the start of the key for every object exported. If you don't
	// specify a value, the default is exportedlogs.
	DestinationPrefix *string

	// Export only log streams that match the provided prefix. If you don't specify a
	// value, no prefix filter is applied.
	LogStreamNamePrefix *string

	// The name of the export task.
	TaskName *string
}

type CreateExportTaskOutput struct {

	// The ID of the export task.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateExportTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateExportTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateExportTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateExportTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "CreateExportTask",
	}
}
