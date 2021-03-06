// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing workflow.
func (c *Client) UpdateWorkflow(ctx context.Context, params *UpdateWorkflowInput, optFns ...func(*Options)) (*UpdateWorkflowOutput, error) {
	stack := middleware.NewStack("UpdateWorkflow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateWorkflowMiddlewares(stack)
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
	addOpUpdateWorkflowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkflow(options.Region), middleware.Before)
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
			OperationName: "UpdateWorkflow",
			Err:           err,
		}
	}
	out := result.(*UpdateWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkflowInput struct {

	// Name of the workflow to be updated.
	//
	// This member is required.
	Name *string

	// A collection of properties to be used as part of each execution of the workflow.
	DefaultRunProperties map[string]*string

	// The description of the workflow.
	Description *string
}

type UpdateWorkflowOutput struct {

	// The name of the workflow which was specified in input.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateWorkflowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWorkflow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWorkflow{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateWorkflow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateWorkflow",
	}
}
