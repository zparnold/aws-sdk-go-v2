// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details of the specified configuration set. For information about
// using configuration sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
// You can execute this operation no more than once per second.
func (c *Client) DescribeConfigurationSet(ctx context.Context, params *DescribeConfigurationSetInput, optFns ...func(*Options)) (*DescribeConfigurationSetOutput, error) {
	if params == nil {
		params = &DescribeConfigurationSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationSet", params, optFns, addOperationDescribeConfigurationSetMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the details of a configuration set. Configuration
// sets enable you to publish email sending events. For information about using
// configuration sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type DescribeConfigurationSetInput struct {

	// The name of the configuration set to describe.
	//
	// This member is required.
	ConfigurationSetName *string

	// A list of configuration set attributes to return.
	ConfigurationSetAttributeNames []types.ConfigurationSetAttribute
}

// Represents the details of a configuration set. Configuration sets enable you to
// publish email sending events. For information about using configuration sets,
// see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type DescribeConfigurationSetOutput struct {

	// The configuration set object associated with the specified configuration set.
	ConfigurationSet *types.ConfigurationSet

	// Specifies whether messages that use the configuration set are required to use
	// Transport Layer Security (TLS).
	DeliveryOptions *types.DeliveryOptions

	// A list of event destinations associated with the configuration set.
	EventDestinations []*types.EventDestination

	// An object that represents the reputation settings for the configuration set.
	ReputationOptions *types.ReputationOptions

	// The name of the custom open and click tracking domain associated with the
	// configuration set.
	TrackingOptions *types.TrackingOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeConfigurationSetMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeConfigurationSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeConfigurationSet{}, middleware.After)
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
	addOpDescribeConfigurationSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeConfigurationSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DescribeConfigurationSet",
	}
}
