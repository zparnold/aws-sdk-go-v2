// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an agent. To specify which agent to delete, use the Amazon Resource Name
// (ARN) of the agent in your request. The operation disassociates the agent from
// your AWS account. However, it doesn't delete the agent virtual machine (VM) from
// your on-premises environment.
func (c *Client) DeleteAgent(ctx context.Context, params *DeleteAgentInput, optFns ...func(*Options)) (*DeleteAgentOutput, error) {
	if params == nil {
		params = &DeleteAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAgent", params, optFns, addOperationDeleteAgentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DeleteAgentRequest
type DeleteAgentInput struct {

	// The Amazon Resource Name (ARN) of the agent to delete. Use the ListAgents
	// operation to return a list of agents for your account and AWS Region.
	//
	// This member is required.
	AgentArn *string
}

type DeleteAgentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAgentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAgent{}, middleware.After)
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
	addOpDeleteAgentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAgent(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAgent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DeleteAgent",
	}
}
