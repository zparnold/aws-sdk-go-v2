// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a dimension that you can use to limit the scope of a metric used in a
// security profile for AWS IoT Device Defender. For example, using a TOPIC_FILTER
// dimension, you can narrow down the scope of the metric only to MQTT topics whose
// name match the pattern specified in the dimension.
func (c *Client) CreateDimension(ctx context.Context, params *CreateDimensionInput, optFns ...func(*Options)) (*CreateDimensionOutput, error) {
	if params == nil {
		params = &CreateDimensionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDimension", params, optFns, addOperationCreateDimensionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDimensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDimensionInput struct {

	// Each dimension must have a unique client request token. If you try to create a
	// new dimension with the same token as a dimension that already exists, an
	// exception occurs. If you omit this value, AWS SDKs will automatically generate a
	// unique client request.
	//
	// This member is required.
	ClientRequestToken *string

	// A unique identifier for the dimension. Choose something that describes the type
	// and value to make it easy to remember what it does.
	//
	// This member is required.
	Name *string

	// Specifies the value or list of values for the dimension. For TOPIC_FILTER
	// dimensions, this is a pattern used to match the MQTT topic (for example,
	// "admin/#").
	//
	// This member is required.
	StringValues []*string

	// Specifies the type of dimension. Supported types: TOPIC_FILTER.
	//
	// This member is required.
	Type types.DimensionType

	// Metadata that can be used to manage the dimension.
	Tags []*types.Tag
}

type CreateDimensionOutput struct {

	// The ARN (Amazon resource name) of the created dimension.
	Arn *string

	// A unique identifier for the dimension.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDimensionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDimension{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDimension{}, middleware.After)
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
	addIdempotencyToken_opCreateDimensionMiddleware(stack, options)
	addOpCreateDimensionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDimension(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateDimension struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDimension) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDimension) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDimensionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDimensionInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDimensionMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateDimension{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDimension(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateDimension",
	}
}
