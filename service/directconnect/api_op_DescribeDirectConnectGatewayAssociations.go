// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the associations between your Direct Connect gateways and virtual private
// gateways. You must specify a Direct Connect gateway, a virtual private gateway,
// or both. If you specify a Direct Connect gateway, the response contains all
// virtual private gateways associated with the Direct Connect gateway. If you
// specify a virtual private gateway, the response contains all Direct Connect
// gateways associated with the virtual private gateway. If you specify both, the
// response contains the association between the Direct Connect gateway and the
// virtual private gateway.
func (c *Client) DescribeDirectConnectGatewayAssociations(ctx context.Context, params *DescribeDirectConnectGatewayAssociationsInput, optFns ...func(*Options)) (*DescribeDirectConnectGatewayAssociationsOutput, error) {
	if params == nil {
		params = &DescribeDirectConnectGatewayAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDirectConnectGatewayAssociations", params, optFns, addOperationDescribeDirectConnectGatewayAssociationsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDirectConnectGatewayAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDirectConnectGatewayAssociationsInput struct {

	// The ID of the associated gateway.
	AssociatedGatewayId *string

	// The ID of the Direct Connect gateway association.
	AssociationId *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. If
	// MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int32

	// The token provided in the previous call to retrieve the next page.
	NextToken *string

	// The ID of the virtual private gateway.
	VirtualGatewayId *string
}

type DescribeDirectConnectGatewayAssociationsOutput struct {

	// Information about the associations.
	DirectConnectGatewayAssociations []*types.DirectConnectGatewayAssociation

	// The token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDirectConnectGatewayAssociationsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDirectConnectGatewayAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDirectConnectGatewayAssociations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAssociations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeDirectConnectGatewayAssociations",
	}
}
