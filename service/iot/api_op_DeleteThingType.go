// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified thing type. You cannot delete a thing type if it has
// things associated with it. To delete a thing type, first mark it as deprecated
// by calling DeprecateThingType (), then remove any associated things by calling
// UpdateThing () to change the thing type on any associated thing, and finally use
// DeleteThingType () to delete the thing type.
func (c *Client) DeleteThingType(ctx context.Context, params *DeleteThingTypeInput, optFns ...func(*Options)) (*DeleteThingTypeOutput, error) {
	stack := middleware.NewStack("DeleteThingType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteThingTypeMiddlewares(stack)
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
	addOpDeleteThingTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteThingType(options.Region), middleware.Before)
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
			OperationName: "DeleteThingType",
			Err:           err,
		}
	}
	out := result.(*DeleteThingTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DeleteThingType operation.
type DeleteThingTypeInput struct {

	// The name of the thing type.
	//
	// This member is required.
	ThingTypeName *string
}

// The output for the DeleteThingType operation.
type DeleteThingTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteThingTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteThingType{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteThingType{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteThingType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteThingType",
	}
}
