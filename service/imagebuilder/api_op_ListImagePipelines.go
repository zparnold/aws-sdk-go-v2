// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of image pipelines.
func (c *Client) ListImagePipelines(ctx context.Context, params *ListImagePipelinesInput, optFns ...func(*Options)) (*ListImagePipelinesOutput, error) {
	if params == nil {
		params = &ListImagePipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImagePipelines", params, optFns, addOperationListImagePipelinesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImagePipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImagePipelinesInput struct {

	// The filters.
	Filters []*types.Filter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string
}

type ListImagePipelinesOutput struct {

	// The list of image pipelines.
	ImagePipelineList []*types.ImagePipeline

	// The next token used for paginated responses. When this is not empty, there are
	// additional elements that the service has not included in this request. Use this
	// token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListImagePipelinesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImagePipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImagePipelines{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListImagePipelines(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListImagePipelines(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListImagePipelines",
	}
}
