// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the name and/or the path of the specified IAM user. You should
// understand the implications of changing an IAM user's path or name. For more
// information, see Renaming an IAM User
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_manage.html#id_users_renaming)
// and Renaming an IAM Group
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_groups_manage_rename.html)
// in the IAM User Guide. To change a user name, the requester must have
// appropriate permissions on both the source object and the target object. For
// example, to change Bob to Robert, the entity making the request must have
// permission on Bob and Robert, or must have permission on all (*). For more
// information about permissions, see Permissions and Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/PermissionsAndPolicies.html).
func (c *Client) UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) {
	stack := middleware.NewStack("UpdateUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateUserMiddlewares(stack)
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
	addOpUpdateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUser(options.Region), middleware.Before)
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
			OperationName: "UpdateUser",
			Err:           err,
		}
	}
	out := result.(*UpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserInput struct {

	// Name of the user to update. If you're changing the name of the user, this is the
	// original user name. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string

	// New path for the IAM user. Include this parameter only if you're changing the
	// user's path. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of either a
	// forward slash (/) by itself or a string that must begin and end with forward
	// slashes. In addition, it can contain any ASCII character from the ! (\u0021)
	// through the DEL character (\u007F), including most punctuation characters,
	// digits, and upper and lowercased letters.
	NewPath *string

	// New name for the user. Include this parameter only if you're changing the user's
	// name. IAM user, group, role, and policy names must be unique within the account.
	// Names are not distinguished by case. For example, you cannot create resources
	// named both "MyResource" and "myresource".
	NewUserName *string
}

type UpdateUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateUser",
	}
}
