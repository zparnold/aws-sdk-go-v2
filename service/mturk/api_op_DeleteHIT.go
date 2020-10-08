// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The DeleteHIT operation is used to delete HIT that is no longer needed. Only the
// Requester who created the HIT can delete it. You can only dispose of HITs that
// are in the Reviewable state, with all of their submitted assignments already
// either approved or rejected. If you call the DeleteHIT operation on a HIT that
// is not in the Reviewable state (for example, that has not expired, or still has
// active assignments), or on a HIT that is Reviewable but without all of its
// submitted assignments already approved or rejected, the service will return an
// error.
//
//     * HITs are automatically disposed of after 120 days.
//
//     * After
// you dispose of a HIT, you can no longer approve the HIT's rejected
// assignments.
//
//     * Disposed HITs are not returned in results for the ListHITs
// operation.
//
//     * Disposing HITs can improve the performance of operations such
// as ListReviewableHITs and ListHITs.
func (c *Client) DeleteHIT(ctx context.Context, params *DeleteHITInput, optFns ...func(*Options)) (*DeleteHITOutput, error) {
	if params == nil {
		params = &DeleteHITInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHIT", params, optFns, addOperationDeleteHITMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteHITInput struct {

	// The ID of the HIT to be deleted.
	//
	// This member is required.
	HITId *string
}

type DeleteHITOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteHITMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteHIT{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteHIT{}, middleware.After)
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
	addOpDeleteHITValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHIT(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteHIT(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "DeleteHIT",
	}
}
