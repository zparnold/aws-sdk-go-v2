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

// Retrieves information about an archive rule.
func (c *Client) GetArchiveRule(ctx context.Context, params *GetArchiveRuleInput, optFns ...func(*Options)) (*GetArchiveRuleOutput, error) {
	if params == nil {
		params = &GetArchiveRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetArchiveRule", params, optFns, addOperationGetArchiveRuleMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetArchiveRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves an archive rule.
type GetArchiveRuleInput struct {

	// The name of the analyzer to retrieve rules from.
	//
	// This member is required.
	AnalyzerName *string

	// The name of the rule to retrieve.
	//
	// This member is required.
	RuleName *string
}

// The response to the request.
type GetArchiveRuleOutput struct {

	// Contains information about an archive rule.
	//
	// This member is required.
	ArchiveRule *types.ArchiveRuleSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetArchiveRuleMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetArchiveRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetArchiveRule{}, middleware.After)
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
	addOpGetArchiveRuleValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
