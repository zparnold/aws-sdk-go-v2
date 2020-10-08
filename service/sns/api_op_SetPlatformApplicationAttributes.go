// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the attributes of the platform application object for the supported push
// notification services, such as APNS and GCM (Firebase Cloud Messaging). For more
// information, see Using Amazon SNS Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). For information
// on configuring attributes for message delivery status, see Using Amazon SNS
// Application Attributes for Message Delivery Status
// (https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html).
func (c *Client) SetPlatformApplicationAttributes(ctx context.Context, params *SetPlatformApplicationAttributesInput, optFns ...func(*Options)) (*SetPlatformApplicationAttributesOutput, error) {
	if params == nil {
		params = &SetPlatformApplicationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetPlatformApplicationAttributes", params, optFns, addOperationSetPlatformApplicationAttributesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*SetPlatformApplicationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetPlatformApplicationAttributes action.
type SetPlatformApplicationAttributesInput struct {

	// A map of the platform application attributes. Attributes in this map include the
	// following:
	//
	//     * PlatformCredential – The credential received from the
	// notification service. For APNS and APNS_SANDBOX, PlatformCredential is private
	// key. For GCM (Firebase Cloud Messaging), PlatformCredential is API key. For ADM,
	// PlatformCredential is client secret.
	//
	//     * PlatformPrincipal – The principal
	// received from the notification service. For APNS and APNS_SANDBOX,
	// PlatformPrincipal is SSL certificate. For GCM (Firebase Cloud Messaging), there
	// is no PlatformPrincipal. For ADM, PlatformPrincipal is client id.
	//
	//     *
	// EventEndpointCreated – Topic ARN to which EndpointCreated event notifications
	// are sent.
	//
	//     * EventEndpointDeleted – Topic ARN to which EndpointDeleted event
	// notifications are sent.
	//
	//     * EventEndpointUpdated – Topic ARN to which
	// EndpointUpdate event notifications are sent.
	//
	//     * EventDeliveryFailure – Topic
	// ARN to which DeliveryFailure event notifications are sent upon Direct Publish
	// delivery failure (permanent) to one of the application's endpoints.
	//
	//     *
	// SuccessFeedbackRoleArn – IAM role ARN used to give Amazon SNS write access to
	// use CloudWatch Logs on your behalf.
	//
	//     * FailureFeedbackRoleArn – IAM role ARN
	// used to give Amazon SNS write access to use CloudWatch Logs on your behalf.
	//
	//
	// * SuccessFeedbackSampleRate – Sample rate percentage (0-100) of successfully
	// delivered messages.
	//
	// This member is required.
	Attributes map[string]*string

	// PlatformApplicationArn for SetPlatformApplicationAttributes action.
	//
	// This member is required.
	PlatformApplicationArn *string
}

type SetPlatformApplicationAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetPlatformApplicationAttributesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetPlatformApplicationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetPlatformApplicationAttributes{}, middleware.After)
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
	addOpSetPlatformApplicationAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetPlatformApplicationAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetPlatformApplicationAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "SetPlatformApplicationAttributes",
	}
}
