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

// Accept a VPC peering connection request. To accept a request, the VPC peering
// connection must be in the pending-acceptance state, and you must be the owner of
// the peer VPC. Use DescribeVpcPeeringConnections () to view your outstanding VPC
// peering connection requests. For an inter-Region VPC peering connection request,
// you must accept the VPC peering connection in the Region of the accepter VPC.
func (c *Client) AcceptVpcPeeringConnection(ctx context.Context, params *AcceptVpcPeeringConnectionInput, optFns ...func(*Options)) (*AcceptVpcPeeringConnectionOutput, error) {
	if params == nil {
		params = &AcceptVpcPeeringConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptVpcPeeringConnection", params, optFns, addOperationAcceptVpcPeeringConnectionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptVpcPeeringConnectionInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the VPC peering connection. You must specify this parameter in the
	// request.
	VpcPeeringConnectionId *string
}

type AcceptVpcPeeringConnectionOutput struct {

	// Information about the VPC peering connection.
	VpcPeeringConnection *types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAcceptVpcPeeringConnectionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAcceptVpcPeeringConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAcceptVpcPeeringConnection{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptVpcPeeringConnection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAcceptVpcPeeringConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AcceptVpcPeeringConnection",
	}
}
