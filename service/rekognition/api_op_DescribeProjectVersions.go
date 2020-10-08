// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists and describes the models in an Amazon Rekognition Custom Labels project.
// You can specify up to 10 model versions in ProjectVersionArns. If you don't
// specify a value, descriptions for all models are returned. This operation
// requires permissions to perform the rekognition:DescribeProjectVersions action.
func (c *Client) DescribeProjectVersions(ctx context.Context, params *DescribeProjectVersionsInput, optFns ...func(*Options)) (*DescribeProjectVersionsOutput, error) {
	if params == nil {
		params = &DescribeProjectVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProjectVersions", params, optFns, addOperationDescribeProjectVersionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProjectVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProjectVersionsInput struct {

	// The Amazon Resource Name (ARN) of the project that contains the models you want
	// to describe.
	//
	// This member is required.
	ProjectArn *string

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	MaxResults *int32

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition Custom Labels returns a pagination token in the
	// response. You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// A list of model version names that you want to describe. You can add up to 10
	// model version names to the list. If you don't specify a value, all model
	// descriptions are returned. A version name is part of a model (ProjectVersion)
	// ARN. For example, my-model.2020-01-21T09.10.15 is the version name in the
	// following ARN.
	// arn:aws:rekognition:us-east-1:123456789012:project/getting-started/version/my-model.2020-01-21T09.10.15/1234567890123.
	VersionNames []*string
}

type DescribeProjectVersionsOutput struct {

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition Custom Labels returns a pagination token in the
	// response. You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// A list of model descriptions. The list is sorted by the creation date and time
	// of the model versions, latest to earliest.
	ProjectVersionDescriptions []*types.ProjectVersionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeProjectVersionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProjectVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProjectVersions{}, middleware.After)
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
	addOpDescribeProjectVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProjectVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeProjectVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DescribeProjectVersions",
	}
}
