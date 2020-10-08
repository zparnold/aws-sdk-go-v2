// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all of the properties of a topic. Topic properties returned might differ
// based on the authorization of the user.
func (c *Client) GetTopicAttributes(ctx context.Context, params *GetTopicAttributesInput, optFns ...func(*Options)) (*GetTopicAttributesOutput, error) {
	if params == nil {
		params = &GetTopicAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTopicAttributes", params, optFns, addOperationGetTopicAttributesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTopicAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for GetTopicAttributes action.
type GetTopicAttributesInput struct {

	// The ARN of the topic whose properties you want to get.
	//
	// This member is required.
	TopicArn *string
}

// Response for GetTopicAttributes action.
type GetTopicAttributesOutput struct {

	// A map of the topic's attributes. Attributes in this map include the following:
	//
	//
	// * DeliveryPolicy – The JSON serialization of the topic's delivery policy.
	//
	//     *
	// DisplayName – The human-readable name used in the From field for notifications
	// to email and email-json endpoints.
	//
	//     * Owner – The AWS account ID of the
	// topic's owner.
	//
	//     * Policy – The JSON serialization of the topic's access
	// control policy.
	//
	//     * SubscriptionsConfirmed – The number of confirmed
	// subscriptions for the topic.
	//
	//     * SubscriptionsDeleted – The number of deleted
	// subscriptions for the topic.
	//
	//     * SubscriptionsPending – The number of
	// subscriptions pending confirmation for the topic.
	//
	//     * TopicArn – The topic's
	// ARN.
	//
	//     * EffectiveDeliveryPolicy – The JSON serialization of the effective
	// delivery policy, taking system defaults into account.
	//
	//     <p>The following
	// attribute applies only to <a
	// href="https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html">server-side-encryption</a>:</p>
	// <ul> <li> <p> <code>KmsMasterKeyId</code> - The ID of an AWS-managed customer
	// master key (CMK) for Amazon SNS or a custom CMK. For more information, see <a
	// href="https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms">Key
	// Terms</a>. For more examples, see <a
	// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters">KeyId</a>
	// in the <i>AWS Key Management Service API Reference</i>.</p> </li> </ul>
	Attributes map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTopicAttributesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetTopicAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetTopicAttributes{}, middleware.After)
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
	addOpGetTopicAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTopicAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetTopicAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "GetTopicAttributes",
	}
}
