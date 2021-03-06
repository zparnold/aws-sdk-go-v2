// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a subscriber. Deleting the last subscriber to a notification also
// deletes the notification.
func (c *Client) DeleteSubscriber(ctx context.Context, params *DeleteSubscriberInput, optFns ...func(*Options)) (*DeleteSubscriberOutput, error) {
	stack := middleware.NewStack("DeleteSubscriber", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteSubscriberMiddlewares(stack)
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
	addOpDeleteSubscriberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSubscriber(options.Region), middleware.Before)
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
			OperationName: "DeleteSubscriber",
			Err:           err,
		}
	}
	out := result.(*DeleteSubscriberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DeleteSubscriber
type DeleteSubscriberInput struct {

	// The accountId that is associated with the budget whose subscriber you want to
	// delete.
	//
	// This member is required.
	AccountId *string

	// The name of the budget whose subscriber you want to delete.
	//
	// This member is required.
	BudgetName *string

	// The notification whose subscriber you want to delete.
	//
	// This member is required.
	Notification *types.Notification

	// The subscriber that you want to delete.
	//
	// This member is required.
	Subscriber *types.Subscriber
}

// Response of DeleteSubscriber
type DeleteSubscriberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteSubscriberMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSubscriber{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSubscriber{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSubscriber(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DeleteSubscriber",
	}
}
