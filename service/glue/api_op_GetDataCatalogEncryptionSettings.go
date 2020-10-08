// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the security configuration for a specified catalog.
func (c *Client) GetDataCatalogEncryptionSettings(ctx context.Context, params *GetDataCatalogEncryptionSettingsInput, optFns ...func(*Options)) (*GetDataCatalogEncryptionSettingsOutput, error) {
	if params == nil {
		params = &GetDataCatalogEncryptionSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataCatalogEncryptionSettings", params, optFns, addOperationGetDataCatalogEncryptionSettingsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataCatalogEncryptionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataCatalogEncryptionSettingsInput struct {

	// The ID of the Data Catalog to retrieve the security configuration for. If none
	// is provided, the AWS account ID is used by default.
	CatalogId *string
}

type GetDataCatalogEncryptionSettingsOutput struct {

	// The requested security configuration.
	DataCatalogEncryptionSettings *types.DataCatalogEncryptionSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDataCatalogEncryptionSettingsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDataCatalogEncryptionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDataCatalogEncryptionSettings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataCatalogEncryptionSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDataCatalogEncryptionSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetDataCatalogEncryptionSettings",
	}
}
