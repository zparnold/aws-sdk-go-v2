// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This API operation deletes an Amazon S3 on Outposts bucket's tags. To delete an
// S3 bucket tags, see DeleteBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html)
// in the Amazon Simple Storage Service API. Deletes the tags from the Outposts
// bucket. For more information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in Amazon
// Simple Storage Service Developer Guide.  <p>To use this operation, you must have
// permission to perform the <code>PutBucketTagging</code> action. By default, the
// bucket owner has this permission and can grant this permission to others. </p>
// <p>All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of outpost-id to be passed with the request and an S3 on
// Outposts endpoint hostname prefix instead of s3-control. For an example of the
// request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint
// hostname prefix and the outpost-id derived using the access point ARN, see the
// <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteBucketTagging.html#API_control_DeleteBucketTagging_Examples">
// Example</a> section below.</p> <p>The following actions are related to
// <code>DeleteBucketTagging</code>:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html">GetBucketTagging</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html">PutBucketTagging</a>
// </p> </li> </ul>
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

	// The AWS account ID of the Outposts bucket tag set to be removed.
	//
	// This member is required.
	AccountId *string

	// The bucket ARN that has the tag set to be removed. For Amazon S3 on Outposts
	// specify the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/. For example, to access the bucket
	// reports through outpost my-outpost owned by account 123456789012 in Region
	// us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
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
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
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
