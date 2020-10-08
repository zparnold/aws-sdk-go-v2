// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new image. This request will create a new image along with all of the
// configured output resources defined in the distribution configuration.
func (c *Client) CreateImage(ctx context.Context, params *CreateImageInput, optFns ...func(*Options)) (*CreateImageOutput, error) {
	if params == nil {
		params = &CreateImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateImage", params, optFns, addOperationCreateImageMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateImageInput struct {

	// The idempotency token used to make this request idempotent.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the image recipe that defines how images are
	// configured, tested, and assessed.
	//
	// This member is required.
	ImageRecipeArn *string

	// The Amazon Resource Name (ARN) of the infrastructure configuration that defines
	// the environment in which your image will be built and tested.
	//
	// This member is required.
	InfrastructureConfigurationArn *string

	// The Amazon Resource Name (ARN) of the distribution configuration that defines
	// and configures the outputs of your pipeline.
	DistributionConfigurationArn *string

	// Collects additional information about the image being created, including the
	// operating system (OS) version and package list. This information is used to
	// enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled *bool

	// The image tests configuration of the image.
	ImageTestsConfiguration *types.ImageTestsConfiguration

	// The tags of the image.
	Tags map[string]*string
}

type CreateImageOutput struct {

	// The idempotency token used to make this request idempotent.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the image that was created by this request.
	ImageBuildVersionArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateImageMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateImage{}, middleware.After)
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
	addIdempotencyToken_opCreateImageMiddleware(stack, options)
	addOpCreateImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateImage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateImage struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateImage) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateImage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateImageInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateImageInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateImageMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateImage{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "CreateImage",
	}
}
