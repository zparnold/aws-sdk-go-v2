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

// Puts the specified proxy configuration to the specified Amazon Chime Voice
// Connector.
func (c *Client) PutVoiceConnectorProxy(ctx context.Context, params *PutVoiceConnectorProxyInput, optFns ...func(*Options)) (*PutVoiceConnectorProxyOutput, error) {
	if params == nil {
		params = &PutVoiceConnectorProxyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutVoiceConnectorProxy", params, optFns, addOperationPutVoiceConnectorProxyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutVoiceConnectorProxyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorProxyInput struct {

	// The default number of minutes allowed for proxy sessions.
	//
	// This member is required.
	DefaultSessionExpiryMinutes *int32

	// The countries for proxy phone numbers to be selected from.
	//
	// This member is required.
	PhoneNumberPoolCountries []*string

	// The Amazon Chime voice connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// When true, stops proxy sessions from being created on the specified Amazon Chime
	// Voice Connector.
	Disabled *bool

	// The phone number to route calls to after a proxy session expires.
	FallBackPhoneNumber *string
}

type PutVoiceConnectorProxyOutput struct {

	// The proxy configuration details.
	Proxy *types.Proxy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutVoiceConnectorProxyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutVoiceConnectorProxy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutVoiceConnectorProxy{}, middleware.After)
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
	addOpPutVoiceConnectorProxyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutVoiceConnectorProxy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutVoiceConnectorProxy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorProxy",
	}
}
