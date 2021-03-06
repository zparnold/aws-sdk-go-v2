// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Terminates the specified provisioned product. This operation does not delete any
// records associated with the provisioned product. You can check the status of
// this request using DescribeRecord ().
func (c *Client) TerminateProvisionedProduct(ctx context.Context, params *TerminateProvisionedProductInput, optFns ...func(*Options)) (*TerminateProvisionedProductOutput, error) {
	stack := middleware.NewStack("TerminateProvisionedProduct", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTerminateProvisionedProductMiddlewares(stack)
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
	addIdempotencyToken_opTerminateProvisionedProductMiddleware(stack, options)
	addOpTerminateProvisionedProductValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateProvisionedProduct(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "TerminateProvisionedProduct",
			Err:           err,
		}
	}
	out := result.(*TerminateProvisionedProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateProvisionedProductInput struct {

	// An idempotency token that uniquely identifies the termination request. This
	// token is only valid during the termination process. After the provisioned
	// product is terminated, subsequent requests to terminate the same provisioned
	// product always return ResourceNotFound.
	//
	// This member is required.
	TerminateToken *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string

	// If set to true, AWS Service Catalog stops managing the specified provisioned
	// product even if it cannot delete the underlying resources.
	IgnoreErrors *bool

	// The identifier of the provisioned product. You cannot specify both
	// ProvisionedProductName and ProvisionedProductId.
	ProvisionedProductId *string

	// The name of the provisioned product. You cannot specify both
	// ProvisionedProductName and ProvisionedProductId.
	ProvisionedProductName *string
}

type TerminateProvisionedProductOutput struct {

	// Information about the result of this request.
	RecordDetail *types.RecordDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTerminateProvisionedProductMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateProvisionedProduct{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateProvisionedProduct{}, middleware.After)
}

type idempotencyToken_initializeOpTerminateProvisionedProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpTerminateProvisionedProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpTerminateProvisionedProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*TerminateProvisionedProductInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *TerminateProvisionedProductInput ")
	}

	if input.TerminateToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.TerminateToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opTerminateProvisionedProductMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpTerminateProvisionedProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opTerminateProvisionedProduct(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "TerminateProvisionedProduct",
	}
}
