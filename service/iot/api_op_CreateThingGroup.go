// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a thing group. This is a control plane operation. See Authorization
// (https://docs.aws.amazon.com/iot/latest/developerguide/iot-authorization.html)
// for information about authorizing control plane actions.
func (c *Client) CreateThingGroup(ctx context.Context, params *CreateThingGroupInput, optFns ...func(*Options)) (*CreateThingGroupOutput, error) {
	if params == nil {
		params = &CreateThingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateThingGroup", params, optFns, addOperationCreateThingGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateThingGroupInput struct {

	// The thing group name to create.
	//
	// This member is required.
	ThingGroupName *string

	// The name of the parent thing group.
	ParentGroupName *string

	// Metadata which can be used to manage the thing group.
	Tags []*types.Tag

	// The thing group properties.
	ThingGroupProperties *types.ThingGroupProperties
}

type CreateThingGroupOutput struct {

	// The thing group ARN.
	ThingGroupArn *string

	// The thing group ID.
	ThingGroupId *string

	// The thing group name.
	ThingGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateThingGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateThingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThingGroup{}, middleware.After)
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
	addOpCreateThingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThingGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateThingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateThingGroup",
	}
}
