// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes statement for the provided action group from the policy.
func (c *Client) RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) {
	if params == nil {
		params = &RemovePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemovePermission", params, optFns, addOperationRemovePermissionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*RemovePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the removePermissionRequest.
type RemovePermissionInput struct {

	// The list of actions that the users and roles can perform on the profiling group.
	//
	// This member is required.
	ActionGroup types.ActionGroup

	// The name of the profiling group.
	//
	// This member is required.
	ProfilingGroupName *string

	// A unique identifier for the current revision of the policy.
	//
	// This member is required.
	RevisionId *string
}

// The structure representing the removePermissionResponse.
type RemovePermissionOutput struct {

	// The resource-based policy.
	//
	// This member is required.
	Policy *string

	// A unique identifier for the current revision of the policy.
	//
	// This member is required.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemovePermissionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemovePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemovePermission{}, middleware.After)
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
	addOpRemovePermissionValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
