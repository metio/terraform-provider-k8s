/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_istio_io_v1beta1

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
	_ datasource.DataSource = &SecurityIstioIoAuthorizationPolicyV1Beta1Manifest{}
)

func NewSecurityIstioIoAuthorizationPolicyV1Beta1Manifest() datasource.DataSource {
	return &SecurityIstioIoAuthorizationPolicyV1Beta1Manifest{}
}

type SecurityIstioIoAuthorizationPolicyV1Beta1Manifest struct{}

type SecurityIstioIoAuthorizationPolicyV1Beta1ManifestData struct {
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
		Action   *string `tfsdk:"action" json:"action,omitempty"`
		Provider *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"provider" json:"provider,omitempty"`
		Rules *[]struct {
			From *[]struct {
				Source *struct {
					IpBlocks             *[]string `tfsdk:"ip_blocks" json:"ipBlocks,omitempty"`
					Namespaces           *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NotIpBlocks          *[]string `tfsdk:"not_ip_blocks" json:"notIpBlocks,omitempty"`
					NotNamespaces        *[]string `tfsdk:"not_namespaces" json:"notNamespaces,omitempty"`
					NotPrincipals        *[]string `tfsdk:"not_principals" json:"notPrincipals,omitempty"`
					NotRemoteIpBlocks    *[]string `tfsdk:"not_remote_ip_blocks" json:"notRemoteIpBlocks,omitempty"`
					NotRequestPrincipals *[]string `tfsdk:"not_request_principals" json:"notRequestPrincipals,omitempty"`
					Principals           *[]string `tfsdk:"principals" json:"principals,omitempty"`
					RemoteIpBlocks       *[]string `tfsdk:"remote_ip_blocks" json:"remoteIpBlocks,omitempty"`
					RequestPrincipals    *[]string `tfsdk:"request_principals" json:"requestPrincipals,omitempty"`
				} `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
			To *[]struct {
				Operation *struct {
					Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
					Methods    *[]string `tfsdk:"methods" json:"methods,omitempty"`
					NotHosts   *[]string `tfsdk:"not_hosts" json:"notHosts,omitempty"`
					NotMethods *[]string `tfsdk:"not_methods" json:"notMethods,omitempty"`
					NotPaths   *[]string `tfsdk:"not_paths" json:"notPaths,omitempty"`
					NotPorts   *[]string `tfsdk:"not_ports" json:"notPorts,omitempty"`
					Paths      *[]string `tfsdk:"paths" json:"paths,omitempty"`
					Ports      *[]string `tfsdk:"ports" json:"ports,omitempty"`
				} `tfsdk:"operation" json:"operation,omitempty"`
			} `tfsdk:"to" json:"to,omitempty"`
			When *[]struct {
				Key       *string   `tfsdk:"key" json:"key,omitempty"`
				NotValues *[]string `tfsdk:"not_values" json:"notValues,omitempty"`
				Values    *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"when" json:"when,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		Selector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_istio_io_authorization_policy_v1beta1_manifest"
}

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Configuration for access control on workloads. See more details at: https://istio.io/docs/reference/config/security/authorization-policy.html",
				MarkdownDescription: "Configuration for access control on workloads. See more details at: https://istio.io/docs/reference/config/security/authorization-policy.html",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ALLOW", "DENY", "AUDIT", "CUSTOM"),
						},
					},

					"provider": schema.SingleNestedAttribute{
						Description:         "Specifies detailed configuration of the CUSTOM action.",
						MarkdownDescription: "Specifies detailed configuration of the CUSTOM action.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Specifies the name of the extension provider.",
								MarkdownDescription: "Specifies the name of the extension provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rules": schema.ListNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.ListNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"source": schema.SingleNestedAttribute{
												Description:         "Source specifies the source of a request.",
												MarkdownDescription: "Source specifies the source of a request.",
												Attributes: map[string]schema.Attribute{
													"ip_blocks": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespaces": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_ip_blocks": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_namespaces": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_principals": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_remote_ip_blocks": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_request_principals": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"principals": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"remote_ip_blocks": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"request_principals": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to": schema.ListNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"operation": schema.SingleNestedAttribute{
												Description:         "Operation specifies the operation of a request.",
												MarkdownDescription: "Operation specifies the operation of a request.",
												Attributes: map[string]schema.Attribute{
													"hosts": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"methods": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_hosts": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_methods": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_paths": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"not_ports": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"paths": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ports": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"when": schema.ListNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The name of an Istio attribute.",
												MarkdownDescription: "The name of an Istio attribute.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"not_values": schema.ListAttribute{
												Description:         "Optional.",
												MarkdownDescription: "Optional.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"values": schema.ListAttribute{
												Description:         "Optional.",
												MarkdownDescription: "Optional.",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "group is the group of the target resource.",
								MarkdownDescription: "group is the group of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is kind of the target resource.",
								MarkdownDescription: "kind is kind of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the target resource.",
								MarkdownDescription: "name is the name of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace is the namespace of the referent.",
								MarkdownDescription: "namespace is the namespace of the referent.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_istio_io_authorization_policy_v1beta1_manifest")

	var model SecurityIstioIoAuthorizationPolicyV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("security.istio.io/v1beta1")
	model.Kind = pointer.String("AuthorizationPolicy")

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
