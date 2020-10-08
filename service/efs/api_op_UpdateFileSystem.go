// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates the throughput mode or the amount of provisioned throughput of an
// existing file system.
func (c *Client) UpdateFileSystem(ctx context.Context, params *UpdateFileSystemInput, optFns ...func(*Options)) (*UpdateFileSystemOutput, error) {
	if params == nil {
		params = &UpdateFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFileSystem", params, optFns, addOperationUpdateFileSystemMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFileSystemInput struct {

	// The ID of the file system that you want to update.
	//
	// This member is required.
	FileSystemId *string

	// (Optional) The amount of throughput, in MiB/s, that you want to provision for
	// your file system. Valid values are 1-1024. Required if ThroughputMode is changed
	// to provisioned on update. If you're not updating the amount of provisioned
	// throughput for your file system, you don't need to provide this value in your
	// request.
	ProvisionedThroughputInMibps *float64

	// (Optional) The throughput mode that you want your file system to use. If you're
	// not updating your throughput mode, you don't need to provide this value in your
	// request. If you are changing the ThroughputMode to provisioned, you must also
	// set a value for ProvisionedThroughputInMibps.
	ThroughputMode types.ThroughputMode
}

// A description of the file system.
type UpdateFileSystemOutput struct {

	// The time that the file system was created, in seconds (since
	// 1970-01-01T00:00:00Z).
	//
	// This member is required.
	CreationTime *time.Time

	// The opaque string specified in the request.
	//
	// This member is required.
	CreationToken *string

	// The ID of the file system, assigned by Amazon EFS.
	//
	// This member is required.
	FileSystemId *string

	// The lifecycle phase of the file system.
	//
	// This member is required.
	LifeCycleState types.LifeCycleState

	// The current number of mount targets that the file system has. For more
	// information, see CreateMountTarget ().
	//
	// This member is required.
	NumberOfMountTargets *int32

	// The AWS account that created the file system. If the file system was created by
	// an IAM user, the parent account to which the user belongs is the owner.
	//
	// This member is required.
	OwnerId *string

	// The performance mode of the file system.
	//
	// This member is required.
	PerformanceMode types.PerformanceMode

	// The latest known metered size (in bytes) of data stored in the file system, in
	// its Value field, and the time at which that size was determined in its Timestamp
	// field. The Timestamp value is the integer number of seconds since
	// 1970-01-01T00:00:00Z. The SizeInBytes value doesn't represent the size of a
	// consistent snapshot of the file system, but it is eventually consistent when
	// there are no writes to the file system. That is, SizeInBytes represents actual
	// size only if the file system is not modified for a period longer than a couple
	// of hours. Otherwise, the value is not the exact size that the file system was at
	// any point in time.
	//
	// This member is required.
	SizeInBytes *types.FileSystemSize

	// The tags associated with the file system, presented as an array of Tag objects.
	//
	// This member is required.
	Tags []*types.Tag

	// A Boolean value that, if true, indicates that the file system is encrypted.
	Encrypted *bool

	// The Amazon Resource Name (ARN) for the EFS file system, in the format
	// arn:aws:elasticfilesystem:region:account-id:file-system/file-system-id . Example
	// with sample data:
	// arn:aws:elasticfilesystem:us-west-2:1111333322228888:file-system/fs-01234567
	FileSystemArn *string

	// The ID of an AWS Key Management Service (AWS KMS) customer master key (CMK) that
	// was used to protect the encrypted file system.
	KmsKeyId *string

	// You can add tags to a file system, including a Name tag. For more information,
	// see CreateFileSystem (). If the file system has a Name tag, Amazon EFS returns
	// the value in this field.
	Name *string

	// The throughput, measured in MiB/s, that you want to provision for a file system.
	// Valid values are 1-1024. Required if ThroughputMode is set to provisioned. The
	// limit on throughput is 1024 MiB/s. You can get these limits increased by
	// contacting AWS Support. For more information, see Amazon EFS Limits That You Can
	// Increase (https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in
	// the Amazon EFS User Guide.
	ProvisionedThroughputInMibps *float64

	// The throughput mode for a file system. There are two throughput modes to choose
	// from for your file system: bursting and provisioned. If you set ThroughputMode
	// to provisioned, you must also set a value for ProvisionedThroughPutInMibps. You
	// can decrease your file system's throughput in Provisioned Throughput mode or
	// change between the throughput modes as long as it’s been more than 24 hours
	// since the last decrease or throughput mode change.
	ThroughputMode types.ThroughputMode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFileSystemMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFileSystem{}, middleware.After)
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
	addOpUpdateFileSystemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFileSystem(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateFileSystem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "UpdateFileSystem",
	}
}
