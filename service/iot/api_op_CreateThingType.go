// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new thing type.
func (c *Client) CreateThingType(ctx context.Context, params *CreateThingTypeInput, optFns ...func(*Options)) (*CreateThingTypeOutput, error) {
	stack := middleware.NewStack("CreateThingType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateThingTypeMiddlewares(stack)
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
	addOpCreateThingTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThingType(options.Region), middleware.Before)
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
			OperationName: "CreateThingType",
			Err:           err,
		}
	}
	out := result.(*CreateThingTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateThingType operation.
type CreateThingTypeInput struct {

	// The name of the thing type.
	//
	// This member is required.
	ThingTypeName *string

	// Metadata which can be used to manage the thing type.
	Tags []*types.Tag

	// The ThingTypeProperties for the thing type to create. It contains information
	// about the new thing type including a description, and a list of searchable thing
	// attribute names.
	ThingTypeProperties *types.ThingTypeProperties
}

// The output of the CreateThingType operation.
type CreateThingTypeOutput struct {

	// The Amazon Resource Name (ARN) of the thing type.
	ThingTypeArn *string

	// The thing type ID.
	ThingTypeId *string

	// The name of the thing type.
	ThingTypeName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateThingTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateThingType{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThingType{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateThingType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateThingType",
	}
}
