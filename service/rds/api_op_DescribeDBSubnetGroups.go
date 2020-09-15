// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of DBSubnetGroup descriptions. If a DBSubnetGroupName is
// specified, the list will contain only the descriptions of the specified
// DBSubnetGroup. For an overview of CIDR ranges, go to the Wikipedia Tutorial
// (http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
func (c *Client) DescribeDBSubnetGroups(ctx context.Context, params *DescribeDBSubnetGroupsInput, optFns ...func(*Options)) (*DescribeDBSubnetGroupsOutput, error) {
	stack := middleware.NewStack("DescribeDBSubnetGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBSubnetGroupsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeDBSubnetGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBSubnetGroups(options.Region), middleware.Before)

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
			OperationName: "DescribeDBSubnetGroups",
			Err:           err,
		}
	}
	out := result.(*DescribeDBSubnetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBSubnetGroupsInput struct {
	// This parameter isn't currently supported.
	Filters []*types.Filter
	// The name of the DB subnet group to return details for.
	DBSubnetGroupName *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// An optional pagination token provided by a previous DescribeDBSubnetGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string
}

// Contains the result of a successful invocation of the DescribeDBSubnetGroups
// action.
type DescribeDBSubnetGroupsOutput struct {
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string
	// A list of DBSubnetGroup instances.
	DBSubnetGroups []*types.DBSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBSubnetGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBSubnetGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBSubnetGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBSubnetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBSubnetGroups",
	}
}
