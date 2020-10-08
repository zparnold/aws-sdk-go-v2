// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified policies or all policies used for SSL negotiation. For
// more information, see Security Policies
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies)
// in the Application Load Balancers Guide.
func (c *Client) DescribeSSLPolicies(ctx context.Context, params *DescribeSSLPoliciesInput, optFns ...func(*Options)) (*DescribeSSLPoliciesOutput, error) {
	if params == nil {
		params = &DescribeSSLPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSSLPolicies", params, optFns, addOperationDescribeSSLPoliciesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSSLPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSSLPoliciesInput struct {

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The names of the policies.
	Names []*string

	// The maximum number of results to return with this call.
	PageSize *int32
}

type DescribeSSLPoliciesOutput struct {

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string

	// Information about the security policies.
	SslPolicies []*types.SslPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSSLPoliciesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSSLPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSSLPolicies{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSSLPolicies(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSSLPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeSSLPolicies",
	}
}
