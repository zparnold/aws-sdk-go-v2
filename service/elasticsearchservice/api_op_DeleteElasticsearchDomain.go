// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Permanently deletes the specified Elasticsearch domain and all of its data. Once
// a domain is deleted, it cannot be recovered.
func (c *Client) DeleteElasticsearchDomain(ctx context.Context, params *DeleteElasticsearchDomainInput, optFns ...func(*Options)) (*DeleteElasticsearchDomainOutput, error) {
	if params == nil {
		params = &DeleteElasticsearchDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteElasticsearchDomain", params, optFns, addOperationDeleteElasticsearchDomainMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteElasticsearchDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteElasticsearchDomain () operation.
// Specifies the name of the Elasticsearch domain that you want to delete.
type DeleteElasticsearchDomainInput struct {

	// The name of the Elasticsearch domain that you want to permanently delete.
	//
	// This member is required.
	DomainName *string
}

// The result of a DeleteElasticsearchDomain request. Contains the status of the
// pending deletion, or no status if the domain and all of its resources have been
// deleted.
type DeleteElasticsearchDomainOutput struct {

	// The status of the Elasticsearch domain being deleted.
	DomainStatus *types.ElasticsearchDomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteElasticsearchDomainMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteElasticsearchDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteElasticsearchDomain{}, middleware.After)
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
	addOpDeleteElasticsearchDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteElasticsearchDomain(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteElasticsearchDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DeleteElasticsearchDomain",
	}
}
