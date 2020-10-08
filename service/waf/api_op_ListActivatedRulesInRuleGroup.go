// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns an array of ActivatedRule () objects.
func (c *Client) ListActivatedRulesInRuleGroup(ctx context.Context, params *ListActivatedRulesInRuleGroupInput, optFns ...func(*Options)) (*ListActivatedRulesInRuleGroupOutput, error) {
	if params == nil {
		params = &ListActivatedRulesInRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListActivatedRulesInRuleGroup", params, optFns, addOperationListActivatedRulesInRuleGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListActivatedRulesInRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListActivatedRulesInRuleGroupInput struct {

	// Specifies the number of ActivatedRules that you want AWS WAF to return for this
	// request. If you have more ActivatedRules than the number that you specify for
	// Limit, the response includes a NextMarker value that you can use to get another
	// batch of ActivatedRules.
	Limit *int32

	// If you specify a value for Limit and you have more ActivatedRules than the value
	// of Limit, AWS WAF returns a NextMarker value in the response that allows you to
	// list another group of ActivatedRules. For the second and subsequent
	// ListActivatedRulesInRuleGroup requests, specify the value of NextMarker from the
	// previous response to get information about another batch of ActivatedRules.
	NextMarker *string

	// The RuleGroupId of the RuleGroup () for which you want to get a list of
	// ActivatedRule () objects.
	RuleGroupId *string
}

type ListActivatedRulesInRuleGroupOutput struct {

	// An array of ActivatedRules objects.
	ActivatedRules []*types.ActivatedRule

	// If you have more ActivatedRules than the number that you specified for Limit in
	// the request, the response includes a NextMarker value. To list more
	// ActivatedRules, submit another ListActivatedRulesInRuleGroup request, and
	// specify the NextMarker value from the response in the NextMarker value in the
	// next request.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListActivatedRulesInRuleGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListActivatedRulesInRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListActivatedRulesInRuleGroup{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListActivatedRulesInRuleGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListActivatedRulesInRuleGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "ListActivatedRulesInRuleGroup",
	}
}
