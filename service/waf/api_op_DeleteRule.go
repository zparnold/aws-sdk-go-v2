// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Permanently deletes a Rule (). You can't delete a Rule if it's still
// used in any WebACL objects or if it still includes any predicates, such as
// ByteMatchSet objects. If you just want to remove a Rule from a WebACL, use
// UpdateWebACL (). To permanently delete a Rule from AWS WAF, perform the
// following steps:
//
//     * Update the Rule to remove predicates, if any. For more
// information, see UpdateRule ().
//
//     * Use GetChangeToken () to get the change
// token that you provide in the ChangeToken parameter of a DeleteRule request.
//
//
// * Submit a DeleteRule request.
func (c *Client) DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error) {
	stack := middleware.NewStack("DeleteRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteRuleMiddlewares(stack)
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
	addOpDeleteRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRule(options.Region), middleware.Before)

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
			OperationName: "DeleteRule",
			Err:           err,
		}
	}
	out := result.(*DeleteRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRuleInput struct {
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
	// The RuleId of the Rule () that you want to delete. RuleId is returned by
	// CreateRule () and by ListRules ().
	RuleId *string
}

type DeleteRuleOutput struct {
	// The ChangeToken that you used to submit the DeleteRule request. You can also use
	// this value to query the status of the request. For more information, see
	// GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "DeleteRule",
	}
}
