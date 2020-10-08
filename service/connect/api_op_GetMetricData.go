// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets historical metric data from the specified Amazon Connect instance. For more
// information, see Historical Metrics Reports
// (https://docs.aws.amazon.com/connect/latest/adminguide/historical-metrics.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) GetMetricData(ctx context.Context, params *GetMetricDataInput, optFns ...func(*Options)) (*GetMetricDataOutput, error) {
	if params == nil {
		params = &GetMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricData", params, optFns, addOperationGetMetricDataMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricDataInput struct {

	// The timestamp, in UNIX Epoch time format, at which to end the reporting interval
	// for the retrieval of historical metrics data. The time must be specified using
	// an interval of 5 minutes, such as 11:00, 11:05, 11:10, and must be later than
	// the start time timestamp. The time range between the start and end time must be
	// less than 24 hours.
	//
	// This member is required.
	EndTime *time.Time

	// The queues, up to 100, or channels, to use to filter the metrics returned.
	// Metric data is retrieved only for the resources associated with the queues or
	// channels included in the filter. You can include both queue IDs and queue ARNs
	// in the same request. The only supported channel is VOICE.
	//
	// This member is required.
	Filters *types.Filters

	// The metrics to retrieve. Specify the name, unit, and statistic for each metric.
	// The following historical metrics are available. For a description of each
	// metric, see Historical Metrics Definitions
	// (https://docs.aws.amazon.com/connect/latest/adminguide/historical-metrics-definitions.html)
	// in the Amazon Connect Administrator Guide. ABANDON_TIME Unit: SECONDS Statistic:
	// AVG AFTER_CONTACT_WORK_TIME Unit: SECONDS Statistic: AVG API_CONTACTS_HANDLED
	// Unit: COUNT Statistic: SUM CALLBACK_CONTACTS_HANDLED Unit: COUNT Statistic: SUM
	// CONTACTS_ABANDONED Unit: COUNT Statistic: SUM CONTACTS_AGENT_HUNG_UP_FIRST Unit:
	// COUNT Statistic: SUM CONTACTS_CONSULTED Unit: COUNT Statistic: SUM
	// CONTACTS_HANDLED Unit: COUNT Statistic: SUM CONTACTS_HANDLED_INCOMING Unit:
	// COUNT Statistic: SUM CONTACTS_HANDLED_OUTBOUND Unit: COUNT Statistic: SUM
	// CONTACTS_HOLD_ABANDONS Unit: COUNT Statistic: SUM CONTACTS_MISSED Unit: COUNT
	// Statistic: SUM CONTACTS_QUEUED Unit: COUNT Statistic: SUM
	// CONTACTS_TRANSFERRED_IN Unit: COUNT Statistic: SUM
	// CONTACTS_TRANSFERRED_IN_FROM_QUEUE Unit: COUNT Statistic: SUM
	// CONTACTS_TRANSFERRED_OUT Unit: COUNT Statistic: SUM
	// CONTACTS_TRANSFERRED_OUT_FROM_QUEUE Unit: COUNT Statistic: SUM HANDLE_TIME Unit:
	// SECONDS Statistic: AVG HOLD_TIME Unit: SECONDS Statistic: AVG
	// INTERACTION_AND_HOLD_TIME Unit: SECONDS Statistic: AVG INTERACTION_TIME Unit:
	// SECONDS Statistic: AVG OCCUPANCY Unit: PERCENT Statistic: AVG QUEUE_ANSWER_TIME
	// Unit: SECONDS Statistic: AVG QUEUED_TIME Unit: SECONDS Statistic: MAX
	// SERVICE_LEVEL Unit: PERCENT Statistic: AVG Threshold: Only "Less than"
	// comparisons are supported, with the following service level thresholds: 15, 20,
	// 25, 30, 45, 60, 90, 120, 180, 240, 300, 600
	//
	// This member is required.
	HistoricalMetrics []*types.HistoricalMetric

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The timestamp, in UNIX Epoch time format, at which to start the reporting
	// interval for the retrieval of historical metrics data. The time must be
	// specified using a multiple of 5 minutes, such as 10:05, 10:10, 10:15. The start
	// time cannot be earlier than 24 hours before the time of the request. Historical
	// metrics are available only for 24 hours.
	//
	// This member is required.
	StartTime *time.Time

	// The grouping applied to the metrics returned. For example, when results are
	// grouped by queue, the metrics returned are grouped by queue. The values returned
	// apply to the metrics for each queue rather than aggregated for all queues. The
	// only supported grouping is QUEUE. If no grouping is specified, a summary of
	// metrics for all queues is returned.
	Groupings []types.Grouping

	// The maximimum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
}

type GetMetricDataOutput struct {

	// Information about the historical metrics. If no grouping is specified, a summary
	// of metric data is returned.
	MetricResults []*types.HistoricalMetricResult

	// If there are additional results, this is the token for the next set of results.
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the token must use the same request parameters as the request
	// that generated the token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMetricDataMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMetricData{}, middleware.After)
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
	addOpGetMetricDataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricData(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMetricData(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "GetMetricData",
	}
}
