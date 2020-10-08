// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a resource that represents the infrastructure where a third-party
// provider is installed. The host is used when you create connections to an
// installed third-party provider type, such as GitHub Enterprise Server. You
// create one host for all connections to that provider. A host created through the
// CLI or the SDK is in `PENDING` status by default. You can make its status
// `AVAILABLE` by setting up the host in the console.
func (c *Client) CreateHost(ctx context.Context, params *CreateHostInput, optFns ...func(*Options)) (*CreateHostOutput, error) {
	if params == nil {
		params = &CreateHostInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHost", params, optFns, addOperationCreateHostMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHostOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHostInput struct {

	// The name of the host to be created. The name must be unique in the calling AWS
	// account.
	//
	// This member is required.
	Name *string

	// The endpoint of the infrastructure to be represented by the host after it is
	// created.
	//
	// This member is required.
	ProviderEndpoint *string

	// The name of the installed provider to be associated with your connection. The
	// host resource represents the infrastructure where your provider type is
	// installed. The valid provider type is GitHub Enterprise Server.
	//
	// This member is required.
	ProviderType types.ProviderType

	// The VPC configuration to be provisioned for the host. A VPC must be configured
	// and the infrastructure to be represented by the host must already be connected
	// to the VPC.
	VpcConfiguration *types.VpcConfiguration
}

type CreateHostOutput struct {

	// The Amazon Resource Name (ARN) of the host to be created.
	HostArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHostMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateHost{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateHost{}, middleware.After)
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
	addOpCreateHostValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHost(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateHost(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-connections",
		OperationName: "CreateHost",
	}
}
