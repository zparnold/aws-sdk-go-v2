// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon Neptune DB cluster. You can use the
// ReplicationSourceIdentifier parameter to create the DB cluster as a Read Replica
// of another DB cluster or Amazon Neptune DB instance. Note that when you create a
// new cluster using CreateDBCluster directly, deletion protection is disabled by
// default (when you create a new production cluster in the console, deletion
// protection is enabled by default). You can only delete a DB cluster if its
// DeletionProtection field is set to false.
func (c *Client) CreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*Options)) (*CreateDBClusterOutput, error) {
	stack := middleware.NewStack("CreateDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBClusterMiddlewares(stack)
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
	addOpCreateDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBCluster(options.Region), middleware.Before)

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
			OperationName: "CreateDBCluster",
			Err:           err,
		}
	}
	out := result.(*CreateDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBClusterInput struct {
	// True to enable mapping of AWS Identity and Access Management (IAM) accounts to
	// database accounts, and otherwise false. Default: false
	EnableIAMDatabaseAuthentication *bool
	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this
	// DB cluster is created as a Read Replica.
	ReplicationSourceIdentifier *string
	// The tags to assign to the new DB cluster.
	Tags []*types.Tag
	// The password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@". Constraints: Must contain
	// from 8 to 41 characters.
	MasterUserPassword *string
	// The DB cluster identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	//     * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//
	// * First character must be a letter.
	//
	//     * Cannot end with a hyphen or contain
	// two consecutive hyphens.
	//
	// Example: my-cluster1
	DBClusterIdentifier *string
	// The number of days for which automated backups are retained. You must specify a
	// minimum value of 1. Default: 1 Constraints:
	//
	//     * Must be a value from 1 to 35
	BackupRetentionPeriod *int32
	// (Not supported by Neptune)
	OptionGroupName *string
	// The name of the DB cluster parameter group to associate with this DB cluster. If
	// this argument is omitted, the default is used. Constraints:
	//
	//     * If supplied,
	// must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string
	// The port number on which the instances in the DB cluster accept connections.
	// Default: 8182
	Port *int32
	// The version number of the database engine to use. Currently, setting this
	// parameter has no effect. Example: 1.0.1
	EngineVersion *string
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool
	// (Not supported by Neptune)
	CharacterSetName *string
	// A DB subnet group to associate with this DB cluster. Constraints: Must match the
	// name of an existing DBSubnetGroup. Must not be default. Example: mySubnetgroup
	DBSubnetGroupName *string
	// A list of EC2 VPC security groups to associate with this DB cluster.
	VpcSecurityGroupIds []*string
	// A list of EC2 Availability Zones that instances in the DB cluster can be created
	// in.
	AvailabilityZones []*string
	// The name of the database engine to be used for this DB cluster. Valid Values:
	// neptune
	Engine *string
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	EnableCloudwatchLogsExports []*string
	// The name for your database of up to 64 alpha-numeric characters. If you do not
	// provide a name, Amazon Neptune will not create a database in the DB cluster you
	// are creating.
	DatabaseName *string
	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is enabled.
	DeletionProtection *bool
	// The name of the master user for the DB cluster. Constraints:
	//
	//     * Must be 1 to
	// 16 letters or numbers.
	//
	//     * First character must be a letter.
	//
	//     * Cannot be
	// a reserved word for the chosen database engine.
	MasterUsername *string
	// This parameter is not currently supported.
	PreSignedUrl *string
	// The AWS KMS key identifier for an encrypted DB cluster. The KMS key identifier
	// is the Amazon Resource Name (ARN) for the KMS encryption key. If you are
	// creating a DB cluster with the same AWS account that owns the KMS encryption key
	// used to encrypt the new DB cluster, then you can use the KMS key alias instead
	// of the ARN for the KMS encryption key. If an encryption key is not specified in
	// KmsKeyId:
	//
	//     * If ReplicationSourceIdentifier identifies an encrypted source,
	// then Amazon Neptune will use the encryption key used to encrypt the source.
	// Otherwise, Amazon Neptune will use your default encryption key.
	//
	//     * If the
	// StorageEncrypted parameter is true and ReplicationSourceIdentifier is not
	// specified, then Amazon Neptune will use your default encryption key.
	//
	// AWS KMS
	// creates the default encryption key for your AWS account. Your AWS account has a
	// different default encryption key for each AWS Region. If you create a Read
	// Replica of an encrypted DB cluster in another AWS Region, you must set KmsKeyId
	// to a KMS key ID that is valid in the destination AWS Region. This key is used to
	// encrypt the Read Replica in that AWS Region.
	KmsKeyId *string
	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region. To see the time blocks available, see  Adjusting the Preferred
	// Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AdjustingTheMaintenanceWindow.html)
	// in the Amazon Neptune User Guide. Constraints:
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
	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. To see the time blocks available,
	// see  Adjusting the Preferred Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AdjustingTheMaintenanceWindow.html)
	// in the Amazon Neptune User Guide. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string
}

type CreateDBClusterOutput struct {
	// Contains the details of an Amazon Neptune DB cluster. This data type is used as
	// a response element in the DescribeDBClusters () action.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBCluster",
	}
}
