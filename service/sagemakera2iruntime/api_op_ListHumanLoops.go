// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakera2iruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about human loops, given the specified parameters. If a
// human loop was deleted, it will not be included.
func (c *Client) ListHumanLoops(ctx context.Context, params *ListHumanLoopsInput, optFns ...func(*Options)) (*ListHumanLoopsOutput, error) {
	if params == nil {
		params = &ListHumanLoopsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHumanLoops", params, optFns, addOperationListHumanLoopsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHumanLoopsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHumanLoopsInput struct {

	// The Amazon Resource Name (ARN) of a flow definition.
	//
	// This member is required.
	FlowDefinitionArn *string

	// (Optional) The timestamp of the date when you want the human loops to begin in
	// ISO 8601 format. For example, 2020-02-24.
	CreationTimeAfter *time.Time

	// (Optional) The timestamp of the date before which you want the human loops to
	// begin in ISO 8601 format. For example, 2020-02-24.
	CreationTimeBefore *time.Time

	// The total number of items to return. If the total number of available items is
	// more than the value specified in MaxResults, then a NextToken is returned in the
	// output. You can use this token to display the next page of results.
	MaxResults *int32

	// A token to display the next page of results.
	NextToken *string

	// Optional. The order for displaying results. Valid values: Ascending and
	// Descending.
	SortOrder types.SortOrder
}

type ListHumanLoopsOutput struct {

	// An array of objects that contain information about the human loops.
	//
	// This member is required.
	HumanLoopSummaries []*types.HumanLoopSummary

	// A token to display the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListHumanLoopsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListHumanLoops{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListHumanLoops{}, middleware.After)
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
	addOpListHumanLoopsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListHumanLoops(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListHumanLoops(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListHumanLoops",
	}
}
