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

// Returns the firewall port states for a specific Amazon Lightsail instance, the
// IP addresses allowed to connect to the instance through the ports, and the
// protocol.
func (c *Client) GetInstancePortStates(ctx context.Context, params *GetInstancePortStatesInput, optFns ...func(*Options)) (*GetInstancePortStatesOutput, error) {
	if params == nil {
		params = &GetInstancePortStatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstancePortStates", params, optFns, addOperationGetInstancePortStatesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstancePortStatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstancePortStatesInput struct {

	// The name of the instance for which to return firewall port states.
	//
	// This member is required.
	InstanceName *string
}

type GetInstancePortStatesOutput struct {

	// An array of objects that describe the firewall port states for the specified
	// instance.
	PortStates []*types.InstancePortState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInstancePortStatesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInstancePortStates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInstancePortStates{}, middleware.After)
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
	addOpGetInstancePortStatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstancePortStates(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetInstancePortStates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetInstancePortStates",
	}
}
