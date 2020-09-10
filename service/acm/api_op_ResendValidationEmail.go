// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resends the email that requests domain ownership validation. The domain owner or
// an authorized representative must approve the ACM certificate before it can be
// issued. The certificate can be approved by clicking a link in the mail to
// navigate to the Amazon certificate approval website and then clicking I Approve.
// However, the validation email can be blocked by spam filters. Therefore, if you
// do not receive the original mail, you can request that the mail be resent within
// 72 hours of requesting the ACM certificate. If more than 72 hours have elapsed
// since your original request or since your last attempt to resend validation
// mail, you must request a new certificate. For more information about setting up
// your contact email addresses, see Configure Email for your Domain
// (https://docs.aws.amazon.com/acm/latest/userguide/setup-email.html).
func (c *Client) ResendValidationEmail(ctx context.Context, params *ResendValidationEmailInput, optFns ...func(*Options)) (*ResendValidationEmailOutput, error) {
	stack := middleware.NewStack("ResendValidationEmail", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResendValidationEmailMiddlewares(stack)
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
	addOpResendValidationEmailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResendValidationEmail(options.Region), middleware.Before)

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
			OperationName: "ResendValidationEmail",
			Err:           err,
		}
	}
	out := result.(*ResendValidationEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResendValidationEmailInput struct {
	// The base validation domain that will act as the suffix of the email addresses
	// that are used to send the emails. This must be the same as the Domain value or a
	// superdomain of the Domain value. For example, if you requested a certificate for
	// site.subdomain.example.com and specify a ValidationDomain of
	// subdomain.example.com, ACM sends email to the domain registrant, technical
	// contact, and administrative contact in WHOIS and the following five addresses:
	//
	//
	// * admin@subdomain.example.com
	//
	//     * administrator@subdomain.example.com
	//
	//     *
	// hostmaster@subdomain.example.com
	//
	//     * postmaster@subdomain.example.com
	//
	//     *
	// webmaster@subdomain.example.com
	ValidationDomain *string
	// String that contains the ARN of the requested certificate. The certificate ARN
	// is generated and returned by the RequestCertificate () action as soon as the
	// request is made. By default, using this parameter causes email to be sent to all
	// top-level domains you specified in the certificate request. The ARN must be of
	// the form:  <p>
	// <code>arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012</code>
	// </p>
	CertificateArn *string
	// The fully qualified domain name (FQDN) of the certificate that needs to be
	// validated.
	Domain *string
}

type ResendValidationEmailOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResendValidationEmailMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResendValidationEmail{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResendValidationEmail{}, middleware.After)
}

func newServiceMetadataMiddleware_opResendValidationEmail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm",
		OperationName: "ResendValidationEmail",
	}
}