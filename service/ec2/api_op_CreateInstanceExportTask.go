// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Exports a running or stopped instance to an Amazon S3 bucket. For information
// about the supported operating systems, image formats, and known limitations for
// the types of instances you can export, see Exporting an Instance as a VM Using
// VM Import/Export
// (https://docs.aws.amazon.com/vm-import/latest/userguide/vmexport.html) in the VM
// Import/Export User Guide.
func (c *Client) CreateInstanceExportTask(ctx context.Context, params *CreateInstanceExportTaskInput, optFns ...func(*Options)) (*CreateInstanceExportTaskOutput, error) {
	if params == nil {
		params = &CreateInstanceExportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstanceExportTask", params, optFns, addOperationCreateInstanceExportTaskMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstanceExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstanceExportTaskInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// A description for the conversion task or the resource being exported. The
	// maximum length is 255 characters.
	Description *string

	// The format and location for an instance export task.
	ExportToS3Task *types.ExportToS3TaskSpecification

	// The tags to apply to the instance export task during creation.
	TagSpecifications []*types.TagSpecification

	// The target virtualization environment.
	TargetEnvironment types.ExportEnvironment
}

type CreateInstanceExportTaskOutput struct {

	// Information about the instance export task.
	ExportTask *types.ExportTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateInstanceExportTaskMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateInstanceExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateInstanceExportTask{}, middleware.After)
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
	addOpCreateInstanceExportTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstanceExportTask(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateInstanceExportTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateInstanceExportTask",
	}
}
