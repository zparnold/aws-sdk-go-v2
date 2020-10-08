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

// Creates or modifies the PublicAccessBlock configuration for an Amazon S3 bucket.
// To use this operation, you must have the s3:PutBucketPublicAccessBlock
// permission. For more information about Amazon S3 permissions, see Specifying
// Permissions in a Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html).
// <important> <p>When Amazon S3 evaluates the <code>PublicAccessBlock</code>
// configuration for a bucket or an object, it checks the
// <code>PublicAccessBlock</code> configuration for both the bucket (or the bucket
// that contains the object) and the bucket owner's account. If the
// <code>PublicAccessBlock</code> configurations are different between the bucket
// and the account, Amazon S3 uses the most restrictive combination of the
// bucket-level and account-level settings.</p> </important> <p>For more
// information about when Amazon S3 considers a bucket or an object public, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status">The
// Meaning of "Public"</a>.</p> <p class="title"> <b>Related Resources</b> </p>
// <ul> <li> <p> <a>GetPublicAccessBlock</a> </p> </li> <li> <p>
// <a>DeletePublicAccessBlock</a> </p> </li> <li> <p> <a>GetBucketPolicyStatus</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html">Using
// Amazon S3 Block Public Access</a> </p> </li> </ul>
func (c *Client) PutPublicAccessBlock(ctx context.Context, params *PutPublicAccessBlockInput, optFns ...func(*Options)) (*PutPublicAccessBlockOutput, error) {
	if params == nil {
		params = &PutPublicAccessBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPublicAccessBlock", params, optFns, addOperationPutPublicAccessBlockMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPublicAccessBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPublicAccessBlockInput struct {

	// The name of the Amazon S3 bucket whose PublicAccessBlock configuration you want
	// to set.
	//
	// This member is required.
	Bucket *string

	// The PublicAccessBlock configuration that you want to apply to this Amazon S3
	// bucket. You can enable the configuration options in any combination. For more
	// information about when Amazon S3 considers a bucket or object public, see The
	// Meaning of "Public"
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// The MD5 hash of the PutPublicAccessBlock request body.
	ContentMD5 *string
}

type PutPublicAccessBlockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutPublicAccessBlockMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutPublicAccessBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutPublicAccessBlock{}, middleware.After)
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
	addOpPutPublicAccessBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutPublicAccessBlock(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutPublicAccessBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutPublicAccessBlock",
	}
}
