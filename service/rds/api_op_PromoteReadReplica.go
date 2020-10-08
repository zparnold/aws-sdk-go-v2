// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Promotes a read replica DB instance to a standalone DB instance.
//
//     * Backup
// duration is a function of the amount of changes to the database since the
// previous backup. If you plan to promote a read replica to a standalone instance,
// we recommend that you enable backups and complete at least one backup prior to
// promotion. In addition, a read replica cannot be promoted to a standalone
// instance when it is in the backing-up status. If you have enabled backups on
// your read replica, configure the automated backup window so that daily backups
// do not interfere with read replica promotion.
//
//     * This command doesn't apply
// to Aurora MySQL and Aurora PostgreSQL.
//
//     </note>
func (c *Client) PromoteReadReplica(ctx context.Context, params *PromoteReadReplicaInput, optFns ...func(*Options)) (*PromoteReadReplicaOutput, error) {
	if params == nil {
		params = &PromoteReadReplicaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PromoteReadReplica", params, optFns, addOperationPromoteReadReplicaMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PromoteReadReplicaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type PromoteReadReplicaInput struct {

	// The DB instance identifier. This value is stored as a lowercase string.
	// Constraints:
	//
	//     * Must match the identifier of an existing read replica DB
	// instance.
	//
	// Example: mydbinstance
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The number of days for which automated backups are retained. Setting this
	// parameter to a positive number enables backups. Setting this parameter to 0
	// disables automated backups. Default: 1 Constraints:
	//
	//     * Must be a value from
	// 0 to 35.
	//
	//     * Can't be set to 0 if the DB instance is a source to read
	// replicas.
	BackupRetentionPeriod *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region. To see the time blocks available, see  Adjusting the Preferred
	// Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AdjustingTheMaintenanceWindow.html)
	// in the Amazon RDS User Guide. Constraints:
	//
	//     * Must be in the format
	// hh24:mi-hh24:mi.
	//
	//     * Must be in Universal Coordinated Time (UTC).
	//
	//     * Must
	// not conflict with the preferred maintenance window.
	//
	//     * Must be at least 30
	// minutes.
	PreferredBackupWindow *string
}

type PromoteReadReplicaOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPromoteReadReplicaMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPromoteReadReplica{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPromoteReadReplica{}, middleware.After)
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
	addOpPromoteReadReplicaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPromoteReadReplica(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPromoteReadReplica(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "PromoteReadReplica",
	}
}
