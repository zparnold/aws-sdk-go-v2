// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
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
// global use. Inserts or deletes SqlInjectionMatchTuple () objects (filters) in a
// SqlInjectionMatchSet (). For each SqlInjectionMatchTuple object, you specify the
// following values:
//
//     * Action: Whether to insert the object into or delete the
// object from the array. To change a SqlInjectionMatchTuple, you delete the
// existing object and add a new one.
//
//     * FieldToMatch: The part of web requests
// that you want AWS WAF to inspect and, if you want AWS WAF to inspect a header or
// custom query parameter, the name of the header or parameter.
//
//     *
// TextTransformation: Which text transformation, if any, to perform on the web
// request before inspecting the request for snippets of malicious SQL code. You
// can only specify a single type of TextTransformation.
//
// You use
// SqlInjectionMatchSet objects to specify which CloudFront requests that you want
// to allow, block, or count. For example, if you're receiving requests that
// contain snippets of SQL code in the query string and you want to block the
// requests, you can create a SqlInjectionMatchSet with the applicable settings,
// and then configure AWS WAF to block the requests. To create and configure a
// SqlInjectionMatchSet, perform the following steps:
//
//     * Submit a
// CreateSqlInjectionMatchSet () request.
//
//     * Use GetChangeToken () to get the
// change token that you provide in the ChangeToken parameter of an UpdateIPSet ()
// request.
//
//     * Submit an UpdateSqlInjectionMatchSet request to specify the
// parts of web requests that you want AWS WAF to inspect for snippets of SQL
// code.
//
// For more information about how to use the AWS WAF API to allow or block
// HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) UpdateSqlInjectionMatchSet(ctx context.Context, params *UpdateSqlInjectionMatchSetInput, optFns ...func(*Options)) (*UpdateSqlInjectionMatchSetOutput, error) {
	stack := middleware.NewStack("UpdateSqlInjectionMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateSqlInjectionMatchSetMiddlewares(stack)
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
	addOpUpdateSqlInjectionMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSqlInjectionMatchSet(options.Region), middleware.Before)
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
			OperationName: "UpdateSqlInjectionMatchSet",
			Err:           err,
		}
	}
	out := result.(*UpdateSqlInjectionMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update a SqlInjectionMatchSet ().
type UpdateSqlInjectionMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string

	// The SqlInjectionMatchSetId of the SqlInjectionMatchSet that you want to update.
	// SqlInjectionMatchSetId is returned by CreateSqlInjectionMatchSet () and by
	// ListSqlInjectionMatchSets ().
	//
	// This member is required.
	SqlInjectionMatchSetId *string

	// An array of SqlInjectionMatchSetUpdate objects that you want to insert into or
	// delete from a SqlInjectionMatchSet (). For more information, see the applicable
	// data types:
	//
	//     * SqlInjectionMatchSetUpdate (): Contains Action and
	// SqlInjectionMatchTuple
	//
	//     * SqlInjectionMatchTuple (): Contains FieldToMatch
	// and TextTransformation
	//
	//     * FieldToMatch (): Contains Data and Type
	//
	// This member is required.
	Updates []*types.SqlInjectionMatchSetUpdate
}

// The response to an UpdateSqlInjectionMatchSets () request.
type UpdateSqlInjectionMatchSetOutput struct {

	// The ChangeToken that you used to submit the UpdateSqlInjectionMatchSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateSqlInjectionMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSqlInjectionMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSqlInjectionMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSqlInjectionMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateSqlInjectionMatchSet",
	}
}
