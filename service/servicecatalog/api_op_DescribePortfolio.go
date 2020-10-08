// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the specified portfolio. A delegated admin is authorized
// to invoke this command.
func (c *Client) DescribePortfolio(ctx context.Context, params *DescribePortfolioInput, optFns ...func(*Options)) (*DescribePortfolioOutput, error) {
	if params == nil {
		params = &DescribePortfolioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePortfolio", params, optFns, addOperationDescribePortfolioMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePortfolioInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	Id *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
}

type DescribePortfolioOutput struct {

	// Information about the associated budgets.
	Budgets []*types.BudgetDetail

	// Information about the portfolio.
	PortfolioDetail *types.PortfolioDetail

	// Information about the TagOptions associated with the portfolio.
	TagOptions []*types.TagOptionDetail

	// Information about the tags associated with the portfolio.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePortfolioMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePortfolio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePortfolio{}, middleware.After)
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
	addOpDescribePortfolioValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePortfolio(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribePortfolio(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribePortfolio",
	}
}
