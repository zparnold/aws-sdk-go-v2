// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Represents a delete integration response.
func (c *Client) DeleteIntegrationResponse(ctx context.Context, params *DeleteIntegrationResponseInput, optFns ...func(*Options)) (*DeleteIntegrationResponseOutput, error) {
	stack := middleware.NewStack("DeleteIntegrationResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteIntegrationResponseMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteIntegrationResponseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIntegrationResponse(options.Region), middleware.Before)

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
			OperationName: "DeleteIntegrationResponse",
			Err:           err,
		}
	}
	out := result.(*DeleteIntegrationResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a delete integration response request.
type DeleteIntegrationResponseInput struct {
	Template *bool
	// [Required] Specifies a delete integration response request's resource
	// identifier.
	ResourceId *string
	Title      *string
	// [Required] Specifies a delete integration response request's HTTP method.
	HttpMethod *string
	// [Required] Specifies a delete integration response request's status code.
	StatusCode       *string
	Name             *string
	TemplateSkipList []*string
	// [Required] The string identifier of the associated RestApi ().
	RestApiId *string
}

type DeleteIntegrationResponseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteIntegrationResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIntegrationResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIntegrationResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteIntegrationResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteIntegrationResponse",
	}
}
