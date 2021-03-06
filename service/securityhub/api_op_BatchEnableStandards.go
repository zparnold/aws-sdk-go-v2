// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the standards specified by the provided StandardsArn. To obtain the ARN
// for a standard, use the DescribeStandards () operation. For more information,
// see the Security Standards
// (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards.html)
// section of the AWS Security Hub User Guide.
func (c *Client) BatchEnableStandards(ctx context.Context, params *BatchEnableStandardsInput, optFns ...func(*Options)) (*BatchEnableStandardsOutput, error) {
	stack := middleware.NewStack("BatchEnableStandards", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchEnableStandardsMiddlewares(stack)
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
	addOpBatchEnableStandardsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchEnableStandards(options.Region), middleware.Before)
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
			OperationName: "BatchEnableStandards",
			Err:           err,
		}
	}
	out := result.(*BatchEnableStandardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchEnableStandardsInput struct {

	// The list of standards checks to enable.
	//
	// This member is required.
	StandardsSubscriptionRequests []*types.StandardsSubscriptionRequest
}

type BatchEnableStandardsOutput struct {

	// The details of the standards subscriptions that were enabled.
	StandardsSubscriptions []*types.StandardsSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchEnableStandardsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchEnableStandards{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchEnableStandards{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchEnableStandards(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchEnableStandards",
	}
}
