// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new MSK cluster.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, addOperationCreateClusterMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// Information about the broker nodes in the cluster.
	//
	// This member is required.
	BrokerNodeGroupInfo *types.BrokerNodeGroupInfo

	// The name of the cluster.
	//
	// This member is required.
	ClusterName *string

	// The version of Apache Kafka.
	//
	// This member is required.
	KafkaVersion *string

	// The number of broker nodes in the cluster.
	//
	// This member is required.
	NumberOfBrokerNodes *int32

	// Includes all client authentication related information.
	ClientAuthentication *types.ClientAuthentication

	// Represents the configuration that you want MSK to use for the brokers in a
	// cluster.
	ConfigurationInfo *types.ConfigurationInfo

	// Includes all encryption-related information.
	EncryptionInfo *types.EncryptionInfo

	// Specifies the level of monitoring for the MSK cluster. The possible values are
	// DEFAULT, PER_BROKER, and PER_TOPIC_PER_BROKER.
	EnhancedMonitoring types.EnhancedMonitoring

	LoggingInfo *types.LoggingInfo

	// The settings for open monitoring.
	OpenMonitoring *types.OpenMonitoringInfo

	// Create tags when creating the cluster.
	Tags map[string]*string
}

type CreateClusterOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The name of the MSK cluster.
	ClusterName *string

	// The state of the cluster. The possible states are CREATING, ACTIVE, and FAILED.
	State types.ClusterState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateClusterMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCluster{}, middleware.After)
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
	addOpCreateClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "CreateCluster",
	}
}
