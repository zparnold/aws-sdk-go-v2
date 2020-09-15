// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies which task set in a service is the primary task set. Any parameters
// that are updated on the primary task set in a service will transition to the
// service. This is used when a service uses the EXTERNAL deployment controller
// type. For more information, see Amazon ECS Deployment Types
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) UpdateServicePrimaryTaskSet(ctx context.Context, params *UpdateServicePrimaryTaskSetInput, optFns ...func(*Options)) (*UpdateServicePrimaryTaskSetOutput, error) {
	stack := middleware.NewStack("UpdateServicePrimaryTaskSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateServicePrimaryTaskSetMiddlewares(stack)
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
	addOpUpdateServicePrimaryTaskSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServicePrimaryTaskSet(options.Region), middleware.Before)

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
			OperationName: "UpdateServicePrimaryTaskSet",
			Err:           err,
		}
	}
	out := result.(*UpdateServicePrimaryTaskSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServicePrimaryTaskSetInput struct {
	// The short name or full Amazon Resource Name (ARN) of the task set to set as the
	// primary task set in the deployment.
	PrimaryTaskSet *string
	// The short name or full Amazon Resource Name (ARN) of the service that the task
	// set exists in.
	Service *string
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service that the task set exists in.
	Cluster *string
}

type UpdateServicePrimaryTaskSetOutput struct {
	// Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an
	// EXTERNAL deployment. An Amazon ECS task set includes details such as the desired
	// number of tasks, how many tasks are running, and whether the task set serves
	// production traffic.
	TaskSet *types.TaskSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateServicePrimaryTaskSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServicePrimaryTaskSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServicePrimaryTaskSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateServicePrimaryTaskSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "UpdateServicePrimaryTaskSet",
	}
}
