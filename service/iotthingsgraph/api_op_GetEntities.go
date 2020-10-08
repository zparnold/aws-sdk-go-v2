// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets definitions of the specified entities. Uses the latest version of the
// user's namespace by default. This API returns the following TDM entities.
//
//     *
// Properties
//
//     * States
//
//     * Events
//
//     * Actions
//
//     * Capabilities
//
//     *
// Mappings
//
//     * Devices
//
//     * Device Models
//
//     * Services
//
// This action
// doesn't return definitions for systems, flows, and deployments.
func (c *Client) GetEntities(ctx context.Context, params *GetEntitiesInput, optFns ...func(*Options)) (*GetEntitiesOutput, error) {
	if params == nil {
		params = &GetEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEntities", params, optFns, addOperationGetEntitiesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEntitiesInput struct {

	// An array of entity IDs. The IDs should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:device:DEVICENAME
	//
	// This member is required.
	Ids []*string

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64
}

type GetEntitiesOutput struct {

	// An array of descriptions for the specified entities.
	Descriptions []*types.EntityDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEntitiesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEntities{}, middleware.After)
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
	addOpGetEntitiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEntities(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetEntities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetEntities",
	}
}
