// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The UpdateHITReviewStatus operation updates the status of a HIT. If the status
// is Reviewable, this operation can update the status to Reviewing, or it can
// revert a Reviewing HIT back to the Reviewable status.
func (c *Client) UpdateHITReviewStatus(ctx context.Context, params *UpdateHITReviewStatusInput, optFns ...func(*Options)) (*UpdateHITReviewStatusOutput, error) {
	if params == nil {
		params = &UpdateHITReviewStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateHITReviewStatus", params, optFns, addOperationUpdateHITReviewStatusMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateHITReviewStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateHITReviewStatusInput struct {

	// The ID of the HIT to update.
	//
	// This member is required.
	HITId *string

	// Specifies how to update the HIT status. Default is False.
	//
	//     * Setting this to
	// false will only transition a HIT from Reviewable to Reviewing
	//
	//     * Setting
	// this to true will only transition a HIT from Reviewing to Reviewable
	Revert *bool
}

type UpdateHITReviewStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateHITReviewStatusMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateHITReviewStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateHITReviewStatus{}, middleware.After)
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
	addOpUpdateHITReviewStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateHITReviewStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateHITReviewStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "UpdateHITReviewStatus",
	}
}
