// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the bandwidth rate limits of a gateway. By default, these limits are not
// set, which means no bandwidth rate limiting is in effect. This operation is
// supported for the stored volume, cached volume and tape gateway types.  <p>This
// operation only returns a value for a bandwidth rate limit only if the limit is
// set. If no limits are set for the gateway, then this operation returns only the
// gateway ARN in the response body. To specify which gateway to describe, use the
// Amazon Resource Name (ARN) of the gateway in your request.</p>
func (c *Client) DescribeBandwidthRateLimit(ctx context.Context, params *DescribeBandwidthRateLimitInput, optFns ...func(*Options)) (*DescribeBandwidthRateLimitOutput, error) {
	if params == nil {
		params = &DescribeBandwidthRateLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBandwidthRateLimit", params, optFns, addOperationDescribeBandwidthRateLimitMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBandwidthRateLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway.
type DescribeBandwidthRateLimitInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

// A JSON object containing the following fields:
type DescribeBandwidthRateLimitOutput struct {

	// The average download bandwidth rate limit in bits per second. This field does
	// not appear in the response if the download rate limit is not set.
	AverageDownloadRateLimitInBitsPerSec *int64

	// The average upload bandwidth rate limit in bits per second. This field does not
	// appear in the response if the upload rate limit is not set.
	AverageUploadRateLimitInBitsPerSec *int64

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBandwidthRateLimitMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBandwidthRateLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBandwidthRateLimit{}, middleware.After)
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
	addOpDescribeBandwidthRateLimitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBandwidthRateLimit(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBandwidthRateLimit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeBandwidthRateLimit",
	}
}
