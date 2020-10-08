// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a network profile.
func (c *Client) CreateNetworkProfile(ctx context.Context, params *CreateNetworkProfileInput, optFns ...func(*Options)) (*CreateNetworkProfileOutput, error) {
	if params == nil {
		params = &CreateNetworkProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNetworkProfile", params, optFns, addOperationCreateNetworkProfileMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNetworkProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNetworkProfileInput struct {

	// The name for the new network profile.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the project for which you want to create a
	// network profile.
	//
	// This member is required.
	ProjectArn *string

	// The description of the network profile.
	Description *string

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	DownlinkBandwidthBits *int64

	// Delay time for all packets to destination in milliseconds as an integer from 0
	// to 2000.
	DownlinkDelayMs *int64

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	DownlinkJitterMs *int64

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent *int32

	// The type of network profile to create. Valid values are listed here.
	Type types.NetworkProfileType

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	UplinkBandwidthBits *int64

	// Delay time for all packets to destination in milliseconds as an integer from 0
	// to 2000.
	UplinkDelayMs *int64

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	UplinkJitterMs *int64

	// Proportion of transmitted packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent *int32
}

type CreateNetworkProfileOutput struct {

	// The network profile that is returned by the create network profile request.
	NetworkProfile *types.NetworkProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNetworkProfileMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNetworkProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNetworkProfile{}, middleware.After)
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
	addOpCreateNetworkProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateNetworkProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateNetworkProfile",
	}
}
