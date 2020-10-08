// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more Stage () resources.
func (c *Client) GetStages(ctx context.Context, params *GetStagesInput, optFns ...func(*Options)) (*GetStagesOutput, error) {
	if params == nil {
		params = &GetStagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStages", params, optFns, addOperationGetStagesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to get information about one or more Stage () resources.
type GetStagesInput struct {

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// The stages' deployment identifiers.
	DeploymentId *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// A list of Stage () resources that are associated with the ApiKey () resource.
// Deploying API in Stages
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/stages.html)
type GetStagesOutput struct {

	// The current page of elements from this collection.
	Item []*types.Stage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetStagesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetStages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStages{}, middleware.After)
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
	addOpGetStagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetStages(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetStages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetStages",
	}
}
