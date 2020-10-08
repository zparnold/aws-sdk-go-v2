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

// Returns a list of DB snapshot attribute names and values for a manual DB
// snapshot. When sharing snapshots with other AWS accounts,
// DescribeDBSnapshotAttributes returns the restore attribute and a list of IDs for
// the AWS accounts that are authorized to copy or restore the manual DB snapshot.
// If all is included in the list of values for the restore attribute, then the
// manual DB snapshot is public and can be copied or restored by all AWS accounts.
// To add or remove access for an AWS account to copy or restore a manual DB
// snapshot, or to make the manual DB snapshot public or private, use the
// ModifyDBSnapshotAttribute API action.
func (c *Client) DescribeDBSnapshotAttributes(ctx context.Context, params *DescribeDBSnapshotAttributesInput, optFns ...func(*Options)) (*DescribeDBSnapshotAttributesOutput, error) {
	if params == nil {
		params = &DescribeDBSnapshotAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBSnapshotAttributes", params, optFns, addOperationDescribeDBSnapshotAttributesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBSnapshotAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBSnapshotAttributesInput struct {

	// The identifier for the DB snapshot to describe the attributes for.
	//
	// This member is required.
	DBSnapshotIdentifier *string
}

type DescribeDBSnapshotAttributesOutput struct {

	// Contains the results of a successful call to the DescribeDBSnapshotAttributes
	// API action. Manual DB snapshot attributes are used to authorize other AWS
	// accounts to copy or restore a manual DB snapshot. For more information, see the
	// ModifyDBSnapshotAttribute API action.
	DBSnapshotAttributesResult *types.DBSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBSnapshotAttributesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBSnapshotAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBSnapshotAttributes{}, middleware.After)
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
	addOpDescribeDBSnapshotAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBSnapshotAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBSnapshotAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBSnapshotAttributes",
	}
}
