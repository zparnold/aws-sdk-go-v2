// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables lifecycle management by creating a new LifecycleConfiguration object. A
// LifecycleConfiguration object defines when files in an Amazon EFS file system
// are automatically transitioned to the lower-cost EFS Infrequent Access (IA)
// storage class. A LifecycleConfiguration applies to all files in a file system.
// Each Amazon EFS file system supports one lifecycle configuration, which applies
// to all files in the file system. If a LifecycleConfiguration object already
// exists for the specified file system, a PutLifecycleConfiguration call modifies
// the existing configuration. A PutLifecycleConfiguration call with an empty
// LifecyclePolicies array in the request body deletes any existing
// LifecycleConfiguration and disables lifecycle management.  <p>In the request,
// specify the following: </p> <ul> <li> <p>The ID for the file system for which
// you are enabling, disabling, or modifying lifecycle management.</p> </li> <li>
// <p>A <code>LifecyclePolicies</code> array of <code>LifecyclePolicy</code>
// objects that define when files are moved to the IA storage class. The array can
// contain only one <code>LifecyclePolicy</code> item.</p> </li> </ul> <p>This
// operation requires permissions for the
// <code>elasticfilesystem:PutLifecycleConfiguration</code> operation.</p> <p>To
// apply a <code>LifecycleConfiguration</code> object to an encrypted file system,
// you need the same AWS Key Management Service (AWS KMS) permissions as when you
// created the encrypted file system. </p>
func (c *Client) PutLifecycleConfiguration(ctx context.Context, params *PutLifecycleConfigurationInput, optFns ...func(*Options)) (*PutLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &PutLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecycleConfiguration", params, optFns, addOperationPutLifecycleConfigurationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleConfigurationInput struct {

	// The ID of the file system for which you are creating the LifecycleConfiguration
	// object (String).
	//
	// This member is required.
	FileSystemId *string

	// An array of LifecyclePolicy objects that define the file system's
	// LifecycleConfiguration object. A LifecycleConfiguration object tells lifecycle
	// management when to transition files from the Standard storage class to the
	// Infrequent Access storage class.
	//
	// This member is required.
	LifecyclePolicies []*types.LifecyclePolicy
}

type PutLifecycleConfigurationOutput struct {

	// An array of lifecycle management policies. Currently, EFS supports a maximum of
	// one policy per file system.
	LifecyclePolicies []*types.LifecyclePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutLifecycleConfigurationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLifecycleConfiguration{}, middleware.After)
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
	addOpPutLifecycleConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutLifecycleConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "PutLifecycleConfiguration",
	}
}
