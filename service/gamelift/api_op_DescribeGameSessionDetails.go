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

// Retrieves properties, including the protection policy in force, for one or more
// game sessions. This action can be used in several ways: (1) provide a
// GameSessionId or GameSessionArn to request details for a specific game session;
// (2) provide either a FleetId or an AliasId to request properties for all game
// sessions running on a fleet. To get game session record(s), specify just one of
// the following: game session ID, fleet ID, or alias ID. You can filter this
// request by game session status. Use the pagination parameters to retrieve
// results as a set of sequential pages. If successful, a GameSessionDetail ()
// object is returned for each session matching the request.
//
//     *
// CreateGameSession ()
//
//     * DescribeGameSessions ()
//
//     *
// DescribeGameSessionDetails ()
//
//     * SearchGameSessions ()
//
//     *
// UpdateGameSession ()
//
//     * GetGameSessionLogUrl ()
//
//     * Game session
// placements
//
//         * StartGameSessionPlacement ()
//
//         *
// DescribeGameSessionPlacement ()
//
//         * StopGameSessionPlacement ()
func (c *Client) DescribeGameSessionDetails(ctx context.Context, params *DescribeGameSessionDetailsInput, optFns ...func(*Options)) (*DescribeGameSessionDetailsOutput, error) {
	if params == nil {
		params = &DescribeGameSessionDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameSessionDetails", params, optFns, addOperationDescribeGameSessionDetailsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameSessionDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeGameSessionDetailsInput struct {

	// A unique identifier for an alias associated with the fleet to retrieve all game
	// sessions for. You can use either the alias ID or ARN value.
	AliasId *string

	// A unique identifier for a fleet to retrieve all game sessions active on the
	// fleet. You can use either the fleet ID or ARN value.
	FleetId *string

	// A unique identifier for the game session to retrieve.
	GameSessionId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// Game session status to filter results on. Possible game session statuses include
	// ACTIVE, TERMINATED, ACTIVATING and TERMINATING (the last two are transitory).
	StatusFilter *string
}

// Represents the returned data in response to a request action.
type DescribeGameSessionDetailsOutput struct {

	// A collection of objects containing game session properties and the protection
	// policy currently in force for each session matching the request.
	GameSessionDetails []*types.GameSessionDetail

	// Token that indicates where to resume retrieving results on the next call to this
	// action. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGameSessionDetailsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameSessionDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameSessionDetails{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameSessionDetails(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeGameSessionDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameSessionDetails",
	}
}
