// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an audit report that lists every time that your CA private key is used.
// The report is saved in the Amazon S3 bucket that you specify on input. The
// IssueCertificate () and RevokeCertificate () actions use the private key.
func (c *Client) CreateCertificateAuthorityAuditReport(ctx context.Context, params *CreateCertificateAuthorityAuditReportInput, optFns ...func(*Options)) (*CreateCertificateAuthorityAuditReportOutput, error) {
	stack := middleware.NewStack("CreateCertificateAuthorityAuditReport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateCertificateAuthorityAuditReportMiddlewares(stack)
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
	addOpCreateCertificateAuthorityAuditReportValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCertificateAuthorityAuditReport(options.Region), middleware.Before)
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
			OperationName: "CreateCertificateAuthorityAuditReport",
			Err:           err,
		}
	}
	out := result.(*CreateCertificateAuthorityAuditReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCertificateAuthorityAuditReportInput struct {

	// The format in which to create the report. This can be either JSON or CSV.
	//
	// This member is required.
	AuditReportResponseFormat types.AuditReportResponseFormat

	// The Amazon Resource Name (ARN) of the CA to be audited. This is of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string

	// The name of the S3 bucket that will contain the audit report.
	//
	// This member is required.
	S3BucketName *string
}

type CreateCertificateAuthorityAuditReportOutput struct {

	// An alphanumeric string that contains a report identifier.
	AuditReportId *string

	// The key that uniquely identifies the report file in your S3 bucket.
	S3Key *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateCertificateAuthorityAuditReportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCertificateAuthorityAuditReport{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCertificateAuthorityAuditReport{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCertificateAuthorityAuditReport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "CreateCertificateAuthorityAuditReport",
	}
}
