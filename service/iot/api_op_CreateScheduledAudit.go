// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a scheduled audit that is run at a specified time interval.
func (c *Client) CreateScheduledAudit(ctx context.Context, params *CreateScheduledAuditInput, optFns ...func(*Options)) (*CreateScheduledAuditOutput, error) {
	if params == nil {
		params = &CreateScheduledAuditInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateScheduledAudit", params, optFns, addOperationCreateScheduledAuditMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateScheduledAuditOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScheduledAuditInput struct {

	// How often the scheduled audit takes place. Can be one of "DAILY", "WEEKLY",
	// "BIWEEKLY" or "MONTHLY". The start time of each audit is determined by the
	// system.
	//
	// This member is required.
	Frequency types.AuditFrequency

	// The name you want to give to the scheduled audit. (Max. 128 chars)
	//
	// This member is required.
	ScheduledAuditName *string

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list of all
	// checks, including those that are enabled or use UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	//
	// This member is required.
	TargetCheckNames []*string

	// The day of the month on which the scheduled audit takes place. Can be "1"
	// through "31" or "LAST". This field is required if the "frequency" parameter is
	// set to "MONTHLY". If days 29-31 are specified, and the month does not have that
	// many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string

	// The day of the week on which the scheduled audit takes place. Can be one of
	// "SUN", "MON", "TUE", "WED", "THU", "FRI", or "SAT". This field is required if
	// the "frequency" parameter is set to "WEEKLY" or "BIWEEKLY".
	DayOfWeek types.DayOfWeek

	// Metadata that can be used to manage the scheduled audit.
	Tags []*types.Tag
}

type CreateScheduledAuditOutput struct {

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateScheduledAuditMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateScheduledAudit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateScheduledAudit{}, middleware.After)
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
	addOpCreateScheduledAuditValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateScheduledAudit(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateScheduledAudit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateScheduledAudit",
	}
}
