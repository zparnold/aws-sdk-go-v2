// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Closes ports for a specific Amazon Lightsail instance. The
// CloseInstancePublicPorts action supports tag-based access control via resource
// tags applied to the resource identified by instanceName. For more information,
// see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CloseInstancePublicPorts(ctx context.Context, params *CloseInstancePublicPortsInput, optFns ...func(*Options)) (*CloseInstancePublicPortsOutput, error) {
	stack := middleware.NewStack("CloseInstancePublicPorts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCloseInstancePublicPortsMiddlewares(stack)
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
	addOpCloseInstancePublicPortsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCloseInstancePublicPorts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CloseInstancePublicPorts",
			Err:           err,
		}
	}
	out := result.(*CloseInstancePublicPortsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CloseInstancePublicPortsInput struct {

	// The name of the instance for which to close ports.
	//
	// This member is required.
	InstanceName *string

	// An object to describe the ports to close for the specified instance.
	//
	// This member is required.
	PortInfo *types.PortInfo
}

type CloseInstancePublicPortsOutput struct {

	// An object that describes the result of the action, such as the status of the
	// request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCloseInstancePublicPortsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCloseInstancePublicPorts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCloseInstancePublicPorts{}, middleware.After)
}

func newServiceMetadataMiddleware_opCloseInstancePublicPorts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CloseInstancePublicPorts",
	}
}
