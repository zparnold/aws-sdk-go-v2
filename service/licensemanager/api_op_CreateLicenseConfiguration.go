// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a license configuration. A license configuration is an abstraction of a
// customer license agreement that can be consumed and enforced by License Manager.
// Components include specifications for the license type (licensing by instance,
// socket, CPU, or vCPU), allowed tenancy (shared tenancy, Dedicated Instance,
// Dedicated Host, or all of these), host affinity (how long a VM must be
// associated with a host), and the number of licenses purchased and used.
func (c *Client) CreateLicenseConfiguration(ctx context.Context, params *CreateLicenseConfigurationInput, optFns ...func(*Options)) (*CreateLicenseConfigurationOutput, error) {
	if params == nil {
		params = &CreateLicenseConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLicenseConfiguration", params, optFns, addOperationCreateLicenseConfigurationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLicenseConfigurationInput struct {

	// Dimension used to track the license inventory.
	//
	// This member is required.
	LicenseCountingType types.LicenseCountingType

	// Name of the license configuration.
	//
	// This member is required.
	Name *string

	// Description of the license configuration.
	Description *string

	// Number of licenses managed by the license configuration.
	LicenseCount *int64

	// Indicates whether hard or soft license enforcement is used. Exceeding a hard
	// limit blocks the launch of new instances.
	LicenseCountHardLimit *bool

	// License rules. The syntax is #name=value (for example,
	// #allowedTenancy=EC2-DedicatedHost). Available rules vary by dimension.
	//
	//     *
	// Cores dimension: allowedTenancy | maximumCores | minimumCores
	//
	//     * Instances
	// dimension: allowedTenancy | maximumCores | minimumCores | maximumSockets |
	// minimumSockets | maximumVcpus | minimumVcpus
	//
	//     * Sockets dimension:
	// allowedTenancy | maximumSockets | minimumSockets
	//
	//     * vCPUs dimension:
	// allowedTenancy | honorVcpuOptimization | maximumVcpus | minimumVcpus
	LicenseRules []*string

	// Product information.
	ProductInformationList []*types.ProductInformation

	// Tags to add to the license configuration.
	Tags []*types.Tag
}

type CreateLicenseConfigurationOutput struct {

	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLicenseConfigurationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLicenseConfiguration{}, middleware.After)
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
	addOpCreateLicenseConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLicenseConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateLicenseConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "CreateLicenseConfiguration",
	}
}
