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

// Deletes the specified association between a VPC and local gateway route table.
func (c *Client) DeleteLocalGatewayRouteTableVpcAssociation(ctx context.Context, params *DeleteLocalGatewayRouteTableVpcAssociationInput, optFns ...func(*Options)) (*DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
	if params == nil {
		params = &DeleteLocalGatewayRouteTableVpcAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLocalGatewayRouteTableVpcAssociation", params, optFns, addOperationDeleteLocalGatewayRouteTableVpcAssociationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLocalGatewayRouteTableVpcAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLocalGatewayRouteTableVpcAssociationInput struct {

	// The ID of the association.
	//
	// This member is required.
	LocalGatewayRouteTableVpcAssociationId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteLocalGatewayRouteTableVpcAssociationOutput struct {

	// Information about the association.
	LocalGatewayRouteTableVpcAssociation *types.LocalGatewayRouteTableVpcAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteLocalGatewayRouteTableVpcAssociationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteLocalGatewayRouteTableVpcAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteLocalGatewayRouteTableVpcAssociation{}, middleware.After)
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
	addOpDeleteLocalGatewayRouteTableVpcAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLocalGatewayRouteTableVpcAssociation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteLocalGatewayRouteTableVpcAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteLocalGatewayRouteTableVpcAssociation",
	}
}
