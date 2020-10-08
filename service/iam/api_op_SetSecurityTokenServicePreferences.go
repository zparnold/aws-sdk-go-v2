// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the specified version of the global endpoint token as the token version
// used for the AWS account. By default, AWS Security Token Service (STS) is
// available as a global service, and all STS requests go to a single endpoint at
// https://sts.amazonaws.com. AWS recommends using Regional STS endpoints to reduce
// latency, build in redundancy, and increase session token availability. For
// information about Regional endpoints for STS, see AWS Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#sts_region) in the AWS
// General Reference. If you make an STS call to the global endpoint, the resulting
// session tokens might be valid in some Regions but not others. It depends on the
// version that is set in this operation. Version 1 tokens are valid only in AWS
// Regions that are available by default. These tokens do not work in manually
// enabled Regions, such as Asia Pacific (Hong Kong). Version 2 tokens are valid in
// all Regions. However, version 2 tokens are longer and might affect systems where
// you temporarily store tokens. For information, see Activating and Deactivating
// STS in an AWS Region
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the IAM User Guide. To view the current session token version, see the
// GlobalEndpointTokenVersion entry in the response of the GetAccountSummary ()
// operation.
func (c *Client) SetSecurityTokenServicePreferences(ctx context.Context, params *SetSecurityTokenServicePreferencesInput, optFns ...func(*Options)) (*SetSecurityTokenServicePreferencesOutput, error) {
	if params == nil {
		params = &SetSecurityTokenServicePreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetSecurityTokenServicePreferences", params, optFns, addOperationSetSecurityTokenServicePreferencesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*SetSecurityTokenServicePreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetSecurityTokenServicePreferencesInput struct {

	// The version of the global endpoint token. Version 1 tokens are valid only in AWS
	// Regions that are available by default. These tokens do not work in manually
	// enabled Regions, such as Asia Pacific (Hong Kong). Version 2 tokens are valid in
	// all Regions. However, version 2 tokens are longer and might affect systems where
	// you temporarily store tokens. For information, see Activating and Deactivating
	// STS in an AWS Region
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
	// in the IAM User Guide.
	//
	// This member is required.
	GlobalEndpointTokenVersion types.GlobalEndpointTokenVersion
}

type SetSecurityTokenServicePreferencesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetSecurityTokenServicePreferencesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetSecurityTokenServicePreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetSecurityTokenServicePreferences{}, middleware.After)
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
	addOpSetSecurityTokenServicePreferencesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetSecurityTokenServicePreferences(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetSecurityTokenServicePreferences(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "SetSecurityTokenServicePreferences",
	}
}
