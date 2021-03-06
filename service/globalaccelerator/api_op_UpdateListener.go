// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update a listener. To see an AWS CLI example of updating listener, scroll down
// to Example.
func (c *Client) UpdateListener(ctx context.Context, params *UpdateListenerInput, optFns ...func(*Options)) (*UpdateListenerOutput, error) {
	stack := middleware.NewStack("UpdateListener", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateListenerMiddlewares(stack)
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
	addOpUpdateListenerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateListener(options.Region), middleware.Before)
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
			OperationName: "UpdateListener",
			Err:           err,
		}
	}
	out := result.(*UpdateListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateListenerInput struct {

	// The Amazon Resource Name (ARN) of the listener to update.
	//
	// This member is required.
	ListenerArn *string

	// Client affinity lets you direct all requests from a user to the same endpoint,
	// if you have stateful applications, regardless of the port and protocol of the
	// client request. Clienty affinity gives you control over whether to always route
	// each client to the same specific endpoint. AWS Global Accelerator uses a
	// consistent-flow hashing algorithm to choose the optimal endpoint for a
	// connection. If client affinity is NONE, Global Accelerator uses the "five-tuple"
	// (5-tuple) properties—source IP address, source port, destination IP address,
	// destination port, and protocol—to select the hash value, and then chooses the
	// best endpoint. However, with this setting, if someone uses different ports to
	// connect to Global Accelerator, their connections might not be always routed to
	// the same endpoint because the hash value changes. If you want a given client to
	// always be routed to the same endpoint, set client affinity to SOURCE_IP instead.
	// When you use the SOURCE_IP setting, Global Accelerator uses the "two-tuple"
	// (2-tuple) properties— source (client) IP address and destination IP address—to
	// select the hash value. The default value is NONE.
	ClientAffinity types.ClientAffinity

	// The updated list of port ranges for the connections from clients to the
	// accelerator.
	PortRanges []*types.PortRange

	// The updated protocol for the connections from clients to the accelerator.
	Protocol types.Protocol
}

type UpdateListenerOutput struct {

	// Information for the updated listener.
	Listener *types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateListenerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateListener{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateListener{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateListener(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "UpdateListener",
	}
}
