// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a dataset.
func (c *Client) UpdateDataSet(ctx context.Context, params *UpdateDataSetInput, optFns ...func(*Options)) (*UpdateDataSetOutput, error) {
	if params == nil {
		params = &UpdateDataSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSet", params, optFns, addOperationUpdateDataSetMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSetInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dataset that you want to update. This ID is unique per AWS Region
	// for each AWS account.
	//
	// This member is required.
	DataSetId *string

	// Indicates whether you want to import the data into SPICE.
	//
	// This member is required.
	ImportMode types.DataSetImportMode

	// The display name for the dataset.
	//
	// This member is required.
	Name *string

	// Declares the physical tables that are available in the underlying data sources.
	//
	// This member is required.
	PhysicalTableMap map[string]*types.PhysicalTable

	// Groupings of columns that work together in certain QuickSight features.
	// Currently, only geospatial hierarchy is supported.
	ColumnGroups []*types.ColumnGroup

	// Configures the combination and transformation of the data from the physical
	// tables.
	LogicalTableMap map[string]*types.LogicalTable

	// The row-level security configuration for the data you want to create.
	RowLevelPermissionDataSet *types.RowLevelPermissionDataSet
}

type UpdateDataSetOutput struct {

	// The Amazon Resource Name (ARN) of the dataset.
	Arn *string

	// The ID for the dataset that you want to create. This ID is unique per AWS Region
	// for each AWS account.
	DataSetId *string

	// The ARN for the ingestion, which is triggered as a result of dataset creation if
	// the import mode is SPICE.
	IngestionArn *string

	// The ID of the ingestion, which is triggered as a result of dataset creation if
	// the import mode is SPICE.
	IngestionId *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDataSetMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataSet{}, middleware.After)
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
	addOpUpdateDataSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDataSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDataSet",
	}
}
