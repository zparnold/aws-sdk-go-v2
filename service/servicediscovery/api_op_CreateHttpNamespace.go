// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an HTTP namespace. Service instances that you register using an HTTP
// namespace can be discovered using a DiscoverInstances request but can't be
// discovered using DNS. For the current limit on the number of namespaces that you
// can create using the same AWS account, see AWS Cloud Map Limits
// (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html) in the
// AWS Cloud Map Developer Guide.
func (c *Client) CreateHttpNamespace(ctx context.Context, params *CreateHttpNamespaceInput, optFns ...func(*Options)) (*CreateHttpNamespaceOutput, error) {
	if params == nil {
		params = &CreateHttpNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHttpNamespace", params, optFns, addOperationCreateHttpNamespaceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHttpNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHttpNamespaceInput struct {

	// The name that you want to assign to this namespace.
	//
	// This member is required.
	Name *string

	// A unique string that identifies the request and that allows failed
	// CreateHttpNamespace requests to be retried without the risk of executing the
	// operation twice. CreatorRequestId can be any unique string, for example, a
	// date/time stamp.
	CreatorRequestId *string

	// A description for the namespace.
	Description *string

	// The tags to add to the namespace. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []*types.Tag
}

type CreateHttpNamespaceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see GetOperation
	// (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHttpNamespaceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHttpNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHttpNamespace{}, middleware.After)
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
	addIdempotencyToken_opCreateHttpNamespaceMiddleware(stack, options)
	addOpCreateHttpNamespaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHttpNamespace(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateHttpNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateHttpNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateHttpNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateHttpNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateHttpNamespaceInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateHttpNamespaceMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateHttpNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateHttpNamespace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "CreateHttpNamespace",
	}
}
