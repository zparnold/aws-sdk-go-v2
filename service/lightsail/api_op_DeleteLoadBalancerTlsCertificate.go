// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an SSL/TLS certificate associated with a Lightsail load balancer. The
// DeleteLoadBalancerTlsCertificate operation supports tag-based access control via
// resource tags applied to the resource identified by load balancer name. For more
// information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) DeleteLoadBalancerTlsCertificate(ctx context.Context, params *DeleteLoadBalancerTlsCertificateInput, optFns ...func(*Options)) (*DeleteLoadBalancerTlsCertificateOutput, error) {
	if params == nil {
		params = &DeleteLoadBalancerTlsCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLoadBalancerTlsCertificate", params, optFns, addOperationDeleteLoadBalancerTlsCertificateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLoadBalancerTlsCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLoadBalancerTlsCertificateInput struct {

	// The SSL/TLS certificate name.
	//
	// This member is required.
	CertificateName *string

	// The load balancer name.
	//
	// This member is required.
	LoadBalancerName *string

	// When true, forces the deletion of an SSL/TLS certificate. There can be two
	// certificates associated with a Lightsail load balancer: the primary and the
	// backup. The force parameter is required when the primary SSL/TLS certificate is
	// in use by an instance attached to the load balancer.
	Force *bool
}

type DeleteLoadBalancerTlsCertificateOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteLoadBalancerTlsCertificateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteLoadBalancerTlsCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteLoadBalancerTlsCertificate{}, middleware.After)
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
	addOpDeleteLoadBalancerTlsCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLoadBalancerTlsCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteLoadBalancerTlsCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DeleteLoadBalancerTlsCertificate",
	}
}
