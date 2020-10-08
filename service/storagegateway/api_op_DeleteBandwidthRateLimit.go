// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the bandwidth rate limits of a gateway. You can delete either the upload
// and download bandwidth rate limit, or you can delete both. If you delete only
// one of the limits, the other limit remains unchanged. To specify which gateway
// to work with, use the Amazon Resource Name (ARN) of the gateway in your request.
// This operation is supported for the stored volume, cached volume and tape
// gateway types.
func (c *Client) DeleteBandwidthRateLimit(ctx context.Context, params *DeleteBandwidthRateLimitInput, optFns ...func(*Options)) (*DeleteBandwidthRateLimitOutput, error) {
	if params == nil {
		params = &DeleteBandwidthRateLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBandwidthRateLimit", params, optFns, addOperationDeleteBandwidthRateLimitMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBandwidthRateLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the following fields:  <ul> <li> <p>
// <a>DeleteBandwidthRateLimitInput$BandwidthType</a> </p> </li> </ul>
type DeleteBandwidthRateLimitInput struct {

	// One of the BandwidthType values that indicates the gateway bandwidth rate limit
	// to delete.  <p>Valid Values: <code>Upload</code> | <code>Download</code> |
	// <code>All</code> </p>
	//
	// This member is required.
	BandwidthType *string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway whose
// bandwidth rate information was deleted.
type DeleteBandwidthRateLimitOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBandwidthRateLimitMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBandwidthRateLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBandwidthRateLimit{}, middleware.After)
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
	addOpDeleteBandwidthRateLimitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBandwidthRateLimit(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBandwidthRateLimit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteBandwidthRateLimit",
	}
}
