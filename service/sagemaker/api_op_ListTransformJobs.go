// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists transform jobs.
func (c *Client) ListTransformJobs(ctx context.Context, params *ListTransformJobsInput, optFns ...func(*Options)) (*ListTransformJobsOutput, error) {
	if params == nil {
		params = &ListTransformJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTransformJobs", params, optFns, addOperationListTransformJobsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTransformJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTransformJobsInput struct {

	// A filter that returns only transform jobs created after the specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only transform jobs created before the specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only transform jobs modified after the specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only transform jobs modified before the specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of transform jobs to return in the response. The default
	// value is 10.
	MaxResults *int32

	// A string in the transform job name. This filter returns only transform jobs
	// whose name contains the specified string.
	NameContains *string

	// If the result of the previous ListTransformJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of transform jobs, use
	// the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.SortBy

	// The sort order for results. The default is Descending.
	SortOrder types.SortOrder

	// A filter that retrieves only transform jobs with a specific status.
	StatusEquals types.TransformJobStatus
}

type ListTransformJobsOutput struct {

	// An array of TransformJobSummary objects.
	//
	// This member is required.
	TransformJobSummaries []*types.TransformJobSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of transform jobs, use it in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTransformJobsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTransformJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTransformJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTransformJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTransformJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTransformJobs",
	}
}
