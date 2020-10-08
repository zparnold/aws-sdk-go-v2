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

// Enables the APNs sandbox channel for an application or updates the status and
// settings of the APNs sandbox channel for an application.
func (c *Client) UpdateApnsSandboxChannel(ctx context.Context, params *UpdateApnsSandboxChannelInput, optFns ...func(*Options)) (*UpdateApnsSandboxChannelOutput, error) {
	if params == nil {
		params = &UpdateApnsSandboxChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApnsSandboxChannel", params, optFns, addOperationUpdateApnsSandboxChannelMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApnsSandboxChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApnsSandboxChannelInput struct {

	// Specifies the status and settings of the APNs (Apple Push Notification service)
	// sandbox channel for an application.
	//
	// This member is required.
	APNSSandboxChannelRequest *types.APNSSandboxChannelRequest

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string
}

type UpdateApnsSandboxChannelOutput struct {

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) sandbox channel for an application.
	//
	// This member is required.
	APNSSandboxChannelResponse *types.APNSSandboxChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateApnsSandboxChannelMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApnsSandboxChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApnsSandboxChannel{}, middleware.After)
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
	addOpUpdateApnsSandboxChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApnsSandboxChannel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateApnsSandboxChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateApnsSandboxChannel",
	}
}
