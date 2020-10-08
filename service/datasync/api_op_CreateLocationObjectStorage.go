// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an endpoint for a self-managed object storage bucket.
func (c *Client) CreateLocationObjectStorage(ctx context.Context, params *CreateLocationObjectStorageInput, optFns ...func(*Options)) (*CreateLocationObjectStorageOutput, error) {
	if params == nil {
		params = &CreateLocationObjectStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationObjectStorage", params, optFns, addOperationCreateLocationObjectStorageMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationObjectStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationObjectStorageRequest
type CreateLocationObjectStorageInput struct {

	// The Amazon Resource Name (ARN) of the agents associated with the self-managed
	// object storage server location.
	//
	// This member is required.
	AgentArns []*string

	// The bucket on the self-managed object storage server that is used to read data
	// from.
	//
	// This member is required.
	BucketName *string

	// The name of the self-managed object storage server. This value is the IP address
	// or Domain Name Service (DNS) name of the object storage server. An agent uses
	// this host name to mount the object storage server in a network.
	//
	// This member is required.
	ServerHostname *string

	// Optional. The access key is used if credentials are required to access the
	// self-managed object storage server.
	AccessKey *string

	// Optional. The secret key is used if credentials are required to access the
	// self-managed object storage server.
	SecretKey *string

	// The port that your self-managed object storage server accepts inbound network
	// traffic on. The server port is set by default to TCP 80 (HTTP) or TCP 443
	// (HTTPS). You can specify a custom port if your self-managed object storage
	// server requires one.
	ServerPort *int32

	// The protocol that the object storage server uses to communicate. Valid values
	// are HTTP or HTTPS.
	ServerProtocol types.ObjectStorageServerProtocol

	// The subdirectory in the self-managed object storage server that is used to read
	// data from.
	Subdirectory *string

	// The key-value pair that represents the tag that you want to add to the location.
	// The value can be an empty string. We recommend using tags to name your
	// resources.
	Tags []*types.TagListEntry
}

// CreateLocationObjectStorageResponse
type CreateLocationObjectStorageOutput struct {

	// The Amazon Resource Name (ARN) of the agents associated with the self-managed
	// object storage server location.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLocationObjectStorageMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationObjectStorage{}, middleware.After)
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
	addOpCreateLocationObjectStorageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationObjectStorage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateLocationObjectStorage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationObjectStorage",
	}
}
