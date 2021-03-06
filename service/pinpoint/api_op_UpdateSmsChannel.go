// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the SMS channel for an application or updates the status and settings of
// the SMS channel for an application.
func (c *Client) UpdateSmsChannel(ctx context.Context, params *UpdateSmsChannelInput, optFns ...func(*Options)) (*UpdateSmsChannelOutput, error) {
	stack := middleware.NewStack("UpdateSmsChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateSmsChannelMiddlewares(stack)
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
	addOpUpdateSmsChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSmsChannel(options.Region), middleware.Before)
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
			OperationName: "UpdateSmsChannel",
			Err:           err,
		}
	}
	out := result.(*UpdateSmsChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSmsChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// Specifies the status and settings of the SMS channel for an application.
	//
	// This member is required.
	SMSChannelRequest *types.SMSChannelRequest
}

type UpdateSmsChannelOutput struct {

	// Provides information about the status and settings of the SMS channel for an
	// application.
	//
	// This member is required.
	SMSChannelResponse *types.SMSChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateSmsChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSmsChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSmsChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSmsChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateSmsChannel",
	}
}
