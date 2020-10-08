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

// Displays a list of categories for all event source types, or, if specified, for
// a specified source type.
func (c *Client) DescribeEventCategories(ctx context.Context, params *DescribeEventCategoriesInput, optFns ...func(*Options)) (*DescribeEventCategoriesOutput, error) {
	if params == nil {
		params = &DescribeEventCategoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventCategories", params, optFns, addOperationDescribeEventCategoriesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventCategoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeEventCategories ().
type DescribeEventCategoriesInput struct {

	// This parameter is not currently supported.
	Filters []*types.Filter

	// The type of source that is generating the events. Valid values: db-instance,
	// db-parameter-group, db-security-group, db-snapshot
	SourceType *string
}

// Represents the output of DescribeEventCategories ().
type DescribeEventCategoriesOutput struct {

	// A list of event category maps.
	EventCategoriesMapList []*types.EventCategoriesMap

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventCategoriesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEventCategories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEventCategories{}, middleware.After)
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
	addOpDescribeEventCategoriesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventCategories(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventCategories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeEventCategories",
	}
}
