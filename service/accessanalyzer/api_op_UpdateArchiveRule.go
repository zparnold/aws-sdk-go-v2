// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the criteria and values for the specified archive rule.
func (c *Client) UpdateArchiveRule(ctx context.Context, params *UpdateArchiveRuleInput, optFns ...func(*Options)) (*UpdateArchiveRuleOutput, error) {
	if params == nil {
		params = &UpdateArchiveRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateArchiveRule", params, optFns, addOperationUpdateArchiveRuleMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateArchiveRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates the specified archive rule.
type UpdateArchiveRuleInput struct {

	// The name of the analyzer to update the archive rules for.
	//
	// This member is required.
	AnalyzerName *string

	// A filter to match for the rules to update. Only rules that match the filter are
	// updated.
	//
	// This member is required.
	Filter map[string]*types.Criterion

	// The name of the rule to update.
	//
	// This member is required.
	RuleName *string

	// A client token.
	ClientToken *string
}

type UpdateArchiveRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateArchiveRuleMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateArchiveRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateArchiveRule{}, middleware.After)
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
	addOpUpdateArchiveRuleValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
