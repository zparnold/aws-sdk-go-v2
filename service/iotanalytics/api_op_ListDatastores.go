// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of data stores.
func (c *Client) ListDatastores(ctx context.Context, params *ListDatastoresInput, optFns ...func(*Options)) (*ListDatastoresOutput, error) {
	if params == nil {
		params = &ListDatastoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatastores", params, optFns, addOperationListDatastoresMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatastoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatastoresInput struct {

	// The maximum number of results to return in this request. The default value is
	// 100.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string
}

type ListDatastoresOutput struct {

	// A list of "DatastoreSummary" objects.
	DatastoreSummaries []*types.DatastoreSummary

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDatastoresMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDatastores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDatastores{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDatastores(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDatastores(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "ListDatastores",
	}
}
