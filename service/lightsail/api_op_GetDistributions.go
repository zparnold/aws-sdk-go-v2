// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about one or more of your Amazon Lightsail content delivery
// network (CDN) distributions.
func (c *Client) GetDistributions(ctx context.Context, params *GetDistributionsInput, optFns ...func(*Options)) (*GetDistributionsOutput, error) {
	if params == nil {
		params = &GetDistributionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDistributions", params, optFns, addOperationGetDistributionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDistributionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDistributionsInput struct {

	// The name of the distribution for which to return information.  <p>Use the
	// <code>GetDistributions</code> action to get a list of distribution names that
	// you can specify.</p> <p>When omitted, the response includes all of your
	// distributions in the AWS Region where the request is made.</p>
	DistributionName *string

	// The token to advance to the next page of results from your request.  <p>To get a
	// page token, perform an initial <code>GetDistributions</code> request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.</p>
	PageToken *string
}

type GetDistributionsOutput struct {

	// An array of objects that describe your distributions.
	Distributions []*types.LightsailDistribution

	// The token to advance to the next page of results from your request.  <p>A next
	// page token is not returned if there are no more results to display.</p> <p>To
	// get the next page of results, perform another <code>GetDistributions</code>
	// request and specify the next page token using the <code>pageToken</code>
	// parameter.</p>
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDistributionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDistributions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDistributions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDistributions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDistributions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetDistributions",
	}
}
