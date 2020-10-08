// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the executions of a maintenance window. This includes information about
// when the maintenance window was scheduled to be active, and information about
// tasks registered and run with the maintenance window.
func (c *Client) DescribeMaintenanceWindowExecutions(ctx context.Context, params *DescribeMaintenanceWindowExecutionsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionsOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowExecutions", params, optFns, addOperationDescribeMaintenanceWindowExecutionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowExecutionsInput struct {

	// The ID of the maintenance window whose executions should be retrieved.
	//
	// This member is required.
	WindowId *string

	// Each entry in the array is a structure containing: Key (string, between 1 and
	// 128 characters) Values (array of strings, each string is between 1 and 256
	// characters) The supported Keys are ExecutedBefore and ExecutedAfter with the
	// value being a date/time string such as 2016-11-04T05:00:00Z.
	Filters []*types.MaintenanceWindowFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeMaintenanceWindowExecutionsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the maintenance window executions.
	WindowExecutions []*types.MaintenanceWindowExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMaintenanceWindowExecutionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowExecutions{}, middleware.After)
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
	addOpDescribeMaintenanceWindowExecutionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowExecutions",
	}
}
