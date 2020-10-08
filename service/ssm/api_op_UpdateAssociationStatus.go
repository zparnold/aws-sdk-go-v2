// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the status of the Systems Manager document associated with the specified
// instance.
func (c *Client) UpdateAssociationStatus(ctx context.Context, params *UpdateAssociationStatusInput, optFns ...func(*Options)) (*UpdateAssociationStatusOutput, error) {
	if params == nil {
		params = &UpdateAssociationStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssociationStatus", params, optFns, addOperationUpdateAssociationStatusMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssociationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssociationStatusInput struct {

	// The association status.
	//
	// This member is required.
	AssociationStatus *types.AssociationStatus

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// The name of the Systems Manager document.
	//
	// This member is required.
	Name *string
}

type UpdateAssociationStatusOutput struct {

	// Information about the association.
	AssociationDescription *types.AssociationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAssociationStatusMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAssociationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAssociationStatus{}, middleware.After)
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
	addOpUpdateAssociationStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssociationStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateAssociationStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateAssociationStatus",
	}
}
