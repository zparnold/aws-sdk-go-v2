// Code generated by smithy-go-codegen DO NOT EDIT.

package costandusagereportservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows you to programatically update your report preferences.
func (c *Client) ModifyReportDefinition(ctx context.Context, params *ModifyReportDefinitionInput, optFns ...func(*Options)) (*ModifyReportDefinitionOutput, error) {
	if params == nil {
		params = &ModifyReportDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyReportDefinition", params, optFns, addOperationModifyReportDefinitionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyReportDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyReportDefinitionInput struct {

	// The definition of AWS Cost and Usage Report. You can specify the report name,
	// time unit, report format, compression format, S3 bucket, additional artifacts,
	// and schema elements in the definition.
	//
	// This member is required.
	ReportDefinition *types.ReportDefinition

	// The name of the report that you want to create. The name must be unique, is case
	// sensitive, and can't include spaces.
	//
	// This member is required.
	ReportName *string
}

type ModifyReportDefinitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyReportDefinitionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyReportDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyReportDefinition{}, middleware.After)
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
	addOpModifyReportDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReportDefinition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyReportDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cur",
		OperationName: "ModifyReportDefinition",
	}
}
