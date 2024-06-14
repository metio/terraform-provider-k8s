/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

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
	_ datasource.DataSource = &KumaIoMeshProxyPatchV1Alpha1Manifest{}
)

func NewKumaIoMeshProxyPatchV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshProxyPatchV1Alpha1Manifest{}
}

type KumaIoMeshProxyPatchV1Alpha1Manifest struct{}

type KumaIoMeshProxyPatchV1Alpha1ManifestData struct {
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
			Kind        *string            `tfsdk:"kind" json:"kind,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Mesh        *string            `tfsdk:"mesh" json:"mesh,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			ProxyTypes  *[]string          `tfsdk:"proxy_types" json:"proxyTypes,omitempty"`
			SectionName *string            `tfsdk:"section_name" json:"sectionName,omitempty"`
			Tags        *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshProxyPatchV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_proxy_patch_v1alpha1_manifest"
}

func (r *KumaIoMeshProxyPatchV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Spec is the specification of the Kuma MeshProxyPatch resource.",
				MarkdownDescription: "Spec is the specification of the Kuma MeshProxyPatch resource.",
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Description:         "Default is a configuration specific to the group of destinationsreferenced in 'targetRef'.",
						MarkdownDescription: "Default is a configuration specific to the group of destinationsreferenced in 'targetRef'.",
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
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy's Clusterresource",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy's Clusterresource",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove", "replace", "move", "copy"),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the cluster to match.",
															MarkdownDescription: "Name of the cluster to match.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched cluster.",
													MarkdownDescription: "Operation to execute on matched cluster.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Add", "Remove", "Patch"),
													},
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http_filter": schema.SingleNestedAttribute{
											Description:         "HTTPFilter is a modification of Envoy HTTP Filteravailable in HTTP Connection Manager in a Listener resource.",
											MarkdownDescription: "HTTPFilter is a modification of Envoy HTTP Filteravailable in HTTP Connection Manager in a Listener resource.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy'sHTTP Filter available in HTTP Connection Manager in a Listener resource.",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy'sHTTP Filter available in HTTP Connection Manager in a Listener resource.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove", "replace", "move", "copy"),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"listener_name": schema.StringAttribute{
															Description:         "Name of the listener to match.",
															MarkdownDescription: "Name of the listener to match.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"listener_tags": schema.MapAttribute{
															Description:         "Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															MarkdownDescription: "Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the HTTP filter. For example 'envoy.filters.http.local_ratelimit'",
															MarkdownDescription: "Name of the HTTP filter. For example 'envoy.filters.http.local_ratelimit'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched listener.",
													MarkdownDescription: "Operation to execute on matched listener.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Remove", "Patch", "AddFirst", "AddBefore", "AddAfter", "AddLast"),
													},
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"listener": schema.SingleNestedAttribute{
											Description:         "Listener is a modification of Envoy's Listener resource.",
											MarkdownDescription: "Listener is a modification of Envoy's Listener resource.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy's Listenerresource",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy's Listenerresource",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove", "replace", "move", "copy"),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the listener to match.",
															MarkdownDescription: "Name of the listener to match.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tags": schema.MapAttribute{
															Description:         "Tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															MarkdownDescription: "Tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
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

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched listener.",
													MarkdownDescription: "Operation to execute on matched listener.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Add", "Remove", "Patch"),
													},
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"network_filter": schema.SingleNestedAttribute{
											Description:         "NetworkFilter is a modification of Envoy Listener's filter.",
											MarkdownDescription: "NetworkFilter is a modification of Envoy Listener's filter.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy Listener'sfilter.",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy Listener'sfilter.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove", "replace", "move", "copy"),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"listener_name": schema.StringAttribute{
															Description:         "Name of the listener to match.",
															MarkdownDescription: "Name of the listener to match.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"listener_tags": schema.MapAttribute{
															Description:         "Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															MarkdownDescription: "Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the network filter. For example 'envoy.filters.network.ratelimit'",
															MarkdownDescription: "Name of the network filter. For example 'envoy.filters.network.ratelimit'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched listener.",
													MarkdownDescription: "Operation to execute on matched listener.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Remove", "Patch", "AddFirst", "AddBefore", "AddAfter", "AddLast"),
													},
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"virtual_host": schema.SingleNestedAttribute{
											Description:         "VirtualHost is a modification of Envoy's VirtualHostreferenced in HTTP Connection Manager in a Listener resource.",
											MarkdownDescription: "VirtualHost is a modification of Envoy's VirtualHostreferenced in HTTP Connection Manager in a Listener resource.",
											Attributes: map[string]schema.Attribute{
												"json_patches": schema.ListNestedAttribute{
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy'sVirtualHost resource",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy'sVirtualHost resource",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from": schema.StringAttribute{
																Description:         "From is a jsonpatch from string, used by move and copy operations.",
																MarkdownDescription: "From is a jsonpatch from string, used by move and copy operations.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"op": schema.StringAttribute{
																Description:         "Op is a jsonpatch operation string.",
																MarkdownDescription: "Op is a jsonpatch operation string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove", "replace", "move", "copy"),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path is a jsonpatch path string.",
																MarkdownDescription: "Path is a jsonpatch path string.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Value must be a valid json value used by replace and add operations.",
																MarkdownDescription: "Value must be a valid json value used by replace and add operations.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match": schema.SingleNestedAttribute{
													Description:         "Match is a set of conditions that have to be matched for modification operation to happen.",
													MarkdownDescription: "Match is a set of conditions that have to be matched for modification operation to happen.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the VirtualHost to match.",
															MarkdownDescription: "Name of the VirtualHost to match.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin": schema.StringAttribute{
															Description:         "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.Here is the list of well-known origins:inbound - resources generated for handling incoming traffic.outbound - resources generated for handling outgoing traffic.transparent - resources generated for transparent proxy functionality.prometheus - resources generated when Prometheus metrics are enabled.direct-access - resources generated for Direct Access functionality.ingress - resources generated for Zone Ingress.egress - resources generated for Zone Egress.gateway - resources generated for MeshGateway.The list is not complete, because policy plugins can introduce new resources.For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"route_configuration_name": schema.StringAttribute{
															Description:         "Name of the RouteConfiguration resource to match.",
															MarkdownDescription: "Name of the RouteConfiguration resource to match.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"operation": schema.StringAttribute{
													Description:         "Operation to execute on matched listener.",
													MarkdownDescription: "Operation to execute on matched listener.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Add", "Remove", "Patch"),
													},
												},

												"value": schema.StringAttribute{
													Description:         "Value of xDS resource in YAML format to add or patch.",
													MarkdownDescription: "Value of xDS resource in YAML format to add or patch.",
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
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on.The resource could be either a real store object or virtual resourcedefined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshServiceSubset", "MeshHTTPRoute"),
								},
							},

							"labels": schema.MapAttribute{
								Description:         "Labels are used to select group of MeshServices that match labels. Either Labels orName and Namespace can be used.",
								MarkdownDescription: "Labels are used to select group of MeshServices that match labels. Either Labels orName and Namespace can be used.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService','MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace specifies the namespace of target resource. If empty only resources in policy namespacewill be targeted.",
								MarkdownDescription: "Namespace specifies the namespace of target resource. If empty only resources in policy namespacewill be targeted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_types": schema.ListAttribute{
								Description:         "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								MarkdownDescription: "ProxyTypes specifies the data plane types that are subject to the policy. When not specified,all data plane types are targeted by the policy.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"section_name": schema.StringAttribute{
								Description:         "SectionName is used to target specific section of resource.For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.",
								MarkdownDescription: "SectionName is used to target specific section of resource.For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds'MeshSubset' and 'MeshServiceSubset'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KumaIoMeshProxyPatchV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_proxy_patch_v1alpha1_manifest")

	var model KumaIoMeshProxyPatchV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshProxyPatch")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
