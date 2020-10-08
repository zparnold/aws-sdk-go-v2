// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all skills in the Alexa skill store by category.
func (c *Client) ListSkillsStoreSkillsByCategory(ctx context.Context, params *ListSkillsStoreSkillsByCategoryInput, optFns ...func(*Options)) (*ListSkillsStoreSkillsByCategoryOutput, error) {
	if params == nil {
		params = &ListSkillsStoreSkillsByCategoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSkillsStoreSkillsByCategory", params, optFns, addOperationListSkillsStoreSkillsByCategoryMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSkillsStoreSkillsByCategoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSkillsStoreSkillsByCategoryInput struct {

	// The category ID for which the skills are being retrieved from the skill store.
	//
	// This member is required.
	CategoryId *int64

	// The maximum number of skills returned per paginated calls.
	MaxResults *int32

	// The tokens used for pagination.
	NextToken *string
}

type ListSkillsStoreSkillsByCategoryOutput struct {

	// The tokens used for pagination.
	NextToken *string

	// The skill store skills.
	SkillsStoreSkills []*types.SkillsStoreSkill

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSkillsStoreSkillsByCategoryMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSkillsStoreSkillsByCategory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSkillsStoreSkillsByCategory{}, middleware.After)
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
	addOpListSkillsStoreSkillsByCategoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSkillsStoreSkillsByCategory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSkillsStoreSkillsByCategory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListSkillsStoreSkillsByCategory",
	}
}
