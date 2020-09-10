// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets detailed information about the specified license configuration.
func (c *Client) GetLicenseConfiguration(ctx context.Context, params *GetLicenseConfigurationInput, optFns ...func(*Options)) (*GetLicenseConfigurationOutput, error) {
	stack := middleware.NewStack("GetLicenseConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetLicenseConfigurationMiddlewares(stack)
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
	addOpGetLicenseConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLicenseConfiguration(options.Region), middleware.Before)

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
			OperationName: "GetLicenseConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLicenseConfigurationInput struct {
	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string
}

type GetLicenseConfigurationOutput struct {
	// Summaries of the managed resources.
	ManagedResourceSummaryList []*types.ManagedResourceSummary
	// Product information.
	ProductInformationList []*types.ProductInformation
	// Number of available licenses.
	LicenseCount *int64
	// Tags for the license configuration.
	Tags []*types.Tag
	// Dimension on which the licenses are counted.
	LicenseCountingType types.LicenseCountingType
	// Account ID of the owner of the license configuration.
	OwnerAccountId *string
	// Name of the license configuration.
	Name *string
	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string
	// Unique ID for the license configuration.
	LicenseConfigurationId *string
	// Summaries of the licenses consumed by resources.
	ConsumedLicenseSummaryList []*types.ConsumedLicenseSummary
	// License configuration status.
	Status *string
	// Description of the license configuration.
	Description *string
	// Number of licenses assigned to resources.
	ConsumedLicenses *int64
	// Automated discovery information.
	AutomatedDiscoveryInformation *types.AutomatedDiscoveryInformation
	// License rules.
	LicenseRules []*string
	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetLicenseConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetLicenseConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLicenseConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetLicenseConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "GetLicenseConfiguration",
	}
}