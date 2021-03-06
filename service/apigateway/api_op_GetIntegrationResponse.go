// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Represents a get integration response.
func (c *Client) GetIntegrationResponse(ctx context.Context, params *GetIntegrationResponseInput, optFns ...func(*Options)) (*GetIntegrationResponseOutput, error) {
	stack := middleware.NewStack("GetIntegrationResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetIntegrationResponseMiddlewares(stack)
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
	addOpGetIntegrationResponseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntegrationResponse(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "GetIntegrationResponse",
			Err:           err,
		}
	}
	out := result.(*GetIntegrationResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a get integration response request.
type GetIntegrationResponseInput struct {

	// [Required] Specifies a get integration response request's HTTP method.
	//
	// This member is required.
	HttpMethod *string

	// [Required] Specifies a get integration response request's resource identifier.
	//
	// This member is required.
	ResourceId *string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// [Required] Specifies a get integration response request's status code.
	//
	// This member is required.
	StatusCode *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents an integration response. The status code must map to an existing
// MethodResponse (), and parameters and templates can be used to transform the
// back-end response. Creating an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type GetIntegrationResponseOutput struct {

	// Specifies how to handle response payload content type conversions. Supported
	// values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following
	// behaviors:
	//
	//     * CONVERT_TO_BINARY: Converts a response payload from a
	// Base64-encoded string to the corresponding binary blob.
	//
	//     * CONVERT_TO_TEXT:
	// Converts a response payload from a binary blob to a Base64-encoded string.
	//
	// If
	// this property is not defined, the response payload will be passed through from
	// the integration response to the method response without modification.
	ContentHandling types.ContentHandlingStrategy

	// A key-value map specifying response parameters that are passed to the method
	// response from the back end. The key is a method response header parameter name
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

	// Specifies the templates used to transform the integration response body.
	// Response templates are represented as a key/value map, with a content-type as
	// the key and a template as the value.
	ResponseTemplates map[string]*string

	// Specifies the regular expression (regex) pattern used to choose an integration
	// response based on the response from the back end. For example, if the success
	// response returns nothing and the error response returns some string, you could
	// use the .+ regex to match error response. However, make sure that the error
	// response does not contain any newline (\n) character in such cases. If the back
	// end is an AWS Lambda function, the AWS Lambda function error header is matched.
	// For all other HTTP and AWS back ends, the HTTP status code is matched.
	SelectionPattern *string

	// Specifies the status code that is used to map the integration response to an
	// existing MethodResponse ().
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetIntegrationResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetIntegrationResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntegrationResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIntegrationResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetIntegrationResponse",
	}
}
