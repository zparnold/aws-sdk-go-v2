// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a metric filter and associates it with the specified log
// group. Metric filters allow you to configure rules to extract metric data from
// log events ingested through PutLogEvents
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html).
// The maximum number of metric filters that can be associated with a log group is
// 100.
func (c *Client) PutMetricFilter(ctx context.Context, params *PutMetricFilterInput, optFns ...func(*Options)) (*PutMetricFilterOutput, error) {
	stack := middleware.NewStack("PutMetricFilter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutMetricFilterMiddlewares(stack)
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
	addOpPutMetricFilterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricFilter(options.Region), middleware.Before)
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
			OperationName: "PutMetricFilter",
			Err:           err,
		}
	}
	out := result.(*PutMetricFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricFilterInput struct {

	// A name for the metric filter.
	//
	// This member is required.
	FilterName *string

	// A filter pattern for extracting metric data out of ingested log events.
	//
	// This member is required.
	FilterPattern *string

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// A collection of information that defines how metric data gets emitted.
	//
	// This member is required.
	MetricTransformations []*types.MetricTransformation
}

type PutMetricFilterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutMetricFilterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutMetricFilter{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutMetricFilter{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutMetricFilter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutMetricFilter",
	}
}
