// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all Elasticsearch instance types that are supported for given
// ElasticsearchVersion
func (c *Client) ListElasticsearchInstanceTypes(ctx context.Context, params *ListElasticsearchInstanceTypesInput, optFns ...func(*Options)) (*ListElasticsearchInstanceTypesOutput, error) {
	stack := middleware.NewStack("ListElasticsearchInstanceTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListElasticsearchInstanceTypesMiddlewares(stack)
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
	addOpListElasticsearchInstanceTypesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListElasticsearchInstanceTypes(options.Region), middleware.Before)

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
			OperationName: "ListElasticsearchInstanceTypes",
			Err:           err,
		}
	}
	out := result.(*ListElasticsearchInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the ListElasticsearchInstanceTypes () operation.
type ListElasticsearchInstanceTypesInput struct {
	// NextToken should be sent in case if earlier API call produced result containing
	// NextToken. It is used for pagination.
	NextToken *string
	// DomainName represents the name of the Domain that we are trying to modify. This
	// should be present only if we are querying for list of available Elasticsearch
	// instance types when modifying existing domain.
	DomainName *string
	// Set this value to limit the number of results returned. Value provided must be
	// greater than 30 else it wont be honored.
	MaxResults *int32
	// Version of Elasticsearch for which list of supported elasticsearch instance
	// types are needed.
	ElasticsearchVersion *string
}

// Container for the parameters returned by ListElasticsearchInstanceTypes ()
// operation.
type ListElasticsearchInstanceTypesOutput struct {
	// List of instance types supported by Amazon Elasticsearch service for given
	// ElasticsearchVersion ()
	ElasticsearchInstanceTypes []types.ESPartitionInstanceType
	// In case if there are more results available NextToken would be present, make
	// further request to the same API with received NextToken to paginate remaining
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListElasticsearchInstanceTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListElasticsearchInstanceTypes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListElasticsearchInstanceTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListElasticsearchInstanceTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListElasticsearchInstanceTypes",
	}
}
