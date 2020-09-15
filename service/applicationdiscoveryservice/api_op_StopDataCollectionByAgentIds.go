// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Instructs the specified agents or connectors to stop collecting data.
func (c *Client) StopDataCollectionByAgentIds(ctx context.Context, params *StopDataCollectionByAgentIdsInput, optFns ...func(*Options)) (*StopDataCollectionByAgentIdsOutput, error) {
	stack := middleware.NewStack("StopDataCollectionByAgentIds", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopDataCollectionByAgentIdsMiddlewares(stack)
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
	addOpStopDataCollectionByAgentIdsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDataCollectionByAgentIds(options.Region), middleware.Before)

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
			OperationName: "StopDataCollectionByAgentIds",
			Err:           err,
		}
	}
	out := result.(*StopDataCollectionByAgentIdsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDataCollectionByAgentIdsInput struct {
	// The IDs of the agents or connectors from which to stop collecting data.
	AgentIds []*string
}

type StopDataCollectionByAgentIdsOutput struct {
	// Information about the agents or connector that were instructed to stop
	// collecting data. Information includes the agent/connector ID, a description of
	// the operation performed, and whether the agent/connector configuration was
	// updated.
	AgentsConfigurationStatus []*types.AgentConfigurationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopDataCollectionByAgentIdsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopDataCollectionByAgentIds{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDataCollectionByAgentIds{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopDataCollectionByAgentIds(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "StopDataCollectionByAgentIds",
	}
}
