// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the status for the specified job. Use this operation to confirm that you
// want to run a job or to cancel an existing job. For more information, see S3
// Batch Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
// Amazon Simple Storage Service Developer Guide. Related actions include:
//
//     *
// CreateJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
//
//
// * ListJobs
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
//
//
// * DescribeJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
//
//
// * UpdateJobStatus
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
func (c *Client) UpdateJobStatus(ctx context.Context, params *UpdateJobStatusInput, optFns ...func(*Options)) (*UpdateJobStatusOutput, error) {
	if params == nil {
		params = &UpdateJobStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJobStatus", params, optFns, addOperationUpdateJobStatusMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobStatusInput struct {

	//
	//
	// This member is required.
	AccountId *string

	// The ID of the job whose status you want to update.
	//
	// This member is required.
	JobId *string

	// The status that you want to move the specified job to.
	//
	// This member is required.
	RequestedJobStatus types.RequestedJobStatus

	// A description of the reason why you want to change the specified job's status.
	// This field can be any string up to the maximum length.
	StatusUpdateReason *string
}

type UpdateJobStatusOutput struct {

	// The ID for the job whose status was updated.
	JobId *string

	// The current status for the specified job.
	Status types.JobStatus

	// The reason that the specified job's status was updated.
	StatusUpdateReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateJobStatusMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateJobStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateJobStatus{}, middleware.After)
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
	addOpUpdateJobStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobStatus(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateJobStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "UpdateJobStatus",
	}
}
