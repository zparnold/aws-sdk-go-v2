// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters the specified AMI. After you deregister an AMI, it can't be used to
// launch new instances; however, it doesn't affect any instances that you've
// already launched from the AMI. You'll continue to incur usage costs for those
// instances until you terminate them. When you deregister an Amazon EBS-backed
// AMI, it doesn't affect the snapshot that was created for the root volume of the
// instance during the AMI creation process. When you deregister an instance
// store-backed AMI, it doesn't affect the files that you uploaded to Amazon S3
// when you created the AMI.
func (c *Client) DeregisterImage(ctx context.Context, params *DeregisterImageInput, optFns ...func(*Options)) (*DeregisterImageOutput, error) {
	stack := middleware.NewStack("DeregisterImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeregisterImageMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeregisterImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterImage(options.Region), middleware.Before)

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
			OperationName: "DeregisterImage",
			Err:           err,
		}
	}
	out := result.(*DeregisterImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeregisterImage.
type DeregisterImageInput struct {
	// The ID of the AMI.
	ImageId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeregisterImageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeregisterImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeregisterImage{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeregisterImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeregisterImage",
	}
}
