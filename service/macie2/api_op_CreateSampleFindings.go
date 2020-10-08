// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates sample findings.
func (c *Client) CreateSampleFindings(ctx context.Context, params *CreateSampleFindingsInput, optFns ...func(*Options)) (*CreateSampleFindingsOutput, error) {
	if params == nil {
		params = &CreateSampleFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSampleFindings", params, optFns, addOperationCreateSampleFindingsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSampleFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSampleFindingsInput struct {

	// An array that lists one or more types of findings to include in the set of
	// sample findings. Currently, the only supported value is
	// Policy:IAMUser/S3BucketEncryptionDisabled.
	FindingTypes []types.FindingType
}

type CreateSampleFindingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSampleFindingsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSampleFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSampleFindings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSampleFindings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateSampleFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateSampleFindings",
	}
}
