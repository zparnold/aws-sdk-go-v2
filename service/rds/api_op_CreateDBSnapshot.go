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

// Creates a DBSnapshot. The source DBInstance must be in "available" state.
func (c *Client) CreateDBSnapshot(ctx context.Context, params *CreateDBSnapshotInput, optFns ...func(*Options)) (*CreateDBSnapshotOutput, error) {
	if params == nil {
		params = &CreateDBSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBSnapshot", params, optFns, addOperationCreateDBSnapshotMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBSnapshotInput struct {

	// The identifier of the DB instance that you want to create the snapshot of.
	// Constraints:
	//
	//     * Must match the identifier of an existing DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The identifier for the DB snapshot. Constraints:
	//
	//     * Can't be null, empty, or
	// blank
	//
	//     * Must contain from 1 to 255 letters, numbers, or hyphens
	//
	//     *
	// First character must be a letter
	//
	//     * Can't end with a hyphen or contain two
	// consecutive hyphens
	//
	// Example: my-snapshot-id
	//
	// This member is required.
	DBSnapshotIdentifier *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []*types.Tag
}

type CreateDBSnapshotOutput struct {

	// Contains the details of an Amazon RDS DB snapshot. This data type is used as a
	// response element in the DescribeDBSnapshots action.
	DBSnapshot *types.DBSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDBSnapshotMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBSnapshot{}, middleware.After)
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
	addOpCreateDBSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBSnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDBSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBSnapshot",
	}
}
