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

// Updates information about an existing private device instance profile.
func (c *Client) UpdateInstanceProfile(ctx context.Context, params *UpdateInstanceProfileInput, optFns ...func(*Options)) (*UpdateInstanceProfileOutput, error) {
	if params == nil {
		params = &UpdateInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInstanceProfile", params, optFns, addOperationUpdateInstanceProfileMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInstanceProfileInput struct {

	// The Amazon Resource Name (ARN) of the instance profile.
	//
	// This member is required.
	Arn *string

	// The updated description for your instance profile.
	Description *string

	// An array of strings that specifies the list of app packages that should not be
	// cleaned up from the device after a test run is over. The list of packages is
	// only considered if you set packageCleanup to true.
	ExcludeAppPackagesFromCleanup []*string

	// The updated name for your instance profile.
	Name *string

	// The updated choice for whether you want to specify package cleanup. The default
	// value is false for private devices.
	PackageCleanup *bool

	// The updated choice for whether you want to reboot the device after use. The
	// default value is true.
	RebootAfterUse *bool
}

type UpdateInstanceProfileOutput struct {

	// An object that contains information about your instance profile.
	InstanceProfile *types.InstanceProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateInstanceProfileMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateInstanceProfile{}, middleware.After)
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
	addOpUpdateInstanceProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInstanceProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateInstanceProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "UpdateInstanceProfile",
	}
}
