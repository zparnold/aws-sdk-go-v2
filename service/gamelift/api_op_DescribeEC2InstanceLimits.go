// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the following information for the specified EC2 instance type:
//
//     *
// Maximum number of instances allowed per AWS account (service limit).
//
//     *
// Current usage for the AWS account.
//
// To learn more about the capabilities of each
// instance type, see Amazon EC2 Instance Types
// (http://aws.amazon.com/ec2/instance-types/). Note that the instance types
// offered may vary depending on the region. Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
//     * CreateFleet ()
//
//     * ListFleets ()
//
//     * DeleteFleet
// ()
//
//     * DescribeFleetAttributes ()
//
//     * UpdateFleetAttributes ()
//
//     *
// StartFleetActions () or StopFleetActions ()
func (c *Client) DescribeEC2InstanceLimits(ctx context.Context, params *DescribeEC2InstanceLimitsInput, optFns ...func(*Options)) (*DescribeEC2InstanceLimitsOutput, error) {
	if params == nil {
		params = &DescribeEC2InstanceLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEC2InstanceLimits", params, optFns, addOperationDescribeEC2InstanceLimitsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEC2InstanceLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeEC2InstanceLimitsInput struct {

	// Name of an EC2 instance type that is supported in Amazon GameLift. A fleet
	// instance type determines the computing resources of each instance in the fleet,
	// including CPU, memory, storage, and networking capacity. Amazon GameLift
	// supports the following EC2 instance types. See Amazon EC2 Instance Types
	// (http://aws.amazon.com/ec2/instance-types/) for detailed descriptions. Leave
	// this parameter blank to retrieve limits for all types.
	EC2InstanceType types.EC2InstanceType
}

// Represents the returned data in response to a request action.
type DescribeEC2InstanceLimitsOutput struct {

	// The maximum number of instances for the specified instance type.
	EC2InstanceLimits []*types.EC2InstanceLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEC2InstanceLimitsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEC2InstanceLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEC2InstanceLimits{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEC2InstanceLimits(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEC2InstanceLimits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeEC2InstanceLimits",
	}
}
