// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a receipt rule. For information about setting up receipt rules, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rules.html).
// You can execute this operation no more than once per second.
func (c *Client) CreateReceiptRule(ctx context.Context, params *CreateReceiptRuleInput, optFns ...func(*Options)) (*CreateReceiptRuleOutput, error) {
	if params == nil {
		params = &CreateReceiptRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReceiptRule", params, optFns, addOperationCreateReceiptRuleMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReceiptRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create a receipt rule. You use receipt rules to receive
// email with Amazon SES. For more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type CreateReceiptRuleInput struct {

	// A data structure that contains the specified rule's name, actions, recipients,
	// domains, enabled status, scan status, and TLS policy.
	//
	// This member is required.
	Rule *types.ReceiptRule

	// The name of the rule set that the receipt rule will be added to.
	//
	// This member is required.
	RuleSetName *string

	// The name of an existing rule after which the new rule will be placed. If this
	// parameter is null, the new rule will be inserted at the beginning of the rule
	// list.
	After *string
}

// An empty element returned on a successful request.
type CreateReceiptRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateReceiptRuleMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateReceiptRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateReceiptRule{}, middleware.After)
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
	addOpCreateReceiptRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReceiptRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateReceiptRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateReceiptRule",
	}
}
