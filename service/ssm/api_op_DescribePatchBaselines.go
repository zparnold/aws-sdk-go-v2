// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the patch baselines in your AWS account.
func (c *Client) DescribePatchBaselines(ctx context.Context, params *DescribePatchBaselinesInput, optFns ...func(*Options)) (*DescribePatchBaselinesOutput, error) {
	if params == nil {
		params = &DescribePatchBaselinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePatchBaselines", params, optFns, addOperationDescribePatchBaselinesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePatchBaselinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePatchBaselinesInput struct {

	// Each element in the array is a structure containing: Key: (string, "NAME_PREFIX"
	// or "OWNER") Value: (array of strings, exactly 1 entry, between 1 and 255
	// characters)
	Filters []*types.PatchOrchestratorFilter

	// The maximum number of patch baselines to return (per page).
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribePatchBaselinesOutput struct {

	// An array of PatchBaselineIdentity elements.
	BaselineIdentities []*types.PatchBaselineIdentity

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePatchBaselinesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePatchBaselines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePatchBaselines{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePatchBaselines(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribePatchBaselines(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribePatchBaselines",
	}
}
