// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables you to provide feedback to Amazon Kendra to improve the performance of
// the service.
func (c *Client) SubmitFeedback(ctx context.Context, params *SubmitFeedbackInput, optFns ...func(*Options)) (*SubmitFeedbackOutput, error) {
	stack := middleware.NewStack("SubmitFeedback", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSubmitFeedbackMiddlewares(stack)
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
	addOpSubmitFeedbackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitFeedback(options.Region), middleware.Before)
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
			OperationName: "SubmitFeedback",
			Err:           err,
		}
	}
	out := result.(*SubmitFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitFeedbackInput struct {

	// The identifier of the index that was queried.
	//
	// This member is required.
	IndexId *string

	// The identifier of the specific query for which you are submitting feedback. The
	// query ID is returned in the response to the operation.
	//
	// This member is required.
	QueryId *string

	// Tells Amazon Kendra that a particular search result link was chosen by the user.
	ClickFeedbackItems []*types.ClickFeedback

	// Provides Amazon Kendra with relevant or not relevant feedback for whether a
	// particular item was relevant to the search.
	RelevanceFeedbackItems []*types.RelevanceFeedback
}

type SubmitFeedbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSubmitFeedbackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSubmitFeedback{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubmitFeedback{}, middleware.After)
}

func newServiceMetadataMiddleware_opSubmitFeedback(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "SubmitFeedback",
	}
}
