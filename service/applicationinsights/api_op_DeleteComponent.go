// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Ungroups a custom component. When you ungroup custom components, all applicable
// monitors that are set up for the component are removed and the instances revert
// to their standalone status.
func (c *Client) DeleteComponent(ctx context.Context, params *DeleteComponentInput, optFns ...func(*Options)) (*DeleteComponentOutput, error) {
	if params == nil {
		params = &DeleteComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteComponent", params, optFns, addOperationDeleteComponentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteComponentInput struct {

	// The name of the component.
	//
	// This member is required.
	ComponentName *string

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string
}

type DeleteComponentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteComponentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteComponent{}, middleware.After)
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
	addOpDeleteComponentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteComponent(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteComponent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "DeleteComponent",
	}
}
