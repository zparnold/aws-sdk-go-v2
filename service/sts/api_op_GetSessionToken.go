// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a set of temporary credentials for an AWS account or IAM user. The
// credentials consist of an access key ID, a secret access key, and a security
// token. Typically, you use GetSessionToken if you want to use MFA to protect
// programmatic calls to specific AWS API operations like Amazon EC2 StopInstances.
// MFA-enabled IAM users would need to call GetSessionToken and submit an MFA code
// that is associated with their MFA device. Using the temporary security
// credentials that are returned from the call, IAM users can then make
// programmatic calls to API operations that require MFA authentication. If you do
// not supply a correct MFA code, then the API returns an access denied error. For
// a comparison of GetSessionToken with the other API operations that produce
// temporary credentials, see Requesting Temporary Security Credentials
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS API operations
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide. Session Duration The GetSessionToken operation must be
// called by using the long-term AWS security credentials of the AWS account root
// user or an IAM user. Credentials that are created by IAM users are valid for the
// duration that you specify. This duration can range from 900 seconds (15 minutes)
// up to a maximum of 129,600 seconds (36 hours), with a default of 43,200 seconds
// (12 hours). Credentials based on account credentials can range from 900 seconds
// (15 minutes) up to 3,600 seconds (1 hour), with a default of 1 hour. Permissions
// The temporary security credentials created by GetSessionToken can be used to
// make API calls to any AWS service with the following exceptions:
//
//     * You
// cannot call any IAM API operations unless MFA authentication information is
// included in the request.
//
//     * You cannot call any STS API except AssumeRole or
// GetCallerIdentity.
//
// We recommend that you do not call GetSessionToken with AWS
// account root user credentials. Instead, follow our best practices
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#create-iam-users)
// by creating one or more IAM users, giving them the necessary permissions, and
// using IAM users for everyday interaction with AWS. The credentials that are
// returned by GetSessionToken are based on permissions associated with the user
// whose credentials were used to call the operation. If GetSessionToken is called
// using AWS account root user credentials, the temporary credentials have root
// user permissions. Similarly, if GetSessionToken is called using the credentials
// of an IAM user, the temporary credentials have the same permissions as the IAM
// user. For more information about using GetSessionToken to create temporary
// credentials, go to Temporary Credentials for Users in Untrusted Environments
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_getsessiontoken)
// in the IAM User Guide.
func (c *Client) GetSessionToken(ctx context.Context, params *GetSessionTokenInput, optFns ...func(*Options)) (*GetSessionTokenOutput, error) {
	if params == nil {
		params = &GetSessionTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSessionToken", params, optFns, addOperationGetSessionTokenMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSessionTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSessionTokenInput struct {

	// The duration, in seconds, that the credentials should remain valid. Acceptable
	// durations for IAM user sessions range from 900 seconds (15 minutes) to 129,600
	// seconds (36 hours), with 43,200 seconds (12 hours) as the default. Sessions for
	// AWS account owners are restricted to a maximum of 3,600 seconds (one hour). If
	// the duration is longer than one hour, the session for AWS account owners
	// defaults to one hour.
	DurationSeconds *int32

	// The identification number of the MFA device that is associated with the IAM user
	// who is making the GetSessionToken call. Specify this value if the IAM user has a
	// policy that requires MFA authentication. The value is either the serial number
	// for a hardware device (such as GAHT12345678) or an Amazon Resource Name (ARN)
	// for a virtual device (such as arn:aws:iam::123456789012:mfa/user). You can find
	// the device for an IAM user by going to the AWS Management Console and viewing
	// the user's security credentials. The regex used to validate this parameter is a
	// string of characters consisting of upper- and lower-case alphanumeric characters
	// with no spaces. You can also include underscores or any of the following
	// characters: =,.@:/-
	SerialNumber *string

	// The value provided by the MFA device, if MFA is required. If any policy requires
	// the IAM user to submit an MFA code, specify this value. If MFA authentication is
	// required, the user must provide a code when requesting a set of temporary
	// security credentials. A user who fails to provide the code receives an "access
	// denied" response when requesting resources that require MFA authentication. The
	// format for this parameter, as described by its regex pattern, is a sequence of
	// six numeric digits.
	TokenCode *string
}

// Contains the response to a successful GetSessionToken () request, including
// temporary AWS credentials that can be used to make AWS requests.
type GetSessionTokenOutput struct {

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token. The size of the security token
	// that STS API operations return is not fixed. We strongly recommend that you make
	// no assumptions about the maximum size.
	Credentials *types.Credentials

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSessionTokenMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetSessionToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetSessionToken{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSessionToken(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSessionToken(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sts",
		OperationName: "GetSessionToken",
	}
}
