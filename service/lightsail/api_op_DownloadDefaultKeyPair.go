// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Downloads the default SSH key pair from the user's account.
func (c *Client) DownloadDefaultKeyPair(ctx context.Context, params *DownloadDefaultKeyPairInput, optFns ...func(*Options)) (*DownloadDefaultKeyPairOutput, error) {
	stack := middleware.NewStack("DownloadDefaultKeyPair", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDownloadDefaultKeyPairMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDownloadDefaultKeyPair(options.Region), middleware.Before)
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
			OperationName: "DownloadDefaultKeyPair",
			Err:           err,
		}
	}
	out := result.(*DownloadDefaultKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DownloadDefaultKeyPairInput struct {
}

type DownloadDefaultKeyPairOutput struct {

	// A base64-encoded RSA private key.
	PrivateKeyBase64 *string

	// A base64-encoded public key of the ssh-rsa type.
	PublicKeyBase64 *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDownloadDefaultKeyPairMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDownloadDefaultKeyPair{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDownloadDefaultKeyPair{}, middleware.After)
}

func newServiceMetadataMiddleware_opDownloadDefaultKeyPair(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DownloadDefaultKeyPair",
	}
}
