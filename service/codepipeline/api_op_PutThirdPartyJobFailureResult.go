// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Represents the failure of a third party job as returned to the pipeline by a job
// worker. Used for partner actions only.
func (c *Client) PutThirdPartyJobFailureResult(ctx context.Context, params *PutThirdPartyJobFailureResultInput, optFns ...func(*Options)) (*PutThirdPartyJobFailureResultOutput, error) {
	if params == nil {
		params = &PutThirdPartyJobFailureResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutThirdPartyJobFailureResult", params, optFns, addOperationPutThirdPartyJobFailureResultMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PutThirdPartyJobFailureResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutThirdPartyJobFailureResult action.
type PutThirdPartyJobFailureResultInput struct {

	// The clientToken portion of the clientId and clientToken pair used to verify that
	// the calling entity is allowed access to the job and its details.
	//
	// This member is required.
	ClientToken *string

	// Represents information about failure details.
	//
	// This member is required.
	FailureDetails *types.FailureDetails

	// The ID of the job that failed. This is the same ID returned from
	// PollForThirdPartyJobs.
	//
	// This member is required.
	JobId *string
}

type PutThirdPartyJobFailureResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutThirdPartyJobFailureResultMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutThirdPartyJobFailureResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutThirdPartyJobFailureResult{}, middleware.After)
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
	addOpPutThirdPartyJobFailureResultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutThirdPartyJobFailureResult(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutThirdPartyJobFailureResult(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PutThirdPartyJobFailureResult",
	}
}
