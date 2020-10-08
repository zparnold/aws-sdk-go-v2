// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of resources of the specified type that have been analyzed by
// the specified analyzer..
func (c *Client) ListAnalyzedResources(ctx context.Context, params *ListAnalyzedResourcesInput, optFns ...func(*Options)) (*ListAnalyzedResourcesOutput, error) {
	if params == nil {
		params = &ListAnalyzedResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnalyzedResources", params, optFns, addOperationListAnalyzedResourcesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnalyzedResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves a list of resources that have been analyzed.
type ListAnalyzedResourcesInput struct {

	// The ARN of the analyzer to retrieve a list of analyzed resources from.
	//
	// This member is required.
	AnalyzerArn *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string

	// The type of resource.
	ResourceType types.ResourceType
}

// The response to the request.
type ListAnalyzedResourcesOutput struct {

	// A list of resources that were analyzed.
	//
	// This member is required.
	AnalyzedResources []*types.AnalyzedResourceSummary

	// A token used for pagination of results returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAnalyzedResourcesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnalyzedResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnalyzedResources{}, middleware.After)
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
	addOpListAnalyzedResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAnalyzedResources(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAnalyzedResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "ListAnalyzedResources",
	}
}
