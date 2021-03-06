// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of strings that identify available versions of a specified
// table.
func (c *Client) GetTableVersions(ctx context.Context, params *GetTableVersionsInput, optFns ...func(*Options)) (*GetTableVersionsOutput, error) {
	stack := middleware.NewStack("GetTableVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTableVersionsMiddlewares(stack)
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
	addOpGetTableVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTableVersions(options.Region), middleware.Before)
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
			OperationName: "GetTableVersions",
			Err:           err,
		}
	}
	out := result.(*GetTableVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTableVersionsInput struct {

	// The database in the catalog in which the table resides. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table. For Hive compatibility, this name is entirely lowercase.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the tables reside. If none is provided, the AWS
	// account ID is used by default.
	CatalogId *string

	// The maximum number of table versions to return in one response.
	MaxResults *int32

	// A continuation token, if this is not the first call.
	NextToken *string
}

type GetTableVersionsOutput struct {

	// A continuation token, if the list of available versions does not include the
	// last one.
	NextToken *string

	// A list of strings identifying available versions of the specified table.
	TableVersions []*types.TableVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTableVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTableVersions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTableVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTableVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetTableVersions",
	}
}
