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

// Stops an Amazon Aurora DB cluster. When you stop a DB cluster, Aurora retains
// the DB cluster's metadata, including its endpoints and DB parameter groups.
// Aurora also retains the transaction logs so you can do a point-in-time restore
// if necessary.  <p>For more information, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-cluster-stop-start.html">
// Stopping and Starting an Aurora Cluster</a> in the <i>Amazon Aurora User
// Guide.</i> </p> <note> <p>This action only applies to Aurora DB clusters.</p>
// </note>
func (c *Client) StopDBCluster(ctx context.Context, params *StopDBClusterInput, optFns ...func(*Options)) (*StopDBClusterOutput, error) {
	if params == nil {
		params = &StopDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDBCluster", params, optFns, addOperationStopDBClusterMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDBClusterInput struct {

	// The DB cluster identifier of the Amazon Aurora DB cluster to be stopped. This
	// parameter is stored as a lowercase string.
	//
	// This member is required.
	DBClusterIdentifier *string
}

type StopDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster. This data type is used as a
	// response element in the DescribeDBClusters, StopDBCluster, and StartDBCluster
	// actions.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopDBClusterMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStopDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStopDBCluster{}, middleware.After)
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
	addOpStopDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDBCluster(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StopDBCluster",
	}
}
