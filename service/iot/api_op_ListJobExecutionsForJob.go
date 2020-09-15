// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the job executions for a job.
func (c *Client) ListJobExecutionsForJob(ctx context.Context, params *ListJobExecutionsForJobInput, optFns ...func(*Options)) (*ListJobExecutionsForJobOutput, error) {
	stack := middleware.NewStack("ListJobExecutionsForJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListJobExecutionsForJobMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListJobExecutionsForJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobExecutionsForJob(options.Region), middleware.Before)

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
			OperationName: "ListJobExecutionsForJob",
			Err:           err,
		}
	}
	out := result.(*ListJobExecutionsForJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobExecutionsForJobInput struct {
	// The maximum number of results to be returned per request.
	MaxResults *int32
	// The unique identifier you assigned to this job when it was created.
	JobId *string
	// The status of the job.
	Status types.JobExecutionStatus
	// The token to retrieve the next set of results.
	NextToken *string
}

type ListJobExecutionsForJobOutput struct {
	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string
	// A list of job execution summaries.
	ExecutionSummaries []*types.JobExecutionSummaryForJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListJobExecutionsForJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListJobExecutionsForJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobExecutionsForJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opListJobExecutionsForJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListJobExecutionsForJob",
	}
}
