// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// You might need to reboot your instance, usually for maintenance reasons. For
// example, if you make certain changes, or if you change the cluster parameter
// group that is associated with the instance, you must reboot the instance for the
// changes to take effect. Rebooting an instance restarts the database engine
// service. Rebooting an instance results in a momentary outage, during which the
// instance status is set to rebooting.
func (c *Client) RebootDBInstance(ctx context.Context, params *RebootDBInstanceInput, optFns ...func(*Options)) (*RebootDBInstanceOutput, error) {
	stack := middleware.NewStack("RebootDBInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRebootDBInstanceMiddlewares(stack)
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
	addOpRebootDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebootDBInstance(options.Region), middleware.Before)
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
			OperationName: "RebootDBInstance",
			Err:           err,
		}
	}
	out := result.(*RebootDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to RebootDBInstance ().
type RebootDBInstanceInput struct {

	// The instance identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	//     * Must match the identifier of an existing DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// When true, the reboot is conducted through a Multi-AZ failover. Constraint: You
	// can't specify true if the instance is not configured for Multi-AZ.
	ForceFailover *bool
}

type RebootDBInstanceOutput struct {

	// Detailed information about an instance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRebootDBInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRebootDBInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRebootDBInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opRebootDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RebootDBInstance",
	}
}
