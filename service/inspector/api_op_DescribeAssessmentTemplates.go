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

// Describes the assessment templates that are specified by the ARNs of the
// assessment templates.
func (c *Client) DescribeAssessmentTemplates(ctx context.Context, params *DescribeAssessmentTemplatesInput, optFns ...func(*Options)) (*DescribeAssessmentTemplatesOutput, error) {
	stack := middleware.NewStack("DescribeAssessmentTemplates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAssessmentTemplatesMiddlewares(stack)
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
	addOpDescribeAssessmentTemplatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssessmentTemplates(options.Region), middleware.Before)
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
			OperationName: "DescribeAssessmentTemplates",
			Err:           err,
		}
	}
	out := result.(*DescribeAssessmentTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssessmentTemplatesInput struct {
	AssessmentTemplateArns []*string
}

type DescribeAssessmentTemplatesOutput struct {

	// Information about the assessment templates.
	//
	// This member is required.
	AssessmentTemplates []*types.AssessmentTemplate

	// Assessment template details that cannot be described. An error code is provided
	// for each failed item.
	//
	// This member is required.
	FailedItems map[string]*types.FailedItemDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAssessmentTemplatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssessmentTemplates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssessmentTemplates{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAssessmentTemplates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DescribeAssessmentTemplates",
	}
}
