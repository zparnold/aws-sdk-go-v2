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

// Gets the License Manager settings for the current Region.
func (c *Client) GetServiceSettings(ctx context.Context, params *GetServiceSettingsInput, optFns ...func(*Options)) (*GetServiceSettingsOutput, error) {
	if params == nil {
		params = &GetServiceSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceSettings", params, optFns, addOperationGetServiceSettingsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceSettingsInput struct {
}

type GetServiceSettingsOutput struct {

	// Indicates whether cross-account discovery has been enabled.
	EnableCrossAccountsDiscovery *bool

	// Amazon Resource Name (ARN) of the AWS resource share. The License Manager master
	// account will provide member accounts with access to this share.
	LicenseManagerResourceShareArn *string

	// Indicates whether AWS Organizations has been integrated with License Manager for
	// cross-account discovery.
	OrganizationConfiguration *types.OrganizationConfiguration

	// Regional S3 bucket path for storing reports, license trail event data, discovery
	// data, and so on.
	S3BucketArn *string

	// SNS topic configured to receive notifications from License Manager.
	SnsTopicArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetServiceSettingsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetServiceSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServiceSettings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetServiceSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "GetServiceSettings",
	}
}
