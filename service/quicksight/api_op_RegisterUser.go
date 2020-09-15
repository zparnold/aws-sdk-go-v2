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

// Creates an Amazon QuickSight user, whose identity is associated with the AWS
// Identity and Access Management (IAM) identity or role specified in the request.
func (c *Client) RegisterUser(ctx context.Context, params *RegisterUserInput, optFns ...func(*Options)) (*RegisterUserOutput, error) {
	stack := middleware.NewStack("RegisterUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRegisterUserMiddlewares(stack)
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
	addOpRegisterUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterUser(options.Region), middleware.Before)

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
			OperationName: "RegisterUser",
			Err:           err,
		}
	}
	out := result.(*RegisterUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterUserInput struct {
	// Amazon QuickSight supports several ways of managing the identity of users. This
	// parameter accepts two values:
	//
	//     * IAM: A user whose identity maps to an
	// existing IAM user or role.
	//
	//     * QUICKSIGHT: A user whose identity is owned and
	// managed internally by Amazon QuickSight.
	IdentityType types.IdentityType
	// You need to use this parameter only when you register one or more users using an
	// assumed IAM role. You don't need to provide the session name for other
	// scenarios, for example when you are registering an IAM user or an Amazon
	// QuickSight user. You can register multiple users using the same IAM role if each
	// user has a different session name. For more information on assuming IAM roles,
	// see assume-role
	// (https://awscli.amazonaws.com/v2/documentation/api/latest/reference/sts/assume-role.html)
	// in the AWS CLI Reference.
	SessionName *string
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn *string
	// The Amazon QuickSight role for the user. The user role can be one of the
	// following:
	//
	//     * READER: A user who has read-only access to dashboards.
	//
	//     *
	// AUTHOR: A user who can create data sources, datasets, analyses, and
	// dashboards.
	//
	//     * ADMIN: A user who is an author, who can also manage Amazon
	// QuickSight settings.
	//
	//     * RESTRICTED_READER: This role isn't currently
	// available for use.
	//
	//     * RESTRICTED_AUTHOR: This role isn't currently available
	// for use.
	UserRole types.UserRole
	// The namespace. Currently, you should set this to default.
	Namespace *string
	// The ID for the AWS account that the user is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string
	// (Enterprise edition only) The name of the custom permissions profile that you
	// want to assign to this user. Currently, custom permissions profile names are
	// assigned to permissions profiles in the QuickSight console. You use this API to
	// assign the named set of permissions to a QuickSight user.  <p>Customizing
	// permissions in the QuickSight UI allows you to control a user's access to the
	// following operations:</p> <ul> <li> <p></p> </li> <li> <p></p> </li> <li>
	// <p></p> </li> <li> <p></p> </li> </ul> <p>QuickSight custom permissions are
	// applied through IAM policies. Therefore, they override the permissions typically
	// granted by assigning QuickSight users to one of the default security cohorts
	// (admin, author, reader) in QuickSight.</p> <p>This feature is available only to
	// QuickSight Enterprise edition subscriptions that use SAML 2.0-Based Federation
	// for Single Sign-On (SSO).</p>
	CustomPermissionsName *string
	// The email address of the user that you want to register.
	Email *string
	// The Amazon QuickSight user name that you want to create for the user you are
	// registering.
	UserName *string
}

type RegisterUserOutput struct {
	// The AWS request ID for this operation.
	RequestId *string
	// The URL the user visits to complete registration and provide a password. This is
	// returned only for users with an identity type of QUICKSIGHT.
	UserInvitationUrl *string
	// The user name.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRegisterUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRegisterUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "RegisterUser",
	}
}
