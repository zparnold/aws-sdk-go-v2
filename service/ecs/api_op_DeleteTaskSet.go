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

// Deletes a specified task set within a service. This is used when a service uses
// the EXTERNAL deployment controller type. For more information, see Amazon ECS
// Deployment Types
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) DeleteTaskSet(ctx context.Context, params *DeleteTaskSetInput, optFns ...func(*Options)) (*DeleteTaskSetOutput, error) {
	stack := middleware.NewStack("DeleteTaskSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteTaskSetMiddlewares(stack)
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
	addOpDeleteTaskSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTaskSet(options.Region), middleware.Before)
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
			OperationName: "DeleteTaskSet",
			Err:           err,
		}
	}
	out := result.(*DeleteTaskSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTaskSetInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service that the task set exists in to delete.
	//
	// This member is required.
	Cluster *string

	// The short name or full Amazon Resource Name (ARN) of the service that hosts the
	// task set to delete.
	//
	// This member is required.
	Service *string

	// The task set ID or full Amazon Resource Name (ARN) of the task set to delete.
	//
	// This member is required.
	TaskSet *string

	// If true, this allows you to delete a task set even if it hasn't been scaled down
	// to zero.
	Force *bool
}

type DeleteTaskSetOutput struct {

	// Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an
	// EXTERNAL deployment. An Amazon ECS task set includes details such as the desired
	// number of tasks, how many tasks are running, and whether the task set serves
	// production traffic.
	TaskSet *types.TaskSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteTaskSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTaskSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTaskSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTaskSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeleteTaskSet",
	}
}
