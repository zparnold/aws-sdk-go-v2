// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows a subscription owner to set an attribute of the subscription to a new
// value.
func (c *Client) SetSubscriptionAttributes(ctx context.Context, params *SetSubscriptionAttributesInput, optFns ...func(*Options)) (*SetSubscriptionAttributesOutput, error) {
	if params == nil {
		params = &SetSubscriptionAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetSubscriptionAttributes", params, optFns, addOperationSetSubscriptionAttributesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*SetSubscriptionAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetSubscriptionAttributes action.
type SetSubscriptionAttributesInput struct {

	// A map of attributes with their corresponding values. The following lists the
	// names, descriptions, and values of the special request parameters that the
	// SetTopicAttributes action uses:
	//
	//     * DeliveryPolicy – The policy that defines
	// how Amazon SNS retries failed deliveries to HTTP/S endpoints.
	//
	//     *
	// FilterPolicy – The simple JSON object that lets your subscriber receive only a
	// subset of messages, rather than receiving every message published to the
	// topic.
	//
	//     * RawMessageDelivery – When set to true, enables raw message
	// delivery to Amazon SQS or HTTP/S endpoints. This eliminates the need for the
	// endpoints to process JSON formatting, which is otherwise created for Amazon SNS
	// metadata.
	//
	//     * RedrivePolicy – When specified, sends undeliverable messages to
	// the specified Amazon SQS dead-letter queue. Messages that can't be delivered due
	// to client errors (for example, when the subscribed endpoint is unreachable) or
	// server errors (for example, when the service that powers the subscribed endpoint
	// becomes unavailable) are held in the dead-letter queue for further analysis or
	// reprocessing.
	//
	// This member is required.
	AttributeName *string

	// The ARN of the subscription to modify.
	//
	// This member is required.
	SubscriptionArn *string

	// The new value for the attribute in JSON format.
	AttributeValue *string
}

type SetSubscriptionAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetSubscriptionAttributesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetSubscriptionAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetSubscriptionAttributes{}, middleware.After)
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
	addOpSetSubscriptionAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetSubscriptionAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetSubscriptionAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "SetSubscriptionAttributes",
	}
}
