// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the status or configuration of a private certificate authority (CA).
// Your private CA must be in the ACTIVE or DISABLED state before you can update
// it. You can disable a private CA that is in the ACTIVE state or make a CA that
// is in the DISABLED state active again.
func (c *Client) UpdateCertificateAuthority(ctx context.Context, params *UpdateCertificateAuthorityInput, optFns ...func(*Options)) (*UpdateCertificateAuthorityOutput, error) {
	if params == nil {
		params = &UpdateCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCertificateAuthority", params, optFns, addOperationUpdateCertificateAuthorityMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCertificateAuthorityInput struct {

	// Amazon Resource Name (ARN) of the private CA that issued the certificate to be
	// revoked. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	// Revocation information for your private CA.
	RevocationConfiguration *types.RevocationConfiguration

	// Status of your private CA.
	Status types.CertificateAuthorityStatus
}

type UpdateCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateCertificateAuthorityMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCertificateAuthority{}, middleware.After)
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
	addOpUpdateCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCertificateAuthority(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "UpdateCertificateAuthority",
	}
}
