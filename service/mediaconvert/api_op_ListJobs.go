// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve a JSON array of up to twenty of your most recently created jobs. This
// array includes in-process, completed, and errored jobs. This will return the
// jobs themselves, not just a list of the jobs. To retrieve the twenty next most
// recent jobs, use the nextToken string returned with the array.
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if params == nil {
		params = &ListJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobs", params, optFns, addOperationListJobsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobsInput struct {

	// Optional. Number of jobs, up to twenty, that will be returned at one time.
	MaxResults *int32

	// Optional. Use this string, provided with the response to a previous request, to
	// request the next batch of jobs.
	NextToken *string

	// Optional. When you request lists of resources, you can specify whether they are
	// sorted in ASCENDING or DESCENDING order. Default varies by resource.
	Order types.Order

	// Optional. Provide a queue name to get back only jobs from that queue.
	Queue *string

	// Optional. A job's status can be SUBMITTED, PROGRESSING, COMPLETE, CANCELED, or
	// ERROR.
	Status types.JobStatus
}

type ListJobsOutput struct {

	// List of jobs
	Jobs []*types.Job

	// Use this string to request the next batch of jobs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "ListJobs",
	}
}
