// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a specified resolver rule, such as the domain name that
// the rule forwards DNS queries for and the ID of the outbound resolver endpoint
// that the rule is associated with.
func (c *Client) GetResolverRule(ctx context.Context, params *GetResolverRuleInput, optFns ...func(*Options)) (*GetResolverRuleOutput, error) {
	if params == nil {
		params = &GetResolverRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResolverRule", params, optFns, addOperationGetResolverRuleMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResolverRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResolverRuleInput struct {

	// The ID of the resolver rule that you want to get information about.
	//
	// This member is required.
	ResolverRuleId *string
}

type GetResolverRuleOutput struct {

	// Information about the resolver rule that you specified in a GetResolverRule
	// request.
	ResolverRule *types.ResolverRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetResolverRuleMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResolverRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResolverRule{}, middleware.After)
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
	addOpGetResolverRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResolverRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetResolverRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "GetResolverRule",
	}
}
