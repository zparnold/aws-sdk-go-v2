// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the association information for customer gateways that are associated with
// devices and links in your global network.
func (c *Client) GetCustomerGatewayAssociations(ctx context.Context, params *GetCustomerGatewayAssociationsInput, optFns ...func(*Options)) (*GetCustomerGatewayAssociationsOutput, error) {
	if params == nil {
		params = &GetCustomerGatewayAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCustomerGatewayAssociations", params, optFns, addOperationGetCustomerGatewayAssociationsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCustomerGatewayAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCustomerGatewayAssociationsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// One or more customer gateway Amazon Resource Names (ARNs). For more information,
	// see Resources Defined by Amazon EC2
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonec2.html#amazonec2-resources-for-iam-policies).
	// The maximum is 10.
	CustomerGatewayArns []*string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type GetCustomerGatewayAssociationsOutput struct {

	// The customer gateway associations.
	CustomerGatewayAssociations []*types.CustomerGatewayAssociation

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCustomerGatewayAssociationsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCustomerGatewayAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCustomerGatewayAssociations{}, middleware.After)
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
	addOpGetCustomerGatewayAssociationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCustomerGatewayAssociations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCustomerGatewayAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetCustomerGatewayAssociations",
	}
}
