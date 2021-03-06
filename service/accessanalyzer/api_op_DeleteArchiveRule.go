// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified archive rule.
func (c *Client) DeleteArchiveRule(ctx context.Context, params *DeleteArchiveRuleInput, optFns ...func(*Options)) (*DeleteArchiveRuleOutput, error) {
	stack := middleware.NewStack("DeleteArchiveRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteArchiveRuleMiddlewares(stack)
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
	addOpDeleteArchiveRuleValidationMiddleware(stack)
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
			OperationName: "DeleteArchiveRule",
			Err:           err,
		}
	}
	out := result.(*DeleteArchiveRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Deletes an archive rule.
type DeleteArchiveRuleInput struct {

	// The name of the analyzer that associated with the archive rule to delete.
	//
	// This member is required.
	AnalyzerName *string

	// The name of the rule to delete.
	//
	// This member is required.
	RuleName *string

	// A client token.
	ClientToken *string
}

type DeleteArchiveRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteArchiveRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteArchiveRule{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteArchiveRule{}, middleware.After)
}
