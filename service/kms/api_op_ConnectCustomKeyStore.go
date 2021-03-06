// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Connects or reconnects a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// to its associated AWS CloudHSM cluster. The custom key store must be connected
// before you can create customer master keys (CMKs) in the key store or use the
// CMKs it contains. You can disconnect and reconnect a custom key store at any
// time. To connect a custom key store, its associated AWS CloudHSM cluster must
// have at least one active HSM. To get the number of active HSMs in a cluster, use
// the DescribeClusters
// (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
// operation. To add HSMs to the cluster, use the CreateHsm
// (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_CreateHsm.html)
// operation. Also, the kmsuser crypto user
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-store-concepts.html#concept-kmsuser)
// (CU) must not be logged into the cluster. This prevents AWS KMS from using this
// account to log in. The connection process can take an extended amount of time to
// complete; up to 20 minutes. This operation starts the connection process, but it
// does not wait for it to complete. When it succeeds, this operation quickly
// returns an HTTP 200 response and a JSON object with no properties. However, this
// response does not indicate that the custom key store is connected. To get the
// connection state of the custom key store, use the DescribeCustomKeyStores ()
// operation. During the connection process, AWS KMS finds the AWS CloudHSM cluster
// that is associated with the custom key store, creates the connection
// infrastructure, connects to the cluster, logs into the AWS CloudHSM client as
// the kmsuser CU, and rotates its password. The ConnectCustomKeyStore operation
// might fail for various reasons. To find the reason, use the
// DescribeCustomKeyStores () operation and see the ConnectionErrorCode in the
// response. For help interpreting the ConnectionErrorCode, see
// CustomKeyStoresListEntry (). To fix the failure, use the
// DisconnectCustomKeyStore () operation to disconnect the custom key store,
// correct the error, use the UpdateCustomKeyStore () operation if necessary, and
// then use ConnectCustomKeyStore again. If you are having trouble connecting or
// disconnecting a custom key store, see Troubleshooting a Custom Key Store
// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) ConnectCustomKeyStore(ctx context.Context, params *ConnectCustomKeyStoreInput, optFns ...func(*Options)) (*ConnectCustomKeyStoreOutput, error) {
	stack := middleware.NewStack("ConnectCustomKeyStore", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpConnectCustomKeyStoreMiddlewares(stack)
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
	addOpConnectCustomKeyStoreValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConnectCustomKeyStore(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "ConnectCustomKeyStore",
			Err:           err,
		}
	}
	out := result.(*ConnectCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConnectCustomKeyStoreInput struct {

	// Enter the key store ID of the custom key store that you want to connect. To find
	// the ID of a custom key store, use the DescribeCustomKeyStores () operation.
	//
	// This member is required.
	CustomKeyStoreId *string
}

type ConnectCustomKeyStoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpConnectCustomKeyStoreMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpConnectCustomKeyStore{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpConnectCustomKeyStore{}, middleware.After)
}

func newServiceMetadataMiddleware_opConnectCustomKeyStore(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ConnectCustomKeyStore",
	}
}
