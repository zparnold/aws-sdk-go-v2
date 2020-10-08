// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets all cache disks that have encountered an error and makes the disks
// available for reconfiguration as cache storage. If your cache disk encounters an
// error, the gateway prevents read and write operations on virtual tapes in the
// gateway. For example, an error can occur when a disk is corrupted or removed
// from the gateway. When a cache is reset, the gateway loses its cache storage. At
// this point, you can reconfigure the disks as cache disks. This operation is only
// supported in the cached volume and tape types.  <important> <p>If the cache disk
// you are resetting contains data that has not been uploaded to Amazon S3 yet,
// that data can be lost. After you reset cache disks, there will be no configured
// cache disks left in the gateway, so you must configure at least one new cache
// disk for your gateway to function properly.</p> </important>
func (c *Client) ResetCache(ctx context.Context, params *ResetCacheInput, optFns ...func(*Options)) (*ResetCacheOutput, error) {
	if params == nil {
		params = &ResetCacheInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetCache", params, optFns, addOperationResetCacheMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetCacheInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

type ResetCacheOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResetCacheMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResetCache{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetCache{}, middleware.After)
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
	addOpResetCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetCache(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResetCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ResetCache",
	}
}
