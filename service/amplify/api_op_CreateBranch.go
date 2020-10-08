// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new branch for an Amplify app.
func (c *Client) CreateBranch(ctx context.Context, params *CreateBranchInput, optFns ...func(*Options)) (*CreateBranchOutput, error) {
	if params == nil {
		params = &CreateBranchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBranch", params, optFns, addOperationCreateBranchMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBranchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the create branch request.
type CreateBranchInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name for the branch.
	//
	// This member is required.
	BranchName *string

	// The Amazon Resource Name (ARN) for a backend environment that is part of an
	// Amplify app.
	BackendEnvironmentArn *string

	// The basic authorization credentials for the branch.
	BasicAuthCredentials *string

	// The build specification (build spec) for the branch.
	BuildSpec *string

	// The description for the branch.
	Description *string

	// The display name for a branch. This is used as the default domain prefix.
	DisplayName *string

	// Enables auto building for the branch.
	EnableAutoBuild *bool

	// Enables basic authorization for the branch.
	EnableBasicAuth *bool

	// Enables notifications for the branch.
	EnableNotification *bool

	// Enables pull request preview for this branch.
	EnablePullRequestPreview *bool

	// The environment variables for the branch.
	EnvironmentVariables map[string]*string

	// The framework for the branch.
	Framework *string

	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName *string

	// Describes the current stage for the branch.
	Stage types.Stage

	// The tag for the branch.
	Tags map[string]*string

	// The content Time To Live (TTL) for the website in seconds.
	Ttl *string
}

// The result structure for create branch request.
type CreateBranchOutput struct {

	// Describes the branch for an Amplify app, which maps to a third-party repository
	// branch.
	//
	// This member is required.
	Branch *types.Branch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBranchMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBranch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBranch{}, middleware.After)
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
	addOpCreateBranchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBranch(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateBranch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "CreateBranch",
	}
}
