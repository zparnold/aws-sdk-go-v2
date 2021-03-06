// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the metrics configurations for the bucket. The metrics configurations are
// only for the request metrics of the bucket and do not provide information on
// daily storage metrics. You can have up to 1,000 configurations per bucket.
// <p>This operation supports list pagination and does not return more than 100
// configurations at a time. Always check the <code>IsTruncated</code> element in
// the response. If there are no more configurations to list,
// <code>IsTruncated</code> is set to false. If there are more configurations to
// list, <code>IsTruncated</code> is set to true, and there is a value in
// <code>NextContinuationToken</code>. You use the
// <code>NextContinuationToken</code> value to continue the pagination of the list
// by passing the value in <code>continuation-token</code> in the request to
// <code>GET</code> the next page.</p> <p>To use this operation, you must have
// permissions to perform the <code>s3:GetMetricsConfiguration</code> action. The
// bucket owner has this permission by default. The bucket owner can grant this
// permission to others. For more information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p>For more information
// about metrics configurations and CloudWatch request metrics, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html">Monitoring
// Metrics with Amazon CloudWatch</a>.</p> <p>The following operations are related
// to <code>ListBucketMetricsConfigurations</code>:</p> <ul> <li> <p>
// <a>PutBucketMetricsConfiguration</a> </p> </li> <li> <p>
// <a>GetBucketMetricsConfiguration</a> </p> </li> <li> <p>
// <a>DeleteBucketMetricsConfiguration</a> </p> </li> </ul>
func (c *Client) ListBucketMetricsConfigurations(ctx context.Context, params *ListBucketMetricsConfigurationsInput, optFns ...func(*Options)) (*ListBucketMetricsConfigurationsOutput, error) {
	stack := middleware.NewStack("ListBucketMetricsConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListBucketMetricsConfigurationsMiddlewares(stack)
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
	addOpListBucketMetricsConfigurationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBucketMetricsConfigurations(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

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
			OperationName: "ListBucketMetricsConfigurations",
			Err:           err,
		}
	}
	out := result.(*ListBucketMetricsConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketMetricsConfigurationsInput struct {

	// The name of the bucket containing the metrics configurations to retrieve.
	//
	// This member is required.
	Bucket *string

	// The marker that is used to continue a metrics configuration listing that has
	// been truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value that
	// Amazon S3 understands.
	ContinuationToken *string
}

type ListBucketMetricsConfigurationsOutput struct {

	// The marker that is used as a starting point for this metrics configuration list
	// response. This value is present if it was sent in the request.
	ContinuationToken *string

	// Indicates whether the returned list of metrics configurations is complete. A
	// value of true indicates that the list is not complete and the
	// NextContinuationToken will be provided for a subsequent request.
	IsTruncated *bool

	// The list of metrics configurations for a bucket.
	MetricsConfigurationList []*types.MetricsConfiguration

	// The marker used to continue a metrics configuration listing that has been
	// truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value that
	// Amazon S3 understands.
	NextContinuationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListBucketMetricsConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListBucketMetricsConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListBucketMetricsConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBucketMetricsConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListBucketMetricsConfigurations",
	}
}
