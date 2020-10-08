// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the tags from the bucket.  <p>To use this operation, you must have
// permission to perform the <code>s3:PutBucketTagging</code> action. By default,
// the bucket owner has this permission and can grant this permission to others.
// </p> <p>The following operations are related to
// <code>DeleteBucketTagging</code>:</p> <ul> <li> <p> <a>GetBucketTagging</a> </p>
// </li> <li> <p> <a>PutBucketTagging</a> </p> </li> </ul>
func (c *Client) DeleteBucketTagging(ctx context.Context, params *DeleteBucketTaggingInput, optFns ...func(*Options)) (*DeleteBucketTaggingOutput, error) {
	if params == nil {
		params = &DeleteBucketTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketTagging", params, optFns, addOperationDeleteBucketTaggingMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketTaggingInput struct {

	// The bucket that has the tag set to be removed.
	//
	// This member is required.
	Bucket *string
}

type DeleteBucketTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBucketTaggingMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketTagging{}, middleware.After)
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
	addOpDeleteBucketTaggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketTagging(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBucketTagging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketTagging",
	}
}
