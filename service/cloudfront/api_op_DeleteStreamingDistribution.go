// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete a streaming distribution. To delete an RTMP distribution using the
// CloudFront API, perform the following steps.  <p> <b>To delete an RTMP
// distribution using the CloudFront API</b>:</p> <ol> <li> <p>Disable the RTMP
// distribution.</p> </li> <li> <p>Submit a <code>GET Streaming Distribution
// Config</code> request to get the current configuration and the <code>Etag</code>
// header for the distribution. </p> </li> <li> <p>Update the XML document that was
// returned in the response to your <code>GET Streaming Distribution Config</code>
// request to change the value of <code>Enabled</code> to <code>false</code>.</p>
// </li> <li> <p>Submit a <code>PUT Streaming Distribution Config</code> request to
// update the configuration for your distribution. In the request body, include the
// XML document that you updated in Step 3. Then set the value of the HTTP
// <code>If-Match</code> header to the value of the <code>ETag</code> header that
// CloudFront returned when you submitted the <code>GET Streaming Distribution
// Config</code> request in Step 2.</p> </li> <li> <p>Review the response to the
// <code>PUT Streaming Distribution Config</code> request to confirm that the
// distribution was successfully disabled.</p> </li> <li> <p>Submit a <code>GET
// Streaming Distribution Config</code> request to confirm that your changes have
// propagated. When propagation is complete, the value of <code>Status</code> is
// <code>Deployed</code>.</p> </li> <li> <p>Submit a <code>DELETE Streaming
// Distribution</code> request. Set the value of the HTTP <code>If-Match</code>
// header to the value of the <code>ETag</code> header that CloudFront returned
// when you submitted the <code>GET Streaming Distribution Config</code> request in
// Step 2.</p> </li> <li> <p>Review the response to your <code>DELETE Streaming
// Distribution</code> request to confirm that the distribution was successfully
// deleted.</p> </li> </ol> <p>For information about deleting a distribution using
// the CloudFront console, see <a
// href="https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/HowToDeleteDistribution.html">Deleting
// a Distribution</a> in the <i>Amazon CloudFront Developer Guide</i>.</p>
func (c *Client) DeleteStreamingDistribution(ctx context.Context, params *DeleteStreamingDistributionInput, optFns ...func(*Options)) (*DeleteStreamingDistributionOutput, error) {
	if params == nil {
		params = &DeleteStreamingDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStreamingDistribution", params, optFns, addOperationDeleteStreamingDistributionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStreamingDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to delete a streaming distribution.
type DeleteStreamingDistributionInput struct {

	// The distribution ID.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when you disabled the streaming
	// distribution. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

type DeleteStreamingDistributionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteStreamingDistributionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteStreamingDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteStreamingDistribution{}, middleware.After)
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
	addOpDeleteStreamingDistributionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStreamingDistribution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteStreamingDistribution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteStreamingDistribution",
	}
}
