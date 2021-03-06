// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists your Amazon EventBridge rules. You can either list all the rules or you
// can provide a prefix to match to the rule names.  <p>ListRules does not list the
// targets of a rule. To see the targets associated with a rule, use
// <a>ListTargetsByRule</a>.</p>
func (c *Client) ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) {
	stack := middleware.NewStack("ListRules", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRulesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRules(options.Region), middleware.Before)
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
			OperationName: "ListRules",
			Err:           err,
		}
	}
	out := result.(*ListRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRulesInput struct {

	// Limits the results to show only the rules associated with the specified event
	// bus.
	EventBusName *string

	// The maximum number of results to return.
	Limit *int32

	// The prefix matching the rule name.
	NamePrefix *string

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string
}

type ListRulesOutput struct {

	// Indicates whether there are additional results to retrieve. If there are no more
	// results, the value is null.
	NextToken *string

	// The rules that match the specified criteria.
	Rules []*types.Rule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRulesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRules{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRules{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListRules",
	}
}
