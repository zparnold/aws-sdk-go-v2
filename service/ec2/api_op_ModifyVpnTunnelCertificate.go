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

// Modifies the VPN tunnel endpoint certificate.
func (c *Client) ModifyVpnTunnelCertificate(ctx context.Context, params *ModifyVpnTunnelCertificateInput, optFns ...func(*Options)) (*ModifyVpnTunnelCertificateOutput, error) {
	if params == nil {
		params = &ModifyVpnTunnelCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpnTunnelCertificate", params, optFns, addOperationModifyVpnTunnelCertificateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpnTunnelCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpnTunnelCertificateInput struct {

	// The ID of the AWS Site-to-Site VPN connection.
	//
	// This member is required.
	VpnConnectionId *string

	// The external IP address of the VPN tunnel.
	//
	// This member is required.
	VpnTunnelOutsideIpAddress *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ModifyVpnTunnelCertificateOutput struct {

	// Describes a VPN connection.
	VpnConnection *types.VpnConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyVpnTunnelCertificateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpnTunnelCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpnTunnelCertificate{}, middleware.After)
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
	addOpModifyVpnTunnelCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpnTunnelCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyVpnTunnelCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpnTunnelCertificate",
	}
}
