// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about an Amazon SageMaker job.
func (c *Client) DescribeAutoMLJob(ctx context.Context, params *DescribeAutoMLJobInput, optFns ...func(*Options)) (*DescribeAutoMLJobOutput, error) {
	stack := middleware.NewStack("DescribeAutoMLJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAutoMLJobMiddlewares(stack)
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
	addOpDescribeAutoMLJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoMLJob(options.Region), middleware.Before)

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
			OperationName: "DescribeAutoMLJob",
			Err:           err,
		}
	}
	out := result.(*DescribeAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoMLJobInput struct {
	// Request information about a job using that job's unique name.
	AutoMLJobName *string
}

type DescribeAutoMLJobOutput struct {
	// Returns the job's creation time.
	CreationTime *time.Time
	// Returns the job's last modified time.
	LastModifiedTime *time.Time
	// Returns the job's BestCandidate.
	BestCandidate *types.AutoMLCandidate
	// Returns the job's end time.
	EndTime *time.Time
	// Returns the name of a job.
	AutoMLJobName *string
	// Returns the job's ARN.
	AutoMLJobArn *string
	// Returns the job's output from GenerateCandidateDefinitionsOnly.
	GenerateCandidateDefinitionsOnly *bool
	// Returns the job's FailureReason.
	FailureReason *string
	// Returns the job's objective.
	AutoMLJobObjective *types.AutoMLJobObjective
	// Returns information on the job's artifacts found in AutoMLJobArtifacts.
	AutoMLJobArtifacts *types.AutoMLJobArtifacts
	// Returns the job's AutoMLJobSecondaryStatus.
	AutoMLJobSecondaryStatus types.AutoMLJobSecondaryStatus
	// This contains ProblemType, AutoMLJobObjective and CompletionCriteria. They're
	// auto-inferred values, if not provided by you. If you do provide them, then
	// they'll be the same as provided.
	ResolvedAttributes *types.ResolvedAttributes
	// Returns the job's output data config.
	OutputDataConfig *types.AutoMLOutputDataConfig
	// Returns the job's problem type.
	ProblemType types.ProblemType
	// Returns the job's AutoMLJobStatus.
	AutoMLJobStatus types.AutoMLJobStatus
	// Returns the job's config.
	AutoMLJobConfig *types.AutoMLJobConfig
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that has read permission to the input data location and write permission to
	// the output data location in Amazon S3.
	RoleArn *string
	// Returns the job's input data config.
	InputDataConfig []*types.AutoMLChannel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAutoMLJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAutoMLJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAutoMLJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAutoMLJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeAutoMLJob",
	}
}
