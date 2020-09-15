// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the class of the channel.
func (c *Client) UpdateChannelClass(ctx context.Context, params *UpdateChannelClassInput, optFns ...func(*Options)) (*UpdateChannelClassOutput, error) {
	stack := middleware.NewStack("UpdateChannelClass", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateChannelClassMiddlewares(stack)
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
	addOpUpdateChannelClassValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChannelClass(options.Region), middleware.Before)

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
			OperationName: "UpdateChannelClass",
			Err:           err,
		}
	}
	out := result.(*UpdateChannelClassOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Channel class that the channel should be updated to.
type UpdateChannelClassInput struct {
	// The channel class that you wish to update this channel to use.
	ChannelClass types.ChannelClass
	// Channel Id of the channel whose class should be updated.
	ChannelId *string
	// A list of output destinations for this channel.
	Destinations []*types.OutputDestination
}

// Placeholder documentation for UpdateChannelClassResponse
type UpdateChannelClassOutput struct {
	// Placeholder documentation for Channel
	Channel *types.Channel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateChannelClassMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateChannelClass{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateChannelClass{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateChannelClass(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateChannelClass",
	}
}
