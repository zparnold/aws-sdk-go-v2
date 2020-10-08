// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Runs deployment or stack commands. For more information, see Deploying Apps
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-deploying.html)
// and Run Stack Commands
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-commands.html).
// Required Permissions: To use this action, an IAM user must have a Deploy or
// Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) {
	if params == nil {
		params = &CreateDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeployment", params, optFns, addOperationCreateDeploymentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeploymentInput struct {

	// A DeploymentCommand object that specifies the deployment command and any
	// associated arguments.
	//
	// This member is required.
	Command *types.DeploymentCommand

	// The stack ID.
	//
	// This member is required.
	StackId *string

	// The app ID. This parameter is required for app deployments, but not for other
	// deployment commands.
	AppId *string

	// A user-defined comment.
	Comment *string

	// A string that contains user-defined, custom JSON. You can use this parameter to
	// override some corresponding default stack configuration JSON values. The string
	// should be in the following format: "{\"key1\": \"value1\", \"key2\":
	// \"value2\",...}" For more information about custom JSON, see Use Custom JSON to
	// Modify the Stack Configuration Attributes
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-json.html)
	// and Overriding Attributes With Custom JSON
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook-json-override.html).
	CustomJson *string

	// The instance IDs for the deployment targets.
	InstanceIds []*string

	// The layer IDs for the deployment targets.
	LayerIds []*string
}

// Contains the response to a CreateDeployment request.
type CreateDeploymentOutput struct {

	// The deployment ID, which can be used with other requests to identify the
	// deployment.
	DeploymentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDeploymentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDeployment{}, middleware.After)
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
	addOpCreateDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeployment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "CreateDeployment",
	}
}
