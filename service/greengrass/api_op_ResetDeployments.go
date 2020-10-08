// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets a group's deployments.
func (c *Client) ResetDeployments(ctx context.Context, params *ResetDeploymentsInput, optFns ...func(*Options)) (*ResetDeploymentsOutput, error) {
	if params == nil {
		params = &ResetDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetDeployments", params, optFns, addOperationResetDeploymentsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Information needed to reset deployments.
type ResetDeploymentsInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// If true, performs a best-effort only core reset.
	Force *bool
}

type ResetDeploymentsOutput struct {

	// The ARN of the deployment.
	DeploymentArn *string

	// The ID of the deployment.
	DeploymentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResetDeploymentsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpResetDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpResetDeployments{}, middleware.After)
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
	addOpResetDeploymentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetDeployments(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResetDeployments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ResetDeployments",
	}
}
