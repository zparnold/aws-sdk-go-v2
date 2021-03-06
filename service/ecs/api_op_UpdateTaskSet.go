// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a task set. This is used when a service uses the EXTERNAL deployment
// controller type. For more information, see Amazon ECS Deployment Types
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) UpdateTaskSet(ctx context.Context, params *UpdateTaskSetInput, optFns ...func(*Options)) (*UpdateTaskSetOutput, error) {
	stack := middleware.NewStack("UpdateTaskSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateTaskSetMiddlewares(stack)
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
	addOpUpdateTaskSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTaskSet(options.Region), middleware.Before)
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
			OperationName: "UpdateTaskSet",
			Err:           err,
		}
	}
	out := result.(*UpdateTaskSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTaskSetInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service that the task set exists in.
	//
	// This member is required.
	Cluster *string

	// A floating-point percentage of the desired number of tasks to place and keep
	// running in the task set.
	//
	// This member is required.
	Scale *types.Scale

	// The short name or full Amazon Resource Name (ARN) of the service that the task
	// set exists in.
	//
	// This member is required.
	Service *string

	// The short name or full Amazon Resource Name (ARN) of the task set to update.
	//
	// This member is required.
	TaskSet *string
}

type UpdateTaskSetOutput struct {

	// Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an
	// EXTERNAL deployment. An Amazon ECS task set includes details such as the desired
	// number of tasks, how many tasks are running, and whether the task set serves
	// production traffic.
	TaskSet *types.TaskSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateTaskSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTaskSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTaskSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTaskSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "UpdateTaskSet",
	}
}
