// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the standards specified by the provided StandardsSubscriptionArns. For
// more information, see Security Standards
// (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards.html)
// section of the AWS Security Hub User Guide.
func (c *Client) BatchDisableStandards(ctx context.Context, params *BatchDisableStandardsInput, optFns ...func(*Options)) (*BatchDisableStandardsOutput, error) {
	if params == nil {
		params = &BatchDisableStandardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDisableStandards", params, optFns, addOperationBatchDisableStandardsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDisableStandardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisableStandardsInput struct {

	// The ARNs of the standards subscriptions to disable.
	//
	// This member is required.
	StandardsSubscriptionArns []*string
}

type BatchDisableStandardsOutput struct {

	// The details of the standards subscriptions that were disabled.
	StandardsSubscriptions []*types.StandardsSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDisableStandardsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDisableStandards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDisableStandards{}, middleware.After)
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
	addOpBatchDisableStandardsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisableStandards(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchDisableStandards(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchDisableStandards",
	}
}
