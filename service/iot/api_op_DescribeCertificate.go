// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the specified certificate.
func (c *Client) DescribeCertificate(ctx context.Context, params *DescribeCertificateInput, optFns ...func(*Options)) (*DescribeCertificateOutput, error) {
	if params == nil {
		params = &DescribeCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCertificate", params, optFns, addOperationDescribeCertificateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeCertificate operation.
type DescribeCertificateInput struct {

	// The ID of the certificate. (The last part of the certificate ARN contains the
	// certificate ID.)
	//
	// This member is required.
	CertificateId *string
}

// The output of the DescribeCertificate operation.
type DescribeCertificateOutput struct {

	// The description of the certificate.
	CertificateDescription *types.CertificateDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCertificateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCertificate{}, middleware.After)
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
	addOpDescribeCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeCertificate",
	}
}
