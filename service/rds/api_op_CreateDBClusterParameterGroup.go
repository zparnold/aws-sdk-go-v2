// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new DB cluster parameter group. Parameters in a DB cluster parameter
// group apply to all of the instances in a DB cluster. A DB cluster parameter
// group is initially created with the default parameters for the database engine
// used by instances in the DB cluster. To provide custom values for any of the
// parameters, you must modify the group after creating it using
// ModifyDBClusterParameterGroup. Once you've created a DB cluster parameter group,
// you need to associate it with your DB cluster using ModifyDBCluster. When you
// associate a new DB cluster parameter group with a running DB cluster, you need
// to reboot the DB instances in the DB cluster without failover for the new DB
// cluster parameter group and associated settings to take effect. After you create
// a DB cluster parameter group, you should wait at least 5 minutes before creating
// your first DB cluster that uses that DB cluster parameter group as the default
// parameter group. This allows Amazon RDS to fully complete the create action
// before the DB cluster parameter group is used as the default for a new DB
// cluster. This is especially important for parameters that are critical when
// creating the default database for a DB cluster, such as the character set for
// the default database defined by the character_set_database parameter. You can
// use the Parameter Groups option of the Amazon RDS console
// (https://console.aws.amazon.com/rds/) or the DescribeDBClusterParameters action
// to verify that your DB cluster parameter group has been created or modified.
// <p>For more information on Amazon Aurora, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html">
// What Is Amazon Aurora?</a> in the <i>Amazon Aurora User Guide.</i> </p> <note>
// <p>This action only applies to Aurora DB clusters.</p> </note>
func (c *Client) CreateDBClusterParameterGroup(ctx context.Context, params *CreateDBClusterParameterGroupInput, optFns ...func(*Options)) (*CreateDBClusterParameterGroupOutput, error) {
	stack := middleware.NewStack("CreateDBClusterParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBClusterParameterGroupMiddlewares(stack)
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
	addOpCreateDBClusterParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBClusterParameterGroup(options.Region), middleware.Before)

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
			OperationName: "CreateDBClusterParameterGroup",
			Err:           err,
		}
	}
	out := result.(*CreateDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBClusterParameterGroupInput struct {
	// Tags to assign to the DB cluster parameter group.
	Tags []*types.Tag
	// The description for the DB cluster parameter group.
	Description *string
	// The DB cluster parameter group family name. A DB cluster parameter group can be
	// associated with one and only one DB cluster parameter group family, and can be
	// applied only to a DB cluster running a database engine and engine version
	// compatible with that DB cluster parameter group family. Aurora MySQL Example:
	// aurora5.6, aurora-mysql5.7 Aurora PostgreSQL Example: aurora-postgresql9.6
	DBParameterGroupFamily *string
	// The name of the DB cluster parameter group. Constraints:
	//
	//     * Must match the
	// name of an existing DB cluster parameter group.
	//
	// This value is stored as a
	// lowercase string.
	DBClusterParameterGroupName *string
}

type CreateDBClusterParameterGroupOutput struct {
	// Contains the details of an Amazon RDS DB cluster parameter group. This data type
	// is used as a response element in the DescribeDBClusterParameterGroups action.
	DBClusterParameterGroup *types.DBClusterParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBClusterParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBClusterParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBClusterParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBClusterParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBClusterParameterGroup",
	}
}
