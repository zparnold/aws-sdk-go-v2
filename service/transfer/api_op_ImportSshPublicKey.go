// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a Secure Shell (SSH) public key to a user account identified by a UserName
// value assigned to the specific file transfer protocol-enabled server, identified
// by ServerId.  <p>The response returns the <code>UserName</code> value, the
// <code>ServerId</code> value, and the name of the
// <code>SshPublicKeyId</code>.</p>
func (c *Client) ImportSshPublicKey(ctx context.Context, params *ImportSshPublicKeyInput, optFns ...func(*Options)) (*ImportSshPublicKeyOutput, error) {
	stack := middleware.NewStack("ImportSshPublicKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpImportSshPublicKeyMiddlewares(stack)
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
	addOpImportSshPublicKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportSshPublicKey(options.Region), middleware.Before)
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
			OperationName: "ImportSshPublicKey",
			Err:           err,
		}
	}
	out := result.(*ImportSshPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportSshPublicKeyInput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server.
	//
	// This member is required.
	ServerId *string

	// The public key portion of an SSH key pair.
	//
	// This member is required.
	SshPublicKeyBody *string

	// The name of the user account that is assigned to one or more file transfer
	// protocol-enabled servers.
	//
	// This member is required.
	UserName *string
}

// Identifies the user, the file transfer protocol-enabled server they belong to,
// and the identifier of the SSH public key associated with that user. A user can
// have more than one key on each server that they are associated with.
type ImportSshPublicKeyOutput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server.
	//
	// This member is required.
	ServerId *string

	// The name given to a public key by the system that was imported.
	//
	// This member is required.
	SshPublicKeyId *string

	// A user name assigned to the ServerID value that you specified.
	//
	// This member is required.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpImportSshPublicKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpImportSshPublicKey{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportSshPublicKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opImportSshPublicKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ImportSshPublicKey",
	}
}
