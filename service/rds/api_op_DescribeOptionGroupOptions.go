// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes all available options.
func (c *Client) DescribeOptionGroupOptions(ctx context.Context, params *DescribeOptionGroupOptionsInput, optFns ...func(*Options)) (*DescribeOptionGroupOptionsOutput, error) {
	if params == nil {
		params = &DescribeOptionGroupOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOptionGroupOptions", params, optFns, addOperationDescribeOptionGroupOptionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOptionGroupOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeOptionGroupOptionsInput struct {

	// A required parameter. Options available for the given engine name are described.
	//
	// This member is required.
	EngineName *string

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// If specified, filters the results to include only options for the specified
	// major engine version.
	MajorEngineVersion *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

//
type DescribeOptionGroupOptionsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// List of available option group options.
	OptionGroupOptions []*types.OptionGroupOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeOptionGroupOptionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOptionGroupOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOptionGroupOptions{}, middleware.After)
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
	addOpDescribeOptionGroupOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOptionGroupOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeOptionGroupOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeOptionGroupOptions",
	}
}
