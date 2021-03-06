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

// Updates the Amazon ECS container agent on a specified container instance.
// Updating the Amazon ECS container agent does not interrupt running tasks or
// services on the container instance. The process for updating the agent differs
// depending on whether your container instance was launched with the Amazon
// ECS-optimized AMI or another operating system. UpdateContainerAgent requires the
// Amazon ECS-optimized AMI or Amazon Linux with the ecs-init service installed and
// running. For help updating the Amazon ECS container agent on other operating
// systems, see Manually Updating the Amazon ECS Container Agent
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html#manually_update_agent)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) UpdateContainerAgent(ctx context.Context, params *UpdateContainerAgentInput, optFns ...func(*Options)) (*UpdateContainerAgentOutput, error) {
	stack := middleware.NewStack("UpdateContainerAgent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateContainerAgentMiddlewares(stack)
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
	addOpUpdateContainerAgentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContainerAgent(options.Region), middleware.Before)
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
			OperationName: "UpdateContainerAgent",
			Err:           err,
		}
	}
	out := result.(*UpdateContainerAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContainerAgentInput struct {

	// The container instance ID or full ARN entries for the container instance on
	// which you would like to update the Amazon ECS container agent.
	//
	// This member is required.
	ContainerInstance *string

	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// container instance is running on. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string
}

type UpdateContainerAgentOutput struct {

	// The container instance for which the container agent was updated.
	ContainerInstance *types.ContainerInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateContainerAgentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateContainerAgent{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateContainerAgent{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateContainerAgent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "UpdateContainerAgent",
	}
}
