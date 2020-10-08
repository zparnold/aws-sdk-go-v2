// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the usage data of a usage plan in a specified time interval.
func (c *Client) GetUsage(ctx context.Context, params *GetUsageInput, optFns ...func(*Options)) (*GetUsageOutput, error) {
	if params == nil {
		params = &GetUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsage", params, optFns, addOperationGetUsageMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get the usage data of a usage plan in a specified time
// interval.
type GetUsageInput struct {

	// [Required] The ending date (e.g., 2016-12-31) of the usage data.
	//
	// This member is required.
	EndDate *string

	// [Required] The starting date (e.g., 2016-01-01) of the usage data.
	//
	// This member is required.
	StartDate *string

	// [Required] The Id of the usage plan associated with the usage data.
	//
	// This member is required.
	UsagePlanId *string

	// The Id of the API key associated with the resultant usage data.
	KeyId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	Name *string

	// The current pagination position in the paged result set.
	Position *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents the usage data of a usage plan. Create and Use Usage Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html),
// Manage Usage in a Usage Plan
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-usage-plans-with-console.html#api-gateway-usage-plan-manage-usage)
type GetUsageOutput struct {

	// The ending date of the usage data.
	EndDate *string

	// The usage data, as daily logs of used and remaining quotas, over the specified
	// time interval indexed over the API keys in a usage plan. For example, {...,
	// "values" : { "{api_key}" : [ [0, 100], [10, 90], [100, 10]]}, where {api_key}
	// stands for an API key value and the daily log entry is of the format [used
	// quota, remaining quota].
	Items map[string][][]*int64

	// The current pagination position in the paged result set.
	Position *string

	// The starting date of the usage data.
	StartDate *string

	// The plan Id associated with this usage data.
	UsagePlanId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUsageMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsage{}, middleware.After)
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
	addOpGetUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsage",
	}
}
