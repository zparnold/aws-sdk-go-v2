// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a vocabulary from Amazon Transcribe.
func (c *Client) DeleteVocabulary(ctx context.Context, params *DeleteVocabularyInput, optFns ...func(*Options)) (*DeleteVocabularyOutput, error) {
	if params == nil {
		params = &DeleteVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVocabulary", params, optFns, addOperationDeleteVocabularyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVocabularyInput struct {

	// The name of the vocabulary to delete.
	//
	// This member is required.
	VocabularyName *string
}

type DeleteVocabularyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteVocabularyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteVocabulary{}, middleware.After)
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
	addOpDeleteVocabularyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVocabulary(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteVocabulary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "DeleteVocabulary",
	}
}
