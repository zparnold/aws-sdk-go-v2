// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve a list of event destinations that are associated with a configuration
// set. In Amazon Pinpoint, events include message sends, deliveries, opens,
// clicks, bounces, and complaints. Event destinations are places that you can send
// information about these events to. For example, you can send event data to
// Amazon SNS to receive notifications when you receive bounces or complaints, or
// you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for
// long-term storage.
func (c *Client) GetConfigurationSetEventDestinations(ctx context.Context, params *GetConfigurationSetEventDestinationsInput, optFns ...func(*Options)) (*GetConfigurationSetEventDestinationsOutput, error) {
	if params == nil {
		params = &GetConfigurationSetEventDestinationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfigurationSetEventDestinations", params, optFns, addOperationGetConfigurationSetEventDestinationsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfigurationSetEventDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain information about the event destinations for a configuration
// set.
type GetConfigurationSetEventDestinationsInput struct {

	// The name of the configuration set that contains the event destination.
	//
	// This member is required.
	ConfigurationSetName *string
}

// Information about an event destination for a configuration set.
type GetConfigurationSetEventDestinationsOutput struct {

	// An array that includes all of the events destinations that have been configured
	// for the configuration set.
	EventDestinations []*types.EventDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetConfigurationSetEventDestinationsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfigurationSetEventDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfigurationSetEventDestinations{}, middleware.After)
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
	addOpGetConfigurationSetEventDestinationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfigurationSetEventDestinations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetConfigurationSetEventDestinations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetConfigurationSetEventDestinations",
	}
}
