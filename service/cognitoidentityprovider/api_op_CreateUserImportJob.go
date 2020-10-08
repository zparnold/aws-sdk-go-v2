// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates the user import job.
func (c *Client) CreateUserImportJob(ctx context.Context, params *CreateUserImportJobInput, optFns ...func(*Options)) (*CreateUserImportJobOutput, error) {
	if params == nil {
		params = &CreateUserImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserImportJob", params, optFns, addOperationCreateUserImportJobMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to create the user import job.
type CreateUserImportJobInput struct {

	// The role ARN for the Amazon CloudWatch Logging role for the user import job.
	//
	// This member is required.
	CloudWatchLogsRoleArn *string

	// The job name for the user import job.
	//
	// This member is required.
	JobName *string

	// The user pool ID for the user pool that the users are being imported into.
	//
	// This member is required.
	UserPoolId *string
}

// Represents the response from the server to the request to create the user import
// job.
type CreateUserImportJobOutput struct {

	// The job object that represents the user import job.
	UserImportJob *types.UserImportJobType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateUserImportJobMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserImportJob{}, middleware.After)
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
	addOpCreateUserImportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserImportJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateUserImportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateUserImportJob",
	}
}
