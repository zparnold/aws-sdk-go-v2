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

// Updates the bundle of your Amazon Lightsail content delivery network (CDN)
// distribution.  <p>A distribution bundle specifies the monthly network transfer
// quota and monthly cost of your dsitribution.</p> <p>Update your distribution's
// bundle if your distribution is going over its monthly network transfer quota and
// is incurring an overage fee.</p> <p>You can update your distribution's bundle
// only one time within your monthly AWS billing cycle. To determine if you can
// update your distribution's bundle, use the <code>GetDistributions</code> action.
// The <code>ableToUpdateBundle</code> parameter in the result will indicate
// whether you can currently update your distribution's bundle.</p>
func (c *Client) UpdateDistributionBundle(ctx context.Context, params *UpdateDistributionBundleInput, optFns ...func(*Options)) (*UpdateDistributionBundleOutput, error) {
	stack := middleware.NewStack("UpdateDistributionBundle", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDistributionBundleMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDistributionBundle(options.Region), middleware.Before)
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
			OperationName: "UpdateDistributionBundle",
			Err:           err,
		}
	}
	out := result.(*UpdateDistributionBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDistributionBundleInput struct {

	// The bundle ID of the new bundle to apply to your distribution. Use the
	// GetDistributionBundles action to get a list of distribution bundle IDs that you
	// can specify.
	BundleId *string

	// The name of the distribution for which to update the bundle.  <p>Use the
	// <code>GetDistributions</code> action to get a list of distribution names that
	// you can specify.</p>
	DistributionName *string
}

type UpdateDistributionBundleOutput struct {

	// Describes the API operation.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDistributionBundleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDistributionBundle{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDistributionBundle{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDistributionBundle(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "UpdateDistributionBundle",
	}
}
