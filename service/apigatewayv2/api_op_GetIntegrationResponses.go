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

// Gets the IntegrationResponses for an Integration.
func (c *Client) GetIntegrationResponses(ctx context.Context, params *GetIntegrationResponsesInput, optFns ...func(*Options)) (*GetIntegrationResponsesOutput, error) {
	stack := middleware.NewStack("GetIntegrationResponses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetIntegrationResponsesMiddlewares(stack)
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
	addOpGetIntegrationResponsesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntegrationResponses(options.Region), middleware.Before)
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
			OperationName: "GetIntegrationResponses",
			Err:           err,
		}
	}
	out := result.(*GetIntegrationResponsesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntegrationResponsesInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The integration ID.
	//
	// This member is required.
	IntegrationId *string

	// The maximum number of elements to be returned for this resource.
	MaxResults *string

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string
}

type GetIntegrationResponsesOutput struct {

	// The elements from this collection.
	Items []*types.IntegrationResponse

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetIntegrationResponsesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetIntegrationResponses{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntegrationResponses{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIntegrationResponses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetIntegrationResponses",
	}
}
