// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of RepositoryAssociationSummary
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html)
// objects that contain summary information about a repository association. You can
// filter the returned list by ProviderType
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-ProviderType),
// Name
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-Name),
// State
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-State),
// and Owner
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-Owner).
func (c *Client) ListRepositoryAssociations(ctx context.Context, params *ListRepositoryAssociationsInput, optFns ...func(*Options)) (*ListRepositoryAssociationsOutput, error) {
	stack := middleware.NewStack("ListRepositoryAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListRepositoryAssociationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositoryAssociations(options.Region), middleware.Before)
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
			OperationName: "ListRepositoryAssociations",
			Err:           err,
		}
	}
	out := result.(*ListRepositoryAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositoryAssociationsInput struct {

	// The maximum number of repository association results returned by
	// ListRepositoryAssociations in paginated output. When this parameter is used,
	// ListRepositoryAssociations only returns maxResults results in a single page with
	// a nextToken response element. The remaining results of the initial request can
	// be seen by sending another ListRepositoryAssociations request with the returned
	// nextToken value. This value can be between 1 and 100. If this parameter is not
	// used, ListRepositoryAssociations returns up to 100 results and a nextToken value
	// if applicable.
	MaxResults *int32

	// List of repository names to use as a filter.
	Names []*string

	// The nextToken value returned from a previous paginated
	// ListRepositoryAssociations request where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value. Treat this token as an
	// opaque identifier that is only used to retrieve the next items in a list and not
	// for other programmatic purposes.
	NextToken *string

	// List of owners to use as a filter. For AWS CodeCommit, it is the name of the
	// CodeCommit account that was used to associate the repository. For other
	// repository source providers, such as Bitbucket, this is name of the account that
	// was used to associate the repository.
	Owners []*string

	// List of provider types to use as a filter.
	ProviderTypes []types.ProviderType

	// List of repository association states to use as a filter. The valid repository
	// association states are:
	//
	//     * Associated: The repository association is
	// complete.
	//
	//     * Associating: CodeGuru Reviewer is:
	//
	//         * Setting up pull
	// request notifications. This is required for pull requests to trigger a CodeGuru
	// Reviewer review. If your repository ProviderType is GitHub or Bitbucket,
	// CodeGuru Reviewer creates webhooks in your repository to trigger CodeGuru
	// Reviewer reviews. If you delete these webhooks, reviews of code in your
	// repository cannot be triggered.
	//
	//         * Setting up source code access. This
	// is required for CodeGuru Reviewer to securely clone code in your repository.
	//
	//
	// * Failed: The repository failed to associate or disassociate.
	//
	//     *
	// Disassociating: CodeGuru Reviewer is removing the repository's pull request
	// notifications and source code access.
	States []types.RepositoryAssociationState
}

type ListRepositoryAssociationsOutput struct {

	// The nextToken value to include in a future ListRecommendations request. When the
	// results of a ListRecommendations request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string

	// A list of repository associations that meet the criteria of the request.
	RepositoryAssociationSummaries []*types.RepositoryAssociationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListRepositoryAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListRepositoryAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListRepositoryAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRepositoryAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "ListRepositoryAssociations",
	}
}
