// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a deployment strategy.
func (c *Client) UpdateDeploymentStrategy(ctx context.Context, params *UpdateDeploymentStrategyInput, optFns ...func(*Options)) (*UpdateDeploymentStrategyOutput, error) {
	stack := middleware.NewStack("UpdateDeploymentStrategy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDeploymentStrategyMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateDeploymentStrategyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDeploymentStrategy(options.Region), middleware.Before)

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
			OperationName: "UpdateDeploymentStrategy",
			Err:           err,
		}
	}
	out := result.(*UpdateDeploymentStrategyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDeploymentStrategyInput struct {
	// A description of the deployment strategy.
	Description *string
	// The algorithm used to define how percentage grows over time. AWS AppConfig
	// supports the following growth types: Linear: For this type, AppConfig processes
	// the deployment by increments of the growth factor evenly distributed over the
	// deployment time. For example, a linear deployment that uses a growth factor of
	// 20 initially makes the configuration available to 20 percent of the targets.
	// After 1/5th of the deployment time has passed, the system updates the percentage
	// to 40 percent. This continues until 100% of the targets are set to receive the
	// deployed configuration.  <p> <b>Exponential</b>: For this type, AppConfig
	// processes the deployment exponentially using the following formula:
	// <code>G*(2^N)</code>. In this formula, <code>G</code> is the growth factor
	// specified by the user and <code>N</code> is the number of steps until the
	// configuration is deployed to all targets. For example, if you specify a growth
	// factor of 2, then the system rolls out the configuration as follows:</p> <p>
	// <code>2*(2^0)</code> </p> <p> <code>2*(2^1)</code> </p> <p> <code>2*(2^2)</code>
	// </p> <p>Expressed numerically, the deployment rolls out as follows: 2% of the
	// targets, 4% of the targets, 8% of the targets, and continues until the
	// configuration has been deployed to all targets.</p>
	GrowthType types.GrowthType
	// Total amount of time for a deployment to last.
	DeploymentDurationInMinutes *int32
	// The deployment strategy ID.
	DeploymentStrategyId *string
	// The amount of time AppConfig monitors for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int32
	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor *float32
}

type UpdateDeploymentStrategyOutput struct {
	// The algorithm used to define how percentage grew over time.
	GrowthType types.GrowthType
	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes *int32
	// Save the deployment strategy to a Systems Manager (SSM) document.
	ReplicateTo types.ReplicateTo
	// The deployment strategy ID.
	Id *string
	// The description of the deployment strategy.
	Description *string
	// The amount of time AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int32
	// The name of the deployment strategy.
	Name *string
	// The percentage of targets that received a deployed configuration during each
	// interval.
	GrowthFactor *float32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDeploymentStrategyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDeploymentStrategy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDeploymentStrategy{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDeploymentStrategy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "UpdateDeploymentStrategy",
	}
}