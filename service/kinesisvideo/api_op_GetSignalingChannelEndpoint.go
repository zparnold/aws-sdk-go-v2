// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides an endpoint for the specified signaling channel to send and receive
// messages. This API uses the SingleMasterChannelEndpointConfiguration input
// parameter, which consists of the Protocols and Role properties. Protocols is
// used to determine the communication mechanism. For example, if you specify WSS
// as the protocol, this API produces a secure websocket endpoint. If you specify
// HTTPS as the protocol, this API generates an HTTPS endpoint. Role determines the
// messaging permissions. A MASTER role results in this API generating an endpoint
// that a client can use to communicate with any of the viewers on the channel. A
// VIEWER role results in this API generating an endpoint that a client can use to
// communicate only with a MASTER.
func (c *Client) GetSignalingChannelEndpoint(ctx context.Context, params *GetSignalingChannelEndpointInput, optFns ...func(*Options)) (*GetSignalingChannelEndpointOutput, error) {
	stack := middleware.NewStack("GetSignalingChannelEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSignalingChannelEndpointMiddlewares(stack)
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
	addOpGetSignalingChannelEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSignalingChannelEndpoint(options.Region), middleware.Before)
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
			OperationName: "GetSignalingChannelEndpoint",
			Err:           err,
		}
	}
	out := result.(*GetSignalingChannelEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSignalingChannelEndpointInput struct {

	// The Amazon Resource Name (ARN) of the signalling channel for which you want to
	// get an endpoint.
	//
	// This member is required.
	ChannelARN *string

	// A structure containing the endpoint configuration for the SINGLE_MASTER channel
	// type.
	SingleMasterChannelEndpointConfiguration *types.SingleMasterChannelEndpointConfiguration
}

type GetSignalingChannelEndpointOutput struct {

	// A list of endpoints for the specified signaling channel.
	ResourceEndpointList []*types.ResourceEndpointListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSignalingChannelEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSignalingChannelEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSignalingChannelEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSignalingChannelEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "GetSignalingChannelEndpoint",
	}
}
