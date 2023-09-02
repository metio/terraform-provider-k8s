/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package app_redislabs_com_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1Manifest{}
)

func NewAppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1Manifest() datasource.DataSource {
	return &AppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1Manifest{}
}

type AppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1Manifest struct{}

type AppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		ApiFqdnUrl   *string `tfsdk:"api_fqdn_url" json:"apiFqdnUrl,omitempty"`
		DbFqdnSuffix *string `tfsdk:"db_fqdn_suffix" json:"dbFqdnSuffix,omitempty"`
		RecName      *string `tfsdk:"rec_name" json:"recName,omitempty"`
		RecNamespace *string `tfsdk:"rec_namespace" json:"recNamespace,omitempty"`
		SecretName   *string `tfsdk:"secret_name" json:"secretName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_redislabs_com_redis_enterprise_remote_cluster_v1alpha1_manifest"
}

func (r *AppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RedisEntepriseRemoteCluster represents a remote participating cluster.",
		MarkdownDescription: "RedisEntepriseRemoteCluster represents a remote participating cluster.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"api_fqdn_url": schema.StringAttribute{
						Description:         "The URL of the cluster, will be used for the active-active database URL.",
						MarkdownDescription: "The URL of the cluster, will be used for the active-active database URL.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_fqdn_suffix": schema.StringAttribute{
						Description:         "The database URL suffix, will be used for the active-active database replication endpoint and replication endpoint SNI.",
						MarkdownDescription: "The database URL suffix, will be used for the active-active database replication endpoint and replication endpoint SNI.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rec_name": schema.StringAttribute{
						Description:         "The name of the REC that the RERC is pointing at",
						MarkdownDescription: "The name of the REC that the RERC is pointing at",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"rec_namespace": schema.StringAttribute{
						Description:         "The namespace of the REC that the RERC is pointing at",
						MarkdownDescription: "The namespace of the REC that the RERC is pointing at",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "The name of the secret containing cluster credentials. Must be of the following format: 'redis-enterprise-<RERC name>'",
						MarkdownDescription: "The name of the secret containing cluster credentials. Must be of the following format: 'redis-enterprise-<RERC name>'",
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

func (r *AppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_redislabs_com_redis_enterprise_remote_cluster_v1alpha1_manifest")

	var model AppRedislabsComRedisEnterpriseRemoteClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("app.redislabs.com/v1alpha1")
	model.Kind = pointer.String("RedisEnterpriseRemoteCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
