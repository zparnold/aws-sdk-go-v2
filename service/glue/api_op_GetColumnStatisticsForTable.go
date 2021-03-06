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

// Retrieves table statistics of columns.
func (c *Client) GetColumnStatisticsForTable(ctx context.Context, params *GetColumnStatisticsForTableInput, optFns ...func(*Options)) (*GetColumnStatisticsForTableOutput, error) {
	stack := middleware.NewStack("GetColumnStatisticsForTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetColumnStatisticsForTableMiddlewares(stack)
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
	addOpGetColumnStatisticsForTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetColumnStatisticsForTable(options.Region), middleware.Before)
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
			OperationName: "GetColumnStatisticsForTable",
			Err:           err,
		}
	}
	out := result.(*GetColumnStatisticsForTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetColumnStatisticsForTableInput struct {

	// A list of the column names.
	//
	// This member is required.
	ColumnNames []*string

	// The name of the catalog database where the partitions reside.
	//
	// This member is required.
	DatabaseName *string

	// The name of the partitions' table.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the partitions in question reside. If none is
	// supplied, the AWS account ID is used by default.
	CatalogId *string
}

type GetColumnStatisticsForTableOutput struct {

	// List of ColumnStatistics that failed to be retrieved.
	ColumnStatisticsList []*types.ColumnStatistics

	// List of ColumnStatistics that failed to be retrieved.
	Errors []*types.ColumnError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetColumnStatisticsForTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetColumnStatisticsForTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetColumnStatisticsForTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetColumnStatisticsForTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetColumnStatisticsForTable",
	}
}
