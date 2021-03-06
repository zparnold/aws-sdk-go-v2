// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a snapshot schedule for a cluster.
func (c *Client) ModifyClusterSnapshotSchedule(ctx context.Context, params *ModifyClusterSnapshotScheduleInput, optFns ...func(*Options)) (*ModifyClusterSnapshotScheduleOutput, error) {
	stack := middleware.NewStack("ModifyClusterSnapshotSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyClusterSnapshotScheduleMiddlewares(stack)
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
	addOpModifyClusterSnapshotScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterSnapshotSchedule(options.Region), middleware.Before)
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
			OperationName: "ModifyClusterSnapshotSchedule",
			Err:           err,
		}
	}
	out := result.(*ModifyClusterSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterSnapshotScheduleInput struct {

	// A unique identifier for the cluster whose snapshot schedule you want to modify.
	//
	// This member is required.
	ClusterIdentifier *string

	// A boolean to indicate whether to remove the assoiciation between the cluster and
	// the schedule.
	DisassociateSchedule *bool

	// A unique alphanumeric identifier for the schedule that you want to associate
	// with the cluster.
	ScheduleIdentifier *string
}

type ModifyClusterSnapshotScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyClusterSnapshotScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterSnapshotSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterSnapshotSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyClusterSnapshotSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterSnapshotSchedule",
	}
}
