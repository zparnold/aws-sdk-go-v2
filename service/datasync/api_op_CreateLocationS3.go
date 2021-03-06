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
)

// Creates an endpoint for an Amazon S3 bucket. For AWS DataSync to access a
// destination S3 bucket, it needs an AWS Identity and Access Management (IAM) role
// that has the required permissions. You can set up the required permissions by
// creating an IAM policy that grants the required permissions and attaching the
// policy to the role. An example of such a policy is shown in the examples
// section.  <p>For more information, see
// https://docs.aws.amazon.com/datasync/latest/userguide/working-with-locations.html#create-s3-location
// in the <i>AWS DataSync User Guide.</i> </p>
func (c *Client) CreateLocationS3(ctx context.Context, params *CreateLocationS3Input, optFns ...func(*Options)) (*CreateLocationS3Output, error) {
	stack := middleware.NewStack("CreateLocationS3", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateLocationS3Middlewares(stack)
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
	addOpCreateLocationS3ValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationS3(options.Region), middleware.Before)
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
			OperationName: "CreateLocationS3",
			Err:           err,
		}
	}
	out := result.(*CreateLocationS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationS3Request
type CreateLocationS3Input struct {

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket.
	//
	// This member is required.
	S3BucketArn *string

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that is used to access an Amazon S3 bucket.  <p>For detailed information
	// about using such a role, see Creating a Location for Amazon S3 in the <i>AWS
	// DataSync User Guide</i>.</p>
	//
	// This member is required.
	S3Config *types.S3Config

	// The Amazon S3 storage class that you want to store your files in when this
	// location is used as a task destination. For more information about S3 storage
	// classes, see Amazon S3 Storage Classes
	// (https://aws.amazon.com/s3/storage-classes/) in the Amazon Simple Storage
	// Service Developer Guide. Some storage classes have behaviors that can affect
	// your S3 storage cost. For detailed information, see using-storage-classes ().
	S3StorageClass types.S3StorageClass

	// A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is used
	// to read data from the S3 source location or write data to the S3 destination.
	Subdirectory *string

	// The key-value pair that represents the tag that you want to add to the location.
	// The value can be an empty string. We recommend using tags to name your
	// resources.
	Tags []*types.TagListEntry
}

// CreateLocationS3Response
type CreateLocationS3Output struct {

	// The Amazon Resource Name (ARN) of the source Amazon S3 bucket location that is
	// created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateLocationS3Middlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationS3{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationS3{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLocationS3(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationS3",
	}
}
