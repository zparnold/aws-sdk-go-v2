// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a previously created document classifier Only those classifiers that are
// in terminated states (IN_ERROR, TRAINED) will be deleted. If an active inference
// job is using the model, a ResourceInUseException will be returned. This is an
// asynchronous action that puts the classifier into a DELETING state, and it is
// then removed by a background job. Once removed, the classifier disappears from
// your account and is no longer available for use.
func (c *Client) DeleteDocumentClassifier(ctx context.Context, params *DeleteDocumentClassifierInput, optFns ...func(*Options)) (*DeleteDocumentClassifierOutput, error) {
	if params == nil {
		params = &DeleteDocumentClassifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDocumentClassifier", params, optFns, addOperationDeleteDocumentClassifierMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDocumentClassifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDocumentClassifierInput struct {

	// The Amazon Resource Name (ARN) that identifies the document classifier.
	//
	// This member is required.
	DocumentClassifierArn *string
}

type DeleteDocumentClassifierOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDocumentClassifierMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDocumentClassifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDocumentClassifier{}, middleware.After)
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
	addOpDeleteDocumentClassifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDocumentClassifier(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDocumentClassifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DeleteDocumentClassifier",
	}
}
