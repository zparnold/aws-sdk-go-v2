// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the rules packages that are specified by the ARNs of the rules
// packages.
func (c *Client) DescribeRulesPackages(ctx context.Context, params *DescribeRulesPackagesInput, optFns ...func(*Options)) (*DescribeRulesPackagesOutput, error) {
	stack := middleware.NewStack("DescribeRulesPackages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeRulesPackagesMiddlewares(stack)
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
	addOpDescribeRulesPackagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRulesPackages(options.Region), middleware.Before)
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
			OperationName: "DescribeRulesPackages",
			Err:           err,
		}
	}
	out := result.(*DescribeRulesPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRulesPackagesInput struct {

	// The ARN that specifies the rules package that you want to describe.
	//
	// This member is required.
	RulesPackageArns []*string

	// The locale that you want to translate a rules package description into.
	Locale types.Locale
}

type DescribeRulesPackagesOutput struct {

	// Rules package details that cannot be described. An error code is provided for
	// each failed item.
	//
	// This member is required.
	FailedItems map[string]*types.FailedItemDetails

	// Information about the rules package.
	//
	// This member is required.
	RulesPackages []*types.RulesPackage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeRulesPackagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRulesPackages{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRulesPackages{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRulesPackages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DescribeRulesPackages",
	}
}
