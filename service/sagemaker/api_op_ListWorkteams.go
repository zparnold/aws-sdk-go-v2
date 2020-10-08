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

// Gets a list of work teams that you have defined in a region. The list may be
// empty if no work team satisfies the filter specified in the NameContains
// parameter.
func (c *Client) ListWorkteams(ctx context.Context, params *ListWorkteamsInput, optFns ...func(*Options)) (*ListWorkteamsOutput, error) {
	if params == nil {
		params = &ListWorkteamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkteams", params, optFns, addOperationListWorkteamsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkteamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkteamsInput struct {

	// The maximum number of work teams to return in each page of the response.
	MaxResults *int32

	// A string in the work team's name. This filter returns only work teams whose name
	// contains the specified string.
	NameContains *string

	// If the result of the previous ListWorkteams request was truncated, the response
	// includes a NextToken. To retrieve the next set of labeling jobs, use the token
	// in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.ListWorkteamsSortByOptions

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder
}

type ListWorkteamsOutput struct {

	// An array of Workteam objects, each describing a work team.
	//
	// This member is required.
	Workteams []*types.Workteam

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of work teams, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWorkteamsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkteams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkteams{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkteams(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListWorkteams(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListWorkteams",
	}
}
