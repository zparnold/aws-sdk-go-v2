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

// Returns the current configuration for one or more requested resources. The
// operation also returns a list of resources that are not processed in the current
// request. If there are no unprocessed resources, the operation returns an empty
// unprocessedResourceKeys list.
//
//     * The API does not return results for deleted
// resources.
//
//     * The API does not return any tags for the requested resources.
// This information is filtered out of the supplementaryConfiguration section of
// the API response.
func (c *Client) BatchGetResourceConfig(ctx context.Context, params *BatchGetResourceConfigInput, optFns ...func(*Options)) (*BatchGetResourceConfigOutput, error) {
	if params == nil {
		params = &BatchGetResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetResourceConfig", params, optFns, addOperationBatchGetResourceConfigMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetResourceConfigInput struct {

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	//
	// This member is required.
	ResourceKeys []*types.ResourceKey
}

type BatchGetResourceConfigOutput struct {

	// A list that contains the current configuration of one or more resources.
	BaseConfigurationItems []*types.BaseConfigurationItem

	// A list of resource keys that were not processed with the current response. The
	// unprocessesResourceKeys value is in the same form as ResourceKeys, so the value
	// can be directly provided to a subsequent BatchGetResourceConfig operation.  If
	// there are no unprocessed resource keys, the response contains an empty
	// unprocessedResourceKeys list. </p>
	UnprocessedResourceKeys []*types.ResourceKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetResourceConfigMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetResourceConfig{}, middleware.After)
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
	addOpBatchGetResourceConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetResourceConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchGetResourceConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "BatchGetResourceConfig",
	}
}
