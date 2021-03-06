// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The UpdateExpirationForHIT operation allows you update the expiration time of a
// HIT. If you update it to a time in the past, the HIT will be immediately
// expired.
func (c *Client) UpdateExpirationForHIT(ctx context.Context, params *UpdateExpirationForHITInput, optFns ...func(*Options)) (*UpdateExpirationForHITOutput, error) {
	stack := middleware.NewStack("UpdateExpirationForHIT", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateExpirationForHITMiddlewares(stack)
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
	addOpUpdateExpirationForHITValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateExpirationForHIT(options.Region), middleware.Before)
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
			OperationName: "UpdateExpirationForHIT",
			Err:           err,
		}
	}
	out := result.(*UpdateExpirationForHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateExpirationForHITInput struct {

	// The date and time at which you want the HIT to expire
	//
	// This member is required.
	ExpireAt *time.Time

	// The HIT to update.
	//
	// This member is required.
	HITId *string
}

type UpdateExpirationForHITOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateExpirationForHITMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateExpirationForHIT{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateExpirationForHIT{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateExpirationForHIT(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "UpdateExpirationForHIT",
	}
}
