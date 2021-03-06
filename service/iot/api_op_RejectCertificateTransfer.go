// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Rejects a pending certificate transfer. After AWS IoT rejects a certificate
// transfer, the certificate status changes from PENDING_TRANSFER to INACTIVE. To
// check for pending certificate transfers, call ListCertificates () to enumerate
// your certificates. This operation can only be called by the transfer
// destination. After it is called, the certificate will be returned to the
// source's account in the INACTIVE state.
func (c *Client) RejectCertificateTransfer(ctx context.Context, params *RejectCertificateTransferInput, optFns ...func(*Options)) (*RejectCertificateTransferOutput, error) {
	stack := middleware.NewStack("RejectCertificateTransfer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRejectCertificateTransferMiddlewares(stack)
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
	addOpRejectCertificateTransferValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRejectCertificateTransfer(options.Region), middleware.Before)
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
			OperationName: "RejectCertificateTransfer",
			Err:           err,
		}
	}
	out := result.(*RejectCertificateTransferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the RejectCertificateTransfer operation.
type RejectCertificateTransferInput struct {

	// The ID of the certificate. (The last part of the certificate ARN contains the
	// certificate ID.)
	//
	// This member is required.
	CertificateId *string

	// The reason the certificate transfer was rejected.
	RejectReason *string
}

type RejectCertificateTransferOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRejectCertificateTransferMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRejectCertificateTransfer{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRejectCertificateTransfer{}, middleware.After)
}

func newServiceMetadataMiddleware_opRejectCertificateTransfer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RejectCertificateTransfer",
	}
}
