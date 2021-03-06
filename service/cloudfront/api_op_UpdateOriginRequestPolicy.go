// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an origin request policy configuration. When you update an origin
// request policy configuration, all the fields are updated with the values
// provided in the request. You cannot update some fields independent of others. To
// update an origin request policy configuration:
//
//     * Use
// GetOriginRequestPolicyConfig to get the current configuration.
//
//     * Locally
// modify the fields in the origin request policy configuration that you want to
// update.
//
//     * Call UpdateOriginRequestPolicy by providing the entire origin
// request policy configuration, including the fields that you modified and those
// that you didn’t.
func (c *Client) UpdateOriginRequestPolicy(ctx context.Context, params *UpdateOriginRequestPolicyInput, optFns ...func(*Options)) (*UpdateOriginRequestPolicyOutput, error) {
	stack := middleware.NewStack("UpdateOriginRequestPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpUpdateOriginRequestPolicyMiddlewares(stack)
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
	addOpUpdateOriginRequestPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateOriginRequestPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "UpdateOriginRequestPolicy",
			Err:           err,
		}
	}
	out := result.(*UpdateOriginRequestPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateOriginRequestPolicyInput struct {

	// The unique identifier for the origin request policy that you are updating. The
	// identifier is returned in a cache behavior’s OriginRequestPolicyId field in the
	// response to GetDistributionConfig.
	//
	// This member is required.
	Id *string

	// An origin request policy configuration.
	//
	// This member is required.
	OriginRequestPolicyConfig *types.OriginRequestPolicyConfig

	// The version of the origin request policy that you are updating. The version is
	// returned in the origin request policy’s ETag field in the response to
	// GetOriginRequestPolicyConfig.
	IfMatch *string
}

type UpdateOriginRequestPolicyOutput struct {

	// The current version of the origin request policy.
	ETag *string

	// An origin request policy.
	OriginRequestPolicy *types.OriginRequestPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpUpdateOriginRequestPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpUpdateOriginRequestPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateOriginRequestPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateOriginRequestPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateOriginRequestPolicy",
	}
}
