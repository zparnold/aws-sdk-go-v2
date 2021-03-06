// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata, such as bucket name, about an Amazon S3 bucket location.
func (c *Client) DescribeLocationS3(ctx context.Context, params *DescribeLocationS3Input, optFns ...func(*Options)) (*DescribeLocationS3Output, error) {
	stack := middleware.NewStack("DescribeLocationS3", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeLocationS3Middlewares(stack)
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
	addOpDescribeLocationS3ValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationS3(options.Region), middleware.Before)
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
			OperationName: "DescribeLocationS3",
			Err:           err,
		}
	}
	out := result.(*DescribeLocationS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationS3Request
type DescribeLocationS3Input struct {

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket location to describe.
	//
	// This member is required.
	LocationArn *string
}

// DescribeLocationS3Response
type DescribeLocationS3Output struct {

	// The time that the Amazon S3 bucket location was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket location.
	LocationArn *string

	// The URL of the Amazon S3 location that was described.
	LocationUri *string

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that is used to access an Amazon S3 bucket.  <p>For detailed information
	// about using such a role, see Creating a Location for Amazon S3 in the <i>AWS
	// DataSync User Guide</i>.</p>
	S3Config *types.S3Config

	// The Amazon S3 storage class that you chose to store your files in when this
	// location is used as a task destination. For more information about S3 storage
	// classes, see Amazon S3 Storage Classes
	// (https://aws.amazon.com/s3/storage-classes/) in the Amazon Simple Storage
	// Service Developer Guide. Some storage classes have behaviors that can affect
	// your S3 storage cost. For detailed information, see using-storage-classes ().
	S3StorageClass types.S3StorageClass

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeLocationS3Middlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationS3{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationS3{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLocationS3(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationS3",
	}
}
