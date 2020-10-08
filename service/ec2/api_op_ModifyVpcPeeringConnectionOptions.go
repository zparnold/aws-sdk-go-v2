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

// Modifies the VPC peering connection options on one side of a VPC peering
// connection. You can do the following:
//
//     * Enable/disable communication over
// the peering connection between an EC2-Classic instance that's linked to your VPC
// (using ClassicLink) and instances in the peer VPC.
//
//     * Enable/disable
// communication over the peering connection between instances in your VPC and an
// EC2-Classic instance that's linked to the peer VPC.
//
//     * Enable/disable the
// ability to resolve public DNS hostnames to private IP addresses when queried
// from instances in the peer VPC.
//
// If the peered VPCs are in the same AWS account,
// you can enable DNS resolution for queries from the local VPC. This ensures that
// queries from the local VPC resolve to private IP addresses in the peer VPC. This
// option is not available if the peered VPCs are in different AWS accounts or
// different Regions. For peered VPCs in different AWS accounts, each AWS account
// owner must initiate a separate request to modify the peering connection options.
// For inter-region peering connections, you must use the Region for the requester
// VPC to modify the requester VPC peering options and the Region for the accepter
// VPC to modify the accepter VPC peering options. To verify which VPCs are the
// accepter and the requester for a VPC peering connection, use the
// DescribeVpcPeeringConnections () command.
func (c *Client) ModifyVpcPeeringConnectionOptions(ctx context.Context, params *ModifyVpcPeeringConnectionOptionsInput, optFns ...func(*Options)) (*ModifyVpcPeeringConnectionOptionsOutput, error) {
	if params == nil {
		params = &ModifyVpcPeeringConnectionOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcPeeringConnectionOptions", params, optFns, addOperationModifyVpcPeeringConnectionOptionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcPeeringConnectionOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcPeeringConnectionOptionsInput struct {

	// The ID of the VPC peering connection.
	//
	// This member is required.
	VpcPeeringConnectionId *string

	// The VPC peering connection options for the accepter VPC.
	AccepterPeeringConnectionOptions *types.PeeringConnectionOptionsRequest

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The VPC peering connection options for the requester VPC.
	RequesterPeeringConnectionOptions *types.PeeringConnectionOptionsRequest
}

type ModifyVpcPeeringConnectionOptionsOutput struct {

	// Information about the VPC peering connection options for the accepter VPC.
	AccepterPeeringConnectionOptions *types.PeeringConnectionOptions

	// Information about the VPC peering connection options for the requester VPC.
	RequesterPeeringConnectionOptions *types.PeeringConnectionOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyVpcPeeringConnectionOptionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcPeeringConnectionOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcPeeringConnectionOptions{}, middleware.After)
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
	addOpModifyVpcPeeringConnectionOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcPeeringConnectionOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyVpcPeeringConnectionOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcPeeringConnectionOptions",
	}
}
