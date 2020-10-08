// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns temporary SSH keys you can use to connect to a specific virtual private
// server, or instance. The get instance access details operation supports
// tag-based access control via resource tags applied to the resource identified by
// instance name. For more information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) GetInstanceAccessDetails(ctx context.Context, params *GetInstanceAccessDetailsInput, optFns ...func(*Options)) (*GetInstanceAccessDetailsOutput, error) {
	if params == nil {
		params = &GetInstanceAccessDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstanceAccessDetails", params, optFns, addOperationGetInstanceAccessDetailsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstanceAccessDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstanceAccessDetailsInput struct {

	// The name of the instance to access.
	//
	// This member is required.
	InstanceName *string

	// The protocol to use to connect to your instance. Defaults to ssh.
	Protocol types.InstanceAccessProtocol
}

type GetInstanceAccessDetailsOutput struct {

	// An array of key-value pairs containing information about a get instance access
	// request.
	AccessDetails *types.InstanceAccessDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInstanceAccessDetailsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInstanceAccessDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInstanceAccessDetails{}, middleware.After)
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
	addOpGetInstanceAccessDetailsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceAccessDetails(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetInstanceAccessDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetInstanceAccessDetails",
	}
}
