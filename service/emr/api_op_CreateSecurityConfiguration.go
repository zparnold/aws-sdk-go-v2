// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a security configuration, which is stored in the service and can be
// specified when a cluster is created.
func (c *Client) CreateSecurityConfiguration(ctx context.Context, params *CreateSecurityConfigurationInput, optFns ...func(*Options)) (*CreateSecurityConfigurationOutput, error) {
	if params == nil {
		params = &CreateSecurityConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityConfiguration", params, optFns, addOperationCreateSecurityConfigurationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityConfigurationInput struct {

	// The name of the security configuration.
	//
	// This member is required.
	Name *string

	// The security configuration details in JSON format. For JSON parameters and
	// examples, see Use Security Configurations to Set Up Cluster Security
	// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-security-configurations.html)
	// in the Amazon EMR Management Guide.
	//
	// This member is required.
	SecurityConfiguration *string
}

type CreateSecurityConfigurationOutput struct {

	// The date and time the security configuration was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The name of the security configuration.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSecurityConfigurationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSecurityConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSecurityConfiguration{}, middleware.After)
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
	addOpCreateSecurityConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateSecurityConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "CreateSecurityConfiguration",
	}
}
