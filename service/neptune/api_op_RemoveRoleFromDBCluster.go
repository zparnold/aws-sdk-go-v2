// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates an Identity and Access Management (IAM) role from a DB cluster.
func (c *Client) RemoveRoleFromDBCluster(ctx context.Context, params *RemoveRoleFromDBClusterInput, optFns ...func(*Options)) (*RemoveRoleFromDBClusterOutput, error) {
	stack := middleware.NewStack("RemoveRoleFromDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRemoveRoleFromDBClusterMiddlewares(stack)
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
	addOpRemoveRoleFromDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveRoleFromDBCluster(options.Region), middleware.Before)
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
			OperationName: "RemoveRoleFromDBCluster",
			Err:           err,
		}
	}
	out := result.(*RemoveRoleFromDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveRoleFromDBClusterInput struct {

	// The name of the DB cluster to disassociate the IAM role from.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The Amazon Resource Name (ARN) of the IAM role to disassociate from the DB
	// cluster, for example arn:aws:iam::123456789012:role/NeptuneAccessRole.
	//
	// This member is required.
	RoleArn *string
}

type RemoveRoleFromDBClusterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRemoveRoleFromDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRemoveRoleFromDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveRoleFromDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveRoleFromDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveRoleFromDBCluster",
	}
}
