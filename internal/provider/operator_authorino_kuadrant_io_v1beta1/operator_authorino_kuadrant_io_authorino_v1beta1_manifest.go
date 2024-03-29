/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_authorino_kuadrant_io_v1beta1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &OperatorAuthorinoKuadrantIoAuthorinoV1Beta1Manifest{}
)

func NewOperatorAuthorinoKuadrantIoAuthorinoV1Beta1Manifest() datasource.DataSource {
	return &OperatorAuthorinoKuadrantIoAuthorinoV1Beta1Manifest{}
}

type OperatorAuthorinoKuadrantIoAuthorinoV1Beta1Manifest struct{}

type OperatorAuthorinoKuadrantIoAuthorinoV1Beta1ManifestData struct {
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
		AuthConfigLabelSelectors *string `tfsdk:"auth_config_label_selectors" json:"authConfigLabelSelectors,omitempty"`
		ClusterWide              *bool   `tfsdk:"cluster_wide" json:"clusterWide,omitempty"`
		EvaluatorCacheSize       *int64  `tfsdk:"evaluator_cache_size" json:"evaluatorCacheSize,omitempty"`
		Healthz                  *struct {
			Port *int64 `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"healthz" json:"healthz,omitempty"`
		Image           *string `tfsdk:"image" json:"image,omitempty"`
		ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		Listener        *struct {
			MaxHttpRequestBodySize *int64 `tfsdk:"max_http_request_body_size" json:"maxHttpRequestBodySize,omitempty"`
			Port                   *int64 `tfsdk:"port" json:"port,omitempty"`
			Ports                  *struct {
				Grpc *int64 `tfsdk:"grpc" json:"grpc,omitempty"`
				Http *int64 `tfsdk:"http" json:"http,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
			Timeout *int64 `tfsdk:"timeout" json:"timeout,omitempty"`
			Tls     *struct {
				CertSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"listener" json:"listener,omitempty"`
		LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		LogMode  *string `tfsdk:"log_mode" json:"logMode,omitempty"`
		Metrics  *struct {
			Deep *bool  `tfsdk:"deep" json:"deep,omitempty"`
			Port *int64 `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		OidcServer *struct {
			Port *int64 `tfsdk:"port" json:"port,omitempty"`
			Tls  *struct {
				CertSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"oidc_server" json:"oidcServer,omitempty"`
		Replicas               *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		SecretLabelSelectors   *string `tfsdk:"secret_label_selectors" json:"secretLabelSelectors,omitempty"`
		SupersedingHostSubsets *bool   `tfsdk:"superseding_host_subsets" json:"supersedingHostSubsets,omitempty"`
		Tracing                *struct {
			Endpoint *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Insecure *bool              `tfsdk:"insecure" json:"insecure,omitempty"`
			Tags     *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"tracing" json:"tracing,omitempty"`
		Volumes *struct {
			DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
			Items       *[]struct {
				ConfigMaps *[]string `tfsdk:"config_maps" json:"configMaps,omitempty"`
				Items      *[]struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"items" json:"items,omitempty"`
				MountPath *string   `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Name      *string   `tfsdk:"name" json:"name,omitempty"`
				Secrets   *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
			} `tfsdk:"items" json:"items,omitempty"`
		} `tfsdk:"volumes" json:"volumes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorAuthorinoKuadrantIoAuthorinoV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_authorino_kuadrant_io_authorino_v1beta1_manifest"
}

func (r *OperatorAuthorinoKuadrantIoAuthorinoV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Authorino is the Schema for the authorinos API",
		MarkdownDescription: "Authorino is the Schema for the authorinos API",
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
				Description:         "AuthorinoSpec defines the desired state of Authorino",
				MarkdownDescription: "AuthorinoSpec defines the desired state of Authorino",
				Attributes: map[string]schema.Attribute{
					"auth_config_label_selectors": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_wide": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"evaluator_cache_size": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"healthz": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"port": schema.Int64Attribute{
								Description:         "Port number of the health/readiness probe endpoints.",
								MarkdownDescription: "Port number of the health/readiness probe endpoints.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "PullPolicy describes a policy for if/when to pull a container image",
						MarkdownDescription: "PullPolicy describes a policy for if/when to pull a container image",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"listener": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"max_http_request_body_size": schema.Int64Attribute{
								Description:         "Maximum payload (request body) size for the auth service (HTTP interface), in bytes.",
								MarkdownDescription: "Maximum payload (request body) size for the auth service (HTTP interface), in bytes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port number of the GRPC interface. DEPRECATED: use 'ports.grpc' instead.",
								MarkdownDescription: "Port number of the GRPC interface. DEPRECATED: use 'ports.grpc' instead.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ports": schema.SingleNestedAttribute{
								Description:         "Port numbers of the GRPC and HTTP auth interfaces.",
								MarkdownDescription: "Port numbers of the GRPC and HTTP auth interfaces.",
								Attributes: map[string]schema.Attribute{
									"grpc": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"http": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": schema.Int64Attribute{
								Description:         "Timeout of the auth service (GRPC and HTTP interfaces), in milliseconds.",
								MarkdownDescription: "Timeout of the auth service (GRPC and HTTP interfaces), in milliseconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS configuration of the auth service (GRPC and HTTP interfaces).",
								MarkdownDescription: "TLS configuration of the auth service (GRPC and HTTP interfaces).",
								Attributes: map[string]schema.Attribute{
									"cert_secret_ref": schema.SingleNestedAttribute{
										Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
										MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"log_level": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_mode": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"deep": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"oidc_server": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cert_secret_ref": schema.SingleNestedAttribute{
										Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
										MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_label_selectors": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"superseding_host_subsets": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tracing": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"insecure": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
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

					"volumes": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"default_mode": schema.Int64Attribute{
								Description:         "Permissions mode.",
								MarkdownDescription: "Permissions mode.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"items": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config_maps": schema.ListAttribute{
											Description:         "Allow multiple configmaps to mount to the same directory",
											MarkdownDescription: "Allow multiple configmaps to mount to the same directory",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"items": schema.ListNestedAttribute{
											Description:         "Mount details",
											MarkdownDescription: "Mount details",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the key to project.",
														MarkdownDescription: "key is the key to project.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"mode": schema.Int64Attribute{
														Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
														MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

										"mount_path": schema.StringAttribute{
											Description:         "An absolute path where to mount it",
											MarkdownDescription: "An absolute path where to mount it",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Volume name",
											MarkdownDescription: "Volume name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secrets": schema.ListAttribute{
											Description:         "Secret mount",
											MarkdownDescription: "Secret mount",
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

func (r *OperatorAuthorinoKuadrantIoAuthorinoV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_authorino_kuadrant_io_authorino_v1beta1_manifest")

	var model OperatorAuthorinoKuadrantIoAuthorinoV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("operator.authorino.kuadrant.io/v1beta1")
	model.Kind = pointer.String("Authorino")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
