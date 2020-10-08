// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the contents of one or more commits in a repository.
func (c *Client) BatchGetCommits(ctx context.Context, params *BatchGetCommitsInput, optFns ...func(*Options)) (*BatchGetCommitsOutput, error) {
	if params == nil {
		params = &BatchGetCommitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetCommits", params, optFns, addOperationBatchGetCommitsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetCommitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetCommitsInput struct {

	// The full commit IDs of the commits to get information about. You must supply the
	// full SHA IDs of each commit. You cannot use shortened SHA IDs.
	//
	// This member is required.
	CommitIds []*string

	// The name of the repository that contains the commits.
	//
	// This member is required.
	RepositoryName *string
}

type BatchGetCommitsOutput struct {

	// An array of commit data type objects, each of which contains information about a
	// specified commit.
	Commits []*types.Commit

	// Returns any commit IDs for which information could not be found. For example, if
	// one of the commit IDs was a shortened SHA ID or that commit was not found in the
	// specified repository, the ID returns an error object with more information.
	Errors []*types.BatchGetCommitsError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetCommitsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetCommits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetCommits{}, middleware.After)
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
	addOpBatchGetCommitsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetCommits(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchGetCommits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "BatchGetCommits",
	}
}
