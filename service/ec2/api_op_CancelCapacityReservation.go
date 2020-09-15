// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels the specified Capacity Reservation, releases the reserved capacity, and
// changes the Capacity Reservation's state to cancelled. Instances running in the
// reserved capacity continue running until you stop them. Stopped instances that
// target the Capacity Reservation can no longer launch. Modify these instances to
// either target a different Capacity Reservation, launch On-Demand Instance
// capacity, or run in any open Capacity Reservation that has matching attributes
// and sufficient capacity.
func (c *Client) CancelCapacityReservation(ctx context.Context, params *CancelCapacityReservationInput, optFns ...func(*Options)) (*CancelCapacityReservationOutput, error) {
	stack := middleware.NewStack("CancelCapacityReservation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCancelCapacityReservationMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCancelCapacityReservationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelCapacityReservation(options.Region), middleware.Before)

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
			OperationName: "CancelCapacityReservation",
			Err:           err,
		}
	}
	out := result.(*CancelCapacityReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelCapacityReservationInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the Capacity Reservation to be cancelled.
	CapacityReservationId *string
}

type CancelCapacityReservationOutput struct {
	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCancelCapacityReservationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCancelCapacityReservation{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCancelCapacityReservation{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelCapacityReservation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CancelCapacityReservation",
	}
}
