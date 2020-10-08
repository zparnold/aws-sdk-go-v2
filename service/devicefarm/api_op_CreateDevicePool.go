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

// Creates a device pool.
func (c *Client) CreateDevicePool(ctx context.Context, params *CreateDevicePoolInput, optFns ...func(*Options)) (*CreateDevicePoolOutput, error) {
	if params == nil {
		params = &CreateDevicePoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDevicePool", params, optFns, addOperationCreateDevicePoolMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDevicePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the create device pool operation.
type CreateDevicePoolInput struct {

	// The device pool's name.
	//
	// This member is required.
	Name *string

	// The ARN of the project for the device pool.
	//
	// This member is required.
	ProjectArn *string

	// The device pool's rules.
	//
	// This member is required.
	Rules []*types.Rule

	// The device pool's description.
	Description *string

	// The number of devices that Device Farm can add to your device pool. Device Farm
	// adds devices that are available and meet the criteria that you assign for the
	// rules parameter. Depending on how many devices meet these constraints, your
	// device pool might contain fewer devices than the value for this parameter. By
	// specifying the maximum number of devices, you can control the costs that you
	// incur by running tests.
	MaxDevices *int32
}

// Represents the result of a create device pool request.
type CreateDevicePoolOutput struct {

	// The newly created device pool.
	DevicePool *types.DevicePool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDevicePoolMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDevicePool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDevicePool{}, middleware.After)
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
	addOpCreateDevicePoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDevicePool(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDevicePool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateDevicePool",
	}
}
