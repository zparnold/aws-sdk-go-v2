// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a snapshot of a block storage disk. You can use snapshots for backups,
// to make copies of disks, and to save data before shutting down a Lightsail
// instance. You can take a snapshot of an attached disk that is in use; however,
// snapshots only capture data that has been written to your disk at the time the
// snapshot command is issued. This may exclude any data that has been cached by
// any applications or the operating system. If you can pause any file systems on
// the disk long enough to take a snapshot, your snapshot should be complete.
// Nevertheless, if you cannot pause all file writes to the disk, you should
// unmount the disk from within the Lightsail instance, issue the create disk
// snapshot command, and then remount the disk to ensure a consistent and complete
// snapshot. You may remount and use your disk while the snapshot status is
// pending. You can also use this operation to create a snapshot of an instance's
// system volume. You might want to do this, for example, to recover data from the
// system volume of a botched instance or to create a backup of the system volume
// like you would for a block storage disk. To create a snapshot of a system
// volume, just define the instance name parameter when issuing the snapshot
// command, and a snapshot of the defined instance's system volume will be created.
// After the snapshot is available, you can create a block storage disk from the
// snapshot and attach it to a running instance to access the data on the disk.
// <p>The <code>create disk snapshot</code> operation supports tag-based access
// control via request tags. For more information, see the <a
// href="https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags">Lightsail
// Dev Guide</a>.</p>
func (c *Client) CreateDiskSnapshot(ctx context.Context, params *CreateDiskSnapshotInput, optFns ...func(*Options)) (*CreateDiskSnapshotOutput, error) {
	stack := middleware.NewStack("CreateDiskSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDiskSnapshotMiddlewares(stack)
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
	addOpCreateDiskSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDiskSnapshot(options.Region), middleware.Before)
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
			OperationName: "CreateDiskSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateDiskSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDiskSnapshotInput struct {

	// The name of the destination disk snapshot (e.g., my-disk-snapshot) based on the
	// source disk.
	//
	// This member is required.
	DiskSnapshotName *string

	// The unique name of the source disk (e.g., Disk-Virginia-1). This parameter
	// cannot be defined together with the instance name parameter. The disk name and
	// instance name parameters are mutually exclusive.
	DiskName *string

	// The unique name of the source instance (e.g., Amazon_Linux-512MB-Virginia-1).
	// When this is defined, a snapshot of the instance's system volume is created.
	// This parameter cannot be defined together with the disk name parameter. The
	// instance name and disk name parameters are mutually exclusive.
	InstanceName *string

	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []*types.Tag
}

type CreateDiskSnapshotOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDiskSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDiskSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDiskSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDiskSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateDiskSnapshot",
	}
}
