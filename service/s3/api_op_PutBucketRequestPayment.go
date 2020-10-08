// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the request payment configuration for a bucket. By default, the bucket
// owner pays for downloads from the bucket. This configuration parameter enables
// the bucket owner (only) to specify that the person requesting the download will
// be charged for the download. For more information, see Requester Pays Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html).
// <p>The following operations are related to
// <code>PutBucketRequestPayment</code>:</p> <ul> <li> <p> <a>CreateBucket</a> </p>
// </li> <li> <p> <a>GetBucketRequestPayment</a> </p> </li> </ul>
func (c *Client) PutBucketRequestPayment(ctx context.Context, params *PutBucketRequestPaymentInput, optFns ...func(*Options)) (*PutBucketRequestPaymentOutput, error) {
	if params == nil {
		params = &PutBucketRequestPaymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketRequestPayment", params, optFns, addOperationPutBucketRequestPaymentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketRequestPaymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketRequestPaymentInput struct {

	// The bucket name.
	//
	// This member is required.
	Bucket *string

	// Container for Payer.
	//
	// This member is required.
	RequestPaymentConfiguration *types.RequestPaymentConfiguration

	// >The base64-encoded 128-bit MD5 digest of the data. You must use this header as
	// a message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see RFC 1864
	// (http://www.ietf.org/rfc/rfc1864.txt).
	ContentMD5 *string
}

type PutBucketRequestPaymentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketRequestPaymentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketRequestPayment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketRequestPayment{}, middleware.After)
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
	addOpPutBucketRequestPaymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketRequestPayment(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutBucketRequestPayment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketRequestPayment",
	}
}
