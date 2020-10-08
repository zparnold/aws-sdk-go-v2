// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete the partition column statistics of a column.
func (c *Client) DeleteColumnStatisticsForPartition(ctx context.Context, params *DeleteColumnStatisticsForPartitionInput, optFns ...func(*Options)) (*DeleteColumnStatisticsForPartitionOutput, error) {
	if params == nil {
		params = &DeleteColumnStatisticsForPartitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteColumnStatisticsForPartition", params, optFns, addOperationDeleteColumnStatisticsForPartitionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteColumnStatisticsForPartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteColumnStatisticsForPartitionInput struct {

	// Name of the column.
	//
	// This member is required.
	ColumnName *string

	// The name of the catalog database where the partitions reside.
	//
	// This member is required.
	DatabaseName *string

	// A list of partition values identifying the partition.
	//
	// This member is required.
	PartitionValues []*string

	// The name of the partitions' table.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the partitions in question reside. If none is
	// supplied, the AWS account ID is used by default.
	CatalogId *string
}

type DeleteColumnStatisticsForPartitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteColumnStatisticsForPartitionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteColumnStatisticsForPartition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteColumnStatisticsForPartition{}, middleware.After)
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
	addOpDeleteColumnStatisticsForPartitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteColumnStatisticsForPartition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteColumnStatisticsForPartition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteColumnStatisticsForPartition",
	}
}
