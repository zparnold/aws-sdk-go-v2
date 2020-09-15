// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Traffic Mirror session. A Traffic Mirror session actively copies
// packets from a Traffic Mirror source to a Traffic Mirror target. Create a
// filter, and then assign it to the session to define a subset of the traffic to
// mirror, for example all TCP traffic. The Traffic Mirror source and the Traffic
// Mirror target (monitoring appliances) can be in the same VPC, or in a different
// VPC connected via VPC peering or a transit gateway. By default, no traffic is
// mirrored. Use CreateTrafficMirrorFilter
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorFilter.htm)
// to create filter rules that specify the traffic to mirror.
func (c *Client) CreateTrafficMirrorSession(ctx context.Context, params *CreateTrafficMirrorSessionInput, optFns ...func(*Options)) (*CreateTrafficMirrorSessionOutput, error) {
	stack := middleware.NewStack("CreateTrafficMirrorSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateTrafficMirrorSessionMiddlewares(stack)
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
	addIdempotencyToken_opCreateTrafficMirrorSessionMiddleware(stack, options)
	addOpCreateTrafficMirrorSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrafficMirrorSession(options.Region), middleware.Before)

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
			OperationName: "CreateTrafficMirrorSession",
			Err:           err,
		}
	}
	out := result.(*CreateTrafficMirrorSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrafficMirrorSessionInput struct {
	// The description of the Traffic Mirror session.
	Description *string
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string
	// The ID of the source network interface.
	NetworkInterfaceId *string
	// The tags to assign to a Traffic Mirror session.
	TagSpecifications []*types.TagSpecification
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The VXLAN ID for the Traffic Mirror session. For more information about the
	// VXLAN protocol, see RFC 7348 (https://tools.ietf.org/html/rfc7348). If you do
	// not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
	VirtualNetworkId *int32
	// The number of bytes in each packet to mirror. These are bytes after the VXLAN
	// header. Do not specify this parameter when you want to mirror the entire packet.
	// To mirror a subset of the packet, set this to the length (in bytes) that you
	// want to mirror. For example, if you set this value to 100, then the first 100
	// bytes that meet the filter criteria are copied to the target. If you do not want
	// to mirror the entire packet, use the PacketLength parameter to specify the
	// number of bytes in each packet to mirror.
	PacketLength *int32
	// The session number determines the order in which sessions are evaluated when an
	// interface is used by multiple sessions. The first session with a matching filter
	// is the one that mirrors the packets. Valid values are 1-32766.
	SessionNumber *int32
	// The ID of the Traffic Mirror target.
	TrafficMirrorTargetId *string
	// The ID of the Traffic Mirror filter.
	TrafficMirrorFilterId *string
}

type CreateTrafficMirrorSessionOutput struct {
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string
	// Information about the Traffic Mirror session.
	TrafficMirrorSession *types.TrafficMirrorSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateTrafficMirrorSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateTrafficMirrorSession{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTrafficMirrorSession{}, middleware.After)
}

type idempotencyToken_initializeOpCreateTrafficMirrorSession struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTrafficMirrorSession) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTrafficMirrorSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTrafficMirrorSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTrafficMirrorSessionInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateTrafficMirrorSessionMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateTrafficMirrorSession{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTrafficMirrorSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTrafficMirrorSession",
	}
}
