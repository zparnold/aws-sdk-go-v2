// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a pending configuration change to a broker.
func (c *Client) UpdateBroker(ctx context.Context, params *UpdateBrokerInput, optFns ...func(*Options)) (*UpdateBrokerOutput, error) {
	if params == nil {
		params = &UpdateBrokerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBroker", params, optFns, addOperationUpdateBrokerMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates the broker using the specified properties.
type UpdateBrokerInput struct {

	// The unique ID that Amazon MQ generates for the broker.
	//
	// This member is required.
	BrokerId *string

	// The authentication strategy used to secure the broker.
	AuthenticationStrategy types.AuthenticationStrategy

	// Enables automatic upgrades to new minor versions for brokers, as Apache releases
	// the versions. The automatic upgrades occur during the maintenance window of the
	// broker or after a manual broker reboot.
	AutoMinorVersionUpgrade *bool

	// A list of information about the configuration.
	Configuration *types.ConfigurationId

	// The version of the broker engine. For a list of supported engine versions, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string

	// The host instance type of the broker to upgrade to. For a list of supported
	// instance types, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide//broker.html#broker-instance-types
	HostInstanceType *string

	// The metadata of the LDAP server used to authenticate and authorize connections
	// to the broker.
	LdapServerMetadata *types.LdapServerMetadataInput

	// Enables Amazon CloudWatch logging for brokers.
	Logs *types.Logs

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []*string
}

type UpdateBrokerOutput struct {

	// The authentication strategy used to secure the broker.
	AuthenticationStrategy types.AuthenticationStrategy

	// The new value of automatic upgrades to new minor version for brokers.
	AutoMinorVersionUpgrade *bool

	// Required. The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// The ID of the updated configuration.
	Configuration *types.ConfigurationId

	// The version of the broker engine to upgrade to. For a list of supported engine
	// versions, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string

	// The host instance type of the broker to upgrade to. For a list of supported
	// instance types, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide//broker.html#broker-instance-types
	HostInstanceType *string

	// The metadata of the LDAP server used to authenticate and authorize connections
	// to the broker.
	LdapServerMetadata *types.LdapServerMetadataOutput

	// The list of information about logs to be enabled for the specified broker.
	Logs *types.Logs

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateBrokerMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBroker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBroker{}, middleware.After)
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
	addOpUpdateBrokerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBroker(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateBroker(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "UpdateBroker",
	}
}
