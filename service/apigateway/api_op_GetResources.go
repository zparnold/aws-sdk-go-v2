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

// Lists information about a collection of Resource () resources.
func (c *Client) GetResources(ctx context.Context, params *GetResourcesInput, optFns ...func(*Options)) (*GetResourcesOutput, error) {
	stack := middleware.NewStack("GetResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetResourcesMiddlewares(stack)
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
	addOpGetResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResources(options.Region), middleware.Before)
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
			OperationName: "GetResources",
			Err:           err,
		}
	}
	out := result.(*GetResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list information about a collection of resources.
type GetResourcesInput struct {

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// A query parameter used to retrieve the specified resources embedded in the
	// returned Resources () resource in the response. This embed parameter value is a
	// list of comma-separated strings. Currently, the request supports only retrieval
	// of the embedded Method () resources this way. The query parameter value must be
	// a single-valued list and contain the "methods" string. For example, GET
	// /restapis/{restapi_id}/resources?embed=methods.
	Embed []*string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	Name *string

	// The current pagination position in the paged result set.
	Position *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a collection of Resource () resources. Create an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type GetResourcesOutput struct {

	// The current page of elements from this collection.
	Items []*types.Resource

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetResources{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetResources",
	}
}
