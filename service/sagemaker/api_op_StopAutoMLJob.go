// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// A method for forcing the termination of a running job.
func (c *Client) StopAutoMLJob(ctx context.Context, params *StopAutoMLJobInput, optFns ...func(*Options)) (*StopAutoMLJobOutput, error) {
	stack := middleware.NewStack("StopAutoMLJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopAutoMLJobMiddlewares(stack)
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
	addOpStopAutoMLJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopAutoMLJob(options.Region), middleware.Before)

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
			OperationName: "StopAutoMLJob",
			Err:           err,
		}
	}
	out := result.(*StopAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopAutoMLJobInput struct {
	// The name of the object you are requesting.
	AutoMLJobName *string
}

type StopAutoMLJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopAutoMLJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopAutoMLJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopAutoMLJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopAutoMLJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StopAutoMLJob",
	}
}
