/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OperatorTigeraIoAmazonCloudIntegrationV1Manifest{}
)

func NewOperatorTigeraIoAmazonCloudIntegrationV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoAmazonCloudIntegrationV1Manifest{}
}

type OperatorTigeraIoAmazonCloudIntegrationV1Manifest struct{}

type OperatorTigeraIoAmazonCloudIntegrationV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AwsRegion                    *string   `tfsdk:"aws_region" json:"awsRegion,omitempty"`
		DefaultPodMetadataAccess     *string   `tfsdk:"default_pod_metadata_access" json:"defaultPodMetadataAccess,omitempty"`
		EnforcedSecurityGroupID      *string   `tfsdk:"enforced_security_group_id" json:"enforcedSecurityGroupID,omitempty"`
		NodeSecurityGroupIDs         *[]string `tfsdk:"node_security_group_i_ds" json:"nodeSecurityGroupIDs,omitempty"`
		PodSecurityGroupID           *string   `tfsdk:"pod_security_group_id" json:"podSecurityGroupID,omitempty"`
		SqsURL                       *string   `tfsdk:"sqs_url" json:"sqsURL,omitempty"`
		TrustEnforcedSecurityGroupID *string   `tfsdk:"trust_enforced_security_group_id" json:"trustEnforcedSecurityGroupID,omitempty"`
		Vpcs                         *[]string `tfsdk:"vpcs" json:"vpcs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoAmazonCloudIntegrationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_amazon_cloud_integration_v1_manifest"
}

func (r *OperatorTigeraIoAmazonCloudIntegrationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AmazonCloudIntegration is the Schema for the amazoncloudintegrations API",
		MarkdownDescription: "AmazonCloudIntegration is the Schema for the amazoncloudintegrations API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "AmazonCloudIntegrationSpec defines the desired state of AmazonCloudIntegration",
				MarkdownDescription: "AmazonCloudIntegrationSpec defines the desired state of AmazonCloudIntegration",
				Attributes: map[string]schema.Attribute{
					"aws_region": schema.StringAttribute{
						Description:         "AWSRegion is the region in which your cluster is located.",
						MarkdownDescription: "AWSRegion is the region in which your cluster is located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_pod_metadata_access": schema.StringAttribute{
						Description:         "DefaultPodMetadataAccess defines what the default behavior will be for accessing the AWS metadata service from a pod. Default: Denied",
						MarkdownDescription: "DefaultPodMetadataAccess defines what the default behavior will be for accessing the AWS metadata service from a pod. Default: Denied",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Allowed", "Denied"),
						},
					},

					"enforced_security_group_id": schema.StringAttribute{
						Description:         "EnforcedSecurityGroupID is the ID of the Security Group which will be applied to all ENIs that are on a host that is also part of the Kubernetes cluster.",
						MarkdownDescription: "EnforcedSecurityGroupID is the ID of the Security Group which will be applied to all ENIs that are on a host that is also part of the Kubernetes cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_security_group_i_ds": schema.ListAttribute{
						Description:         "NodeSecurityGroupIDs is a list of Security Group IDs that all nodes and masters will be in.",
						MarkdownDescription: "NodeSecurityGroupIDs is a list of Security Group IDs that all nodes and masters will be in.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_security_group_id": schema.StringAttribute{
						Description:         "PodSecurityGroupID is the ID of the Security Group which all pods should be placed in by default.",
						MarkdownDescription: "PodSecurityGroupID is the ID of the Security Group which all pods should be placed in by default.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sqs_url": schema.StringAttribute{
						Description:         "SQSURL is the SQS URL needed to access the Simple Queue Service.",
						MarkdownDescription: "SQSURL is the SQS URL needed to access the Simple Queue Service.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"trust_enforced_security_group_id": schema.StringAttribute{
						Description:         "TrustEnforcedSecurityGroupID is the ID of the Security Group which will be applied to all ENIs in the VPC.",
						MarkdownDescription: "TrustEnforcedSecurityGroupID is the ID of the Security Group which will be applied to all ENIs in the VPC.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vpcs": schema.ListAttribute{
						Description:         "VPCS is a list of VPC IDs to monitor for ENIs and Security Groups, only one is supported.",
						MarkdownDescription: "VPCS is a list of VPC IDs to monitor for ENIs and Security Groups, only one is supported.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *OperatorTigeraIoAmazonCloudIntegrationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_amazon_cloud_integration_v1_manifest")

	var model OperatorTigeraIoAmazonCloudIntegrationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("AmazonCloudIntegration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
