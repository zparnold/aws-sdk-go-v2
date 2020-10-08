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

// Creates a robot application.
func (c *Client) CreateRobotApplication(ctx context.Context, params *CreateRobotApplicationInput, optFns ...func(*Options)) (*CreateRobotApplicationOutput, error) {
	if params == nil {
		params = &CreateRobotApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRobotApplication", params, optFns, addOperationCreateRobotApplicationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRobotApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRobotApplicationInput struct {

	// The name of the robot application.
	//
	// This member is required.
	Name *string

	// The robot software suite (ROS distribuition) used by the robot application.
	//
	// This member is required.
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The sources of the robot application.
	//
	// This member is required.
	Sources []*types.SourceConfig

	// A map that contains tag keys and tag values that are attached to the robot
	// application.
	Tags map[string]*string
}

type CreateRobotApplicationOutput struct {

	// The Amazon Resource Name (ARN) of the robot application.
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

	// The list of all tags added to the robot application.
	Tags map[string]*string

	// The version of the robot application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRobotApplicationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRobotApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRobotApplication{}, middleware.After)
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
	addOpCreateRobotApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRobotApplication(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateRobotApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateRobotApplication",
	}
}
