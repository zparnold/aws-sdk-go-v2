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
// global use. Returns an array of XssMatchSet () objects.
func (c *Client) ListXssMatchSets(ctx context.Context, params *ListXssMatchSetsInput, optFns ...func(*Options)) (*ListXssMatchSetsOutput, error) {
	if params == nil {
		params = &ListXssMatchSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListXssMatchSets", params, optFns, addOperationListXssMatchSetsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListXssMatchSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to list the XssMatchSet () objects created by the current AWS account.
type ListXssMatchSetsInput struct {

	// Specifies the number of XssMatchSet () objects that you want AWS WAF to return
	// for this request. If you have more XssMatchSet objects than the number you
	// specify for Limit, the response includes a NextMarker value that you can use to
	// get another batch of Rules.
	Limit *int32

	// If you specify a value for Limit and you have more XssMatchSet () objects than
	// the value of Limit, AWS WAF returns a NextMarker value in the response that
	// allows you to list another group of XssMatchSets. For the second and subsequent
	// ListXssMatchSets requests, specify the value of NextMarker from the previous
	// response to get information about another batch of XssMatchSets.
	NextMarker *string
}

// The response to a ListXssMatchSets () request.
type ListXssMatchSetsOutput struct {

	// If you have more XssMatchSet () objects than the number that you specified for
	// Limit in the request, the response includes a NextMarker value. To list more
	// XssMatchSet objects, submit another ListXssMatchSets request, and specify the
	// NextMarker value from the response in the NextMarker value in the next request.
	NextMarker *string

	// An array of XssMatchSetSummary () objects.
	XssMatchSets []*types.XssMatchSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListXssMatchSetsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListXssMatchSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListXssMatchSets{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListXssMatchSets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListXssMatchSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "ListXssMatchSets",
	}
}
