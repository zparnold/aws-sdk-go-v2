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

// Returns the optimization findings for an account. For example, it returns the
// number of Amazon EC2 instances in an account that are under-provisioned,
// over-provisioned, or optimized. It also returns the number of Auto Scaling
// groups in an account that are not optimized, or optimized.
func (c *Client) GetRecommendationSummaries(ctx context.Context, params *GetRecommendationSummariesInput, optFns ...func(*Options)) (*GetRecommendationSummariesOutput, error) {
	stack := middleware.NewStack("GetRecommendationSummaries", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpGetRecommendationSummariesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendationSummaries(options.Region), middleware.Before)
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
			OperationName: "GetRecommendationSummaries",
			Err:           err,
		}
	}
	out := result.(*GetRecommendationSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecommendationSummariesInput struct {

	// The IDs of the AWS accounts for which to return recommendation summaries. If
	// your account is the master account of an organization, use this parameter to
	// specify the member accounts for which you want to return recommendation
	// summaries. Only one account ID can be specified per request.
	AccountIds []*string

	// The maximum number of recommendation summaries to return with a single request.
	// To retrieve the remaining results, make another request with the returned
	// NextToken value.
	MaxResults *int32

	// The token to advance to the next page of recommendation summaries.
	NextToken *string
}

type GetRecommendationSummariesOutput struct {

	// The token to use to advance to the next page of recommendation summaries. This
	// value is null when there are no more pages of recommendation summaries to
	// return.
	NextToken *string

	// An array of objects that summarize a recommendation.
	RecommendationSummaries []*types.RecommendationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpGetRecommendationSummariesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpGetRecommendationSummaries{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRecommendationSummaries{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRecommendationSummaries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetRecommendationSummaries",
	}
}
