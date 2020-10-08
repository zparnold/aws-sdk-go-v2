// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modify the version-specific settings of a Lambda function.  <p>When you update a
// function, Lambda provisions an instance of the function and its supporting
// resources. If your function connects to a VPC, this process can take a minute.
// During this time, you can't modify the function, but you can still invoke it.
// The <code>LastUpdateStatus</code>, <code>LastUpdateStatusReason</code>, and
// <code>LastUpdateStatusReasonCode</code> fields in the response from
// <a>GetFunctionConfiguration</a> indicate when the update is complete and the
// function is processing events with the new configuration. For more information,
// see <a
// href="https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html">Function
// States</a>.</p> <p>These settings can vary between versions of a function and
// are locked when you publish a version. You can't modify the configuration of a
// published version, only the unpublished version.</p> <p>To configure function
// concurrency, use <a>PutFunctionConcurrency</a>. To grant invoke permissions to
// an account or AWS service, use <a>AddPermission</a>.</p>
func (c *Client) UpdateFunctionConfiguration(ctx context.Context, params *UpdateFunctionConfigurationInput, optFns ...func(*Options)) (*UpdateFunctionConfigurationOutput, error) {
	if params == nil {
		params = &UpdateFunctionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFunctionConfiguration", params, optFns, addOperationUpdateFunctionConfigurationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFunctionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFunctionConfigurationInput struct {

	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// my-function.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN -
	// 123456789012:function:my-function.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it is limited to 64 characters
	// in length.
	//
	// This member is required.
	FunctionName *string

	// A dead letter queue configuration that specifies the queue or topic where Lambda
	// sends asynchronous events when they fail processing. For more information, see
	// Dead Letter Queues
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq).
	DeadLetterConfig *types.DeadLetterConfig

	// A description of the function.
	Description *string

	// Environment variables that are accessible from function code during execution.
	Environment *types.Environment

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []*types.FileSystemConfig

	// The name of the method within your code that Lambda calls to execute your
	// function. The format includes the file name. It can also include namespaces and
	// other qualifiers, depending on the runtime. For more information, see
	// Programming Model
	// (https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html).
	Handler *string

	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt
	// your function's environment variables. If it's not provided, AWS Lambda uses a
	// default service key.
	KMSKeyArn *string

	// A list of function layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add
	// to the function's execution environment. Specify each layer by its ARN,
	// including the version.
	Layers []*string

	// The amount of memory that your function has access to. Increasing the function's
	// memory also increases its CPU allocation. The default value is 128 MB. The value
	// must be a multiple of 64 MB.
	MemorySize *int32

	// Only update the function if the revision ID matches the ID that's specified. Use
	// this option to avoid modifying a function that has changed since you last read
	// it.
	RevisionId *string

	// The Amazon Resource Name (ARN) of the function's execution role.
	Role *string

	// The identifier of the function's runtime
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).
	Runtime types.Runtime

	// The amount of time that Lambda allows a function to run before stopping it. The
	// default is 3 seconds. The maximum allowed value is 900 seconds.
	Timeout *int32

	// Set Mode to Active to sample and trace a subset of incoming requests with AWS
	// X-Ray.
	TracingConfig *types.TracingConfig

	// For network connectivity to AWS resources in a VPC, specify a list of security
	// groups and subnets in the VPC. When you connect a function to a VPC, it can only
	// access resources and the internet through that VPC. For more information, see
	// VPC Settings
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).
	VpcConfig *types.VpcConfig
}

// Details about a function's configuration.
type UpdateFunctionConfigurationOutput struct {

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize *int64

	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig

	// The function's description.
	Description *string

	// The function's environment variables.
	Environment *types.EnvironmentResponse

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []*types.FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin executing your function.
	Handler *string

	// The KMS key that's used to encrypt the function's environment variables. This
	// key is only returned if you've configured a customer managed CMK.
	KMSKeyArn *string

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode

	// The function's  layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []*types.Layer

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string

	// The memory that's allocated to the function.
	MemorySize *int32

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The runtime environment for the Lambda function.
	Runtime types.Runtime

	// The current state of the function. When the state is Inactive, you can
	// reactivate the function by invoking it.
	State types.State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating, you
	// can't invoke or modify the function.
	StateReasonCode types.StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's AWS X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFunctionConfigurationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFunctionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFunctionConfiguration{}, middleware.After)
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
	addOpUpdateFunctionConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFunctionConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateFunctionConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "UpdateFunctionConfiguration",
	}
}
