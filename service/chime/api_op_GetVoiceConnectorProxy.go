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

// Gets the proxy configuration details for the specified Amazon Chime Voice
// Connector.
func (c *Client) GetVoiceConnectorProxy(ctx context.Context, params *GetVoiceConnectorProxyInput, optFns ...func(*Options)) (*GetVoiceConnectorProxyOutput, error) {
	stack := middleware.NewStack("GetVoiceConnectorProxy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetVoiceConnectorProxyMiddlewares(stack)
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
	addOpGetVoiceConnectorProxyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVoiceConnectorProxy(options.Region), middleware.Before)

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
			OperationName: "GetVoiceConnectorProxy",
			Err:           err,
		}
	}
	out := result.(*GetVoiceConnectorProxyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVoiceConnectorProxyInput struct {
	// The Amazon Chime voice connector ID.
	VoiceConnectorId *string
}

type GetVoiceConnectorProxyOutput struct {
	// The proxy configuration details.
	Proxy *types.Proxy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetVoiceConnectorProxyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetVoiceConnectorProxy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVoiceConnectorProxy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVoiceConnectorProxy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetVoiceConnectorProxy",
	}
}
