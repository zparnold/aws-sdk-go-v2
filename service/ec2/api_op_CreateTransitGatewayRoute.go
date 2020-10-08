// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a static route for the specified transit gateway route table.
func (c *Client) CreateTransitGatewayRoute(ctx context.Context, params *CreateTransitGatewayRouteInput, optFns ...func(*Options)) (*CreateTransitGatewayRouteOutput, error) {
	if params == nil {
		params = &CreateTransitGatewayRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransitGatewayRoute", params, optFns, addOperationCreateTransitGatewayRouteMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransitGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitGatewayRouteInput struct {

	// The CIDR range used for destination matches. Routing decisions are based on the
	// most specific match.
	//
	// This member is required.
	DestinationCidrBlock *string

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Indicates whether to drop traffic that matches this route.
	Blackhole *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the attachment.
	TransitGatewayAttachmentId *string
}

type CreateTransitGatewayRouteOutput struct {

	// Information about the route.
	Route *types.TransitGatewayRoute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTransitGatewayRouteMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTransitGatewayRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTransitGatewayRoute{}, middleware.After)
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
	addOpCreateTransitGatewayRouteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitGatewayRoute(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateTransitGatewayRoute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTransitGatewayRoute",
	}
}
