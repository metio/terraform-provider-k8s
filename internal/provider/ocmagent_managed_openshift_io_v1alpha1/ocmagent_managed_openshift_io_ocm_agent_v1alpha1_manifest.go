/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ocmagent_managed_openshift_io_v1alpha1

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
	_ datasource.DataSource = &OcmagentManagedOpenshiftIoOcmAgentV1Alpha1Manifest{}
)

func NewOcmagentManagedOpenshiftIoOcmAgentV1Alpha1Manifest() datasource.DataSource {
	return &OcmagentManagedOpenshiftIoOcmAgentV1Alpha1Manifest{}
}

type OcmagentManagedOpenshiftIoOcmAgentV1Alpha1Manifest struct{}

type OcmagentManagedOpenshiftIoOcmAgentV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AgentConfig *struct {
			OcmBaseUrl *string   `tfsdk:"ocm_base_url" json:"ocmBaseUrl,omitempty"`
			Services   *[]string `tfsdk:"services" json:"services,omitempty"`
		} `tfsdk:"agent_config" json:"agentConfig,omitempty"`
		FleetMode     *bool   `tfsdk:"fleet_mode" json:"fleetMode,omitempty"`
		OcmAgentImage *string `tfsdk:"ocm_agent_image" json:"ocmAgentImage,omitempty"`
		Replicas      *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		TokenSecret   *string `tfsdk:"token_secret" json:"tokenSecret,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OcmagentManagedOpenshiftIoOcmAgentV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ocmagent_managed_openshift_io_ocm_agent_v1alpha1_manifest"
}

func (r *OcmagentManagedOpenshiftIoOcmAgentV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OcmAgent is the Schema for the ocmagents API",
		MarkdownDescription: "OcmAgent is the Schema for the ocmagents API",
		Attributes: map[string]schema.Attribute{
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

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
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
				Description:         "OcmAgentSpec defines the desired state of OcmAgent",
				MarkdownDescription: "OcmAgentSpec defines the desired state of OcmAgent",
				Attributes: map[string]schema.Attribute{
					"agent_config": schema.SingleNestedAttribute{
						Description:         "AgentConfig refers to OCM agent config fields separated",
						MarkdownDescription: "AgentConfig refers to OCM agent config fields separated",
						Attributes: map[string]schema.Attribute{
							"ocm_base_url": schema.StringAttribute{
								Description:         "OcmBaseUrl defines the OCM api endpoint for OCM agent to access",
								MarkdownDescription: "OcmBaseUrl defines the OCM api endpoint for OCM agent to access",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"services": schema.ListAttribute{
								Description:         "Services defines the supported OCM services, eg, service_log, cluster_management",
								MarkdownDescription: "Services defines the supported OCM services, eg, service_log, cluster_management",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"fleet_mode": schema.BoolAttribute{
						Description:         "FleetMode indicates if the OCM agent is running in fleet mode, default to false",
						MarkdownDescription: "FleetMode indicates if the OCM agent is running in fleet mode, default to false",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocm_agent_image": schema.StringAttribute{
						Description:         "OcmAgentImage defines the image which will be used by the OCM Agent",
						MarkdownDescription: "OcmAgentImage defines the image which will be used by the OCM Agent",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas defines the replica count for the OCM Agent service",
						MarkdownDescription: "Replicas defines the replica count for the OCM Agent service",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"token_secret": schema.StringAttribute{
						Description:         "TokenSecret points to the secret name which stores the access token to OCM server",
						MarkdownDescription: "TokenSecret points to the secret name which stores the access token to OCM server",
						Required:            true,
						Optional:            false,
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

func (r *OcmagentManagedOpenshiftIoOcmAgentV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ocmagent_managed_openshift_io_ocm_agent_v1alpha1_manifest")

	var model OcmagentManagedOpenshiftIoOcmAgentV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ocmagent.managed.openshift.io/v1alpha1")
	model.Kind = pointer.String("OcmAgent")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
