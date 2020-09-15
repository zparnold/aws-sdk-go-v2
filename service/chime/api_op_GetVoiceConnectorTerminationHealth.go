// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the last time a SIP OPTIONS ping was received from
// your SIP infrastructure for the specified Amazon Chime Voice Connector.
func (c *Client) GetVoiceConnectorTerminationHealth(ctx context.Context, params *GetVoiceConnectorTerminationHealthInput, optFns ...func(*Options)) (*GetVoiceConnectorTerminationHealthOutput, error) {
	stack := middleware.NewStack("GetVoiceConnectorTerminationHealth", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetVoiceConnectorTerminationHealthMiddlewares(stack)
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
	addOpGetVoiceConnectorTerminationHealthValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVoiceConnectorTerminationHealth(options.Region), middleware.Before)

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
			OperationName: "GetVoiceConnectorTerminationHealth",
			Err:           err,
		}
	}
	out := result.(*GetVoiceConnectorTerminationHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVoiceConnectorTerminationHealthInput struct {
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
}

type GetVoiceConnectorTerminationHealthOutput struct {
	// The termination health details.
	TerminationHealth *types.TerminationHealth

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetVoiceConnectorTerminationHealthMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetVoiceConnectorTerminationHealth{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVoiceConnectorTerminationHealth{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVoiceConnectorTerminationHealth(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetVoiceConnectorTerminationHealth",
	}
}
