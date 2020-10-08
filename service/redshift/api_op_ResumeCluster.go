// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resumes a paused cluster.
func (c *Client) ResumeCluster(ctx context.Context, params *ResumeClusterInput, optFns ...func(*Options)) (*ResumeClusterOutput, error) {
	if params == nil {
		params = &ResumeClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResumeCluster", params, optFns, addOperationResumeClusterMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ResumeClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeClusterInput struct {

	// The identifier of the cluster to be resumed.
	//
	// This member is required.
	ClusterIdentifier *string
}

type ResumeClusterOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResumeClusterMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResumeCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResumeCluster{}, middleware.After)
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
	addOpResumeClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResumeCluster(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResumeCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ResumeCluster",
	}
}
