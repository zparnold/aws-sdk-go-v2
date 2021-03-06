// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a transit gateway peering attachment.
func (c *Client) DeleteTransitGatewayPeeringAttachment(ctx context.Context, params *DeleteTransitGatewayPeeringAttachmentInput, optFns ...func(*Options)) (*DeleteTransitGatewayPeeringAttachmentOutput, error) {
	stack := middleware.NewStack("DeleteTransitGatewayPeeringAttachment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeleteTransitGatewayPeeringAttachmentMiddlewares(stack)
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
	addOpDeleteTransitGatewayPeeringAttachmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTransitGatewayPeeringAttachment(options.Region), middleware.Before)
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
			OperationName: "DeleteTransitGatewayPeeringAttachment",
			Err:           err,
		}
	}
	out := result.(*DeleteTransitGatewayPeeringAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTransitGatewayPeeringAttachmentInput struct {

	// The ID of the transit gateway peering attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteTransitGatewayPeeringAttachmentOutput struct {

	// The transit gateway peering attachment.
	TransitGatewayPeeringAttachment *types.TransitGatewayPeeringAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeleteTransitGatewayPeeringAttachmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeleteTransitGatewayPeeringAttachment{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteTransitGatewayPeeringAttachment{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTransitGatewayPeeringAttachment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteTransitGatewayPeeringAttachment",
	}
}
