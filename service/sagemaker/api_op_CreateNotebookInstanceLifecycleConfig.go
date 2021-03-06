// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a lifecycle configuration that you can associate with a notebook
// instance. A lifecycle configuration is a collection of shell scripts that run
// when you create or start a notebook instance. Each lifecycle configuration
// script has a limit of 16384 characters. The value of the $PATH environment
// variable that is available to both scripts is /sbin:bin:/usr/sbin:/usr/bin. View
// CloudWatch Logs for notebook instance lifecycle configurations in log group
// /aws/sagemaker/NotebookInstances in log stream
// [notebook-instance-name]/[LifecycleConfigHook]. Lifecycle configuration scripts
// cannot run for longer than 5 minutes. If a script runs for longer than 5
// minutes, it fails and the notebook instance is not created or started. For
// information about notebook instance lifestyle configurations, see Step 2.1:
// (Optional) Customize a Notebook Instance
// (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html).
func (c *Client) CreateNotebookInstanceLifecycleConfig(ctx context.Context, params *CreateNotebookInstanceLifecycleConfigInput, optFns ...func(*Options)) (*CreateNotebookInstanceLifecycleConfigOutput, error) {
	stack := middleware.NewStack("CreateNotebookInstanceLifecycleConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateNotebookInstanceLifecycleConfigMiddlewares(stack)
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
	addOpCreateNotebookInstanceLifecycleConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNotebookInstanceLifecycleConfig(options.Region), middleware.Before)
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
			OperationName: "CreateNotebookInstanceLifecycleConfig",
			Err:           err,
		}
	}
	out := result.(*CreateNotebookInstanceLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNotebookInstanceLifecycleConfigInput struct {

	// The name of the lifecycle configuration.
	//
	// This member is required.
	NotebookInstanceLifecycleConfigName *string

	// A shell script that runs only once, when you create a notebook instance. The
	// shell script must be a base64-encoded string.
	OnCreate []*types.NotebookInstanceLifecycleHook

	// A shell script that runs every time you start a notebook instance, including
	// when you create the notebook instance. The shell script must be a base64-encoded
	// string.
	OnStart []*types.NotebookInstanceLifecycleHook
}

type CreateNotebookInstanceLifecycleConfigOutput struct {

	// The Amazon Resource Name (ARN) of the lifecycle configuration.
	NotebookInstanceLifecycleConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateNotebookInstanceLifecycleConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNotebookInstanceLifecycleConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNotebookInstanceLifecycleConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateNotebookInstanceLifecycleConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateNotebookInstanceLifecycleConfig",
	}
}
