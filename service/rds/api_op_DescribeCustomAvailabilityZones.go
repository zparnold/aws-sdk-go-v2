// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about custom Availability Zones (AZs). A custom AZ is an
// on-premises AZ that is integrated with a VMware vSphere cluster. For more
// information about RDS on VMware, see the  RDS on VMware User Guide.
// (https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
func (c *Client) DescribeCustomAvailabilityZones(ctx context.Context, params *DescribeCustomAvailabilityZonesInput, optFns ...func(*Options)) (*DescribeCustomAvailabilityZonesOutput, error) {
	stack := middleware.NewStack("DescribeCustomAvailabilityZones", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeCustomAvailabilityZonesMiddlewares(stack)
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
	addOpDescribeCustomAvailabilityZonesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomAvailabilityZones(options.Region), middleware.Before)
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
			OperationName: "DescribeCustomAvailabilityZones",
			Err:           err,
		}
	}
	out := result.(*DescribeCustomAvailabilityZonesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomAvailabilityZonesInput struct {

	// The custom AZ identifier. If this parameter is specified, information from only
	// the specific custom AZ is returned.
	CustomAvailabilityZoneId *string

	// A filter that specifies one or more custom AZs to describe.
	Filters []*types.Filter

	// An optional pagination token provided by a previous
	// DescribeCustomAvailabilityZones request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeCustomAvailabilityZonesOutput struct {

	// The list of CustomAvailabilityZone () objects for the AWS account.
	CustomAvailabilityZones []*types.CustomAvailabilityZone

	// An optional pagination token provided by a previous
	// DescribeCustomAvailabilityZones request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeCustomAvailabilityZonesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCustomAvailabilityZones{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCustomAvailabilityZones{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCustomAvailabilityZones(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeCustomAvailabilityZones",
	}
}
