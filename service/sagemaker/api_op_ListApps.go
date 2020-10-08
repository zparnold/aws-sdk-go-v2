// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists apps.
func (c *Client) ListApps(ctx context.Context, params *ListAppsInput, optFns ...func(*Options)) (*ListAppsOutput, error) {
	if params == nil {
		params = &ListAppsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApps", params, optFns, addOperationListAppsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppsInput struct {

	// A parameter to search for the domain ID.
	DomainIdEquals *string

	// Returns a list up to a specified limit.
	MaxResults *int32

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy types.AppSortKey

	// The sort order for the results. The default is Ascending.
	SortOrder types.SortOrder

	// A parameter to search by user profile name.
	UserProfileNameEquals *string
}

type ListAppsOutput struct {

	// The list of apps.
	Apps []*types.AppDetails

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAppsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApps{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListApps(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListApps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListApps",
	}
}
