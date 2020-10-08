// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes all versions of a solution and the Solution object itself. Before
// deleting a solution, you must delete all campaigns based on the solution. To
// determine what campaigns are using the solution, call ListCampaigns () and
// supply the Amazon Resource Name (ARN) of the solution. You can't delete a
// solution if an associated SolutionVersion is in the CREATE PENDING or IN
// PROGRESS state. For more information on solutions, see CreateSolution ().
func (c *Client) DeleteSolution(ctx context.Context, params *DeleteSolutionInput, optFns ...func(*Options)) (*DeleteSolutionOutput, error) {
	if params == nil {
		params = &DeleteSolutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSolution", params, optFns, addOperationDeleteSolutionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSolutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSolutionInput struct {

	// The ARN of the solution to delete.
	//
	// This member is required.
	SolutionArn *string
}

type DeleteSolutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteSolutionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSolution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSolution{}, middleware.After)
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
	addOpDeleteSolutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSolution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteSolution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DeleteSolution",
	}
}
