// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Domain used by SageMaker Studio. A domain consists of an associated
// directory, a list of authorized users, and a variety of security, application,
// policy, and Amazon Virtual Private Cloud (VPC) configurations. An AWS account is
// limited to one domain per region. Users within a domain can share notebook files
// and other artifacts with each other. When a domain is created, an Amazon Elastic
// File System (EFS) volume is also created for use by all of the users within the
// domain. Each user receives a private home directory within the EFS for
// notebooks, Git repositories, and data files. All traffic between the domain and
// the EFS volume is communicated through the specified subnet IDs. All other
// traffic goes over the Internet through an Amazon SageMaker system VPC. The EFS
// traffic uses the NFS/TCP protocol over port 2049. NFS traffic over TCP on port
// 2049 needs to be allowed in both inbound and outbound rules in order to launch a
// SageMaker Studio app successfully.
func (c *Client) CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) {
	stack := middleware.NewStack("CreateDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDomainMiddlewares(stack)
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
	addOpCreateDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomain(options.Region), middleware.Before)

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
			OperationName: "CreateDomain",
			Err:           err,
		}
	}
	out := result.(*CreateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainInput struct {
	// The ID of the Amazon Virtual Private Cloud (VPC) to use for communication with
	// the EFS volume.
	VpcId *string
	// The default user settings.
	DefaultUserSettings *types.UserSettings
	// The mode of authentication that members use to access the domain.
	AuthMode types.AuthMode
	// The VPC subnets to use for communication with the EFS volume.
	SubnetIds []*string
	// The AWS Key Management Service (KMS) encryption key ID. Encryption with a
	// customer master key (CMK) is not supported.
	HomeEfsFileSystemKmsKeyId *string
	// Tags to associated with the Domain. Each tag consists of a key and an optional
	// value. Tag keys must be unique per resource. Tags are searchable using the
	// Search () API.
	Tags []*types.Tag
	// A name for the domain.
	DomainName *string
}

type CreateDomainOutput struct {
	// The URL to the created domain.
	Url *string
	// The Amazon Resource Name (ARN) of the created domain.
	DomainArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateDomain",
	}
}
