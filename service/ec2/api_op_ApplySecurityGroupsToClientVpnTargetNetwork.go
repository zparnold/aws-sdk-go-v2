// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Applies a security group to the association between the target network and the
// Client VPN endpoint. This action replaces the existing security groups with the
// specified security groups.
func (c *Client) ApplySecurityGroupsToClientVpnTargetNetwork(ctx context.Context, params *ApplySecurityGroupsToClientVpnTargetNetworkInput, optFns ...func(*Options)) (*ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	if params == nil {
		params = &ApplySecurityGroupsToClientVpnTargetNetworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApplySecurityGroupsToClientVpnTargetNetwork", params, optFns, addOperationApplySecurityGroupsToClientVpnTargetNetworkMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ApplySecurityGroupsToClientVpnTargetNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApplySecurityGroupsToClientVpnTargetNetworkInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The IDs of the security groups to apply to the associated target network. Up to
	// 5 security groups can be applied to an associated target network.
	//
	// This member is required.
	SecurityGroupIds []*string

	// The ID of the VPC in which the associated target network is located.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ApplySecurityGroupsToClientVpnTargetNetworkOutput struct {

	// The IDs of the applied security groups.
	SecurityGroupIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationApplySecurityGroupsToClientVpnTargetNetworkMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpApplySecurityGroupsToClientVpnTargetNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpApplySecurityGroupsToClientVpnTargetNetwork{}, middleware.After)
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
	addOpApplySecurityGroupsToClientVpnTargetNetworkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opApplySecurityGroupsToClientVpnTargetNetwork(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opApplySecurityGroupsToClientVpnTargetNetwork(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ApplySecurityGroupsToClientVpnTargetNetwork",
	}
}
