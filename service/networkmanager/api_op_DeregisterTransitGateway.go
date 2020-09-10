// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters a transit gateway from your global network. This action does not
// delete your transit gateway, or modify any of its attachments. This action
// removes any customer gateway associations.
func (c *Client) DeregisterTransitGateway(ctx context.Context, params *DeregisterTransitGatewayInput, optFns ...func(*Options)) (*DeregisterTransitGatewayOutput, error) {
	stack := middleware.NewStack("DeregisterTransitGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeregisterTransitGatewayMiddlewares(stack)
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
	addOpDeregisterTransitGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterTransitGateway(options.Region), middleware.Before)

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
			OperationName: "DeregisterTransitGateway",
			Err:           err,
		}
	}
	out := result.(*DeregisterTransitGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTransitGatewayInput struct {
	// The ID of the global network.
	GlobalNetworkId *string
	// The Amazon Resource Name (ARN) of the transit gateway.
	TransitGatewayArn *string
}

type DeregisterTransitGatewayOutput struct {
	// The transit gateway registration information.
	TransitGatewayRegistration *types.TransitGatewayRegistration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeregisterTransitGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterTransitGateway{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterTransitGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterTransitGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "DeregisterTransitGateway",
	}
}