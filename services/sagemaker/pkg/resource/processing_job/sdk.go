// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package processing_job

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/sagemaker/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.ProcessingJob{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeProcessingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeProcessingJob", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "ValidationException" && strings.HasPrefix(awsErr.Message(), "Could not find requested job") {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AppSpecification != nil {
		f0 := &svcapitypes.AppSpecification{}
		if resp.AppSpecification.ContainerArguments != nil {
			f0f0 := []*string{}
			for _, f0f0iter := range resp.AppSpecification.ContainerArguments {
				var f0f0elem string
				f0f0elem = *f0f0iter
				f0f0 = append(f0f0, &f0f0elem)
			}
			f0.ContainerArguments = f0f0
		}
		if resp.AppSpecification.ContainerEntrypoint != nil {
			f0f1 := []*string{}
			for _, f0f1iter := range resp.AppSpecification.ContainerEntrypoint {
				var f0f1elem string
				f0f1elem = *f0f1iter
				f0f1 = append(f0f1, &f0f1elem)
			}
			f0.ContainerEntrypoint = f0f1
		}
		if resp.AppSpecification.ImageUri != nil {
			f0.ImageURI = resp.AppSpecification.ImageUri
		}
		ko.Spec.AppSpecification = f0
	}
	if resp.Environment != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range resp.Environment {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		ko.Spec.Environment = f3
	}
	if resp.ExperimentConfig != nil {
		f5 := &svcapitypes.ExperimentConfig{}
		if resp.ExperimentConfig.ExperimentName != nil {
			f5.ExperimentName = resp.ExperimentConfig.ExperimentName
		}
		if resp.ExperimentConfig.TrialComponentDisplayName != nil {
			f5.TrialComponentDisplayName = resp.ExperimentConfig.TrialComponentDisplayName
		}
		if resp.ExperimentConfig.TrialName != nil {
			f5.TrialName = resp.ExperimentConfig.TrialName
		}
		ko.Spec.ExperimentConfig = f5
	}
	if resp.FailureReason != nil {
		ko.Status.FailureReason = resp.FailureReason
	}
	if resp.NetworkConfig != nil {
		f9 := &svcapitypes.NetworkConfig{}
		if resp.NetworkConfig.EnableInterContainerTrafficEncryption != nil {
			f9.EnableInterContainerTrafficEncryption = resp.NetworkConfig.EnableInterContainerTrafficEncryption
		}
		if resp.NetworkConfig.EnableNetworkIsolation != nil {
			f9.EnableNetworkIsolation = resp.NetworkConfig.EnableNetworkIsolation
		}
		if resp.NetworkConfig.VpcConfig != nil {
			f9f2 := &svcapitypes.VPCConfig{}
			if resp.NetworkConfig.VpcConfig.SecurityGroupIds != nil {
				f9f2f0 := []*string{}
				for _, f9f2f0iter := range resp.NetworkConfig.VpcConfig.SecurityGroupIds {
					var f9f2f0elem string
					f9f2f0elem = *f9f2f0iter
					f9f2f0 = append(f9f2f0, &f9f2f0elem)
				}
				f9f2.SecurityGroupIDs = f9f2f0
			}
			if resp.NetworkConfig.VpcConfig.Subnets != nil {
				f9f2f1 := []*string{}
				for _, f9f2f1iter := range resp.NetworkConfig.VpcConfig.Subnets {
					var f9f2f1elem string
					f9f2f1elem = *f9f2f1iter
					f9f2f1 = append(f9f2f1, &f9f2f1elem)
				}
				f9f2.Subnets = f9f2f1
			}
			f9.VPCConfig = f9f2
		}
		ko.Spec.NetworkConfig = f9
	}
	if resp.ProcessingInputs != nil {
		f11 := []*svcapitypes.ProcessingInput{}
		for _, f11iter := range resp.ProcessingInputs {
			f11elem := &svcapitypes.ProcessingInput{}
			if f11iter.InputName != nil {
				f11elem.InputName = f11iter.InputName
			}
			if f11iter.S3Input != nil {
				f11elemf1 := &svcapitypes.ProcessingS3Input{}
				if f11iter.S3Input.LocalPath != nil {
					f11elemf1.LocalPath = f11iter.S3Input.LocalPath
				}
				if f11iter.S3Input.S3CompressionType != nil {
					f11elemf1.S3CompressionType = f11iter.S3Input.S3CompressionType
				}
				if f11iter.S3Input.S3DataDistributionType != nil {
					f11elemf1.S3DataDistributionType = f11iter.S3Input.S3DataDistributionType
				}
				if f11iter.S3Input.S3DataType != nil {
					f11elemf1.S3DataType = f11iter.S3Input.S3DataType
				}
				if f11iter.S3Input.S3InputMode != nil {
					f11elemf1.S3InputMode = f11iter.S3Input.S3InputMode
				}
				if f11iter.S3Input.S3Uri != nil {
					f11elemf1.S3URI = f11iter.S3Input.S3Uri
				}
				f11elem.S3Input = f11elemf1
			}
			f11 = append(f11, f11elem)
		}
		ko.Spec.ProcessingInputs = f11
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ProcessingJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ProcessingJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.ProcessingJobName != nil {
		ko.Spec.ProcessingJobName = resp.ProcessingJobName
	}
	if resp.ProcessingJobStatus != nil {
		ko.Status.ProcessingJobStatus = resp.ProcessingJobStatus
	}
	if resp.ProcessingOutputConfig != nil {
		f15 := &svcapitypes.ProcessingOutputConfig{}
		if resp.ProcessingOutputConfig.KmsKeyId != nil {
			f15.KMSKeyID = resp.ProcessingOutputConfig.KmsKeyId
		}
		if resp.ProcessingOutputConfig.Outputs != nil {
			f15f1 := []*svcapitypes.ProcessingOutput{}
			for _, f15f1iter := range resp.ProcessingOutputConfig.Outputs {
				f15f1elem := &svcapitypes.ProcessingOutput{}
				if f15f1iter.OutputName != nil {
					f15f1elem.OutputName = f15f1iter.OutputName
				}
				if f15f1iter.S3Output != nil {
					f15f1elemf1 := &svcapitypes.ProcessingS3Output{}
					if f15f1iter.S3Output.LocalPath != nil {
						f15f1elemf1.LocalPath = f15f1iter.S3Output.LocalPath
					}
					if f15f1iter.S3Output.S3UploadMode != nil {
						f15f1elemf1.S3UploadMode = f15f1iter.S3Output.S3UploadMode
					}
					if f15f1iter.S3Output.S3Uri != nil {
						f15f1elemf1.S3URI = f15f1iter.S3Output.S3Uri
					}
					f15f1elem.S3Output = f15f1elemf1
				}
				f15f1 = append(f15f1, f15f1elem)
			}
			f15.Outputs = f15f1
		}
		ko.Spec.ProcessingOutputConfig = f15
	}
	if resp.ProcessingResources != nil {
		f16 := &svcapitypes.ProcessingResources{}
		if resp.ProcessingResources.ClusterConfig != nil {
			f16f0 := &svcapitypes.ProcessingClusterConfig{}
			if resp.ProcessingResources.ClusterConfig.InstanceCount != nil {
				f16f0.InstanceCount = resp.ProcessingResources.ClusterConfig.InstanceCount
			}
			if resp.ProcessingResources.ClusterConfig.InstanceType != nil {
				f16f0.InstanceType = resp.ProcessingResources.ClusterConfig.InstanceType
			}
			if resp.ProcessingResources.ClusterConfig.VolumeKmsKeyId != nil {
				f16f0.VolumeKMSKeyID = resp.ProcessingResources.ClusterConfig.VolumeKmsKeyId
			}
			if resp.ProcessingResources.ClusterConfig.VolumeSizeInGB != nil {
				f16f0.VolumeSizeInGB = resp.ProcessingResources.ClusterConfig.VolumeSizeInGB
			}
			f16.ClusterConfig = f16f0
		}
		ko.Spec.ProcessingResources = f16
	}
	if resp.RoleArn != nil {
		ko.Spec.RoleARN = resp.RoleArn
	}
	if resp.StoppingCondition != nil {
		f19 := &svcapitypes.ProcessingStoppingCondition{}
		if resp.StoppingCondition.MaxRuntimeInSeconds != nil {
			f19.MaxRuntimeInSeconds = resp.StoppingCondition.MaxRuntimeInSeconds
		}
		ko.Spec.StoppingCondition = f19
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.ProcessingJobName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeProcessingJobInput, error) {
	res := &svcsdk.DescribeProcessingJobInput{}

	if r.ko.Spec.ProcessingJobName != nil {
		res.SetProcessingJobName(*r.ko.Spec.ProcessingJobName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateProcessingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateProcessingJob", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ProcessingJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ProcessingJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateProcessingJobInput, error) {
	res := &svcsdk.CreateProcessingJobInput{}

	if r.ko.Spec.AppSpecification != nil {
		f0 := &svcsdk.AppSpecification{}
		if r.ko.Spec.AppSpecification.ContainerArguments != nil {
			f0f0 := []*string{}
			for _, f0f0iter := range r.ko.Spec.AppSpecification.ContainerArguments {
				var f0f0elem string
				f0f0elem = *f0f0iter
				f0f0 = append(f0f0, &f0f0elem)
			}
			f0.SetContainerArguments(f0f0)
		}
		if r.ko.Spec.AppSpecification.ContainerEntrypoint != nil {
			f0f1 := []*string{}
			for _, f0f1iter := range r.ko.Spec.AppSpecification.ContainerEntrypoint {
				var f0f1elem string
				f0f1elem = *f0f1iter
				f0f1 = append(f0f1, &f0f1elem)
			}
			f0.SetContainerEntrypoint(f0f1)
		}
		if r.ko.Spec.AppSpecification.ImageURI != nil {
			f0.SetImageUri(*r.ko.Spec.AppSpecification.ImageURI)
		}
		res.SetAppSpecification(f0)
	}
	if r.ko.Spec.Environment != nil {
		f1 := map[string]*string{}
		for f1key, f1valiter := range r.ko.Spec.Environment {
			var f1val string
			f1val = *f1valiter
			f1[f1key] = &f1val
		}
		res.SetEnvironment(f1)
	}
	if r.ko.Spec.ExperimentConfig != nil {
		f2 := &svcsdk.ExperimentConfig{}
		if r.ko.Spec.ExperimentConfig.ExperimentName != nil {
			f2.SetExperimentName(*r.ko.Spec.ExperimentConfig.ExperimentName)
		}
		if r.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil {
			f2.SetTrialComponentDisplayName(*r.ko.Spec.ExperimentConfig.TrialComponentDisplayName)
		}
		if r.ko.Spec.ExperimentConfig.TrialName != nil {
			f2.SetTrialName(*r.ko.Spec.ExperimentConfig.TrialName)
		}
		res.SetExperimentConfig(f2)
	}
	if r.ko.Spec.NetworkConfig != nil {
		f3 := &svcsdk.NetworkConfig{}
		if r.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption != nil {
			f3.SetEnableInterContainerTrafficEncryption(*r.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption)
		}
		if r.ko.Spec.NetworkConfig.EnableNetworkIsolation != nil {
			f3.SetEnableNetworkIsolation(*r.ko.Spec.NetworkConfig.EnableNetworkIsolation)
		}
		if r.ko.Spec.NetworkConfig.VPCConfig != nil {
			f3f2 := &svcsdk.VpcConfig{}
			if r.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs != nil {
				f3f2f0 := []*string{}
				for _, f3f2f0iter := range r.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs {
					var f3f2f0elem string
					f3f2f0elem = *f3f2f0iter
					f3f2f0 = append(f3f2f0, &f3f2f0elem)
				}
				f3f2.SetSecurityGroupIds(f3f2f0)
			}
			if r.ko.Spec.NetworkConfig.VPCConfig.Subnets != nil {
				f3f2f1 := []*string{}
				for _, f3f2f1iter := range r.ko.Spec.NetworkConfig.VPCConfig.Subnets {
					var f3f2f1elem string
					f3f2f1elem = *f3f2f1iter
					f3f2f1 = append(f3f2f1, &f3f2f1elem)
				}
				f3f2.SetSubnets(f3f2f1)
			}
			f3.SetVpcConfig(f3f2)
		}
		res.SetNetworkConfig(f3)
	}
	if r.ko.Spec.ProcessingInputs != nil {
		f4 := []*svcsdk.ProcessingInput{}
		for _, f4iter := range r.ko.Spec.ProcessingInputs {
			f4elem := &svcsdk.ProcessingInput{}
			if f4iter.InputName != nil {
				f4elem.SetInputName(*f4iter.InputName)
			}
			if f4iter.S3Input != nil {
				f4elemf1 := &svcsdk.ProcessingS3Input{}
				if f4iter.S3Input.LocalPath != nil {
					f4elemf1.SetLocalPath(*f4iter.S3Input.LocalPath)
				}
				if f4iter.S3Input.S3CompressionType != nil {
					f4elemf1.SetS3CompressionType(*f4iter.S3Input.S3CompressionType)
				}
				if f4iter.S3Input.S3DataDistributionType != nil {
					f4elemf1.SetS3DataDistributionType(*f4iter.S3Input.S3DataDistributionType)
				}
				if f4iter.S3Input.S3DataType != nil {
					f4elemf1.SetS3DataType(*f4iter.S3Input.S3DataType)
				}
				if f4iter.S3Input.S3InputMode != nil {
					f4elemf1.SetS3InputMode(*f4iter.S3Input.S3InputMode)
				}
				if f4iter.S3Input.S3URI != nil {
					f4elemf1.SetS3Uri(*f4iter.S3Input.S3URI)
				}
				f4elem.SetS3Input(f4elemf1)
			}
			f4 = append(f4, f4elem)
		}
		res.SetProcessingInputs(f4)
	}
	if r.ko.Spec.ProcessingJobName != nil {
		res.SetProcessingJobName(*r.ko.Spec.ProcessingJobName)
	}
	if r.ko.Spec.ProcessingOutputConfig != nil {
		f6 := &svcsdk.ProcessingOutputConfig{}
		if r.ko.Spec.ProcessingOutputConfig.KMSKeyID != nil {
			f6.SetKmsKeyId(*r.ko.Spec.ProcessingOutputConfig.KMSKeyID)
		}
		if r.ko.Spec.ProcessingOutputConfig.Outputs != nil {
			f6f1 := []*svcsdk.ProcessingOutput{}
			for _, f6f1iter := range r.ko.Spec.ProcessingOutputConfig.Outputs {
				f6f1elem := &svcsdk.ProcessingOutput{}
				if f6f1iter.OutputName != nil {
					f6f1elem.SetOutputName(*f6f1iter.OutputName)
				}
				if f6f1iter.S3Output != nil {
					f6f1elemf1 := &svcsdk.ProcessingS3Output{}
					if f6f1iter.S3Output.LocalPath != nil {
						f6f1elemf1.SetLocalPath(*f6f1iter.S3Output.LocalPath)
					}
					if f6f1iter.S3Output.S3UploadMode != nil {
						f6f1elemf1.SetS3UploadMode(*f6f1iter.S3Output.S3UploadMode)
					}
					if f6f1iter.S3Output.S3URI != nil {
						f6f1elemf1.SetS3Uri(*f6f1iter.S3Output.S3URI)
					}
					f6f1elem.SetS3Output(f6f1elemf1)
				}
				f6f1 = append(f6f1, f6f1elem)
			}
			f6.SetOutputs(f6f1)
		}
		res.SetProcessingOutputConfig(f6)
	}
	if r.ko.Spec.ProcessingResources != nil {
		f7 := &svcsdk.ProcessingResources{}
		if r.ko.Spec.ProcessingResources.ClusterConfig != nil {
			f7f0 := &svcsdk.ProcessingClusterConfig{}
			if r.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount != nil {
				f7f0.SetInstanceCount(*r.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount)
			}
			if r.ko.Spec.ProcessingResources.ClusterConfig.InstanceType != nil {
				f7f0.SetInstanceType(*r.ko.Spec.ProcessingResources.ClusterConfig.InstanceType)
			}
			if r.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID != nil {
				f7f0.SetVolumeKmsKeyId(*r.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID)
			}
			if r.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB != nil {
				f7f0.SetVolumeSizeInGB(*r.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB)
			}
			f7.SetClusterConfig(f7f0)
		}
		res.SetProcessingResources(f7)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}
	if r.ko.Spec.StoppingCondition != nil {
		f9 := &svcsdk.ProcessingStoppingCondition{}
		if r.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil {
			f9.SetMaxRuntimeInSeconds(*r.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
		}
		res.SetStoppingCondition(f9)
	}
	if r.ko.Spec.Tags != nil {
		f10 := []*svcsdk.Tag{}
		for _, f10iter := range r.ko.Spec.Tags {
			f10elem := &svcsdk.Tag{}
			if f10iter.Key != nil {
				f10elem.SetKey(*f10iter.Key)
			}
			if f10iter.Value != nil {
				f10elem.SetValue(*f10iter.Value)
			}
			f10 = append(f10, f10elem)
		}
		res.SetTags(f10)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.StopProcessingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "StopProcessingJob", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.StopProcessingJobInput, error) {
	res := &svcsdk.StopProcessingJobInput{}

	if r.ko.Spec.ProcessingJobName != nil {
		res.SetProcessingJobName(*r.ko.Spec.ProcessingJobName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.ProcessingJob,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	if terminalCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
