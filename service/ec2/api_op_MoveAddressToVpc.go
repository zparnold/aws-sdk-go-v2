// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Moves an Elastic IP address from the EC2-Classic platform to the EC2-VPC
// platform. The Elastic IP address must be allocated to your account for more than
// 24 hours, and it must not be associated with an instance. After the Elastic IP
// address is moved, it is no longer available for use in the EC2-Classic platform,
// unless you move it back using the RestoreAddressToClassic () request. You cannot
// move an Elastic IP address that was originally allocated for use in the EC2-VPC
// platform to the EC2-Classic platform.
func (c *Client) MoveAddressToVpc(ctx context.Context, params *MoveAddressToVpcInput, optFns ...func(*Options)) (*MoveAddressToVpcOutput, error) {
	if params == nil {
		params = &MoveAddressToVpcInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MoveAddressToVpc", params, optFns, addOperationMoveAddressToVpcMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*MoveAddressToVpcOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MoveAddressToVpcInput struct {

	// The Elastic IP address.
	//
	// This member is required.
	PublicIp *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type MoveAddressToVpcOutput struct {

	// The allocation ID for the Elastic IP address.
	AllocationId *string

	// The status of the move of the IP address.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationMoveAddressToVpcMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpMoveAddressToVpc{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpMoveAddressToVpc{}, middleware.After)
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
	addOpMoveAddressToVpcValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opMoveAddressToVpc(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opMoveAddressToVpc(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "MoveAddressToVpc",
	}
}
