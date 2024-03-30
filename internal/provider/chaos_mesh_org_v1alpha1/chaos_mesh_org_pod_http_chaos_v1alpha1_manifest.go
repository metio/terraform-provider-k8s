/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

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
	_ datasource.DataSource = &ChaosMeshOrgPodHttpChaosV1Alpha1Manifest{}
)

func NewChaosMeshOrgPodHttpChaosV1Alpha1Manifest() datasource.DataSource {
	return &ChaosMeshOrgPodHttpChaosV1Alpha1Manifest{}
}

type ChaosMeshOrgPodHttpChaosV1Alpha1Manifest struct{}

type ChaosMeshOrgPodHttpChaosV1Alpha1ManifestData struct {
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
		Rules *[]struct {
			Actions *struct {
				Abort *bool   `tfsdk:"abort" json:"abort,omitempty"`
				Delay *string `tfsdk:"delay" json:"delay,omitempty"`
				Patch *struct {
					Body *struct {
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"body" json:"body,omitempty"`
					Headers *[]string `tfsdk:"headers" json:"headers,omitempty"`
					Queries *[]string `tfsdk:"queries" json:"queries,omitempty"`
				} `tfsdk:"patch" json:"patch,omitempty"`
				Replace *struct {
					Body    *string            `tfsdk:"body" json:"body,omitempty"`
					Code    *int64             `tfsdk:"code" json:"code,omitempty"`
					Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					Method  *string            `tfsdk:"method" json:"method,omitempty"`
					Path    *string            `tfsdk:"path" json:"path,omitempty"`
					Queries *map[string]string `tfsdk:"queries" json:"queries,omitempty"`
				} `tfsdk:"replace" json:"replace,omitempty"`
			} `tfsdk:"actions" json:"actions,omitempty"`
			Port     *int64 `tfsdk:"port" json:"port,omitempty"`
			Selector *struct {
				Code             *int64             `tfsdk:"code" json:"code,omitempty"`
				Method           *string            `tfsdk:"method" json:"method,omitempty"`
				Path             *string            `tfsdk:"path" json:"path,omitempty"`
				Port             *int64             `tfsdk:"port" json:"port,omitempty"`
				Request_headers  *map[string]string `tfsdk:"request_headers" json:"request_headers,omitempty"`
				Response_headers *map[string]string `tfsdk:"response_headers" json:"response_headers,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
			Target *string `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		Tls *struct {
			CaName          *string `tfsdk:"ca_name" json:"caName,omitempty"`
			CertName        *string `tfsdk:"cert_name" json:"certName,omitempty"`
			KeyName         *string `tfsdk:"key_name" json:"keyName,omitempty"`
			SecretName      *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			SecretNamespace *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_pod_http_chaos_v1alpha1_manifest"
}

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PodHttpChaos is the Schema for the podhttpchaos API",
		MarkdownDescription: "PodHttpChaos is the Schema for the podhttpchaos API",
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
				Description:         "PodHttpChaosSpec defines the desired state of PodHttpChaos.",
				MarkdownDescription: "PodHttpChaosSpec defines the desired state of PodHttpChaos.",
				Attributes: map[string]schema.Attribute{
					"rules": schema.ListNestedAttribute{
						Description:         "Rules are a list of injection rule for http request.",
						MarkdownDescription: "Rules are a list of injection rule for http request.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"actions": schema.SingleNestedAttribute{
									Description:         "Actions contains rules to inject target.",
									MarkdownDescription: "Actions contains rules to inject target.",
									Attributes: map[string]schema.Attribute{
										"abort": schema.BoolAttribute{
											Description:         "Abort is a rule to abort a http session.",
											MarkdownDescription: "Abort is a rule to abort a http session.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"delay": schema.StringAttribute{
											Description:         "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											MarkdownDescription: "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"patch": schema.SingleNestedAttribute{
											Description:         "Patch is a rule to patch some contents in target.",
											MarkdownDescription: "Patch is a rule to patch some contents in target.",
											Attributes: map[string]schema.Attribute{
												"body": schema.SingleNestedAttribute{
													Description:         "Body is a rule to patch message body of target.",
													MarkdownDescription: "Body is a rule to patch message body of target.",
													Attributes: map[string]schema.Attribute{
														"type": schema.StringAttribute{
															Description:         "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
															MarkdownDescription: "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the patch contents.",
															MarkdownDescription: "Value is the patch contents.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"headers": schema.ListAttribute{
													Description:         "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
													MarkdownDescription: "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"queries": schema.ListAttribute{
													Description:         "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
													MarkdownDescription: "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
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

										"replace": schema.SingleNestedAttribute{
											Description:         "Replace is a rule to replace some contents in target.",
											MarkdownDescription: "Replace is a rule to replace some contents in target.",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "Body is a rule to replace http message body in target.",
													MarkdownDescription: "Body is a rule to replace http message body in target.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														validators.Base64Validator(),
													},
												},

												"code": schema.Int64Attribute{
													Description:         "Code is a rule to replace http status code in response.",
													MarkdownDescription: "Code is a rule to replace http status code in response.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.MapAttribute{
													Description:         "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
													MarkdownDescription: "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"method": schema.StringAttribute{
													Description:         "Method is a rule to replace http method in request.",
													MarkdownDescription: "Method is a rule to replace http method in request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Path is rule to to replace uri path in http request.",
													MarkdownDescription: "Path is rule to to replace uri path in http request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"queries": schema.MapAttribute{
													Description:         "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
													MarkdownDescription: "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
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
									Required: true,
									Optional: false,
									Computed: false,
								},

								"port": schema.Int64Attribute{
									Description:         "Port represents the target port to be proxy of.",
									MarkdownDescription: "Port represents the target port to be proxy of.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"selector": schema.SingleNestedAttribute{
									Description:         "Selector contains the rules to select target.",
									MarkdownDescription: "Selector contains the rules to select target.",
									Attributes: map[string]schema.Attribute{
										"code": schema.Int64Attribute{
											Description:         "Code is a rule to select target by http status code in response.",
											MarkdownDescription: "Code is a rule to select target by http status code in response.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"method": schema.StringAttribute{
											Description:         "Method is a rule to select target by http method in request.",
											MarkdownDescription: "Method is a rule to select target by http method in request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path is a rule to select target by uri path in http request.",
											MarkdownDescription: "Path is a rule to select target by uri path in http request.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "Port is a rule to select server listening on specific port.",
											MarkdownDescription: "Port is a rule to select server listening on specific port.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"request_headers": schema.MapAttribute{
											Description:         "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
											MarkdownDescription: "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"response_headers": schema.MapAttribute{
											Description:         "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
											MarkdownDescription: "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
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

								"source": schema.StringAttribute{
									Description:         "Source represents the source of current rules",
									MarkdownDescription: "Source represents the source of current rules",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target": schema.StringAttribute{
									Description:         "Target is the object to be selected and injected, <Request|Response>.",
									MarkdownDescription: "Target is the object to be selected and injected, <Request|Response>.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS is the tls config, will be override if there are multiple HTTPChaos experiments are applied",
						MarkdownDescription: "TLS is the tls config, will be override if there are multiple HTTPChaos experiments are applied",
						Attributes: map[string]schema.Attribute{
							"ca_name": schema.StringAttribute{
								Description:         "CAName represents the data name of ca file in secret, 'ca.crt' for example",
								MarkdownDescription: "CAName represents the data name of ca file in secret, 'ca.crt' for example",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cert_name": schema.StringAttribute{
								Description:         "CertName represents the data name of cert file in secret, 'tls.crt' for example",
								MarkdownDescription: "CertName represents the data name of cert file in secret, 'tls.crt' for example",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"key_name": schema.StringAttribute{
								Description:         "KeyName represents the data name of key file in secret, 'tls.key' for example",
								MarkdownDescription: "KeyName represents the data name of key file in secret, 'tls.key' for example",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName represents the name of required secret resource",
								MarkdownDescription: "SecretName represents the name of required secret resource",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"secret_namespace": schema.StringAttribute{
								Description:         "SecretNamespace represents the namespace of required secret resource",
								MarkdownDescription: "SecretNamespace represents the namespace of required secret resource",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ChaosMeshOrgPodHttpChaosV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_pod_http_chaos_v1alpha1_manifest")

	var model ChaosMeshOrgPodHttpChaosV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("PodHttpChaos")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
