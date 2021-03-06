// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Model for an API.
func (c *Client) CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) {
	stack := middleware.NewStack("CreateModel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateModelMiddlewares(stack)
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
	addOpCreateModelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModel(options.Region), middleware.Before)
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
			OperationName: "CreateModel",
			Err:           err,
		}
	}
	out := result.(*CreateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new Model.
type CreateModelInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The name of the model. Must be alphanumeric.
	//
	// This member is required.
	Name *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 model.
	//
	// This member is required.
	Schema *string

	// The content-type for the model, for example, "application/json".
	ContentType *string

	// The description of the model.
	Description *string
}

type CreateModelOutput struct {

	// The content-type for the model, for example, "application/json".
	ContentType *string

	// The description of the model.
	Description *string

	// The model identifier.
	ModelId *string

	// The name of the model. Must be alphanumeric.
	Name *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 model.
	Schema *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateModelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateModel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateModel{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateModel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateModel",
	}
}
