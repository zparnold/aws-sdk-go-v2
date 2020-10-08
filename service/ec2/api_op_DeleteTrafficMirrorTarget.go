// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified Traffic Mirror target. You cannot delete a Traffic Mirror
// target that is in use by a Traffic Mirror session.
func (c *Client) DeleteTrafficMirrorTarget(ctx context.Context, params *DeleteTrafficMirrorTargetInput, optFns ...func(*Options)) (*DeleteTrafficMirrorTargetOutput, error) {
	if params == nil {
		params = &DeleteTrafficMirrorTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTrafficMirrorTarget", params, optFns, addOperationDeleteTrafficMirrorTargetMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTrafficMirrorTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTrafficMirrorTargetInput struct {

	// The ID of the Traffic Mirror target.
	//
	// This member is required.
	TrafficMirrorTargetId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteTrafficMirrorTargetOutput struct {

	// The ID of the deleted Traffic Mirror target.
	TrafficMirrorTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTrafficMirrorTargetMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteTrafficMirrorTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteTrafficMirrorTarget{}, middleware.After)
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
	addOpDeleteTrafficMirrorTargetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTrafficMirrorTarget(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTrafficMirrorTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteTrafficMirrorTarget",
	}
}
