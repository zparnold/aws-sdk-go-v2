// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attempts to stop an ongoing deployment.
func (c *Client) StopDeployment(ctx context.Context, params *StopDeploymentInput, optFns ...func(*Options)) (*StopDeploymentOutput, error) {
	if params == nil {
		params = &StopDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDeployment", params, optFns, addOperationStopDeploymentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a StopDeployment operation.
type StopDeploymentInput struct {

	// The unique ID of a deployment.
	//
	// This member is required.
	DeploymentId *string

	// Indicates, when a deployment is stopped, whether instances that have been
	// updated should be rolled back to the previous version of the application
	// revision.
	AutoRollbackEnabled *bool
}

// Represents the output of a StopDeployment operation.
type StopDeploymentOutput struct {

	// The status of the stop deployment operation:
	//
	//     * Pending: The stop operation
	// is pending.
	//
	//     * Succeeded: The stop operation was successful.
	Status types.StopStatus

	// An accompanying status message.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopDeploymentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDeployment{}, middleware.After)
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
	addOpStopDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDeployment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "StopDeployment",
	}
}
