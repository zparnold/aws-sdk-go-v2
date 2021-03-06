// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes recommendation export jobs created in the last seven days.  <p>Use the
// <code>ExportAutoScalingGroupRecommendations</code> or
// <code>ExportEC2InstanceRecommendations</code> actions to request an export of
// your recommendations. Then use the <code>DescribeRecommendationExportJobs</code>
// action to view your export jobs.</p>
func (c *Client) DescribeRecommendationExportJobs(ctx context.Context, params *DescribeRecommendationExportJobsInput, optFns ...func(*Options)) (*DescribeRecommendationExportJobsOutput, error) {
	stack := middleware.NewStack("DescribeRecommendationExportJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeRecommendationExportJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecommendationExportJobs(options.Region), middleware.Before)
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
			OperationName: "DescribeRecommendationExportJobs",
			Err:           err,
		}
	}
	out := result.(*DescribeRecommendationExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecommendationExportJobsInput struct {

	// An array of objects that describe a filter to return a more specific list of
	// export jobs.
	Filters []*types.JobFilter

	// The identification numbers of the export jobs to return.  <p>An export job ID is
	// returned when you create an export using the
	// <code>ExportAutoScalingGroupRecommendations</code> or
	// <code>ExportEC2InstanceRecommendations</code> actions.</p> <p>All export jobs
	// created in the last seven days are returned if this parameter is omitted.</p>
	JobIds []*string

	// The maximum number of export jobs to return with a single request. To retrieve
	// the remaining results, make another request with the returned NextToken value.
	MaxResults *int32

	// The token to advance to the next page of export jobs.
	NextToken *string
}

type DescribeRecommendationExportJobsOutput struct {

	// The token to use to advance to the next page of export jobs. This value is null
	// when there are no more pages of export jobs to return.
	NextToken *string

	// An array of objects that describe recommendation export jobs.
	RecommendationExportJobs []*types.RecommendationExportJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeRecommendationExportJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRecommendationExportJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRecommendationExportJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRecommendationExportJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "DescribeRecommendationExportJobs",
	}
}
