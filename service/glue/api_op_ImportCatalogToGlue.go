// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Imports an existing Amazon Athena Data Catalog to AWS Glue
func (c *Client) ImportCatalogToGlue(ctx context.Context, params *ImportCatalogToGlueInput, optFns ...func(*Options)) (*ImportCatalogToGlueOutput, error) {
	if params == nil {
		params = &ImportCatalogToGlueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportCatalogToGlue", params, optFns, addOperationImportCatalogToGlueMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportCatalogToGlueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportCatalogToGlueInput struct {

	// The ID of the catalog to import. Currently, this should be the AWS account ID.
	CatalogId *string
}

type ImportCatalogToGlueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportCatalogToGlueMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportCatalogToGlue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportCatalogToGlue{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportCatalogToGlue(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opImportCatalogToGlue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ImportCatalogToGlue",
	}
}
