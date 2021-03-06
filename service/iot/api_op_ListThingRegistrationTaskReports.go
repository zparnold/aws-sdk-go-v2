// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Information about the thing registration tasks.
func (c *Client) ListThingRegistrationTaskReports(ctx context.Context, params *ListThingRegistrationTaskReportsInput, optFns ...func(*Options)) (*ListThingRegistrationTaskReportsOutput, error) {
	stack := middleware.NewStack("ListThingRegistrationTaskReports", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListThingRegistrationTaskReportsMiddlewares(stack)
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
	addOpListThingRegistrationTaskReportsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThingRegistrationTaskReports(options.Region), middleware.Before)
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
			OperationName: "ListThingRegistrationTaskReports",
			Err:           err,
		}
	}
	out := result.(*ListThingRegistrationTaskReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingRegistrationTaskReportsInput struct {

	// The type of task report.
	//
	// This member is required.
	ReportType types.ReportType

	// The id of the task.
	//
	// This member is required.
	TaskId *string

	// The maximum number of results to return per request.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string
}

type ListThingRegistrationTaskReportsOutput struct {

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The type of task report.
	ReportType types.ReportType

	// Links to the task resources.
	ResourceLinks []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListThingRegistrationTaskReportsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListThingRegistrationTaskReports{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingRegistrationTaskReports{}, middleware.After)
}

func newServiceMetadataMiddleware_opListThingRegistrationTaskReports(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingRegistrationTaskReports",
	}
}
