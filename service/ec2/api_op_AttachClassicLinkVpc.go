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

// Links an EC2-Classic instance to a ClassicLink-enabled VPC through one or more
// of the VPC's security groups. You cannot link an EC2-Classic instance to more
// than one VPC at a time. You can only link an instance that's in the running
// state. An instance is automatically unlinked from a VPC when it's stopped - you
// can link it to the VPC again when you restart it. After you've linked an
// instance, you cannot change the VPC security groups that are associated with it.
// To change the security groups, you must first unlink the instance, and then link
// it again. Linking your instance to a VPC is sometimes referred to as attaching
// your instance.
func (c *Client) AttachClassicLinkVpc(ctx context.Context, params *AttachClassicLinkVpcInput, optFns ...func(*Options)) (*AttachClassicLinkVpcOutput, error) {
	stack := middleware.NewStack("AttachClassicLinkVpc", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAttachClassicLinkVpcMiddlewares(stack)
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
	addOpAttachClassicLinkVpcValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachClassicLinkVpc(options.Region), middleware.Before)

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
			OperationName: "AttachClassicLinkVpc",
			Err:           err,
		}
	}
	out := result.(*AttachClassicLinkVpcOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachClassicLinkVpcInput struct {
	// The ID of one or more of the VPC's security groups. You cannot specify security
	// groups from a different VPC.
	Groups []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of an EC2-Classic instance to link to the ClassicLink-enabled VPC.
	InstanceId *string
	// The ID of a ClassicLink-enabled VPC.
	VpcId *string
}

type AttachClassicLinkVpcOutput struct {
	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAttachClassicLinkVpcMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAttachClassicLinkVpc{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAttachClassicLinkVpc{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachClassicLinkVpc(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AttachClassicLinkVpc",
	}
}
