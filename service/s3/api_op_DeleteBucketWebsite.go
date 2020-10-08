// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation removes the website configuration for a bucket. Amazon S3 returns
// a 200 OK response upon successfully deleting a website configuration on the
// specified bucket. You will get a 200 OK response if the website configuration
// you are trying to delete does not exist on the bucket. Amazon S3 returns a 404
// response if the bucket specified in the request does not exist.  <p>This DELETE
// operation requires the <code>S3:DeleteBucketWebsite</code> permission. By
// default, only the bucket owner can delete the website configuration attached to
// a bucket. However, bucket owners can grant other users permission to delete the
// website configuration by writing a bucket policy granting them the
// <code>S3:DeleteBucketWebsite</code> permission. </p> <p>For more information
// about hosting websites, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html">Hosting
// Websites on Amazon S3</a>. </p> <p>The following operations are related to
// <code>DeleteBucketWebsite</code>:</p> <ul> <li> <p> <a>GetBucketWebsite</a> </p>
// </li> <li> <p> <a>PutBucketWebsite</a> </p> </li> </ul>
func (c *Client) DeleteBucketWebsite(ctx context.Context, params *DeleteBucketWebsiteInput, optFns ...func(*Options)) (*DeleteBucketWebsiteOutput, error) {
	if params == nil {
		params = &DeleteBucketWebsiteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketWebsite", params, optFns, addOperationDeleteBucketWebsiteMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketWebsiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketWebsiteInput struct {

	// The bucket name for which you want to remove the website configuration.
	//
	// This member is required.
	Bucket *string
}

type DeleteBucketWebsiteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBucketWebsiteMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketWebsite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketWebsite{}, middleware.After)
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
	addOpDeleteBucketWebsiteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketWebsite(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBucketWebsite(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketWebsite",
	}
}
