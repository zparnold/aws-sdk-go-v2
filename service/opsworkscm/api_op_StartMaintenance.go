// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Manually starts server maintenance. This command can be useful if an earlier
// maintenance attempt failed, and the underlying cause of maintenance failure has
// been resolved. The server is in an UNDER_MAINTENANCE state while maintenance is
// in progress. Maintenance can only be started on servers in HEALTHY and UNHEALTHY
// states. Otherwise, an InvalidStateException is thrown. A
// ResourceNotFoundException is thrown when the server does not exist. A
// ValidationException is raised when parameters of the request are not valid.
func (c *Client) StartMaintenance(ctx context.Context, params *StartMaintenanceInput, optFns ...func(*Options)) (*StartMaintenanceOutput, error) {
	if params == nil {
		params = &StartMaintenanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMaintenance", params, optFns, addOperationStartMaintenanceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMaintenanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMaintenanceInput struct {

	// The name of the server on which to run maintenance.
	//
	// This member is required.
	ServerName *string

	// Engine attributes that are specific to the server on which you want to run
	// maintenance. Attributes accepted in a StartMaintenance request for Chef
	//
	//     *
	// CHEF_MAJOR_UPGRADE: If a Chef Automate server is eligible for upgrade to Chef
	// Automate 2, add this engine attribute to a StartMaintenance request and set the
	// value to true to upgrade the server to Chef Automate 2. For more information,
	// see Upgrade an AWS OpsWorks for Chef Automate Server to Chef Automate 2
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/opscm-a2upgrade.html).
	EngineAttributes []*types.EngineAttribute
}

type StartMaintenanceOutput struct {

	// Contains the response to a StartMaintenance request.
	Server *types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartMaintenanceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMaintenance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMaintenance{}, middleware.After)
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
	addOpStartMaintenanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartMaintenance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartMaintenance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "StartMaintenance",
	}
}
