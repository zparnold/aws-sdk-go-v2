// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified NAT gateway. Deleting a NAT gateway disassociates its
// Elastic IP address, but does not release the address from your account. Deleting
// a NAT gateway does not delete any NAT gateway routes in your route tables.
func (c *Client) DeleteNatGateway(ctx context.Context, params *DeleteNatGatewayInput, optFns ...func(*Options)) (*DeleteNatGatewayOutput, error) {
	if params == nil {
		params = &DeleteNatGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteNatGateway", params, optFns, addOperationDeleteNatGatewayMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteNatGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteNatGatewayInput struct {

	// The ID of the NAT gateway.
	//
	// This member is required.
	NatGatewayId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteNatGatewayOutput struct {

	// The ID of the NAT gateway.
	NatGatewayId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteNatGatewayMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteNatGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteNatGateway{}, middleware.After)
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
	addOpDeleteNatGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteNatGateway(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteNatGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteNatGateway",
	}
}
