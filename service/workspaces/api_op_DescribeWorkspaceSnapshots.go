// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the snapshots for the specified WorkSpace.
func (c *Client) DescribeWorkspaceSnapshots(ctx context.Context, params *DescribeWorkspaceSnapshotsInput, optFns ...func(*Options)) (*DescribeWorkspaceSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeWorkspaceSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaceSnapshots", params, optFns, addOperationDescribeWorkspaceSnapshotsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspaceSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceSnapshotsInput struct {

	// The identifier of the WorkSpace.
	//
	// This member is required.
	WorkspaceId *string
}

type DescribeWorkspaceSnapshotsOutput struct {

	// Information about the snapshots that can be used to rebuild a WorkSpace. These
	// snapshots include the user volume.
	RebuildSnapshots []*types.Snapshot

	// Information about the snapshots that can be used to restore a WorkSpace. These
	// snapshots include both the root volume and the user volume.
	RestoreSnapshots []*types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeWorkspaceSnapshotsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceSnapshots{}, middleware.After)
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
	addOpDescribeWorkspaceSnapshotsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceSnapshots(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeWorkspaceSnapshots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceSnapshots",
	}
}
