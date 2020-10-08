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

// Deletes the specified transit gateway route table. You must disassociate the
// route table from any transit gateway route tables before you can delete it.
func (c *Client) DeleteTransitGatewayRouteTable(ctx context.Context, params *DeleteTransitGatewayRouteTableInput, optFns ...func(*Options)) (*DeleteTransitGatewayRouteTableOutput, error) {
	if params == nil {
		params = &DeleteTransitGatewayRouteTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTransitGatewayRouteTable", params, optFns, addOperationDeleteTransitGatewayRouteTableMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTransitGatewayRouteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTransitGatewayRouteTableInput struct {

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteTransitGatewayRouteTableOutput struct {

	// Information about the deleted transit gateway route table.
	TransitGatewayRouteTable *types.TransitGatewayRouteTable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTransitGatewayRouteTableMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteTransitGatewayRouteTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteTransitGatewayRouteTable{}, middleware.After)
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
	addOpDeleteTransitGatewayRouteTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTransitGatewayRouteTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTransitGatewayRouteTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteTransitGatewayRouteTable",
	}
}
