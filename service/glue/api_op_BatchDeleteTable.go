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

// Deletes multiple tables at once. After completing this operation, you no longer
// have access to the table versions and partitions that belong to the deleted
// table. AWS Glue deletes these "orphaned" resources asynchronously in a timely
// manner, at the discretion of the service. To ensure the immediate deletion of
// all related resources, before calling BatchDeleteTable, use DeleteTableVersion
// or BatchDeleteTableVersion, and DeletePartition or BatchDeletePartition, to
// delete any resources that belong to the table.
func (c *Client) BatchDeleteTable(ctx context.Context, params *BatchDeleteTableInput, optFns ...func(*Options)) (*BatchDeleteTableOutput, error) {
	if params == nil {
		params = &BatchDeleteTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteTable", params, optFns, addOperationBatchDeleteTableMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteTableInput struct {

	// The name of the catalog database in which the tables to delete reside. For Hive
	// compatibility, this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// A list of the table to delete.
	//
	// This member is required.
	TablesToDelete []*string

	// The ID of the Data Catalog where the table resides. If none is provided, the AWS
	// account ID is used by default.
	CatalogId *string
}

type BatchDeleteTableOutput struct {

	// A list of errors encountered in attempting to delete the specified tables.
	Errors []*types.TableError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeleteTableMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteTable{}, middleware.After)
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
	addOpBatchDeleteTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchDeleteTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchDeleteTable",
	}
}
