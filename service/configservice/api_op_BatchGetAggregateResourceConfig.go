// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current configuration items for resources that are present in your
// AWS Config aggregator. The operation also returns a list of resources that are
// not processed in the current request. If there are no unprocessed resources, the
// operation returns an empty unprocessedResourceIdentifiers list.  <note> <ul>
// <li> <p>The API does not return results for deleted resources.</p> </li> <li>
// <p> The API does not return tags and relationships.</p> </li> </ul> </note>
func (c *Client) BatchGetAggregateResourceConfig(ctx context.Context, params *BatchGetAggregateResourceConfigInput, optFns ...func(*Options)) (*BatchGetAggregateResourceConfigOutput, error) {
	if params == nil {
		params = &BatchGetAggregateResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetAggregateResourceConfig", params, optFns, addOperationBatchGetAggregateResourceConfigMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetAggregateResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetAggregateResourceConfigInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// A list of aggregate ResourceIdentifiers objects.
	//
	// This member is required.
	ResourceIdentifiers []*types.AggregateResourceIdentifier
}

type BatchGetAggregateResourceConfigOutput struct {

	// A list that contains the current configuration of one or more resources.
	BaseConfigurationItems []*types.BaseConfigurationItem

	// A list of resource identifiers that were not processed with current scope. The
	// list is empty if all the resources are processed.
	UnprocessedResourceIdentifiers []*types.AggregateResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetAggregateResourceConfigMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetAggregateResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetAggregateResourceConfig{}, middleware.After)
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
	addOpBatchGetAggregateResourceConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetAggregateResourceConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchGetAggregateResourceConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "BatchGetAggregateResourceConfig",
	}
}
