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

// Returns the versioning state of a bucket. To retrieve the versioning state of a
// bucket, you must be the bucket owner.  <p>This implementation also returns the
// MFA Delete status of the versioning state. If the MFA Delete status is
// <code>enabled</code>, the bucket owner must use an authentication device to
// change the versioning state of the bucket.</p> <p>The following operations are
// related to <code>GetBucketVersioning</code>:</p> <ul> <li> <p> <a>GetObject</a>
// </p> </li> <li> <p> <a>PutObject</a> </p> </li> <li> <p> <a>DeleteObject</a>
// </p> </li> </ul>
func (c *Client) GetBucketVersioning(ctx context.Context, params *GetBucketVersioningInput, optFns ...func(*Options)) (*GetBucketVersioningOutput, error) {
	if params == nil {
		params = &GetBucketVersioningInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketVersioning", params, optFns, addOperationGetBucketVersioningMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketVersioningOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketVersioningInput struct {

	// The name of the bucket for which to get the versioning information.
	//
	// This member is required.
	Bucket *string
}

type GetBucketVersioningOutput struct {

	// Specifies whether MFA delete is enabled in the bucket versioning configuration.
	// This element is only returned if the bucket has been configured with MFA delete.
	// If the bucket has never been so configured, this element is not returned.
	MFADelete types.MFADeleteStatus

	// The versioning state of the bucket.
	Status types.BucketVersioningStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketVersioningMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketVersioning{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketVersioning{}, middleware.After)
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
	addOpGetBucketVersioningValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketVersioning(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketVersioning(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketVersioning",
	}
}
