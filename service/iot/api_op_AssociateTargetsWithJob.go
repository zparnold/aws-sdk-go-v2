// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a group with a continuous job. The following criteria must be met:
//
//
// * The job must have been created with the targetSelection field set to
// "CONTINUOUS".
//
//     * The job status must currently be "IN_PROGRESS".
//
//     * The
// total number of targets associated with a job must not exceed 100.
func (c *Client) AssociateTargetsWithJob(ctx context.Context, params *AssociateTargetsWithJobInput, optFns ...func(*Options)) (*AssociateTargetsWithJobOutput, error) {
	stack := middleware.NewStack("AssociateTargetsWithJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAssociateTargetsWithJobMiddlewares(stack)
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
	addOpAssociateTargetsWithJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTargetsWithJob(options.Region), middleware.Before)
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
			OperationName: "AssociateTargetsWithJob",
			Err:           err,
		}
	}
	out := result.(*AssociateTargetsWithJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTargetsWithJobInput struct {

	// The unique identifier you assigned to this job when it was created.
	//
	// This member is required.
	JobId *string

	// A list of thing group ARNs that define the targets of the job.
	//
	// This member is required.
	Targets []*string

	// An optional comment string describing why the job was associated with the
	// targets.
	Comment *string
}

type AssociateTargetsWithJobOutput struct {

	// A short text description of the job.
	Description *string

	// An ARN identifying the job.
	JobArn *string

	// The unique identifier you assigned to this job when it was created.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAssociateTargetsWithJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAssociateTargetsWithJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateTargetsWithJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateTargetsWithJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "AssociateTargetsWithJob",
	}
}
