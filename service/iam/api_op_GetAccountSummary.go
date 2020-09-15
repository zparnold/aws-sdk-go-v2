// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about IAM entity usage and IAM quotas in the AWS account.
// The number and size of IAM resources in an AWS account are limited. For more
// information, see IAM and STS Quotas
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in
// the IAM User Guide.
func (c *Client) GetAccountSummary(ctx context.Context, params *GetAccountSummaryInput, optFns ...func(*Options)) (*GetAccountSummaryOutput, error) {
	stack := middleware.NewStack("GetAccountSummary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetAccountSummaryMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccountSummary(options.Region), middleware.Before)

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
			OperationName: "GetAccountSummary",
			Err:           err,
		}
	}
	out := result.(*GetAccountSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccountSummaryInput struct {
}

// Contains the response to a successful GetAccountSummary () request.
type GetAccountSummaryOutput struct {
	// A set of key–value pairs containing information about IAM entity usage and IAM
	// quotas.
	SummaryMap map[string]*int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetAccountSummaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetAccountSummary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetAccountSummary{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccountSummary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetAccountSummary",
	}
}
