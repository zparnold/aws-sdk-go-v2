// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This action is part of Amazon GameLift FleetIQ with game server groups, which is
// in preview release and is subject to change. Locates an available game server
// and temporarily reserves it to host gameplay and players. This action is called
// by a game client or client service (such as a matchmaker) to request hosting
// resources for a new game session. In response, GameLift FleetIQ searches for an
// available game server in the specified game server group, places the game server
// in "claimed" status for 60 seconds, and returns connection information back to
// the requester so that players can connect to the game server. There are two ways
// you can claim a game server. For the first option, you provide a game server
// group ID only, which prompts GameLift FleetIQ to search for an available game
// server in the specified group and claim it. With this option, GameLift FleetIQ
// attempts to consolidate gameplay on as few instances as possible to minimize
// hosting costs. For the second option, you request a specific game server by its
// ID. This option results in a less efficient claiming process because it does not
// take advantage of consolidation and may fail if the requested game server is
// unavailable. To claim a game server, identify a game server group and
// (optionally) a game server ID. If your game requires that game data be provided
// to the game server at the start of a game, such as a game map or player
// information, you can provide it in your claim request. When a game server is
// successfully claimed, connection information is returned. A claimed game
// server's utilization status remains AVAILABLE, while the claim status is set to
// CLAIMED for up to 60 seconds. This time period allows the game server to be
// prompted to update its status to UTILIZED (using UpdateGameServer ()). If the
// game server's status is not updated within 60 seconds, the game server reverts
// to unclaimed status and is available to be claimed by another request. If you
// try to claim a specific game server, this request will fail in the following
// cases: (1) if the game server utilization status is UTILIZED, (2) if the game
// server claim status is CLAIMED, or (3) if the instance that the game server is
// running on is flagged as draining. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
// Related operations
//
//     * RegisterGameServer ()
//
//     * ListGameServers ()
//
//     *
// ClaimGameServer ()
//
//     * DescribeGameServer ()
//
//     * UpdateGameServer ()
//
//
// * DeregisterGameServer ()
func (c *Client) ClaimGameServer(ctx context.Context, params *ClaimGameServerInput, optFns ...func(*Options)) (*ClaimGameServerOutput, error) {
	stack := middleware.NewStack("ClaimGameServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpClaimGameServerMiddlewares(stack)
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
	addOpClaimGameServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opClaimGameServer(options.Region), middleware.Before)
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
			OperationName: "ClaimGameServer",
			Err:           err,
		}
	}
	out := result.(*ClaimGameServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ClaimGameServerInput struct {

	// An identifier for the game server group. When claiming a specific game server,
	// this is the game server group whether the game server is located. When
	// requesting that GameLift FleetIQ locate an available game server, this is the
	// game server group to search on. You can use either the GameServerGroup () name
	// or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// A set of custom game server properties, formatted as a single string value, to
	// be passed to the claimed game server.
	GameServerData *string

	// A custom string that uniquely identifies the game server to claim. If this
	// parameter is left empty, GameLift FleetIQ searches for an available game server
	// in the specified game server group.
	GameServerId *string
}

type ClaimGameServerOutput struct {

	// Object that describes the newly claimed game server resource.
	GameServer *types.GameServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpClaimGameServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpClaimGameServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpClaimGameServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opClaimGameServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ClaimGameServer",
	}
}
