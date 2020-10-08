// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Instructs the specified agents or connectors to start collecting data.
func (c *Client) StartDataCollectionByAgentIds(ctx context.Context, params *StartDataCollectionByAgentIdsInput, optFns ...func(*Options)) (*StartDataCollectionByAgentIdsOutput, error) {
	if params == nil {
		params = &StartDataCollectionByAgentIdsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDataCollectionByAgentIds", params, optFns, addOperationStartDataCollectionByAgentIdsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDataCollectionByAgentIdsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDataCollectionByAgentIdsInput struct {

	// The IDs of the agents or connectors from which to start collecting data. If you
	// send a request to an agent/connector ID that you do not have permission to
	// contact, according to your AWS account, the service does not throw an exception.
	// Instead, it returns the error in the Description field. If you send a request to
	// multiple agents/connectors and you do not have permission to contact some of
	// those agents/connectors, the system does not throw an exception. Instead, the
	// system shows Failed in the Description field.
	//
	// This member is required.
	AgentIds []*string
}

type StartDataCollectionByAgentIdsOutput struct {

	// Information about agents or the connector that were instructed to start
	// collecting data. Information includes the agent/connector ID, a description of
	// the operation performed, and whether the agent/connector configuration was
	// updated.
	AgentsConfigurationStatus []*types.AgentConfigurationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartDataCollectionByAgentIdsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDataCollectionByAgentIds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDataCollectionByAgentIds{}, middleware.After)
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
	addOpStartDataCollectionByAgentIdsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartDataCollectionByAgentIds(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartDataCollectionByAgentIds(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "StartDataCollectionByAgentIds",
	}
}
