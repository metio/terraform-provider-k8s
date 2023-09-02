/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

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
	_ datasource.DataSource              = &KumaIoMeshProxyPatchV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KumaIoMeshProxyPatchV1Alpha1DataSource{}
)

func NewKumaIoMeshProxyPatchV1Alpha1DataSource() datasource.DataSource {
	return &KumaIoMeshProxyPatchV1Alpha1DataSource{}
}

type KumaIoMeshProxyPatchV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KumaIoMeshProxyPatchV1Alpha1DataSourceData struct {
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
		Default *struct {
			AppendModifications *[]struct {
				Cluster *struct {
					JsonPatches *[]struct {
						From  *string            `tfsdk:"from" json:"from,omitempty"`
						Op    *string            `tfsdk:"op" json:"op,omitempty"`
						Path  *string            `tfsdk:"path" json:"path,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"json_patches" json:"jsonPatches,omitempty"`
					Match *struct {
						Name   *string `tfsdk:"name" json:"name,omitempty"`
						Origin *string `tfsdk:"origin" json:"origin,omitempty"`
					} `tfsdk:"match" json:"match,omitempty"`
					Operation *string `tfsdk:"operation" json:"operation,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"cluster" json:"cluster,omitempty"`
				HttpFilter *struct {
					JsonPatches *[]struct {
						From  *string            `tfsdk:"from" json:"from,omitempty"`
						Op    *string            `tfsdk:"op" json:"op,omitempty"`
						Path  *string            `tfsdk:"path" json:"path,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"json_patches" json:"jsonPatches,omitempty"`
					Match *struct {
						ListenerName *string            `tfsdk:"listener_name" json:"listenerName,omitempty"`
						ListenerTags *map[string]string `tfsdk:"listener_tags" json:"listenerTags,omitempty"`
						Name         *string            `tfsdk:"name" json:"name,omitempty"`
						Origin       *string            `tfsdk:"origin" json:"origin,omitempty"`
					} `tfsdk:"match" json:"match,omitempty"`
					Operation *string `tfsdk:"operation" json:"operation,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"http_filter" json:"httpFilter,omitempty"`
				Listener *struct {
					JsonPatches *[]struct {
						From  *string            `tfsdk:"from" json:"from,omitempty"`
						Op    *string            `tfsdk:"op" json:"op,omitempty"`
						Path  *string            `tfsdk:"path" json:"path,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"json_patches" json:"jsonPatches,omitempty"`
					Match *struct {
						Name   *string            `tfsdk:"name" json:"name,omitempty"`
						Origin *string            `tfsdk:"origin" json:"origin,omitempty"`
						Tags   *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
					} `tfsdk:"match" json:"match,omitempty"`
					Operation *string `tfsdk:"operation" json:"operation,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"listener" json:"listener,omitempty"`
				NetworkFilter *struct {
					JsonPatches *[]struct {
						From  *string            `tfsdk:"from" json:"from,omitempty"`
						Op    *string            `tfsdk:"op" json:"op,omitempty"`
						Path  *string            `tfsdk:"path" json:"path,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"json_patches" json:"jsonPatches,omitempty"`
					Match *struct {
						ListenerName *string            `tfsdk:"listener_name" json:"listenerName,omitempty"`
						ListenerTags *map[string]string `tfsdk:"listener_tags" json:"listenerTags,omitempty"`
						Name         *string            `tfsdk:"name" json:"name,omitempty"`
						Origin       *string            `tfsdk:"origin" json:"origin,omitempty"`
					} `tfsdk:"match" json:"match,omitempty"`
					Operation *string `tfsdk:"operation" json:"operation,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"network_filter" json:"networkFilter,omitempty"`
				VirtualHost *struct {
					JsonPatches *[]struct {
						From  *string            `tfsdk:"from" json:"from,omitempty"`
						Op    *string            `tfsdk:"op" json:"op,omitempty"`
						Path  *string            `tfsdk:"path" json:"path,omitempty"`
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"json_patches" json:"jsonPatches,omitempty"`
					Match *struct {
						Name                   *string `tfsdk:"name" json:"name,omitempty"`
						Origin                 *string `tfsdk:"origin" json:"origin,omitempty"`
						RouteConfigurationName *string `tfsdk:"route_configuration_name" json:"routeConfigurationName,omitempty"`
					} `tfsdk:"match" json:"match,omitempty"`
					Operation *string `tfsdk:"operation" json:"operation,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"virtual_host" json:"virtualHost,omitempty"`
			} `tfsdk:"append_modifications" json:"appendModifications,omitempty"`
		} `tfsdk:"default" json:"default,omitempty"`
		TargetRef *struct {
			Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
			Mesh *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name *string            `tfsdk:"name" json:"name,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshProxyPatchV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_proxy_patch_v1alpha1"
}

func (r *KumaIoMeshProxyPatchV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Spec is the specification of the Kuma MeshProxyPatch resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshProxyPatch resource.",
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Description:         "Default is a configuration specific to the group of destinations referenced in 'targetRef'.",
						MarkdownDescription: "Default is a configuration specific to the group of destinations referenced in 'targetRef'.",
						Attributes: map[string]schema.Attribute{
							"append_modifications": schema.ListNestedAttribute{
								Description:         "AppendModifications is a list of modifications applied on the selected proxy.",
								MarkdownDescription: "AppendModifications is a list of modifications applied on the selected proxy.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster": schema.SingleNestedAttribute{
											Description:         "Cluster is a modification of Envoy's Cluster resource.",
											MarkdownDescription: "Cluster is a modification of Envoy's Cluster resource.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy's Cluster resource",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy's Cluster resource",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
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

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the cluster to match.",
															MarkdownDescription: "Name of the cluster to match.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched cluster.",
													MarkdownDescription: "Operation to execute on matched cluster.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"http_filter": schema.SingleNestedAttribute{
											Description:         "HTTPFilter is a modification of Envoy HTTP Filter available in HTTP Connection Manager in a Listener resource.",
											MarkdownDescription: "HTTPFilter is a modification of Envoy HTTP Filter available in HTTP Connection Manager in a Listener resource.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy's HTTP Filter available in HTTP Connection Manager in a Listener resource.",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy's HTTP Filter available in HTTP Connection Manager in a Listener resource.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
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

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"listener_name": schema.StringAttribute{
															Description:         "Name of the listener to match.",
															MarkdownDescription: "Name of the listener to match.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"listener_tags": schema.MapAttribute{
															Description:         "Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															MarkdownDescription: "Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the HTTP filter. For example 'envoy.filters.http.local_ratelimit'",
															MarkdownDescription: "Name of the HTTP filter. For example 'envoy.filters.http.local_ratelimit'",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched listener.",
													MarkdownDescription: "Operation to execute on matched listener.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"listener": schema.SingleNestedAttribute{
											Description:         "Listener is a modification of Envoy's Listener resource.",
											MarkdownDescription: "Listener is a modification of Envoy's Listener resource.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy's Listener resource",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy's Listener resource",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
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

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the listener to match.",
															MarkdownDescription: "Name of the listener to match.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"tags": schema.MapAttribute{
															Description:         "Tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															MarkdownDescription: "Tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched listener.",
													MarkdownDescription: "Operation to execute on matched listener.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"network_filter": schema.SingleNestedAttribute{
											Description:         "NetworkFilter is a modification of Envoy Listener's filter.",
											MarkdownDescription: "NetworkFilter is a modification of Envoy Listener's filter.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy Listener's filter.",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy Listener's filter.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
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

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"listener_name": schema.StringAttribute{
															Description:         "Name of the listener to match.",
															MarkdownDescription: "Name of the listener to match.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"listener_tags": schema.MapAttribute{
															Description:         "Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															MarkdownDescription: "Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the network filter. For example 'envoy.filters.network.ratelimit'",
															MarkdownDescription: "Name of the network filter. For example 'envoy.filters.network.ratelimit'",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched listener.",
													MarkdownDescription: "Operation to execute on matched listener.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"virtual_host": schema.SingleNestedAttribute{
											Description:         "VirtualHost is a modification of Envoy's VirtualHost referenced in HTTP Connection Manager in a Listener resource.",
											MarkdownDescription: "VirtualHost is a modification of Envoy's VirtualHost referenced in HTTP Connection Manager in a Listener resource.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy's VirtualHost resource",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy's VirtualHost resource",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
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

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the VirtualHost to match.",
															MarkdownDescription: "Name of the VirtualHost to match.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"route_configuration_name": schema.StringAttribute{
															Description:         "Name of the RouteConfiguration resource to match.",
															MarkdownDescription: "Name of the RouteConfiguration resource to match.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched listener.",
													MarkdownDescription: "Operation to execute on matched listener.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
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

func (r *KumaIoMeshProxyPatchV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KumaIoMeshProxyPatchV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kuma_io_mesh_proxy_patch_v1alpha1")

	var data KumaIoMeshProxyPatchV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "MeshProxyPatch"}).
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

	var readResponse KumaIoMeshProxyPatchV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("kuma.io/v1alpha1")
	data.Kind = pointer.String("MeshProxyPatch")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
