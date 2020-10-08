// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a certificate from your private CA. The ARN of the certificate is
// returned when you call the IssueCertificate () action. You must specify both the
// ARN of your private CA and the ARN of the issued certificate when calling the
// GetCertificate action. You can retrieve the certificate if it is in the ISSUED
// state. You can call the CreateCertificateAuthorityAuditReport () action to
// create a report that contains information about all of the certificates issued
// and revoked by your private CA.
func (c *Client) GetCertificate(ctx context.Context, params *GetCertificateInput, optFns ...func(*Options)) (*GetCertificateOutput, error) {
	if params == nil {
		params = &GetCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCertificate", params, optFns, addOperationGetCertificateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificateInput struct {

	// The ARN of the issued certificate. The ARN contains the certificate serial
	// number and must be in the following form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/286535153982981100925020015808220737245
	//
	// This member is required.
	CertificateArn *string

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority (). This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string
}

type GetCertificateOutput struct {

	// The base64 PEM-encoded certificate specified by the CertificateArn parameter.
	Certificate *string

	// The base64 PEM-encoded certificate chain that chains up to the on-premises root
	// CA certificate that you used to sign your private CA certificate.
	CertificateChain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCertificateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificate{}, middleware.After)
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
	addOpGetCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "GetCertificate",
	}
}
