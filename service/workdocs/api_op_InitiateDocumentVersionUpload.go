// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new document object and version object. The client specifies the
// parent folder ID and name of the document to upload. The ID is optionally
// specified when creating a new version of an existing document. This is the first
// step to upload a document. Next, upload the document to the URL returned from
// the call, and then call UpdateDocumentVersion (). To cancel the document upload,
// call AbortDocumentVersionUpload ().
func (c *Client) InitiateDocumentVersionUpload(ctx context.Context, params *InitiateDocumentVersionUploadInput, optFns ...func(*Options)) (*InitiateDocumentVersionUploadOutput, error) {
	if params == nil {
		params = &InitiateDocumentVersionUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InitiateDocumentVersionUpload", params, optFns, addOperationInitiateDocumentVersionUploadMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*InitiateDocumentVersionUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InitiateDocumentVersionUploadInput struct {

	// The ID of the parent folder.
	//
	// This member is required.
	ParentFolderId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The timestamp when the content of the document was originally created.
	ContentCreatedTimestamp *time.Time

	// The timestamp when the content of the document was modified.
	ContentModifiedTimestamp *time.Time

	// The content type of the document.
	ContentType *string

	// The size of the document, in bytes.
	DocumentSizeInBytes *int64

	// The ID of the document.
	Id *string

	// The name of the document.
	Name *string
}

type InitiateDocumentVersionUploadOutput struct {

	// The document metadata.
	Metadata *types.DocumentMetadata

	// The upload metadata.
	UploadMetadata *types.UploadMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInitiateDocumentVersionUploadMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInitiateDocumentVersionUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInitiateDocumentVersionUpload{}, middleware.After)
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
	addOpInitiateDocumentVersionUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateDocumentVersionUpload(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opInitiateDocumentVersionUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "InitiateDocumentVersionUpload",
	}
}
