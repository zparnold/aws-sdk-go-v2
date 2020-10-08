// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates a robot application.
func (c *Client) UpdateRobotApplication(ctx context.Context, params *UpdateRobotApplicationInput, optFns ...func(*Options)) (*UpdateRobotApplicationOutput, error) {
	if params == nil {
		params = &UpdateRobotApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRobotApplication", params, optFns, addOperationUpdateRobotApplicationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRobotApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRobotApplicationInput struct {

	// The application information for the robot application.
	//
	// This member is required.
	Application *string

	// The robot software suite (ROS distribution) used by the robot application.
	//
	// This member is required.
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The sources of the robot application.
	//
	// This member is required.
	Sources []*types.SourceConfig

	// The revision id for the robot application.
	CurrentRevisionId *string
}

type UpdateRobotApplicationOutput struct {

	// The Amazon Resource Name (ARN) of the updated robot application.
	Arn *string

	// The time, in milliseconds since the epoch, when the robot application was last
	// updated.
	LastUpdatedAt *time.Time

	// The name of the robot application.
	Name *string

	// The revision id of the robot application.
	RevisionId *string

	// The robot software suite (ROS distribution) used by the robot application.
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The sources of the robot application.
	Sources []*types.Source

	// The version of the robot application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRobotApplicationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRobotApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRobotApplication{}, middleware.After)
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
	addOpUpdateRobotApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRobotApplication(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateRobotApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "UpdateRobotApplication",
	}
}
