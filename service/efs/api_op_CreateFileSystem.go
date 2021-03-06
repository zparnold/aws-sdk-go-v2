// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new, empty file system. The operation requires a creation token in the
// request that Amazon EFS uses to ensure idempotent creation (calling the
// operation with same creation token has no effect). If a file system does not
// currently exist that is owned by the caller's AWS account with the specified
// creation token, this operation does the following:
//
//     * Creates a new, empty
// file system. The file system will have an Amazon EFS assigned ID, and an initial
// lifecycle state creating.
//
//     * Returns with the description of the created
// file system.
//
// Otherwise, this operation returns a FileSystemAlreadyExists error
// with the ID of the existing file system. For basic use cases, you can use a
// randomly generated UUID for the creation token. The idempotent operation allows
// you to retry a CreateFileSystem call without risk of creating an extra file
// system. This can happen when an initial call fails in a way that leaves it
// uncertain whether or not a file system was actually created. An example might be
// that a transport level timeout occurred or your connection was reset. As long as
// you use the same creation token, if the initial call had succeeded in creating a
// file system, the client can learn of its existence from the
// FileSystemAlreadyExists error.  <note> <p>The <code>CreateFileSystem</code> call
// returns while the file system's lifecycle state is still <code>creating</code>.
// You can check the file system creation status by calling the
// <a>DescribeFileSystems</a> operation, which among other things returns the file
// system state.</p> </note> <p>This operation also takes an optional
// <code>PerformanceMode</code> parameter that you choose for your file system. We
// recommend <code>generalPurpose</code> performance mode for most file systems.
// File systems using the <code>maxIO</code> performance mode can scale to higher
// levels of aggregate throughput and operations per second with a tradeoff of
// slightly higher latencies for most file operations. The performance mode can't
// be changed after the file system has been created. For more information, see <a
// href="https://docs.aws.amazon.com/efs/latest/ug/performance.html#performancemodes.html">Amazon
// EFS: Performance Modes</a>.</p> <p>After the file system is fully created,
// Amazon EFS sets its lifecycle state to <code>available</code>, at which point
// you can create one or more mount targets for the file system in your VPC. For
// more information, see <a>CreateMountTarget</a>. You mount your Amazon EFS file
// system on an EC2 instances in your VPC by using the mount target. For more
// information, see <a
// href="https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html">Amazon EFS:
// How it Works</a>. </p> <p> This operation requires permissions for the
// <code>elasticfilesystem:CreateFileSystem</code> action. </p>
func (c *Client) CreateFileSystem(ctx context.Context, params *CreateFileSystemInput, optFns ...func(*Options)) (*CreateFileSystemOutput, error) {
	stack := middleware.NewStack("CreateFileSystem", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateFileSystemMiddlewares(stack)
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
	addIdempotencyToken_opCreateFileSystemMiddleware(stack, options)
	addOpCreateFileSystemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFileSystem(options.Region), middleware.Before)
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
			OperationName: "CreateFileSystem",
			Err:           err,
		}
	}
	out := result.(*CreateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFileSystemInput struct {

	// A string of up to 64 ASCII characters. Amazon EFS uses this to ensure idempotent
	// creation.
	//
	// This member is required.
	CreationToken *string

	// A Boolean value that, if true, creates an encrypted file system. When creating
	// an encrypted file system, you have the option of specifying
	// CreateFileSystemRequest$KmsKeyId () for an existing AWS Key Management Service
	// (AWS KMS) customer master key (CMK). If you don't specify a CMK, then the
	// default CMK for Amazon EFS, /aws/elasticfilesystem, is used to protect the
	// encrypted file system.
	Encrypted *bool

	// The ID of the AWS KMS CMK to be used to protect the encrypted file system. This
	// parameter is only required if you want to use a nondefault CMK. If this
	// parameter is not specified, the default CMK for Amazon EFS is used. This ID can
	// be in one of the following formats:
	//
	//     * Key ID - A unique identifier of the
	// key, for example 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//     * ARN - An Amazon
	// Resource Name (ARN) for the key, for example
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//
	// * Key alias - A previously created display name for a key, for example
	// alias/projectKey1.
	//
	//     * Key alias ARN - An ARN for a key alias, for example
	// arn:aws:kms:us-west-2:444455556666:alias/projectKey1.
	//
	// If KmsKeyId is specified,
	// the CreateFileSystemRequest$Encrypted () parameter must be set to true. EFS
	// accepts only symmetric CMKs. You cannot use asymmetric CMKs with EFS file
	// systems.
	KmsKeyId *string

	// The performance mode of the file system. We recommend generalPurpose performance
	// mode for most file systems. File systems using the maxIO performance mode can
	// scale to higher levels of aggregate throughput and operations per second with a
	// tradeoff of slightly higher latencies for most file operations. The performance
	// mode can't be changed after the file system has been created.
	PerformanceMode types.PerformanceMode

	// The throughput, measured in MiB/s, that you want to provision for a file system
	// that you're creating. Valid values are 1-1024. Required if ThroughputMode is set
	// to provisioned. The upper limit for throughput is 1024 MiB/s. You can get this
	// limit increased by contacting AWS Support. For more information, see Amazon EFS
	// Limits That You Can Increase
	// (https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the
	// Amazon EFS User Guide.
	ProvisionedThroughputInMibps *float64

	// A value that specifies to create one or more tags associated with the file
	// system. Each tag is a user-defined key-value pair. Name your file system on
	// creation by including a "Key":"Name","Value":"{value}" key-value pair.
	Tags []*types.Tag

	// The throughput mode for the file system to be created. There are two throughput
	// modes to choose from for your file system: bursting and provisioned. If you set
	// ThroughputMode to provisioned, you must also set a value for
	// ProvisionedThroughPutInMibps. You can decrease your file system's throughput in
	// Provisioned Throughput mode or change between the throughput modes as long as
	// it’s been more than 24 hours since the last decrease or throughput mode change.
	// For more, see Specifying Throughput with Provisioned Mode
	// (https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput)
	// in the Amazon EFS User Guide.
	ThroughputMode types.ThroughputMode
}

// A description of the file system.
type CreateFileSystemOutput struct {

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

func addawsRestjson1_serdeOpCreateFileSystemMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateFileSystem{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFileSystem{}, middleware.After)
}

type idempotencyToken_initializeOpCreateFileSystem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFileSystem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFileSystemInput ")
	}

	if input.CreationToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreationToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFileSystemMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateFileSystem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFileSystem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "CreateFileSystem",
	}
}
