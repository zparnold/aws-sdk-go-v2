// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a permission with a resource share.
func (c *Client) AssociateResourceSharePermission(ctx context.Context, params *AssociateResourceSharePermissionInput, optFns ...func(*Options)) (*AssociateResourceSharePermissionOutput, error) {
	if params == nil {
		params = &AssociateResourceSharePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateResourceSharePermission", params, optFns, addOperationAssociateResourceSharePermissionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateResourceSharePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResourceSharePermissionInput struct {

	// The ARN of the AWS RAM permission to associate with the resource share.
	//
	// This member is required.
	PermissionArn *string

	// The Amazon Resource Name (ARN) of the resource share.
	//
	// This member is required.
	ResourceShareArn *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// Indicates whether the permission should replace the permissions that are
	// currently associated with the resource share. Use true to replace the current
	// permissions. Use false to add the permission to the current permission.
	Replace *bool
}

type AssociateResourceSharePermissionOutput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// Indicates whether the request succeeded.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateResourceSharePermissionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateResourceSharePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateResourceSharePermission{}, middleware.After)
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
	addOpAssociateResourceSharePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResourceSharePermission(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateResourceSharePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "AssociateResourceSharePermission",
	}
}
