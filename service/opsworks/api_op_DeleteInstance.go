// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a specified instance, which terminates the associated Amazon EC2
// instance. You must stop an instance before you can delete it. For more
// information, see Deleting Instances
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-delete.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DeleteInstance(ctx context.Context, params *DeleteInstanceInput, optFns ...func(*Options)) (*DeleteInstanceOutput, error) {
	if params == nil {
		params = &DeleteInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteInstance", params, optFns, addOperationDeleteInstanceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInstanceInput struct {

	// The instance ID.
	//
	// This member is required.
	InstanceId *string

	// Whether to delete the instance Elastic IP address.
	DeleteElasticIp *bool

	// Whether to delete the instance's Amazon EBS volumes.
	DeleteVolumes *bool
}

type DeleteInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteInstanceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteInstance{}, middleware.After)
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
	addOpDeleteInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DeleteInstance",
	}
}
