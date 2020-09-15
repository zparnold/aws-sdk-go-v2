// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a new job template. For information about job templates see the User
// Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
func (c *Client) CreateJobTemplate(ctx context.Context, params *CreateJobTemplateInput, optFns ...func(*Options)) (*CreateJobTemplateOutput, error) {
	stack := middleware.NewStack("CreateJobTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateJobTemplateMiddlewares(stack)
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
	addOpCreateJobTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJobTemplate(options.Region), middleware.Before)

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
			OperationName: "CreateJobTemplate",
			Err:           err,
		}
	}
	out := result.(*CreateJobTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobTemplateInput struct {
	// The name of the job template you are creating.
	Name *string
	// Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch
	// Events. Set the interval, in seconds, between status updates. MediaConvert sends
	// an update at this interval from the time the service begins processing your job
	// to the time it completes the transcode or encounters an error.
	StatusUpdateInterval types.StatusUpdateInterval
	// Optional. The queue that jobs created from this template are assigned to. If you
	// don't specify this, jobs will go to the default queue.
	Queue *string
	// Optional. A description of the job template you are creating.
	Description *string
	// Accelerated transcoding can significantly speed up jobs with long, visually
	// complex content. Outputs that use this feature incur pro-tier pricing. For
	// information about feature limitations, see the AWS Elemental MediaConvert User
	// Guide.
	AccelerationSettings *types.AccelerationSettings
	// Optional. A category for the job template you are creating
	Category *string
	// The tags that you want to add to the resource. You can tag resources with a
	// key-value pair or with only a key.
	Tags map[string]*string
	// Specify the relative priority for this job. In any given queue, the service
	// begins processing the job with the highest value first. When more than one job
	// has the same priority, the service begins processing the job that you submitted
	// first. If you don't specify a priority, the service uses the default value 0.
	Priority *int32
	// Optional. Use queue hopping to avoid overly long waits in the backlog of the
	// queue that you submit your job to. Specify an alternate queue and the maximum
	// time that your job will wait in the initial queue before hopping. For more
	// information about this feature, see the AWS Elemental MediaConvert User Guide.
	HopDestinations []*types.HopDestination
	// JobTemplateSettings contains all the transcode settings saved in the template
	// that will be applied to jobs created from it.
	Settings *types.JobTemplateSettings
}

type CreateJobTemplateOutput struct {
	// A job template is a pre-made set of encoding instructions that you can use to
	// quickly create a job.
	JobTemplate *types.JobTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateJobTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateJobTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJobTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateJobTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "CreateJobTemplate",
	}
}
