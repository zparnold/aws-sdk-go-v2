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

// Creates an internet gateway for use with a VPC. After creating the internet
// gateway, you attach it to a VPC using AttachInternetGateway (). For more
// information about your VPC and internet gateway, see the Amazon Virtual Private
// Cloud User Guide (https://docs.aws.amazon.com/vpc/latest/userguide/).
func (c *Client) CreateInternetGateway(ctx context.Context, params *CreateInternetGatewayInput, optFns ...func(*Options)) (*CreateInternetGatewayOutput, error) {
	if params == nil {
		params = &CreateInternetGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInternetGateway", params, optFns, addOperationCreateInternetGatewayMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInternetGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInternetGatewayInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags to assign to the internet gateway.
	TagSpecifications []*types.TagSpecification
}

type CreateInternetGatewayOutput struct {

	// Information about the internet gateway.
	InternetGateway *types.InternetGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateInternetGatewayMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateInternetGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateInternetGateway{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInternetGateway(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateInternetGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateInternetGateway",
	}
}
