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

// Deletes an existing MethodResponse () resource.
func (c *Client) DeleteMethodResponse(ctx context.Context, params *DeleteMethodResponseInput, optFns ...func(*Options)) (*DeleteMethodResponseOutput, error) {
	stack := middleware.NewStack("DeleteMethodResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteMethodResponseMiddlewares(stack)
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
	addOpDeleteMethodResponseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMethodResponse(options.Region), middleware.Before)

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
			OperationName: "DeleteMethodResponse",
			Err:           err,
		}
	}
	out := result.(*DeleteMethodResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete an existing MethodResponse () resource.
type DeleteMethodResponseInput struct {
	Title *string
	// [Required] The HTTP verb of the Method () resource.
	HttpMethod *string
	// [Required] The status code identifier for the MethodResponse () resource.
	StatusCode *string
	// [Required] The Resource () identifier for the MethodResponse () resource.
	ResourceId *string
	Template   *bool
	// [Required] The string identifier of the associated RestApi ().
	RestApiId        *string
	Name             *string
	TemplateSkipList []*string
}

type DeleteMethodResponseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteMethodResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMethodResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMethodResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteMethodResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteMethodResponse",
	}
}
