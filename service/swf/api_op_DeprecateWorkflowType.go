// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deprecates the specified workflow type. After a workflow type has been
// deprecated, you cannot create new executions of that type. Executions that were
// started before the type was deprecated continues to run. A deprecated workflow
// type may still be used when calling visibility actions. This operation is
// eventually consistent. The results are best effort and may not exactly reflect
// recent updates and changes. Access Control You can use IAM policies to control
// this action's access to Amazon SWF resources as follows:
//
//     * Use a Resource
// element with the domain name to limit the action to only specified domains.
//
//
// * Use an Action element to allow or deny permission to call this action.
//
//     *
// Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
//         * workflowType.name: String constraint. The key is
// swf:workflowType.name.
//
//         * workflowType.version: String constraint. The
// key is swf:workflowType.version.
//
// If the caller doesn't have sufficient
// permissions to invoke the action, or the parameter values fall outside the
// specified constraints, the action fails. The associated event attribute's cause
// parameter is set to OPERATION_NOT_PERMITTED. For details and example IAM
// policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) DeprecateWorkflowType(ctx context.Context, params *DeprecateWorkflowTypeInput, optFns ...func(*Options)) (*DeprecateWorkflowTypeOutput, error) {
	stack := middleware.NewStack("DeprecateWorkflowType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDeprecateWorkflowTypeMiddlewares(stack)
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
	addOpDeprecateWorkflowTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeprecateWorkflowType(options.Region), middleware.Before)
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
			OperationName: "DeprecateWorkflowType",
			Err:           err,
		}
	}
	out := result.(*DeprecateWorkflowTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeprecateWorkflowTypeInput struct {

	// The name of the domain in which the workflow type is registered.
	//
	// This member is required.
	Domain *string

	// The workflow type to deprecate.
	//
	// This member is required.
	WorkflowType *types.WorkflowType
}

type DeprecateWorkflowTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDeprecateWorkflowTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDeprecateWorkflowType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeprecateWorkflowType{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeprecateWorkflowType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "DeprecateWorkflowType",
	}
}
