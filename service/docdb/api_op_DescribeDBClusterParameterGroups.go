// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of DBClusterParameterGroup descriptions. If a
// DBClusterParameterGroupName parameter is specified, the list contains only the
// description of the specified cluster parameter group.
func (c *Client) DescribeDBClusterParameterGroups(ctx context.Context, params *DescribeDBClusterParameterGroupsInput, optFns ...func(*Options)) (*DescribeDBClusterParameterGroupsOutput, error) {
	if params == nil {
		params = &DescribeDBClusterParameterGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterParameterGroups", params, optFns, addOperationDescribeDBClusterParameterGroupsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBClusterParameterGroups ().
type DescribeDBClusterParameterGroupsInput struct {

	// The name of a specific cluster parameter group to return details for.
	// Constraints:
	//
	//     * If provided, must match the name of an existing
	// DBClusterParameterGroup.
	DBClusterParameterGroupName *string

	// This parameter is not currently supported.
	Filters []*types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Represents the output of DBClusterParameterGroups ().
type DescribeDBClusterParameterGroupsOutput struct {

	// A list of cluster parameter groups.
	DBClusterParameterGroups []*types.DBClusterParameterGroup

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBClusterParameterGroupsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterParameterGroups{}, middleware.After)
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
	addOpDescribeDBClusterParameterGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterParameterGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBClusterParameterGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterParameterGroups",
	}
}
