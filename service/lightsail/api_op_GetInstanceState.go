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

// Returns the state of a specific instance. Works on one instance at a time.
func (c *Client) GetInstanceState(ctx context.Context, params *GetInstanceStateInput, optFns ...func(*Options)) (*GetInstanceStateOutput, error) {
	if params == nil {
		params = &GetInstanceStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstanceState", params, optFns, addOperationGetInstanceStateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstanceStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstanceStateInput struct {

	// The name of the instance to get state information about.
	//
	// This member is required.
	InstanceName *string
}

type GetInstanceStateOutput struct {

	// The state of the instance.
	State *types.InstanceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInstanceStateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInstanceState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInstanceState{}, middleware.After)
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
	addOpGetInstanceStateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceState(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetInstanceState(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetInstanceState",
	}
}
