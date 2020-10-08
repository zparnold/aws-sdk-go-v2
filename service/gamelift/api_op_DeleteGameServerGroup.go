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

// This action is part of Amazon GameLift FleetIQ with game server groups, which is
// in preview release and is subject to change. Terminates a game server group and
// permanently deletes the game server group record. You have several options for
// how these resources are impacted when deleting the game server group. Depending
// on the type of delete action selected, this action may affect three types of
// resources: the game server group, the corresponding Auto Scaling group, and all
// game servers currently running in the group. To delete a game server group,
// identify the game server group to delete and specify the type of delete action
// to initiate. Game server groups can only be deleted if they are in ACTIVE or
// ERROR status. If the delete request is successful, a series of actions are
// kicked off. The game server group status is changed to DELETE_SCHEDULED, which
// prevents new game servers from being registered and stops autoscaling activity.
// Once all game servers in the game server group are de-registered, GameLift
// FleetIQ can begin deleting resources. If any of the delete actions fail, the
// game server group is placed in ERROR status. GameLift FleetIQ emits delete
// events to Amazon CloudWatch. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
// Related operations
//
//     * CreateGameServerGroup ()
//
//     * ListGameServerGroups
// ()
//
//     * DescribeGameServerGroup ()
//
//     * UpdateGameServerGroup ()
//
//     *
// DeleteGameServerGroup ()
//
//     * ResumeGameServerGroup ()
//
//     *
// SuspendGameServerGroup ()
func (c *Client) DeleteGameServerGroup(ctx context.Context, params *DeleteGameServerGroupInput, optFns ...func(*Options)) (*DeleteGameServerGroupOutput, error) {
	if params == nil {
		params = &DeleteGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGameServerGroup", params, optFns, addOperationDeleteGameServerGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGameServerGroupInput struct {

	// The unique identifier of the game server group to delete. Use either the
	// GameServerGroup () name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The type of delete to perform. Options include:
	//
	//     * SAFE_DELETE – Terminates
	// the game server group and EC2 Auto Scaling group only when it has no game
	// servers that are in IN_USE status.
	//
	//     * FORCE_DELETE – Terminates the game
	// server group, including all active game servers regardless of their utilization
	// status, and the EC2 Auto Scaling group.
	//
	//     * RETAIN – Does a safe delete of
	// the game server group but retains the EC2 Auto Scaling group as is.
	DeleteOption types.GameServerGroupDeleteOption
}

type DeleteGameServerGroupOutput struct {

	// An object that describes the deleted game server group resource, with status
	// updated to DELETE_SCHEDULED.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGameServerGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteGameServerGroup{}, middleware.After)
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
	addOpDeleteGameServerGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGameServerGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteGameServerGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteGameServerGroup",
	}
}
