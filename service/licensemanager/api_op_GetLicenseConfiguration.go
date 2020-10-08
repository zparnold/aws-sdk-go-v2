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

// Gets detailed information about the specified license configuration.
func (c *Client) GetLicenseConfiguration(ctx context.Context, params *GetLicenseConfigurationInput, optFns ...func(*Options)) (*GetLicenseConfigurationOutput, error) {
	if params == nil {
		params = &GetLicenseConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLicenseConfiguration", params, optFns, addOperationGetLicenseConfigurationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLicenseConfigurationInput struct {

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string
}

type GetLicenseConfigurationOutput struct {

	// Automated discovery information.
	AutomatedDiscoveryInformation *types.AutomatedDiscoveryInformation

	// Summaries of the licenses consumed by resources.
	ConsumedLicenseSummaryList []*types.ConsumedLicenseSummary

	// Number of licenses assigned to resources.
	ConsumedLicenses *int64

	// Description of the license configuration.
	Description *string

	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string

	// Unique ID for the license configuration.
	LicenseConfigurationId *string

	// Number of available licenses.
	LicenseCount *int64

	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool

	// Dimension on which the licenses are counted.
	LicenseCountingType types.LicenseCountingType

	// License rules.
	LicenseRules []*string

	// Summaries of the managed resources.
	ManagedResourceSummaryList []*types.ManagedResourceSummary

	// Name of the license configuration.
	Name *string

	// Account ID of the owner of the license configuration.
	OwnerAccountId *string

	// Product information.
	ProductInformationList []*types.ProductInformation

	// License configuration status.
	Status *string

	// Tags for the license configuration.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLicenseConfigurationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLicenseConfiguration{}, middleware.After)
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
	addOpGetLicenseConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLicenseConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetLicenseConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "GetLicenseConfiguration",
	}
}
