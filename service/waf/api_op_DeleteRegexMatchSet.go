// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
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
// global use. Permanently deletes a RegexMatchSet (). You can't delete a
// RegexMatchSet if it's still used in any Rules or if it still includes any
// RegexMatchTuples objects (any filters). If you just want to remove a
// RegexMatchSet from a Rule, use UpdateRule (). To permanently delete a
// RegexMatchSet, perform the following steps:
//
//     * Update the RegexMatchSet to
// remove filters, if any. For more information, see UpdateRegexMatchSet ().
//
//     *
// Use GetChangeToken () to get the change token that you provide in the
// ChangeToken parameter of a DeleteRegexMatchSet request.
//
//     * Submit a
// DeleteRegexMatchSet request.
func (c *Client) DeleteRegexMatchSet(ctx context.Context, params *DeleteRegexMatchSetInput, optFns ...func(*Options)) (*DeleteRegexMatchSetOutput, error) {
	stack := middleware.NewStack("DeleteRegexMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteRegexMatchSetMiddlewares(stack)
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
	addOpDeleteRegexMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRegexMatchSet(options.Region), middleware.Before)
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
			OperationName: "DeleteRegexMatchSet",
			Err:           err,
		}
	}
	out := result.(*DeleteRegexMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRegexMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string

	// The RegexMatchSetId of the RegexMatchSet () that you want to delete.
	// RegexMatchSetId is returned by CreateRegexMatchSet () and by ListRegexMatchSets
	// ().
	//
	// This member is required.
	RegexMatchSetId *string
}

type DeleteRegexMatchSetOutput struct {

	// The ChangeToken that you used to submit the DeleteRegexMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteRegexMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRegexMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRegexMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRegexMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "DeleteRegexMatchSet",
	}
}
