// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates and starts a new SPICE ingestion on a dataset  <p>Any ingestions
// operating on tagged datasets inherit the same tags automatically for use in
// access control. For an example, see <a
// href="https://aws.amazon.com/premiumsupport/knowledge-center/iam-ec2-resource-tags/">How
// do I create an IAM policy to control access to Amazon EC2 resources using
// tags?</a> in the AWS Knowledge Center. Tags are visible on the tagged dataset,
// but not on the ingestion resource.</p>
func (c *Client) CreateIngestion(ctx context.Context, params *CreateIngestionInput, optFns ...func(*Options)) (*CreateIngestionOutput, error) {
	stack := middleware.NewStack("CreateIngestion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateIngestionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateIngestionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIngestion(options.Region), middleware.Before)

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
			OperationName: "CreateIngestion",
			Err:           err,
		}
	}
	out := result.(*CreateIngestionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIngestionInput struct {
	// The ID of the dataset used in the ingestion.
	DataSetId *string
	// An ID for the ingestion.
	IngestionId *string
	// The AWS account ID.
	AwsAccountId *string
}

type CreateIngestionOutput struct {
	// An ID for the ingestion.
	IngestionId *string
	// The AWS request ID for this operation.
	RequestId *string
	// The ingestion status.
	IngestionStatus types.IngestionStatus
	// The Amazon Resource Name (ARN) for the data ingestion.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateIngestionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateIngestion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIngestion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateIngestion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateIngestion",
	}
}
