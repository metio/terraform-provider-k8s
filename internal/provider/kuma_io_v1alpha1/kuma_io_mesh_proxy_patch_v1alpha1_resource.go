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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &KumaIoMeshProxyPatchV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &KumaIoMeshProxyPatchV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &KumaIoMeshProxyPatchV1Alpha1Resource{}
)

func NewKumaIoMeshProxyPatchV1Alpha1Resource() resource.Resource {
	return &KumaIoMeshProxyPatchV1Alpha1Resource{}
}

type KumaIoMeshProxyPatchV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type KumaIoMeshProxyPatchV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

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

func (r *KumaIoMeshProxyPatchV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_proxy_patch_v1alpha1"
}

func (r *KumaIoMeshProxyPatchV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
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
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
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
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy's Listener resource",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy's Listener resource",
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
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
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
													Description:         "JsonPatches specifies list of jsonpatches to apply to on Envoy Listener's filter.",
													MarkdownDescription: "JsonPatches specifies list of jsonpatches to apply to on Envoy Listener's filter.",
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
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
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
															Description:         "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
															MarkdownDescription: "Origin is the name of the component or plugin that generated the resource.  Here is the list of well-known origins: inbound - resources generated for handling incoming traffic. outbound - resources generated for handling outgoing traffic. transparent - resources generated for transparent proxy functionality. prometheus - resources generated when Prometheus metrics are enabled. direct-access - resources generated for Direct Access functionality. ingress - resources generated for Zone Ingress. egress - resources generated for Zone Egress. gateway - resources generated for MeshGateway.  The list is not complete, because policy plugins can introduce new resources. For example MeshTrace plugin can create Cluster with 'mesh-trace' origin.",
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
						Description:         "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						MarkdownDescription: "TargetRef is a reference to the resource the policy takes an effect on. The resource could be either a real store object or virtual resource defined inplace.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshServiceSubset", "MeshHTTPRoute"),
								},
							},

							"mesh": schema.StringAttribute{
								Description:         "Mesh is reserved for future use to identify cross mesh resources.",
								MarkdownDescription: "Mesh is reserved for future use to identify cross mesh resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								MarkdownDescription: "Name of the referenced resource. Can only be used with kinds: 'MeshService', 'MeshServiceSubset' and 'MeshGatewayRoute'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
								MarkdownDescription: "Tags used to select a subset of proxies by tags. Can only be used with kinds 'MeshSubset' and 'MeshServiceSubset'",
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

func (r *KumaIoMeshProxyPatchV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KumaIoMeshProxyPatchV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kuma_io_mesh_proxy_patch_v1alpha1")

	var model KumaIoMeshProxyPatchV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshProxyPatch")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "MeshProxyPatch"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KumaIoMeshProxyPatchV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KumaIoMeshProxyPatchV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_proxy_patch_v1alpha1")

	var data KumaIoMeshProxyPatchV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
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

	var readResponse KumaIoMeshProxyPatchV1Alpha1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *KumaIoMeshProxyPatchV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kuma_io_mesh_proxy_patch_v1alpha1")

	var model KumaIoMeshProxyPatchV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshProxyPatch")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "MeshProxyPatch"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KumaIoMeshProxyPatchV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KumaIoMeshProxyPatchV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kuma_io_mesh_proxy_patch_v1alpha1")

	var data KumaIoMeshProxyPatchV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "MeshProxyPatch"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *KumaIoMeshProxyPatchV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
