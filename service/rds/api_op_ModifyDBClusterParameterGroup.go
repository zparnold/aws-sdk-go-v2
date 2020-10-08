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

// Modifies the parameters of a DB cluster parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName, ParameterValue, and
// ApplyMethod. A maximum of 20 parameters can be modified in a single request. For
// more information on Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. Changes to dynamic parameters are applied
// immediately. Changes to static parameters require a reboot without failover to
// the DB cluster associated with the parameter group before the change can take
// effect. After you create a DB cluster parameter group, you should wait at least
// 5 minutes before creating your first DB cluster that uses that DB cluster
// parameter group as the default parameter group. This allows Amazon RDS to fully
// complete the create action before the parameter group is used as the default for
// a new DB cluster. This is especially important for parameters that are critical
// when creating the default database for a DB cluster, such as the character set
// for the default database defined by the character_set_database parameter. You
// can use the Parameter Groups option of the Amazon RDS console
// (https://console.aws.amazon.com/rds/) or the DescribeDBClusterParameters action
// to verify that your DB cluster parameter group has been created or modified. If
// the modified DB cluster parameter group is used by an Aurora Serverless cluster,
// Aurora applies the update immediately. The cluster restart might interrupt your
// workload. In that case, your application must reopen any connections and retry
// any transactions that were active when the parameter changes took effect. This
// action only applies to Aurora DB clusters.
func (c *Client) ModifyDBClusterParameterGroup(ctx context.Context, params *ModifyDBClusterParameterGroupInput, optFns ...func(*Options)) (*ModifyDBClusterParameterGroupOutput, error) {
	if params == nil {
		params = &ModifyDBClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterParameterGroup", params, optFns, addOperationModifyDBClusterParameterGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyDBClusterParameterGroupInput struct {

	// The name of the DB cluster parameter group to modify.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// A list of parameters in the DB cluster parameter group to modify.
	//
	// This member is required.
	Parameters []*types.Parameter
}

//
type ModifyDBClusterParameterGroupOutput struct {

	// The name of the DB cluster parameter group. Constraints:
	//
	//     * Must be 1 to 255
	// letters or numbers.
	//
	//     * First character must be a letter
	//
	//     * Can't end
	// with a hyphen or contain two consecutive hyphens
	//
	// This value is stored as a
	// lowercase string.
	DBClusterParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDBClusterParameterGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterParameterGroup{}, middleware.After)
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
	addOpModifyDBClusterParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyDBClusterParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterParameterGroup",
	}
}
