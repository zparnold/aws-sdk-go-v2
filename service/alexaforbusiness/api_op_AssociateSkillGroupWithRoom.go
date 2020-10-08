// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a skill group with a given room. This enables all skills in the
// associated skill group on all devices in the room.
func (c *Client) AssociateSkillGroupWithRoom(ctx context.Context, params *AssociateSkillGroupWithRoomInput, optFns ...func(*Options)) (*AssociateSkillGroupWithRoomOutput, error) {
	if params == nil {
		params = &AssociateSkillGroupWithRoomInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateSkillGroupWithRoom", params, optFns, addOperationAssociateSkillGroupWithRoomMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateSkillGroupWithRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSkillGroupWithRoomInput struct {

	// The ARN of the room with which to associate the skill group. Required.
	RoomArn *string

	// The ARN of the skill group to associate with a room. Required.
	SkillGroupArn *string
}

type AssociateSkillGroupWithRoomOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateSkillGroupWithRoomMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateSkillGroupWithRoom{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateSkillGroupWithRoom{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSkillGroupWithRoom(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateSkillGroupWithRoom(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "AssociateSkillGroupWithRoom",
	}
}
