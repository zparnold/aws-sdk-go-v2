// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the current ApiKeys () resource.
func (c *Client) GetApiKeys(ctx context.Context, params *GetApiKeysInput, optFns ...func(*Options)) (*GetApiKeysOutput, error) {
	if params == nil {
		params = &GetApiKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApiKeys", params, optFns, addOperationGetApiKeysMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApiKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about the current ApiKeys () resource.
type GetApiKeysInput struct {

	// The identifier of a customer in AWS Marketplace or an external system, such as a
	// developer portal.
	CustomerId *string

	// A boolean flag to specify whether (true) or not (false) the result contains key
	// values.
	IncludeValues *bool

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	Name *string

	// The name of queried API keys.
	NameQuery *string

	// The current pagination position in the paged result set.
	Position *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a collection of API keys as represented by an ApiKeys () resource.
// Use API Keys
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-api-keys.html)
type GetApiKeysOutput struct {

	// The current page of elements from this collection.
	Items []*types.ApiKey

	// The current pagination position in the paged result set.
	Position *string

	// A list of warning messages logged during the import of API keys when the
	// failOnWarnings option is set to true.
	Warnings []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetApiKeysMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApiKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApiKeys{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetApiKeys(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetApiKeys(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetApiKeys",
	}
}
