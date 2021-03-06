// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Rejects an offer to share the specified portfolio.
func (c *Client) RejectPortfolioShare(ctx context.Context, params *RejectPortfolioShareInput, optFns ...func(*Options)) (*RejectPortfolioShareOutput, error) {
	stack := middleware.NewStack("RejectPortfolioShare", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRejectPortfolioShareMiddlewares(stack)
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
	addOpRejectPortfolioShareValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRejectPortfolioShare(options.Region), middleware.Before)
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
			OperationName: "RejectPortfolioShare",
			Err:           err,
		}
	}
	out := result.(*RejectPortfolioShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RejectPortfolioShareInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string

	// The type of shared portfolios to reject. The default is to reject imported
	// portfolios.
	//
	//     * AWS_ORGANIZATIONS - Reject portfolios shared by the master
	// account of your organization.
	//
	//     * IMPORTED - Reject imported portfolios.
	//
	//
	// * AWS_SERVICECATALOG - Not supported. (Throws ResourceNotFoundException.)
	//
	// For
	// example, aws servicecatalog reject-portfolio-share --portfolio-id
	// "port-2qwzkwxt3y5fk" --portfolio-share-type AWS_ORGANIZATIONS
	PortfolioShareType types.PortfolioShareType
}

type RejectPortfolioShareOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRejectPortfolioShareMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRejectPortfolioShare{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRejectPortfolioShare{}, middleware.After)
}

func newServiceMetadataMiddleware_opRejectPortfolioShare(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "RejectPortfolioShare",
	}
}
