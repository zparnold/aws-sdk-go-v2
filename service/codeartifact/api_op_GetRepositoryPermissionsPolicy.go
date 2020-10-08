// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the resource policy that is set on a repository.
func (c *Client) GetRepositoryPermissionsPolicy(ctx context.Context, params *GetRepositoryPermissionsPolicyInput, optFns ...func(*Options)) (*GetRepositoryPermissionsPolicyOutput, error) {
	if params == nil {
		params = &GetRepositoryPermissionsPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRepositoryPermissionsPolicy", params, optFns, addOperationGetRepositoryPermissionsPolicyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRepositoryPermissionsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRepositoryPermissionsPolicyInput struct {

	// The name of the domain containing the repository whose associated resource
	// policy is to be retrieved.
	//
	// This member is required.
	Domain *string

	// The name of the repository whose associated resource policy is to be retrieved.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
}

type GetRepositoryPermissionsPolicyOutput struct {

	// The returned resource policy.
	Policy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRepositoryPermissionsPolicyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRepositoryPermissionsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRepositoryPermissionsPolicy{}, middleware.After)
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
	addOpGetRepositoryPermissionsPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRepositoryPermissionsPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRepositoryPermissionsPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "GetRepositoryPermissionsPolicy",
	}
}
