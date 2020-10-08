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

// Gets a list of HyperParameterTuningJobSummary () objects that describe the
// hyperparameter tuning jobs launched in your account.
func (c *Client) ListHyperParameterTuningJobs(ctx context.Context, params *ListHyperParameterTuningJobsInput, optFns ...func(*Options)) (*ListHyperParameterTuningJobsOutput, error) {
	if params == nil {
		params = &ListHyperParameterTuningJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHyperParameterTuningJobs", params, optFns, addOperationListHyperParameterTuningJobsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHyperParameterTuningJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHyperParameterTuningJobsInput struct {

	// A filter that returns only tuning jobs that were created after the specified
	// time.
	CreationTimeAfter *time.Time

	// A filter that returns only tuning jobs that were created before the specified
	// time.
	CreationTimeBefore *time.Time

	// A filter that returns only tuning jobs that were modified after the specified
	// time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only tuning jobs that were modified before the specified
	// time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of tuning jobs to return. The default value is 10.
	MaxResults *int32

	// A string in the tuning job name. This filter returns only tuning jobs whose name
	// contains the specified string.
	NameContains *string

	// If the result of the previous ListHyperParameterTuningJobs request was
	// truncated, the response includes a NextToken. To retrieve the next set of tuning
	// jobs, use the token in the next request.
	NextToken *string

	// The field to sort results by. The default is Name.
	SortBy types.HyperParameterTuningJobSortByOptions

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that returns only tuning jobs with the specified status.
	StatusEquals types.HyperParameterTuningJobStatus
}

type ListHyperParameterTuningJobsOutput struct {

	// A list of HyperParameterTuningJobSummary () objects that describe the tuning
	// jobs that the ListHyperParameterTuningJobs request returned.
	//
	// This member is required.
	HyperParameterTuningJobSummaries []*types.HyperParameterTuningJobSummary

	// If the result of this ListHyperParameterTuningJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of tuning jobs, use the
	// token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListHyperParameterTuningJobsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHyperParameterTuningJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHyperParameterTuningJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListHyperParameterTuningJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListHyperParameterTuningJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListHyperParameterTuningJobs",
	}
}
