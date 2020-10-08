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

// Attaches a Transport Layer Security (TLS) certificate to your load balancer. TLS
// is just an updated, more secure version of Secure Socket Layer (SSL). Once you
// create and validate your certificate, you can attach it to your load balancer.
// You can also use this API to rotate the certificates on your account. Use the
// AttachLoadBalancerTlsCertificate action with the non-attached certificate, and
// it will replace the existing one and become the attached certificate. The
// AttachLoadBalancerTlsCertificate operation supports tag-based access control via
// resource tags applied to the resource identified by load balancer name. For more
// information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) AttachLoadBalancerTlsCertificate(ctx context.Context, params *AttachLoadBalancerTlsCertificateInput, optFns ...func(*Options)) (*AttachLoadBalancerTlsCertificateOutput, error) {
	if params == nil {
		params = &AttachLoadBalancerTlsCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachLoadBalancerTlsCertificate", params, optFns, addOperationAttachLoadBalancerTlsCertificateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachLoadBalancerTlsCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachLoadBalancerTlsCertificateInput struct {

	// The name of your SSL/TLS certificate.
	//
	// This member is required.
	CertificateName *string

	// The name of the load balancer to which you want to associate the SSL/TLS
	// certificate.
	//
	// This member is required.
	LoadBalancerName *string
}

type AttachLoadBalancerTlsCertificateOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request. These SSL/TLS certificates are only usable by Lightsail load balancers.
	// You can't get the certificate and use it for another purpose.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachLoadBalancerTlsCertificateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachLoadBalancerTlsCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachLoadBalancerTlsCertificate{}, middleware.After)
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
	addOpAttachLoadBalancerTlsCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachLoadBalancerTlsCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAttachLoadBalancerTlsCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "AttachLoadBalancerTlsCertificate",
	}
}
