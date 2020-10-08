// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified Systems Manager document.
func (c *Client) DescribeDocument(ctx context.Context, params *DescribeDocumentInput, optFns ...func(*Options)) (*DescribeDocumentOutput, error) {
	if params == nil {
		params = &DescribeDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDocument", params, optFns, addOperationDescribeDocumentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDocumentInput struct {

	// The name of the Systems Manager document.
	//
	// This member is required.
	Name *string

	// The document version for which you want information. Can be a specific version
	// or the default version.
	DocumentVersion *string

	// An optional field specifying the version of the artifact associated with the
	// document. For example, "Release 12, Update 6". This value is unique across all
	// versions of a document, and cannot be changed.
	VersionName *string
}

type DescribeDocumentOutput struct {

	// Information about the Systems Manager document.
	Document *types.DocumentDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDocumentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDocument{}, middleware.After)
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
	addOpDescribeDocumentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDocument(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDocument(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeDocument",
	}
}
