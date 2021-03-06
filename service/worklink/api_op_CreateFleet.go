// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a fleet. A fleet consists of resources and the configuration that
// delivers associated websites to authorized users who download and set up the
// Amazon WorkLink app.
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	stack := middleware.NewStack("CreateFleet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateFleetMiddlewares(stack)
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
	addOpCreateFleetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before)
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
			OperationName: "CreateFleet",
			Err:           err,
		}
	}
	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetInput struct {

	// A unique name for the fleet.
	//
	// This member is required.
	FleetName *string

	// The fleet name to display.
	DisplayName *string

	// The option to optimize for better performance by routing traffic through the
	// closest AWS Region to users, which may be outside of your home Region.
	OptimizeForEndUserLocation *bool

	// The tags to add to the resource. A tag is a key-value pair.
	Tags map[string]*string
}

type CreateFleetOutput struct {

	// The Amazon Resource Name (ARN) of the fleet.
	FleetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateFleetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateFleet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFleet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateFleet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "CreateFleet",
	}
}
