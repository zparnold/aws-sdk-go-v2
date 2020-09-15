// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the targets registered with the maintenance window.
func (c *Client) DescribeMaintenanceWindowTargets(ctx context.Context, params *DescribeMaintenanceWindowTargetsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowTargetsOutput, error) {
	stack := middleware.NewStack("DescribeMaintenanceWindowTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeMaintenanceWindowTargetsMiddlewares(stack)
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
	addOpDescribeMaintenanceWindowTargetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowTargets(options.Region), middleware.Before)

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
			OperationName: "DescribeMaintenanceWindowTargets",
			Err:           err,
		}
	}
	out := result.(*DescribeMaintenanceWindowTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowTargetsInput struct {
	// Optional filters that can be used to narrow down the scope of the returned
	// window targets. The supported filter keys are Type, WindowTargetId and
	// OwnerInformation.
	Filters []*types.MaintenanceWindowFilter
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
	// The ID of the maintenance window whose targets should be retrieved.
	WindowId *string
}

type DescribeMaintenanceWindowTargetsOutput struct {
	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string
	// Information about the targets in the maintenance window.
	Targets []*types.MaintenanceWindowTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeMaintenanceWindowTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowTargets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowTargets",
	}
}
