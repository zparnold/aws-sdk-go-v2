// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the field-level encryption configuration information.
func (c *Client) GetFieldLevelEncryption(ctx context.Context, params *GetFieldLevelEncryptionInput, optFns ...func(*Options)) (*GetFieldLevelEncryptionOutput, error) {
	stack := middleware.NewStack("GetFieldLevelEncryption", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetFieldLevelEncryptionMiddlewares(stack)
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
	addOpGetFieldLevelEncryptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFieldLevelEncryption(options.Region), middleware.Before)

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
			OperationName: "GetFieldLevelEncryption",
			Err:           err,
		}
	}
	out := result.(*GetFieldLevelEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFieldLevelEncryptionInput struct {
	// Request the ID for the field-level encryption configuration information.
	Id *string
}

type GetFieldLevelEncryptionOutput struct {
	// Return the field-level encryption configuration information.
	FieldLevelEncryption *types.FieldLevelEncryption
	// The current version of the field level encryption configuration. For example:
	// E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetFieldLevelEncryptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetFieldLevelEncryption{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetFieldLevelEncryption{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFieldLevelEncryption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetFieldLevelEncryption",
	}
}