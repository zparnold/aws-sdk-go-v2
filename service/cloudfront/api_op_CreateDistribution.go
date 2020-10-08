// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new web distribution. You create a CloudFront distribution to tell
// CloudFront where you want content to be delivered from, and the details about
// how to track and manage content delivery. Send a POST request to the /CloudFront
// API version/distribution/distribution ID resource. When you update a
// distribution, there are more required fields than when you create a
// distribution. When you update your distribution by using UpdateDistribution
// (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html),
// follow the steps included in the documentation to get the current configuration
// and then make your updates. This helps to make sure that you include all of the
// required fields. To view a summary, see Required Fields for Create Distribution
// and Update Distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-overview-required-fields.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateDistribution(ctx context.Context, params *CreateDistributionInput, optFns ...func(*Options)) (*CreateDistributionOutput, error) {
	if params == nil {
		params = &CreateDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDistribution", params, optFns, addOperationCreateDistributionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create a new distribution.
type CreateDistributionInput struct {

	// The distribution's configuration information.
	//
	// This member is required.
	DistributionConfig *types.DistributionConfig
}

// The returned result of the corresponding request.
type CreateDistributionOutput struct {

	// The distribution's information.
	Distribution *types.Distribution

	// The current version of the distribution created.
	ETag *string

	// The fully qualified URI of the new distribution resource just created.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDistributionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateDistribution{}, middleware.After)
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
	addOpCreateDistributionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDistribution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDistribution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateDistribution",
	}
}
