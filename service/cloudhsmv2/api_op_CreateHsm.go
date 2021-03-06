// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new hardware security module (HSM) in the specified AWS CloudHSM
// cluster.
func (c *Client) CreateHsm(ctx context.Context, params *CreateHsmInput, optFns ...func(*Options)) (*CreateHsmOutput, error) {
	stack := middleware.NewStack("CreateHsm", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateHsmMiddlewares(stack)
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
	addOpCreateHsmValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHsm(options.Region), middleware.Before)
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
			OperationName: "CreateHsm",
			Err:           err,
		}
	}
	out := result.(*CreateHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHsmInput struct {

	// The Availability Zone where you are creating the HSM. To find the cluster's
	// Availability Zones, use DescribeClusters ().
	//
	// This member is required.
	AvailabilityZone *string

	// The identifier (ID) of the HSM's cluster. To find the cluster ID, use
	// DescribeClusters ().
	//
	// This member is required.
	ClusterId *string

	// The HSM's IP address. If you specify an IP address, use an available address
	// from the subnet that maps to the Availability Zone where you are creating the
	// HSM. If you don't specify an IP address, one is chosen for you from that subnet.
	IpAddress *string
}

type CreateHsmOutput struct {

	// Information about the HSM that was created.
	Hsm *types.Hsm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateHsmMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHsm{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHsm{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateHsm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "CreateHsm",
	}
}
