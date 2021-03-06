// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new device in a global network. If you specify both a site ID and a
// location, the location of the site is used for visualization in the Network
// Manager console.
func (c *Client) CreateDevice(ctx context.Context, params *CreateDeviceInput, optFns ...func(*Options)) (*CreateDeviceOutput, error) {
	stack := middleware.NewStack("CreateDevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDeviceMiddlewares(stack)
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
	addOpCreateDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDevice(options.Region), middleware.Before)
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
			OperationName: "CreateDevice",
			Err:           err,
		}
	}
	out := result.(*CreateDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeviceInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// A description of the device. Length Constraints: Maximum length of 256
	// characters.
	Description *string

	// The location of the device.
	Location *types.Location

	// The model of the device. Length Constraints: Maximum length of 128 characters.
	Model *string

	// The serial number of the device. Length Constraints: Maximum length of 128
	// characters.
	SerialNumber *string

	// The ID of the site.
	SiteId *string

	// The tags to apply to the resource during creation.
	Tags []*types.Tag

	// The type of the device.
	Type *string

	// The vendor of the device. Length Constraints: Maximum length of 128 characters.
	Vendor *string
}

type CreateDeviceOutput struct {

	// Information about the device.
	Device *types.Device

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDevice{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "CreateDevice",
	}
}
