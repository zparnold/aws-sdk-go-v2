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

// Assigns attributes (key and value pairs) to the findings that are specified by
// the ARNs of the findings.
func (c *Client) AddAttributesToFindings(ctx context.Context, params *AddAttributesToFindingsInput, optFns ...func(*Options)) (*AddAttributesToFindingsOutput, error) {
	stack := middleware.NewStack("AddAttributesToFindings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddAttributesToFindingsMiddlewares(stack)
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
	addOpAddAttributesToFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddAttributesToFindings(options.Region), middleware.Before)
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
			OperationName: "AddAttributesToFindings",
			Err:           err,
		}
	}
	out := result.(*AddAttributesToFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddAttributesToFindingsInput struct {

	// The array of attributes that you want to assign to specified findings.
	//
	// This member is required.
	Attributes []*types.Attribute

	// The ARNs that specify the findings that you want to assign attributes to.
	//
	// This member is required.
	FindingArns []*string
}

type AddAttributesToFindingsOutput struct {

	// Attribute details that cannot be described. An error code is provided for each
	// failed item.
	//
	// This member is required.
	FailedItems map[string]*types.FailedItemDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddAttributesToFindingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddAttributesToFindings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddAttributesToFindings{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddAttributesToFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "AddAttributesToFindings",
	}
}
