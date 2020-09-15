// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a channel's configuration. This does not affect an ongoing stream of
// this channel. You must stop and restart the stream for the changes to take
// effect.
func (c *Client) UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) {
	stack := middleware.NewStack("UpdateChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateChannelMiddlewares(stack)
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
	addOpUpdateChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChannel(options.Region), middleware.Before)

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
			OperationName: "UpdateChannel",
			Err:           err,
		}
	}
	out := result.(*UpdateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateChannelInput struct {
	// Channel latency mode. Default: LOW.
	LatencyMode types.ChannelLatencyMode
	// ARN of the channel to be updated.
	Arn *string
	// Channel name.
	Name *string
	// Channel type, which determines the allowable resolution and bitrate. If you
	// exceed the allowable resolution or bitrate, the stream probably will disconnect
	// immediately. Valid values:
	//
	//     * STANDARD: Multiple qualities are generated
	// from the original input, to automatically give viewers the best experience for
	// their devices and network conditions. Vertical resolution can be up to 1080 and
	// bitrate can be up to 8.5 Mbps.
	//
	//     * BASIC: Amazon IVS delivers the original
	// input to viewers. The viewer’s video-quality choice is limited to the original
	// input. Vertical resolution can be up to 480 and bitrate can be up to 1.5
	// Mbps.
	//
	// Default: STANDARD.
	Type types.ChannelType
}

type UpdateChannelOutput struct {
	// Object specifying a channel.
	Channel *types.Channel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "UpdateChannel",
	}
}
