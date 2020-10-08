// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the default engine and system parameter information for the cluster
// database engine.
func (c *Client) DescribeEngineDefaultClusterParameters(ctx context.Context, params *DescribeEngineDefaultClusterParametersInput, optFns ...func(*Options)) (*DescribeEngineDefaultClusterParametersOutput, error) {
	if params == nil {
		params = &DescribeEngineDefaultClusterParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEngineDefaultClusterParameters", params, optFns, addOperationDescribeEngineDefaultClusterParametersMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEngineDefaultClusterParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEngineDefaultClusterParametersInput struct {

	// The name of the DB cluster parameter group family to return engine parameter
	// information for.
	//
	// This member is required.
	DBParameterGroupFamily *string

	// This parameter is not currently supported.
	Filters []*types.Filter

	// An optional pagination token provided by a previous
	// DescribeEngineDefaultClusterParameters request. If this parameter is specified,
	// the response includes only records beyond the marker, up to the value specified
	// by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeEngineDefaultClusterParametersOutput struct {

	// Contains the result of a successful invocation of the
	// DescribeEngineDefaultParameters () action.
	EngineDefaults *types.EngineDefaults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEngineDefaultClusterParametersMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEngineDefaultClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEngineDefaultClusterParameters{}, middleware.After)
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
	addOpDescribeEngineDefaultClusterParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEngineDefaultClusterParameters(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEngineDefaultClusterParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeEngineDefaultClusterParameters",
	}
}
