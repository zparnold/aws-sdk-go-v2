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

// Flushes a stage's cache.
func (c *Client) FlushStageCache(ctx context.Context, params *FlushStageCacheInput, optFns ...func(*Options)) (*FlushStageCacheOutput, error) {
	stack := middleware.NewStack("FlushStageCache", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpFlushStageCacheMiddlewares(stack)
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
	addOpFlushStageCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opFlushStageCache(options.Region), middleware.Before)
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
			OperationName: "FlushStageCache",
			Err:           err,
		}
	}
	out := result.(*FlushStageCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to flush a stage's cache.
type FlushStageCacheInput struct {

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// [Required] The name of the stage to flush its cache.
	//
	// This member is required.
	StageName *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

type FlushStageCacheOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpFlushStageCacheMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpFlushStageCache{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpFlushStageCache{}, middleware.After)
}

func newServiceMetadataMiddleware_opFlushStageCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "FlushStageCache",
	}
}
