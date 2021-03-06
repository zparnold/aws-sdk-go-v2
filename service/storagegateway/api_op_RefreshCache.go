// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Refreshes the cache for the specified file share. This operation finds objects
// in the Amazon S3 bucket that were added, removed, or replaced since the gateway
// last listed the bucket's contents and cached the results. This operation is only
// supported in the file gateway type. You can subscribe to be notified through an
// Amazon CloudWatch event when your RefreshCache operation completes. For more
// information, see Getting notified about file operations
// (https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-notification)
// in the AWS Storage Gateway User Guide.  <p>When this API is called, it only
// initiates the refresh operation. When the API call completes and returns a
// success code, it doesn't necessarily mean that the file refresh has completed.
// You should use the refresh-complete notification to determine that the operation
// has completed before you check for new files on the gateway file share. You can
// subscribe to be notified through an CloudWatch event when your
// <code>RefreshCache</code> operation completes.</p> <p>Throttle limit: This API
// is asynchronous so the gateway will accept no more than two refreshes at any
// time. We recommend using the refresh-complete CloudWatch event notification
// before issuing additional requests. For more information, see <a
// href="https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-notification">Getting
// notified about file operations</a> in the <i>AWS Storage Gateway User
// Guide</i>.</p> <p>If you invoke the RefreshCache API when two requests are
// already being processed, any new request will cause an
// <code>InvalidGatewayRequestException</code> error because too many requests were
// sent to the server.</p> <p>For more information, see <a
// href="https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-notification">Getting
// notified about file operations</a> in the <i>AWS Storage Gateway User
// Guide</i>.</p>
func (c *Client) RefreshCache(ctx context.Context, params *RefreshCacheInput, optFns ...func(*Options)) (*RefreshCacheOutput, error) {
	stack := middleware.NewStack("RefreshCache", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRefreshCacheMiddlewares(stack)
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
	addOpRefreshCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRefreshCache(options.Region), middleware.Before)
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
			OperationName: "RefreshCache",
			Err:           err,
		}
	}
	out := result.(*RefreshCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// RefreshCacheInput
type RefreshCacheInput struct {

	// The Amazon Resource Name (ARN) of the file share you want to refresh.
	//
	// This member is required.
	FileShareARN *string

	// A comma-separated list of the paths of folders to refresh in the cache. The
	// default is ["/"]. The default refreshes objects and folders at the root of the
	// Amazon S3 bucket. If Recursive is set to true, the entire S3 bucket that the
	// file share has access to is refreshed.
	FolderList []*string

	// A value that specifies whether to recursively refresh folders in the cache. The
	// refresh includes folders that were in the cache the last time the gateway listed
	// the folder's contents. If this value set to true, each folder that is listed in
	// FolderList is recursively updated. Otherwise, subfolders listed in FolderList
	// are not refreshed. Only objects that are in folders listed directly under
	// FolderList are found and used for the update. The default is true.  <p>Valid
	// Values: <code>true</code> | <code>false</code> </p>
	Recursive *bool
}

// RefreshCacheOutput
type RefreshCacheOutput struct {

	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string

	// The randomly generated ID of the notification that was sent. This ID is in UUID
	// format.
	NotificationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRefreshCacheMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRefreshCache{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRefreshCache{}, middleware.After)
}

func newServiceMetadataMiddleware_opRefreshCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "RefreshCache",
	}
}
