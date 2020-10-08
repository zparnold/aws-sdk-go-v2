// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Archives GuardDuty findings that are specified by the list of finding IDs. Only
// the master account can archive findings. Member accounts don't have permission
// to archive findings from their accounts.
func (c *Client) ArchiveFindings(ctx context.Context, params *ArchiveFindingsInput, optFns ...func(*Options)) (*ArchiveFindingsOutput, error) {
	if params == nil {
		params = &ArchiveFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ArchiveFindings", params, optFns, addOperationArchiveFindingsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ArchiveFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ArchiveFindingsInput struct {

	// The ID of the detector that specifies the GuardDuty service whose findings you
	// want to archive.
	//
	// This member is required.
	DetectorId *string

	// The IDs of the findings that you want to archive.
	//
	// This member is required.
	FindingIds []*string
}

type ArchiveFindingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationArchiveFindingsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpArchiveFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpArchiveFindings{}, middleware.After)
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
	addOpArchiveFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opArchiveFindings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opArchiveFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ArchiveFindings",
	}
}
