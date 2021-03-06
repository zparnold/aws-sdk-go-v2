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

// Cancels a pending transfer for the specified certificate. Note Only the transfer
// source account can use this operation to cancel a transfer. (Transfer
// destinations can use RejectCertificateTransfer () instead.) After transfer, AWS
// IoT returns the certificate to the source account in the INACTIVE state. After
// the destination account has accepted the transfer, the transfer cannot be
// cancelled. After a certificate transfer is cancelled, the status of the
// certificate changes from PENDING_TRANSFER to INACTIVE.
func (c *Client) CancelCertificateTransfer(ctx context.Context, params *CancelCertificateTransferInput, optFns ...func(*Options)) (*CancelCertificateTransferOutput, error) {
	stack := middleware.NewStack("CancelCertificateTransfer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCancelCertificateTransferMiddlewares(stack)
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
	addOpCancelCertificateTransferValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelCertificateTransfer(options.Region), middleware.Before)
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
			OperationName: "CancelCertificateTransfer",
			Err:           err,
		}
	}
	out := result.(*CancelCertificateTransferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CancelCertificateTransfer operation.
type CancelCertificateTransferInput struct {

	// The ID of the certificate. (The last part of the certificate ARN contains the
	// certificate ID.)
	//
	// This member is required.
	CertificateId *string
}

type CancelCertificateTransferOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCancelCertificateTransferMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCancelCertificateTransfer{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelCertificateTransfer{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelCertificateTransfer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CancelCertificateTransfer",
	}
}
