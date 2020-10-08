// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves all aliases for this AWS account. You can filter the result set by
// alias name and/or routing strategy type. Use the pagination parameters to
// retrieve results in sequential pages. Returned aliases are not listed in any
// particular order.
//
//     * CreateAlias ()
//
//     * ListAliases ()
//
//     *
// DescribeAlias ()
//
//     * UpdateAlias ()
//
//     * DeleteAlias ()
//
//     * ResolveAlias
// ()
func (c *Client) ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if params == nil {
		params = &ListAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAliases", params, optFns, addOperationListAliasesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type ListAliasesInput struct {

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A descriptive label that is associated with an alias. Alias names do not need to
	// be unique.
	Name *string

	// A token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// The routing type to filter results on. Use this parameter to retrieve only
	// aliases with a certain routing type. To retrieve all aliases, leave this
	// parameter empty. Possible routing types include the following:
	//
	//     * SIMPLE --
	// The alias resolves to one specific fleet. Use this type when routing to active
	// fleets.
	//
	//     * TERMINAL -- The alias does not resolve to a fleet but instead can
	// be used to display a message to the user. A terminal alias throws a
	// TerminalRoutingStrategyException with the RoutingStrategy () message embedded.
	RoutingStrategyType types.RoutingStrategyType
}

// Represents the returned data in response to a request action.
type ListAliasesOutput struct {

	// A collection of alias resources that match the request parameters.
	Aliases []*types.Alias

	// A token that indicates where to resume retrieving results on the next call to
	// this action. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAliasesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAliases{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAliases(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAliases(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListAliases",
	}
}
