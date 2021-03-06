// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables ClassicLink DNS support for a VPC. If disabled, DNS hostnames resolve
// to public IP addresses when addressed between a linked EC2-Classic instance and
// instances in the VPC to which it's linked. For more information, see ClassicLink
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
// the Amazon Elastic Compute Cloud User Guide. You must specify a VPC ID in the
// request.
func (c *Client) DisableVpcClassicLinkDnsSupport(ctx context.Context, params *DisableVpcClassicLinkDnsSupportInput, optFns ...func(*Options)) (*DisableVpcClassicLinkDnsSupportOutput, error) {
	stack := middleware.NewStack("DisableVpcClassicLinkDnsSupport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisableVpcClassicLinkDnsSupportMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableVpcClassicLinkDnsSupport(options.Region), middleware.Before)
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
			OperationName: "DisableVpcClassicLinkDnsSupport",
			Err:           err,
		}
	}
	out := result.(*DisableVpcClassicLinkDnsSupportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableVpcClassicLinkDnsSupportInput struct {

	// The ID of the VPC.
	VpcId *string
}

type DisableVpcClassicLinkDnsSupportOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisableVpcClassicLinkDnsSupportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisableVpcClassicLinkDnsSupport{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisableVpcClassicLinkDnsSupport{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableVpcClassicLinkDnsSupport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableVpcClassicLinkDnsSupport",
	}
}
