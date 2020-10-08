// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the status of your service-linked role deletion. After you use the
// DeleteServiceLinkedRole () API operation to submit a service-linked role for
// deletion, you can use the DeletionTaskId parameter in
// GetServiceLinkedRoleDeletionStatus to check the status of the deletion. If the
// deletion fails, this operation returns the reason that it failed, if that
// information is returned by the service.
func (c *Client) GetServiceLinkedRoleDeletionStatus(ctx context.Context, params *GetServiceLinkedRoleDeletionStatusInput, optFns ...func(*Options)) (*GetServiceLinkedRoleDeletionStatusOutput, error) {
	if params == nil {
		params = &GetServiceLinkedRoleDeletionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceLinkedRoleDeletionStatus", params, optFns, addOperationGetServiceLinkedRoleDeletionStatusMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceLinkedRoleDeletionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceLinkedRoleDeletionStatusInput struct {

	// The deletion task identifier. This identifier is returned by the
	// DeleteServiceLinkedRole () operation in the format task/aws-service-role///.
	//
	// This member is required.
	DeletionTaskId *string
}

type GetServiceLinkedRoleDeletionStatusOutput struct {

	// The status of the deletion.
	//
	// This member is required.
	Status types.DeletionTaskStatusType

	// An object that contains details about the reason the deletion failed.
	Reason *types.DeletionTaskFailureReasonType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetServiceLinkedRoleDeletionStatusMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetServiceLinkedRoleDeletionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetServiceLinkedRoleDeletionStatus{}, middleware.After)
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
	addOpGetServiceLinkedRoleDeletionStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceLinkedRoleDeletionStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetServiceLinkedRoleDeletionStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetServiceLinkedRoleDeletionStatus",
	}
}
