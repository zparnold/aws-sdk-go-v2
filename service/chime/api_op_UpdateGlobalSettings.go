// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates global settings for the administrator's AWS account, such as Amazon
// Chime Business Calling and Amazon Chime Voice Connector settings.
func (c *Client) UpdateGlobalSettings(ctx context.Context, params *UpdateGlobalSettingsInput, optFns ...func(*Options)) (*UpdateGlobalSettingsOutput, error) {
	if params == nil {
		params = &UpdateGlobalSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGlobalSettings", params, optFns, addOperationUpdateGlobalSettingsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGlobalSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGlobalSettingsInput struct {

	// The Amazon Chime Business Calling settings.
	//
	// This member is required.
	BusinessCalling *types.BusinessCallingSettings

	// The Amazon Chime Voice Connector settings.
	//
	// This member is required.
	VoiceConnector *types.VoiceConnectorSettings
}

type UpdateGlobalSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGlobalSettingsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGlobalSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGlobalSettings{}, middleware.After)
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
	addOpUpdateGlobalSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGlobalSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateGlobalSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateGlobalSettings",
	}
}
