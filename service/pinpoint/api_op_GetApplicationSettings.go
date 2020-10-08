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

// Retrieves information about the settings for an application.
func (c *Client) GetApplicationSettings(ctx context.Context, params *GetApplicationSettingsInput, optFns ...func(*Options)) (*GetApplicationSettingsOutput, error) {
	if params == nil {
		params = &GetApplicationSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplicationSettings", params, optFns, addOperationGetApplicationSettingsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationSettingsInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string
}

type GetApplicationSettingsOutput struct {

	// Provides information about an application, including the default settings for an
	// application.
	//
	// This member is required.
	ApplicationSettingsResource *types.ApplicationSettingsResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetApplicationSettingsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApplicationSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplicationSettings{}, middleware.After)
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
	addOpGetApplicationSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplicationSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetApplicationSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetApplicationSettings",
	}
}
