// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the data catalog that has the specified name.
func (c *Client) UpdateDataCatalog(ctx context.Context, params *UpdateDataCatalogInput, optFns ...func(*Options)) (*UpdateDataCatalogOutput, error) {
	if params == nil {
		params = &UpdateDataCatalogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataCatalog", params, optFns, addOperationUpdateDataCatalogMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataCatalogInput struct {

	// The name of the data catalog to update. The catalog name must be unique for the
	// AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or
	// hyphen characters.
	//
	// This member is required.
	Name *string

	// Specifies the type of data catalog to update. Specify LAMBDA for a federated
	// catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	//
	// This member is required.
	Type types.DataCatalogType

	// New or modified text that describes the data catalog.
	Description *string

	// Specifies the Lambda function or functions to use for updating the data catalog.
	// This is a mapping whose values depend on the catalog type.
	//
	//     * For the HIVE
	// data catalog type, use the following syntax. The metadata-function parameter is
	// required. The sdk-version parameter is optional and defaults to the currently
	// supported version. metadata-function=lambda_arn, sdk-version=version_number
	//
	//
	// * For the LAMBDA data catalog type, use one of the following sets of required
	// parameters, but not both.
	//
	//         * If you have one Lambda function that
	// processes metadata and another for reading the actual data, use the following
	// syntax. Both parameters are required. metadata-function=lambda_arn,
	// record-function=lambda_arn
	//
	//         * If you have a composite Lambda function
	// that processes both metadata and data, use the following syntax to specify your
	// Lambda function. function=lambda_arn
	//
	//     * The GLUE type has no parameters.
	Parameters map[string]*string
}

type UpdateDataCatalogOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDataCatalogMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDataCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDataCatalog{}, middleware.After)
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
	addOpUpdateDataCatalogValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataCatalog(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDataCatalog(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "UpdateDataCatalog",
	}
}
