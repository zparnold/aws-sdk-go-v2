// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The CreatePreset operation creates a preset with settings that you specify.
// Elastic Transcoder checks the CreatePreset settings to ensure that they meet
// Elastic Transcoder requirements and to determine whether they comply with H.264
// standards. If your settings are not valid for Elastic Transcoder, Elastic
// Transcoder returns an HTTP 400 response (ValidationException) and does not
// create the preset. If the settings are valid for Elastic Transcoder but aren't
// strictly compliant with the H.264 standard, Elastic Transcoder creates the
// preset and returns a warning message in the response. This helps you determine
// whether your settings comply with the H.264 standard while giving you greater
// flexibility with respect to the video that Elastic Transcoder produces. Elastic
// Transcoder uses the H.264 video-compression format. For more information, see
// the International Telecommunication Union publication Recommendation ITU-T
// H.264: Advanced video coding for generic audiovisual services.
func (c *Client) CreatePreset(ctx context.Context, params *CreatePresetInput, optFns ...func(*Options)) (*CreatePresetOutput, error) {
	stack := middleware.NewStack("CreatePreset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreatePresetMiddlewares(stack)
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
	addOpCreatePresetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePreset(options.Region), middleware.Before)
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
			OperationName: "CreatePreset",
			Err:           err,
		}
	}
	out := result.(*CreatePresetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The CreatePresetRequest structure.
type CreatePresetInput struct {

	// The container type for the output file. Valid values include flac, flv, fmp4,
	// gif, mp3, mp4, mpg, mxf, oga, ogg, ts, and webm.
	//
	// This member is required.
	Container *string

	// The name of the preset. We recommend that the name be unique within the AWS
	// account, but uniqueness is not enforced.
	//
	// This member is required.
	Name *string

	// A section of the request body that specifies the audio parameters.
	Audio *types.AudioParameters

	// A description of the preset.
	Description *string

	// A section of the request body that specifies the thumbnail parameters, if any.
	Thumbnails *types.Thumbnails

	// A section of the request body that specifies the video parameters.
	Video *types.VideoParameters
}

// The CreatePresetResponse structure.
type CreatePresetOutput struct {

	// A section of the response body that provides information about the preset that
	// is created.
	Preset *types.Preset

	// If the preset settings don't comply with the standards for the video codec but
	// Elastic Transcoder created the preset, this message explains the reason the
	// preset settings don't meet the standard. Elastic Transcoder created the preset
	// because the settings might produce acceptable output.
	Warning *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreatePresetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreatePreset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePreset{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePreset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "CreatePreset",
	}
}
