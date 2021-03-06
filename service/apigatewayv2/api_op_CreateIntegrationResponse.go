// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an IntegrationResponses.
func (c *Client) CreateIntegrationResponse(ctx context.Context, params *CreateIntegrationResponseInput, optFns ...func(*Options)) (*CreateIntegrationResponseOutput, error) {
	stack := middleware.NewStack("CreateIntegrationResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateIntegrationResponseMiddlewares(stack)
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
	addOpCreateIntegrationResponseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIntegrationResponse(options.Region), middleware.Before)
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
			OperationName: "CreateIntegrationResponse",
			Err:           err,
		}
	}
	out := result.(*CreateIntegrationResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new IntegrationResponse resource to represent an integration response.
type CreateIntegrationResponseInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The integration ID.
	//
	// This member is required.
	IntegrationId *string

	// The integration response key.
	//
	// This member is required.
	IntegrationResponseKey *string

	// Specifies how to handle response payload content type conversions. Supported
	// values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following behaviors:
	// CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string to
	// the corresponding binary blob. CONVERT_TO_TEXT: Converts a response payload from
	// a binary blob to a Base64-encoded string. If this property is not defined, the
	// response payload will be passed through from the integration response to the
	// route response or method response without modification.
	ContentHandlingStrategy types.ContentHandlingStrategy

	// A key-value map specifying response parameters that are passed to the method
	// response from the backend. The key is a method response header parameter name
	// and the mapped value is an integration response header value, a static value
	// enclosed within a pair of single quotes, or a JSON expression from the
	// integration response body. The mapping key must match the pattern of
	// method.response.header.{name}, where {name} is a valid and unique header name.
	// The mapped non-static value must match the pattern of
	// integration.response.header.{name} or
	// integration.response.body.{JSON-expression}, where {name} is a valid and unique
	// response header name and {JSON-expression} is a valid JSON expression without
	// the $ prefix.
	ResponseParameters map[string]*string

	// The collection of response templates for the integration response as a
	// string-to-string map of key-value pairs. Response templates are represented as a
	// key/value map, with a content-type as the key and a template as the value.
	ResponseTemplates map[string]*string

	// The template selection expression for the integration response. Supported only
	// for WebSocket APIs.
	TemplateSelectionExpression *string
}

type CreateIntegrationResponseOutput struct {

	// Supported only for WebSocket APIs. Specifies how to handle response payload
	// content type conversions. Supported values are CONVERT_TO_BINARY and
	// CONVERT_TO_TEXT, with the following behaviors: CONVERT_TO_BINARY: Converts a
	// response payload from a Base64-encoded string to the corresponding binary blob.
	// CONVERT_TO_TEXT: Converts a response payload from a binary blob to a
	// Base64-encoded string. If this property is not defined, the response payload
	// will be passed through from the integration response to the route response or
	// method response without modification.
	ContentHandlingStrategy types.ContentHandlingStrategy

	// The integration response ID.
	IntegrationResponseId *string

	// The integration response key.
	IntegrationResponseKey *string

	// A key-value map specifying response parameters that are passed to the method
	// response from the backend. The key is a method response header parameter name
	// and the mapped value is an integration response header value, a static value
	// enclosed within a pair of single quotes, or a JSON expression from the
	// integration response body. The mapping key must match the pattern of
	// method.response.header.{name}, where name is a valid and unique header name. The
	// mapped non-static value must match the pattern of
	// integration.response.header.{name} or
	// integration.response.body.{JSON-expression}, where name is a valid and unique
	// response header name and JSON-expression is a valid JSON expression without the
	// $ prefix.
	ResponseParameters map[string]*string

	// The collection of response templates for the integration response as a
	// string-to-string map of key-value pairs. Response templates are represented as a
	// key/value map, with a content-type as the key and a template as the value.
	ResponseTemplates map[string]*string

	// The template selection expressions for the integration response.
	TemplateSelectionExpression *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateIntegrationResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateIntegrationResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIntegrationResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateIntegrationResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateIntegrationResponse",
	}
}
