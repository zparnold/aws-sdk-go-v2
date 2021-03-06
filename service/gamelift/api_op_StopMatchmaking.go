// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels a matchmaking ticket or match backfill ticket that is currently being
// processed. To stop the matchmaking operation, specify the ticket ID. If
// successful, work on the ticket is stopped, and the ticket status is changed to
// CANCELLED. This call is also used to turn off automatic backfill for an
// individual game session. This is for game sessions that are created with a
// matchmaking configuration that has automatic backfill enabled. The ticket ID is
// included in the MatchmakerData of an updated game session object, which is
// provided to the game server. If the action is successful, the service sends back
// an empty JSON struct with the HTTP 200 response (not an empty HTTP body). Learn
// more  Add FlexMatch to a Game Client
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-client.html)
// Related operations
//
//     * StartMatchmaking ()
//
//     * DescribeMatchmaking ()
//
//
// * StopMatchmaking ()
//
//     * AcceptMatch ()
//
//     * StartMatchBackfill ()
func (c *Client) StopMatchmaking(ctx context.Context, params *StopMatchmakingInput, optFns ...func(*Options)) (*StopMatchmakingOutput, error) {
	stack := middleware.NewStack("StopMatchmaking", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopMatchmakingMiddlewares(stack)
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
	addOpStopMatchmakingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopMatchmaking(options.Region), middleware.Before)
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
			OperationName: "StopMatchmaking",
			Err:           err,
		}
	}
	out := result.(*StopMatchmakingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type StopMatchmakingInput struct {

	// A unique identifier for a matchmaking ticket.
	//
	// This member is required.
	TicketId *string
}

type StopMatchmakingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopMatchmakingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopMatchmaking{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopMatchmaking{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopMatchmaking(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StopMatchmaking",
	}
}
