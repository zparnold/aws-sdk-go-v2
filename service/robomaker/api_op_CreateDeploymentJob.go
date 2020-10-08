// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deploys a specific version of a robot application to robots in a fleet. The
// robot application must have a numbered applicationVersion for consistency
// reasons. To create a new version, use CreateRobotApplicationVersion or see
// Creating a Robot Application Version
// (https://docs.aws.amazon.com/robomaker/latest/dg/create-robot-application-version.html).
// After 90 days, deployment jobs expire and will be deleted. They will no longer
// be accessible.
func (c *Client) CreateDeploymentJob(ctx context.Context, params *CreateDeploymentJobInput, optFns ...func(*Options)) (*CreateDeploymentJobOutput, error) {
	if params == nil {
		params = &CreateDeploymentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeploymentJob", params, optFns, addOperationCreateDeploymentJobMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeploymentJobInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientRequestToken *string

	// The deployment application configuration.
	//
	// This member is required.
	DeploymentApplicationConfigs []*types.DeploymentApplicationConfig

	// The Amazon Resource Name (ARN) of the fleet to deploy.
	//
	// This member is required.
	Fleet *string

	// The requested deployment configuration.
	DeploymentConfig *types.DeploymentConfig

	// A map that contains tag keys and tag values that are attached to the deployment
	// job.
	Tags map[string]*string
}

type CreateDeploymentJobOutput struct {

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time

	// The deployment application configuration.
	DeploymentApplicationConfigs []*types.DeploymentApplicationConfig

	// The deployment configuration.
	DeploymentConfig *types.DeploymentConfig

	// The failure code of the simulation job if it failed: BadPermissionError  <p>AWS
	// Greengrass requires a service-level role permission to access other services.
	// The role must include the <a
	// href="https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSGreengrassResourceAccessRolePolicy$jsonEditor">
	// <code>AWSGreengrassResourceAccessRolePolicy</code> managed policy</a>. </p>
	// </dd> <dt>ExtractingBundleFailure</dt> <dd> <p>The robot application could not
	// be extracted from the bundle.</p> </dd> <dt>FailureThresholdBreached</dt> <dd>
	// <p>The percentage of robots that could not be updated exceeded the percentage
	// set for the deployment.</p> </dd> <dt>GreengrassDeploymentFailed</dt> <dd>
	// <p>The robot application could not be deployed to the robot.</p> </dd>
	// <dt>GreengrassGroupVersionDoesNotExist</dt> <dd> <p>The AWS Greengrass group or
	// version associated with a robot is missing.</p> </dd>
	// <dt>InternalServerError</dt> <dd> <p>An internal error has occurred. Retry your
	// request, but if the problem persists, contact us with details.</p> </dd>
	// <dt>MissingRobotApplicationArchitecture</dt> <dd> <p>The robot application does
	// not have a source that matches the architecture of the robot.</p> </dd>
	// <dt>MissingRobotDeploymentResource</dt> <dd> <p>One or more of the resources
	// specified for the robot application are missing. For example, does the robot
	// application have the correct launch package and launch file?</p> </dd>
	// <dt>PostLaunchFileFailure</dt> <dd> <p>The post-launch script failed.</p> </dd>
	// <dt>PreLaunchFileFailure</dt> <dd> <p>The pre-launch script failed.</p> </dd>
	// <dt>ResourceNotFound</dt> <dd> <p>One or more deployment resources are missing.
	// For example, do robot application source bundles still exist? </p> </dd>
	// <dt>RobotDeploymentNoResponse</dt> <dd> <p>There is no response from the robot.
	// It might not be powered on or connected to the internet.</p> </dd> </dl>
	FailureCode types.DeploymentJobErrorCode

	// The failure reason of the deployment job if it failed.
	FailureReason *string

	// The target fleet for the deployment job.
	Fleet *string

	// The status of the deployment job.
	Status types.DeploymentStatus

	// The list of all tags added to the deployment job.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDeploymentJobMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeploymentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeploymentJob{}, middleware.After)
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
	addIdempotencyToken_opCreateDeploymentJobMiddleware(stack, options)
	addOpCreateDeploymentJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeploymentJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateDeploymentJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDeploymentJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDeploymentJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDeploymentJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDeploymentJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDeploymentJobMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateDeploymentJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDeploymentJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateDeploymentJob",
	}
}
