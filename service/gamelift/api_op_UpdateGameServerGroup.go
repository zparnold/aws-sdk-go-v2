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
// in preview release and is subject to change. Updates GameLift FleetIQ-specific
// properties for a game server group. These properties include instance
// rebalancing and game server protection. Many Auto Scaling group properties are
// updated directly. These include autoscaling policies, minimum/maximum/desired
// instance counts, and launch template. To update the game server group, specify
// the game server group ID and provide the updated values. Updated properties are
// validated to ensure that GameLift FleetIQ can continue to perform its core
// instance rebalancing activity. When you change Auto Scaling group properties
// directly and the changes cause errors with GameLift FleetIQ activities, an alert
// is sent. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)Updating
// a GameLift FleetIQ-Linked Auto Scaling Group
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-asgroups.html)
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
func (c *Client) UpdateGameServerGroup(ctx context.Context, params *UpdateGameServerGroupInput, optFns ...func(*Options)) (*UpdateGameServerGroupOutput, error) {
	if params == nil {
		params = &UpdateGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGameServerGroup", params, optFns, addOperationUpdateGameServerGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGameServerGroupInput struct {

	// The unique identifier of the game server group to update. Use either the
	// GameServerGroup () name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The fallback balancing method to use for the game server group when Spot
	// instances in a Region become unavailable or are not viable for game hosting.
	// Once triggered, this method remains active until Spot instances can once again
	// be used. Method options include:
	//
	//     * SPOT_ONLY -- If Spot instances are
	// unavailable, the game server group provides no hosting capacity. No new
	// instances are started, and the existing nonviable Spot instances are terminated
	// (once current gameplay ends) and not replaced.
	//
	//     * SPOT_PREFERRED -- If Spot
	// instances are unavailable, the game server group continues to provide hosting
	// capacity by using On-Demand instances. Existing nonviable Spot instances are
	// terminated (once current gameplay ends) and replaced with new On-Demand
	// instances.
	BalancingStrategy types.BalancingStrategy

	// A flag that indicates whether instances in the game server group are protected
	// from early termination. Unprotected instances that have active game servers
	// running may by terminated during a scale-down event, causing players to be
	// dropped from the game. Protected instances cannot be terminated while there are
	// active game servers running. An exception to this is Spot Instances, which may
	// be terminated by AWS regardless of protection status. This property is set to
	// NO_PROTECTION by default.
	GameServerProtectionPolicy types.GameServerProtectionPolicy

	// An updated list of EC2 instance types to use when creating instances in the
	// group. The instance definition must specify instance types that are supported by
	// GameLift FleetIQ, and must include at least two instance types. This updated
	// list replaces the entire current list of instance definitions for the game
	// server group. For more information on instance types, see EC2 Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon EC2 User Guide..
	InstanceDefinitions []*types.InstanceDefinition

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) for an IAM
	// role that allows Amazon GameLift to access your EC2 Auto Scaling groups. The
	// submitted role is validated to ensure that it contains the necessary permissions
	// for game server groups.
	RoleArn *string
}

type UpdateGameServerGroupOutput struct {

	// An object that describes the game server group resource with updated properties.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGameServerGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGameServerGroup{}, middleware.After)
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
	addOpUpdateGameServerGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGameServerGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateGameServerGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateGameServerGroup",
	}
}
