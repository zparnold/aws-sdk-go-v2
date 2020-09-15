// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	smithy "github.com/awslabs/smithy-go"
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
	stack := middleware.NewStack("GetEntities", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetEntitiesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetEntitiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEntities(options.Region), middleware.Before)

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
			OperationName: "GetEntities",
			Err:           err,
		}
	}
	out := result.(*GetEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEntitiesInput struct {
	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64
	// An array of entity IDs. The IDs should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:device:DEVICENAME
	Ids []*string
}

type GetEntitiesOutput struct {
	// An array of descriptions for the specified entities.
	Descriptions []*types.EntityDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetEntitiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetEntities{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEntities{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetEntities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetEntities",
	}
}