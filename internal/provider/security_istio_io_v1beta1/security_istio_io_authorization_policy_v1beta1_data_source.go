/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_istio_io_v1beta1

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
	_ datasource.DataSource              = &SecurityIstioIoAuthorizationPolicyV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &SecurityIstioIoAuthorizationPolicyV1Beta1DataSource{}
)

func NewSecurityIstioIoAuthorizationPolicyV1Beta1DataSource() datasource.DataSource {
	return &SecurityIstioIoAuthorizationPolicyV1Beta1DataSource{}
}

type SecurityIstioIoAuthorizationPolicyV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SecurityIstioIoAuthorizationPolicyV1Beta1DataSourceData struct {
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

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_istio_io_authorization_policy_v1beta1"
}

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Configuration for access control on workloads. See more details at: https://istio.io/docs/reference/config/security/authorization-policy.html",
				MarkdownDescription: "Configuration for access control on workloads. See more details at: https://istio.io/docs/reference/config/security/authorization-policy.html",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"provider": schema.SingleNestedAttribute{
						Description:         "Specifies detailed configuration of the CUSTOM action.",
						MarkdownDescription: "Specifies detailed configuration of the CUSTOM action.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Specifies the name of the extension provider.",
								MarkdownDescription: "Specifies the name of the extension provider.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
														Optional:            false,
														Computed:            true,
													},

													"namespaces": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_ip_blocks": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_namespaces": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_principals": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_remote_ip_blocks": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_request_principals": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"principals": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"remote_ip_blocks": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"request_principals": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
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
														Optional:            false,
														Computed:            true,
													},

													"methods": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_hosts": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_methods": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_paths": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"not_ports": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"paths": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"ports": schema.ListAttribute{
														Description:         "Optional.",
														MarkdownDescription: "Optional.",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
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
												Optional:            false,
												Computed:            true,
											},

											"not_values": schema.ListAttribute{
												Description:         "Optional.",
												MarkdownDescription: "Optional.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"values": schema.ListAttribute{
												Description:         "Optional.",
												MarkdownDescription: "Optional.",
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
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "group is the group of the target resource.",
								MarkdownDescription: "group is the group of the target resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is kind of the target resource.",
								MarkdownDescription: "kind is kind of the target resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the target resource.",
								MarkdownDescription: "name is the name of the target resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace is the namespace of the referent.",
								MarkdownDescription: "namespace is the namespace of the referent.",
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

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SecurityIstioIoAuthorizationPolicyV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_security_istio_io_authorization_policy_v1beta1")

	var data SecurityIstioIoAuthorizationPolicyV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "security.istio.io", Version: "v1beta1", Resource: "AuthorizationPolicy"}).
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

	var readResponse SecurityIstioIoAuthorizationPolicyV1Beta1DataSourceData
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
	data.ApiVersion = pointer.String("security.istio.io/v1beta1")
	data.Kind = pointer.String("AuthorizationPolicy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
