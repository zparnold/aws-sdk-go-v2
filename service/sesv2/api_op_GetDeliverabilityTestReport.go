// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve the results of a predictive inbox placement test.
func (c *Client) GetDeliverabilityTestReport(ctx context.Context, params *GetDeliverabilityTestReportInput, optFns ...func(*Options)) (*GetDeliverabilityTestReportOutput, error) {
	stack := middleware.NewStack("GetDeliverabilityTestReport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDeliverabilityTestReportMiddlewares(stack)
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
	addOpGetDeliverabilityTestReportValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeliverabilityTestReport(options.Region), middleware.Before)
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
			OperationName: "GetDeliverabilityTestReport",
			Err:           err,
		}
	}
	out := result.(*GetDeliverabilityTestReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve the results of a predictive inbox placement test.
type GetDeliverabilityTestReportInput struct {

	// A unique string that identifies the predictive inbox placement test.
	//
	// This member is required.
	ReportId *string
}

// The results of the predictive inbox placement test.
type GetDeliverabilityTestReportOutput struct {

	// An object that contains the results of the predictive inbox placement test.
	//
	// This member is required.
	DeliverabilityTestReport *types.DeliverabilityTestReport

	// An object that describes how the test email was handled by several email
	// providers, including Gmail, Hotmail, Yahoo, AOL, and others.
	//
	// This member is required.
	IspPlacements []*types.IspPlacement

	// An object that specifies how many test messages that were sent during the
	// predictive inbox placement test were delivered to recipients' inboxes, how many
	// were sent to recipients' spam folders, and how many weren't delivered.
	//
	// This member is required.
	OverallPlacement *types.PlacementStatistics

	// An object that contains the message that you sent when you performed this
	// predictive inbox placement test.
	Message *string

	// An array of objects that define the tags (keys and values) that are associated
	// with the predictive inbox placement test.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDeliverabilityTestReportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDeliverabilityTestReport{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeliverabilityTestReport{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDeliverabilityTestReport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDeliverabilityTestReport",
	}
}
