// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves (queries) statistical data and other information about one or more S3
// buckets that Amazon Macie monitors and analyzes.
func (c *Client) DescribeBuckets(ctx context.Context, params *DescribeBucketsInput, optFns ...func(*Options)) (*DescribeBucketsOutput, error) {
	if params == nil {
		params = &DescribeBucketsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBuckets", params, optFns, addOperationDescribeBucketsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBucketsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBucketsInput struct {

	// The criteria to use to filter the query results.
	Criteria map[string]*types.BucketCriteriaAdditionalProperties

	// The maximum number of items to include in each page of the response. The default
	// value is 50.
	MaxResults *int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string

	// The criteria to use to sort the query results.
	SortCriteria *types.BucketSortCriteria
}

type DescribeBucketsOutput struct {

	// An array of objects, one for each bucket that meets the filter criteria
	// specified in the request.
	Buckets []*types.BucketMetadata

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBucketsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBuckets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBuckets{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBuckets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBuckets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "DescribeBuckets",
	}
}
