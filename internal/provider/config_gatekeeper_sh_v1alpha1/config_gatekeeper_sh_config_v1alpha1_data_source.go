/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_gatekeeper_sh_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &ConfigGatekeeperShConfigV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigGatekeeperShConfigV1Alpha1DataSource{}
)

func NewConfigGatekeeperShConfigV1Alpha1DataSource() datasource.DataSource {
	return &ConfigGatekeeperShConfigV1Alpha1DataSource{}
}

type ConfigGatekeeperShConfigV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ConfigGatekeeperShConfigV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Match *[]struct {
			ExcludedNamespaces *[]string `tfsdk:"excluded_namespaces" json:"excludedNamespaces,omitempty"`
			Processes          *[]string `tfsdk:"processes" json:"processes,omitempty"`
		} `tfsdk:"match" json:"match,omitempty"`
		Readiness *struct {
			StatsEnabled *bool `tfsdk:"stats_enabled" json:"statsEnabled,omitempty"`
		} `tfsdk:"readiness" json:"readiness,omitempty"`
		Sync *struct {
			SyncOnly *[]struct {
				Group   *string `tfsdk:"group" json:"group,omitempty"`
				Kind    *string `tfsdk:"kind" json:"kind,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"sync_only" json:"syncOnly,omitempty"`
		} `tfsdk:"sync" json:"sync,omitempty"`
		Validation *struct {
			Traces *[]struct {
				Dump *string `tfsdk:"dump" json:"dump,omitempty"`
				Kind *struct {
					Group   *string `tfsdk:"group" json:"group,omitempty"`
					Kind    *string `tfsdk:"kind" json:"kind,omitempty"`
					Version *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"kind" json:"kind,omitempty"`
				User *string `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"traces" json:"traces,omitempty"`
		} `tfsdk:"validation" json:"validation,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigGatekeeperShConfigV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_gatekeeper_sh_config_v1alpha1"
}

func (r *ConfigGatekeeperShConfigV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Config is the Schema for the configs API.",
		MarkdownDescription: "Config is the Schema for the configs API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ConfigSpec defines the desired state of Config.",
				MarkdownDescription: "ConfigSpec defines the desired state of Config.",
				Attributes: map[string]schema.Attribute{
					"match": schema.ListNestedAttribute{
						Description:         "Configuration for namespace exclusion",
						MarkdownDescription: "Configuration for namespace exclusion",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"excluded_namespaces": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"processes": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"readiness": schema.SingleNestedAttribute{
						Description:         "Configuration for readiness tracker",
						MarkdownDescription: "Configuration for readiness tracker",
						Attributes: map[string]schema.Attribute{
							"stats_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"sync": schema.SingleNestedAttribute{
						Description:         "Configuration for syncing k8s objects",
						MarkdownDescription: "Configuration for syncing k8s objects",
						Attributes: map[string]schema.Attribute{
							"sync_only": schema.ListNestedAttribute{
								Description:         "If non-empty, only entries on this list will be replicated into OPA",
								MarkdownDescription: "If non-empty, only entries on this list will be replicated into OPA",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"group": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"kind": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"validation": schema.SingleNestedAttribute{
						Description:         "Configuration for validation",
						MarkdownDescription: "Configuration for validation",
						Attributes: map[string]schema.Attribute{
							"traces": schema.ListNestedAttribute{
								Description:         "List of requests to trace. Both 'user' and 'kinds' must be specified",
								MarkdownDescription: "List of requests to trace. Both 'user' and 'kinds' must be specified",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dump": schema.StringAttribute{
											Description:         "Also dump the state of OPA with the trace. Set to 'All' to dump everything.",
											MarkdownDescription: "Also dump the state of OPA with the trace. Set to 'All' to dump everything.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"kind": schema.SingleNestedAttribute{
											Description:         "Only trace requests of the following GroupVersionKind",
											MarkdownDescription: "Only trace requests of the following GroupVersionKind",
											Attributes: map[string]schema.Attribute{
												"group": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"kind": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"version": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"user": schema.StringAttribute{
											Description:         "Only trace requests from the specified user",
											MarkdownDescription: "Only trace requests from the specified user",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ConfigGatekeeperShConfigV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ConfigGatekeeperShConfigV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_config_gatekeeper_sh_config_v1alpha1")

	var data ConfigGatekeeperShConfigV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.gatekeeper.sh", Version: "v1alpha1", Resource: "Config"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ConfigGatekeeperShConfigV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("config.gatekeeper.sh/v1alpha1")
	data.Kind = pointer.String("Config")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
