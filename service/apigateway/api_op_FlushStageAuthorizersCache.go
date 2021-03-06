// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Flushes all authorizer cache entries on a stage.
func (c *Client) FlushStageAuthorizersCache(ctx context.Context, params *FlushStageAuthorizersCacheInput, optFns ...func(*Options)) (*FlushStageAuthorizersCacheOutput, error) {
	stack := middleware.NewStack("FlushStageAuthorizersCache", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpFlushStageAuthorizersCacheMiddlewares(stack)
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
	addOpFlushStageAuthorizersCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opFlushStageAuthorizersCache(options.Region), middleware.Before)
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
			OperationName: "FlushStageAuthorizersCache",
			Err:           err,
		}
	}
	out := result.(*FlushStageAuthorizersCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to flush authorizer cache entries on a specified stage.
type FlushStageAuthorizersCacheInput struct {

	// The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// The name of the stage to flush.
	//
	// This member is required.
	StageName *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

type FlushStageAuthorizersCacheOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpFlushStageAuthorizersCacheMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpFlushStageAuthorizersCache{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpFlushStageAuthorizersCache{}, middleware.After)
}

func newServiceMetadataMiddleware_opFlushStageAuthorizersCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "FlushStageAuthorizersCache",
	}
}
