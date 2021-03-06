// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the current snapshot for the patch baseline the instance uses. This
// API is primarily used by the AWS-RunPatchBaseline Systems Manager document.
func (c *Client) GetDeployablePatchSnapshotForInstance(ctx context.Context, params *GetDeployablePatchSnapshotForInstanceInput, optFns ...func(*Options)) (*GetDeployablePatchSnapshotForInstanceOutput, error) {
	stack := middleware.NewStack("GetDeployablePatchSnapshotForInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDeployablePatchSnapshotForInstanceMiddlewares(stack)
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
	addOpGetDeployablePatchSnapshotForInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeployablePatchSnapshotForInstance(options.Region), middleware.Before)
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
			OperationName: "GetDeployablePatchSnapshotForInstance",
			Err:           err,
		}
	}
	out := result.(*GetDeployablePatchSnapshotForInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeployablePatchSnapshotForInstanceInput struct {

	// The ID of the instance for which the appropriate patch snapshot should be
	// retrieved.
	//
	// This member is required.
	InstanceId *string

	// The user-defined snapshot ID.
	//
	// This member is required.
	SnapshotId *string
}

type GetDeployablePatchSnapshotForInstanceOutput struct {

	// The ID of the instance.
	InstanceId *string

	// Returns the specific operating system (for example Windows Server 2012 or Amazon
	// Linux 2015.09) on the instance for the specified patch snapshot.
	Product *string

	// A pre-signed Amazon S3 URL that can be used to download the patch snapshot.
	SnapshotDownloadUrl *string

	// The user-defined snapshot ID.
	SnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDeployablePatchSnapshotForInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeployablePatchSnapshotForInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeployablePatchSnapshotForInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDeployablePatchSnapshotForInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetDeployablePatchSnapshotForInstance",
	}
}
