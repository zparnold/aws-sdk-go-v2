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

// Lists available reserved DB instance offerings.
func (c *Client) DescribeReservedDBInstancesOfferings(ctx context.Context, params *DescribeReservedDBInstancesOfferingsInput, optFns ...func(*Options)) (*DescribeReservedDBInstancesOfferingsOutput, error) {
	stack := middleware.NewStack("DescribeReservedDBInstancesOfferings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeReservedDBInstancesOfferingsMiddlewares(stack)
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
	addOpDescribeReservedDBInstancesOfferingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedDBInstancesOfferings(options.Region), middleware.Before)

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
			OperationName: "DescribeReservedDBInstancesOfferings",
			Err:           err,
		}
	}
	out := result.(*DescribeReservedDBInstancesOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReservedDBInstancesOfferingsInput struct {
	// The DB instance class filter value. Specify this parameter to show only the
	// available offerings matching the specified DB instance class.
	DBInstanceClass *string
	// The offering identifier filter value. Specify this parameter to show only the
	// available offering that matches the specified reservation identifier. Example:
	// 438012d3-4052-4cc7-b2e3-8d3372e0e706
	ReservedDBInstancesOfferingId *string
	// A value that indicates whether to show only those reservations that support
	// Multi-AZ.
	MultiAZ *bool
	// This parameter isn't currently supported.
	Filters []*types.Filter
	// Product description filter value. Specify this parameter to show only the
	// available offerings that contain the specified product description. The results
	// show offerings that partially match the filter value.
	ProductDescription *string
	// Duration filter value, specified in years or seconds. Specify this parameter to
	// show only reservations for this duration. Valid Values: 1 | 3 | 31536000 |
	// 94608000
	Duration *string
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string
	// The offering type filter value. Specify this parameter to show only the
	// available offerings matching the specified offering type. Valid Values: "Partial
	// Upfront" | "All Upfront" | "No Upfront"
	OfferingType *string
	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Contains the result of a successful invocation of the
// DescribeReservedDBInstancesOfferings action.
type DescribeReservedDBInstancesOfferingsOutput struct {
	// A list of reserved DB instance offerings.
	ReservedDBInstancesOfferings []*types.ReservedDBInstancesOffering
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeReservedDBInstancesOfferingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReservedDBInstancesOfferings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReservedDBInstancesOfferings{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeReservedDBInstancesOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeReservedDBInstancesOfferings",
	}
}
