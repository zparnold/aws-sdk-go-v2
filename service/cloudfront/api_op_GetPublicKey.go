// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the public key information.
func (c *Client) GetPublicKey(ctx context.Context, params *GetPublicKeyInput, optFns ...func(*Options)) (*GetPublicKeyOutput, error) {
	stack := middleware.NewStack("GetPublicKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetPublicKeyMiddlewares(stack)
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
	addOpGetPublicKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPublicKey(options.Region), middleware.Before)
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
			OperationName: "GetPublicKey",
			Err:           err,
		}
	}
	out := result.(*GetPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPublicKeyInput struct {

	// Request the ID for the public key.
	//
	// This member is required.
	Id *string
}

type GetPublicKeyOutput struct {

	// The current version of the public key. For example: E2QWRUHAPOMQZL.
	ETag *string

	// Return the public key.
	PublicKey *types.PublicKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetPublicKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetPublicKey{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetPublicKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPublicKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetPublicKey",
	}
}
