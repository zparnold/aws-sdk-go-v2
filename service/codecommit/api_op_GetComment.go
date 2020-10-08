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

// Returns the content of a comment made on a change, file, or commit in a
// repository. Reaction counts might include numbers from user identities who were
// deleted after the reaction was made. For a count of reactions from active
// identities, use GetCommentReactions.
func (c *Client) GetComment(ctx context.Context, params *GetCommentInput, optFns ...func(*Options)) (*GetCommentOutput, error) {
	if params == nil {
		params = &GetCommentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComment", params, optFns, addOperationGetCommentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCommentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommentInput struct {

	// The unique, system-generated ID of the comment. To get this ID, use
	// GetCommentsForComparedCommit () or GetCommentsForPullRequest ().
	//
	// This member is required.
	CommentId *string
}

type GetCommentOutput struct {

	// The contents of the comment.
	Comment *types.Comment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCommentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComment{}, middleware.After)
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
	addOpGetCommentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetComment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetComment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetComment",
	}
}
