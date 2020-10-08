// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Puts an Api resource.
func (c *Client) ReimportApi(ctx context.Context, params *ReimportApiInput, optFns ...func(*Options)) (*ReimportApiOutput, error) {
	if params == nil {
		params = &ReimportApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReimportApi", params, optFns, addOperationReimportApiMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ReimportApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ReimportApiInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The OpenAPI definition. Supported only for HTTP APIs.
	//
	// This member is required.
	Body *string

	// Specifies how to interpret the base path of the API during import. Valid values
	// are ignore, prepend, and split. The default value is ignore. To learn more, see
	// Set the OpenAPI basePath Property
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html).
	// Supported only for HTTP APIs.
	Basepath *string

	// Specifies whether to rollback the API creation when a warning is encountered. By
	// default, API creation continues if a warning is encountered.
	FailOnWarnings *bool
}

type ReimportApiOutput struct {

	// The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com. The
	// stage name is typically appended to this URI to form a complete path to a
	// deployed API stage.
	ApiEndpoint *string

	// The API ID.
	ApiId *string

	// An API key selection expression. Supported only for WebSocket APIs. See API Key
	// Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression *string

	// A CORS configuration. Supported only for HTTP APIs.
	CorsConfiguration *types.Cors

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The description of the API.
	Description *string

	// Avoid validating models when creating a deployment. Supported only for WebSocket
	// APIs.
	DisableSchemaValidation *bool

	// The validation information during API import. This may include particular
	// properties of your OpenAPI definition which are ignored during import. Supported
	// only for HTTP APIs.
	ImportInfo []*string

	// The name of the API.
	Name *string

	// The API protocol.
	ProtocolType types.ProtocolType

	// The route selection expression for the API. For HTTP APIs, the
	// routeSelectionExpression must be ${request.method} ${request.path}. If not
	// provided, this will be the default for HTTP APIs. This property is required for
	// WebSocket APIs.
	RouteSelectionExpression *string

	// A collection of tags associated with the API.
	Tags map[string]*string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationReimportApiMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpReimportApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpReimportApi{}, middleware.After)
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
	addOpReimportApiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReimportApi(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opReimportApi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "ReimportApi",
	}
}
