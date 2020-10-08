// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the APNs channel for an application and deletes any existing settings
// for the channel.
func (c *Client) DeleteApnsChannel(ctx context.Context, params *DeleteApnsChannelInput, optFns ...func(*Options)) (*DeleteApnsChannelOutput, error) {
	if params == nil {
		params = &DeleteApnsChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApnsChannel", params, optFns, addOperationDeleteApnsChannelMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApnsChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApnsChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string
}

type DeleteApnsChannelOutput struct {

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) channel for an application.
	//
	// This member is required.
	APNSChannelResponse *types.APNSChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteApnsChannelMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteApnsChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteApnsChannel{}, middleware.After)
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
	addOpDeleteApnsChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApnsChannel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteApnsChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "DeleteApnsChannel",
	}
}
