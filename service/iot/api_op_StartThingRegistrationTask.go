// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a bulk thing provisioning task.
func (c *Client) StartThingRegistrationTask(ctx context.Context, params *StartThingRegistrationTaskInput, optFns ...func(*Options)) (*StartThingRegistrationTaskOutput, error) {
	if params == nil {
		params = &StartThingRegistrationTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartThingRegistrationTask", params, optFns, addOperationStartThingRegistrationTaskMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StartThingRegistrationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartThingRegistrationTaskInput struct {

	// The S3 bucket that contains the input file.
	//
	// This member is required.
	InputFileBucket *string

	// The name of input file within the S3 bucket. This file contains a newline
	// delimited JSON file. Each line contains the parameter values to provision one
	// device (thing).
	//
	// This member is required.
	InputFileKey *string

	// The IAM role ARN that grants permission the input file.
	//
	// This member is required.
	RoleArn *string

	// The provisioning template.
	//
	// This member is required.
	TemplateBody *string
}

type StartThingRegistrationTaskOutput struct {

	// The bulk thing provisioning task ID.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartThingRegistrationTaskMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartThingRegistrationTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartThingRegistrationTask{}, middleware.After)
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
	addOpStartThingRegistrationTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartThingRegistrationTask(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartThingRegistrationTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "StartThingRegistrationTask",
	}
}
