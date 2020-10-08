// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specifies the Key Management Service (KMS) customer master key (CMK) to be used
// to encrypt content in Amazon Fraud Detector.
func (c *Client) PutKMSEncryptionKey(ctx context.Context, params *PutKMSEncryptionKeyInput, optFns ...func(*Options)) (*PutKMSEncryptionKeyOutput, error) {
	if params == nil {
		params = &PutKMSEncryptionKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutKMSEncryptionKey", params, optFns, addOperationPutKMSEncryptionKeyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutKMSEncryptionKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutKMSEncryptionKeyInput struct {

	// The KMS encryption key ARN.
	//
	// This member is required.
	KmsEncryptionKeyArn *string
}

type PutKMSEncryptionKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutKMSEncryptionKeyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutKMSEncryptionKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutKMSEncryptionKey{}, middleware.After)
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
	addOpPutKMSEncryptionKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutKMSEncryptionKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutKMSEncryptionKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "PutKMSEncryptionKey",
	}
}
