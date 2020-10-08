// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the resource query of a group.
func (c *Client) UpdateGroupQuery(ctx context.Context, params *UpdateGroupQueryInput, optFns ...func(*Options)) (*UpdateGroupQueryOutput, error) {
	if params == nil {
		params = &UpdateGroupQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGroupQuery", params, optFns, addOperationUpdateGroupQueryMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGroupQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGroupQueryInput struct {

	// The resource query to determine which AWS resources are members of this resource
	// group.
	//
	// This member is required.
	ResourceQuery *types.ResourceQuery

	// The name or the ARN of the resource group to query.
	Group *string

	// Don't use this parameter. Use Group instead.
	GroupName *string
}

type UpdateGroupQueryOutput struct {

	// The updated resource query associated with the resource group after the update.
	GroupQuery *types.GroupQuery

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGroupQueryMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGroupQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGroupQuery{}, middleware.After)
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
	addOpUpdateGroupQueryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGroupQuery(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateGroupQuery(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "UpdateGroupQuery",
	}
}
