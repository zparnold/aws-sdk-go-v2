// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the maximum number of hosted zones that you can associate with the
// specified reusable delegation set. For the default limit, see Limits
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To request a higher limit, open a case
// (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53).
func (c *Client) GetReusableDelegationSetLimit(ctx context.Context, params *GetReusableDelegationSetLimitInput, optFns ...func(*Options)) (*GetReusableDelegationSetLimitOutput, error) {
	stack := middleware.NewStack("GetReusableDelegationSetLimit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetReusableDelegationSetLimitMiddlewares(stack)
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
	addOpGetReusableDelegationSetLimitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetReusableDelegationSetLimit(options.Region), middleware.Before)

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
			OperationName: "GetReusableDelegationSetLimit",
			Err:           err,
		}
	}
	out := result.(*GetReusableDelegationSetLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to create a hosted
// zone.
type GetReusableDelegationSetLimitInput struct {
	// The ID of the delegation set that you want to get the limit for.
	DelegationSetId *string
	// Specify MAX_ZONES_BY_REUSABLE_DELEGATION_SET to get the maximum number of hosted
	// zones that you can associate with the specified reusable delegation set.
	Type types.ReusableDelegationSetLimitType
}

// A complex type that contains the requested limit.
type GetReusableDelegationSetLimitOutput struct {
	// The current setting for the limit on hosted zones that you can associate with
	// the specified reusable delegation set.
	Limit *types.ReusableDelegationSetLimit
	// The current number of hosted zones that you can associate with the specified
	// reusable delegation set.
	Count *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetReusableDelegationSetLimitMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetReusableDelegationSetLimit{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetReusableDelegationSetLimit{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetReusableDelegationSetLimit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetReusableDelegationSetLimit",
	}
}
