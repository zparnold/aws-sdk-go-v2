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

// Returns the request payment configuration of a bucket. To use this version of
// the operation, you must be the bucket owner. For more information, see Requester
// Pays Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html).
// <p>The following operations are related to
// <code>GetBucketRequestPayment</code>:</p> <ul> <li> <p> <a>ListObjects</a> </p>
// </li> </ul>
func (c *Client) GetBucketRequestPayment(ctx context.Context, params *GetBucketRequestPaymentInput, optFns ...func(*Options)) (*GetBucketRequestPaymentOutput, error) {
	if params == nil {
		params = &GetBucketRequestPaymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketRequestPayment", params, optFns, addOperationGetBucketRequestPaymentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketRequestPaymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketRequestPaymentInput struct {

	// The name of the bucket for which to get the payment request configuration
	//
	// This member is required.
	Bucket *string
}

type GetBucketRequestPaymentOutput struct {

	// Specifies who pays for the download and request fees.
	Payer types.Payer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketRequestPaymentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketRequestPayment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketRequestPayment{}, middleware.After)
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
	addOpGetBucketRequestPaymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketRequestPayment(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketRequestPayment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketRequestPayment",
	}
}
