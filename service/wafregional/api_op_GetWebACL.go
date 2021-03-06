// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
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
// global use. Returns the WebACL () that is specified by WebACLId.
func (c *Client) GetWebACL(ctx context.Context, params *GetWebACLInput, optFns ...func(*Options)) (*GetWebACLOutput, error) {
	stack := middleware.NewStack("GetWebACL", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetWebACLMiddlewares(stack)
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
	addOpGetWebACLValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetWebACL(options.Region), middleware.Before)
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
			OperationName: "GetWebACL",
			Err:           err,
		}
	}
	out := result.(*GetWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWebACLInput struct {

	// The WebACLId of the WebACL () that you want to get. WebACLId is returned by
	// CreateWebACL () and by ListWebACLs ().
	//
	// This member is required.
	WebACLId *string
}

type GetWebACLOutput struct {

	// Information about the WebACL () that you specified in the GetWebACL request. For
	// more information, see the following topics:
	//
	//     * WebACL (): Contains
	// DefaultAction, MetricName, Name, an array of Rule objects, and WebACLId
	//
	//     *
	// DefaultAction (Data type is WafAction ()): Contains Type
	//
	//     * Rules: Contains
	// an array of ActivatedRule objects, which contain Action, Priority, and RuleId
	//
	//
	// * Action: Contains Type
	WebACL *types.WebACL

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetWebACLMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetWebACL{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWebACL{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetWebACL(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetWebACL",
	}
}
