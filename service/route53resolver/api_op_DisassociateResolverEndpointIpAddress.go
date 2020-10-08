// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes IP addresses from an inbound or an outbound resolver endpoint. If you
// want to remove more than one IP address, submit one
// DisassociateResolverEndpointIpAddress request for each IP address. To add an IP
// address to an endpoint, see AssociateResolverEndpointIpAddress ().
func (c *Client) DisassociateResolverEndpointIpAddress(ctx context.Context, params *DisassociateResolverEndpointIpAddressInput, optFns ...func(*Options)) (*DisassociateResolverEndpointIpAddressOutput, error) {
	if params == nil {
		params = &DisassociateResolverEndpointIpAddressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateResolverEndpointIpAddress", params, optFns, addOperationDisassociateResolverEndpointIpAddressMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateResolverEndpointIpAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateResolverEndpointIpAddressInput struct {

	// The IPv4 address that you want to remove from a resolver endpoint.
	//
	// This member is required.
	IpAddress *types.IpAddressUpdate

	// The ID of the resolver endpoint that you want to disassociate an IP address
	// from.
	//
	// This member is required.
	ResolverEndpointId *string
}

type DisassociateResolverEndpointIpAddressOutput struct {

	// The response to an DisassociateResolverEndpointIpAddress request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateResolverEndpointIpAddressMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateResolverEndpointIpAddress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateResolverEndpointIpAddress{}, middleware.After)
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
	addOpDisassociateResolverEndpointIpAddressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateResolverEndpointIpAddress(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateResolverEndpointIpAddress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "DisassociateResolverEndpointIpAddress",
	}
}
