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

// Creates a thing record in the registry. If this call is made multiple times
// using the same thing name and configuration, the call will succeed. If this call
// is made with the same thing name but different configuration a
// ResourceAlreadyExistsException is thrown. This is a control plane operation. See
// Authorization
// (https://docs.aws.amazon.com/iot/latest/developerguide/iot-authorization.html)
// for information about authorizing control plane actions.
func (c *Client) CreateThing(ctx context.Context, params *CreateThingInput, optFns ...func(*Options)) (*CreateThingOutput, error) {
	stack := middleware.NewStack("CreateThing", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateThingMiddlewares(stack)
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
	addOpCreateThingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThing(options.Region), middleware.Before)
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
			OperationName: "CreateThing",
			Err:           err,
		}
	}
	out := result.(*CreateThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateThing operation.
type CreateThingInput struct {

	// The name of the thing to create. You can't change a thing's name after you
	// create it. To change a thing's name, you must create a new thing, give it the
	// new name, and then delete the old thing.
	//
	// This member is required.
	ThingName *string

	// The attribute payload, which consists of up to three name/value pairs in a JSON
	// document. For example: {\"attributes\":{\"string1\":\"string2\"}}
	AttributePayload *types.AttributePayload

	// The name of the billing group the thing will be added to.
	BillingGroupName *string

	// The name of the thing type associated with the new thing.
	ThingTypeName *string
}

// The output of the CreateThing operation.
type CreateThingOutput struct {

	// The ARN of the new thing.
	ThingArn *string

	// The thing ID.
	ThingId *string

	// The name of the new thing.
	ThingName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateThingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateThing{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThing{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateThing",
	}
}
