// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops advertising an address range that is provisioned as an address pool. You
// can perform this operation at most once every 10 seconds, even if you specify
// different address ranges each time. It can take a few minutes before traffic to
// the specified addresses stops routing to AWS because of BGP propagation delays.
func (c *Client) WithdrawByoipCidr(ctx context.Context, params *WithdrawByoipCidrInput, optFns ...func(*Options)) (*WithdrawByoipCidrOutput, error) {
	stack := middleware.NewStack("WithdrawByoipCidr", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpWithdrawByoipCidrMiddlewares(stack)
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
	addOpWithdrawByoipCidrValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opWithdrawByoipCidr(options.Region), middleware.Before)

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
			OperationName: "WithdrawByoipCidr",
			Err:           err,
		}
	}
	out := result.(*WithdrawByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type WithdrawByoipCidrInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The address range, in CIDR notation.
	Cidr *string
}

type WithdrawByoipCidrOutput struct {
	// Information about the address pool.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpWithdrawByoipCidrMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpWithdrawByoipCidr{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpWithdrawByoipCidr{}, middleware.After)
}

func newServiceMetadataMiddleware_opWithdrawByoipCidr(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "WithdrawByoipCidr",
	}
}
