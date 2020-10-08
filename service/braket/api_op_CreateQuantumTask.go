// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a quantum task.
func (c *Client) CreateQuantumTask(ctx context.Context, params *CreateQuantumTaskInput, optFns ...func(*Options)) (*CreateQuantumTaskOutput, error) {
	if params == nil {
		params = &CreateQuantumTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateQuantumTask", params, optFns, addOperationCreateQuantumTaskMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateQuantumTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQuantumTaskInput struct {

	// The action associated with the task.
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Action *string

	// The client token associated with the request.
	//
	// This member is required.
	ClientToken *string

	// The ARN of the device to run the task on.
	//
	// This member is required.
	DeviceArn *string

	// The S3 bucket to store task result files in.
	//
	// This member is required.
	OutputS3Bucket *string

	// The key prefix for the location in the S3 bucket to store task results in.
	//
	// This member is required.
	OutputS3KeyPrefix *string

	// The number of shots to use for the task.
	//
	// This member is required.
	Shots *int64

	// The parameters for the device to run the task on.
	// This value conforms to the media type: application/json
	DeviceParameters *string
}

type CreateQuantumTaskOutput struct {

	// The ARN of the task created by the request.
	//
	// This member is required.
	QuantumTaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateQuantumTaskMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateQuantumTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateQuantumTask{}, middleware.After)
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
	addOpCreateQuantumTaskValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
