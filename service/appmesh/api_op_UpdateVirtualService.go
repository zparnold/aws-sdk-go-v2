// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing virtual service in a specified service mesh.
func (c *Client) UpdateVirtualService(ctx context.Context, params *UpdateVirtualServiceInput, optFns ...func(*Options)) (*UpdateVirtualServiceOutput, error) {
	if params == nil {
		params = &UpdateVirtualServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVirtualService", params, optFns, addOperationUpdateVirtualServiceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVirtualServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type UpdateVirtualServiceInput struct {

	// The name of the service mesh that the virtual service resides in.
	//
	// This member is required.
	MeshName *string

	// The new virtual service specification to apply. This overwrites the existing
	// data.
	//
	// This member is required.
	Spec *types.VirtualServiceSpec

	// The name of the virtual service to update.
	//
	// This member is required.
	VirtualServiceName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type UpdateVirtualServiceOutput struct {

	// A full description of the virtual service that was updated.
	//
	// This member is required.
	VirtualService *types.VirtualServiceData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateVirtualServiceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVirtualService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVirtualService{}, middleware.After)
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
	addOpUpdateVirtualServiceValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
