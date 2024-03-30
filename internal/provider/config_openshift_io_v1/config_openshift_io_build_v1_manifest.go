/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoBuildV1Manifest{}
)

func NewConfigOpenshiftIoBuildV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoBuildV1Manifest{}
}

type ConfigOpenshiftIoBuildV1Manifest struct{}

type ConfigOpenshiftIoBuildV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AdditionalTrustedCA *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"additional_trusted_ca" json:"additionalTrustedCA,omitempty"`
		BuildDefaults *struct {
			DefaultProxy *struct {
				HttpProxy          *string   `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
				HttpsProxy         *string   `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
				NoProxy            *string   `tfsdk:"no_proxy" json:"noProxy,omitempty"`
				ReadinessEndpoints *[]string `tfsdk:"readiness_endpoints" json:"readinessEndpoints,omitempty"`
				TrustedCA          *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
			} `tfsdk:"default_proxy" json:"defaultProxy,omitempty"`
			Env *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
						Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"env" json:"env,omitempty"`
			GitProxy *struct {
				HttpProxy          *string   `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
				HttpsProxy         *string   `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
				NoProxy            *string   `tfsdk:"no_proxy" json:"noProxy,omitempty"`
				ReadinessEndpoints *[]string `tfsdk:"readiness_endpoints" json:"readinessEndpoints,omitempty"`
				TrustedCA          *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
			} `tfsdk:"git_proxy" json:"gitProxy,omitempty"`
			ImageLabels *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"image_labels" json:"imageLabels,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"build_defaults" json:"buildDefaults,omitempty"`
		BuildOverrides *struct {
			ForcePull   *bool `tfsdk:"force_pull" json:"forcePull,omitempty"`
			ImageLabels *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"image_labels" json:"imageLabels,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations  *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"build_overrides" json:"buildOverrides,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoBuildV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_build_v1_manifest"
}

func (r *ConfigOpenshiftIoBuildV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Build configures the behavior of OpenShift builds for the entire cluster. This includes default settings that can be overridden in BuildConfig objects, and overrides which are applied to all builds.  The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Build configures the behavior of OpenShift builds for the entire cluster. This includes default settings that can be overridden in BuildConfig objects, and overrides which are applied to all builds.  The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "Spec holds user-settable values for the build controller configuration",
				MarkdownDescription: "Spec holds user-settable values for the build controller configuration",
				Attributes: map[string]schema.Attribute{
					"additional_trusted_ca": schema.SingleNestedAttribute{
						Description:         "AdditionalTrustedCA is a reference to a ConfigMap containing additional CAs that should be trusted for image pushes and pulls during builds. The namespace for this config map is openshift-config.  DEPRECATED: Additional CAs for image pull and push should be set on image.config.openshift.io/cluster instead.",
						MarkdownDescription: "AdditionalTrustedCA is a reference to a ConfigMap containing additional CAs that should be trusted for image pushes and pulls during builds. The namespace for this config map is openshift-config.  DEPRECATED: Additional CAs for image pull and push should be set on image.config.openshift.io/cluster instead.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the metadata.name of the referenced config map",
								MarkdownDescription: "name is the metadata.name of the referenced config map",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"build_defaults": schema.SingleNestedAttribute{
						Description:         "BuildDefaults controls the default information for Builds",
						MarkdownDescription: "BuildDefaults controls the default information for Builds",
						Attributes: map[string]schema.Attribute{
							"default_proxy": schema.SingleNestedAttribute{
								Description:         "DefaultProxy contains the default proxy settings for all build operations, including image pull/push and source download.  Values can be overrode by setting the 'HTTP_PROXY', 'HTTPS_PROXY', and 'NO_PROXY' environment variables in the build config's strategy.",
								MarkdownDescription: "DefaultProxy contains the default proxy settings for all build operations, including image pull/push and source download.  Values can be overrode by setting the 'HTTP_PROXY', 'HTTPS_PROXY', and 'NO_PROXY' environment variables in the build config's strategy.",
								Attributes: map[string]schema.Attribute{
									"http_proxy": schema.StringAttribute{
										Description:         "httpProxy is the URL of the proxy for HTTP requests.  Empty means unset and will not result in an env var.",
										MarkdownDescription: "httpProxy is the URL of the proxy for HTTP requests.  Empty means unset and will not result in an env var.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"https_proxy": schema.StringAttribute{
										Description:         "httpsProxy is the URL of the proxy for HTTPS requests.  Empty means unset and will not result in an env var.",
										MarkdownDescription: "httpsProxy is the URL of the proxy for HTTPS requests.  Empty means unset and will not result in an env var.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"no_proxy": schema.StringAttribute{
										Description:         "noProxy is a comma-separated list of hostnames and/or CIDRs and/or IPs for which the proxy should not be used. Empty means unset and will not result in an env var.",
										MarkdownDescription: "noProxy is a comma-separated list of hostnames and/or CIDRs and/or IPs for which the proxy should not be used. Empty means unset and will not result in an env var.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"readiness_endpoints": schema.ListAttribute{
										Description:         "readinessEndpoints is a list of endpoints used to verify readiness of the proxy.",
										MarkdownDescription: "readinessEndpoints is a list of endpoints used to verify readiness of the proxy.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"trusted_ca": schema.SingleNestedAttribute{
										Description:         "trustedCA is a reference to a ConfigMap containing a CA certificate bundle. The trustedCA field should only be consumed by a proxy validator. The validator is responsible for reading the certificate bundle from the required key 'ca-bundle.crt', merging it with the system default trust bundle, and writing the merged trust bundle to a ConfigMap named 'trusted-ca-bundle' in the 'openshift-config-managed' namespace. Clients that expect to make proxy connections must use the trusted-ca-bundle for all HTTPS requests to the proxy, and may use the trusted-ca-bundle for non-proxy HTTPS requests as well.  The namespace for the ConfigMap referenced by trustedCA is 'openshift-config'. Here is an example ConfigMap (in yaml):  apiVersion: v1 kind: ConfigMap metadata: name: user-ca-bundle namespace: openshift-config data: ca-bundle.crt: | -----BEGIN CERTIFICATE----- Custom CA certificate bundle. -----END CERTIFICATE-----",
										MarkdownDescription: "trustedCA is a reference to a ConfigMap containing a CA certificate bundle. The trustedCA field should only be consumed by a proxy validator. The validator is responsible for reading the certificate bundle from the required key 'ca-bundle.crt', merging it with the system default trust bundle, and writing the merged trust bundle to a ConfigMap named 'trusted-ca-bundle' in the 'openshift-config-managed' namespace. Clients that expect to make proxy connections must use the trusted-ca-bundle for all HTTPS requests to the proxy, and may use the trusted-ca-bundle for non-proxy HTTPS requests as well.  The namespace for the ConfigMap referenced by trustedCA is 'openshift-config'. Here is an example ConfigMap (in yaml):  apiVersion: v1 kind: ConfigMap metadata: name: user-ca-bundle namespace: openshift-config data: ca-bundle.crt: | -----BEGIN CERTIFICATE----- Custom CA certificate bundle. -----END CERTIFICATE-----",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",
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

							"env": schema.ListNestedAttribute{
								Description:         "Env is a set of default environment variables that will be applied to the build if the specified variables do not exist on the build",
								MarkdownDescription: "Env is a set of default environment variables that will be applied to the build if the specified variables do not exist on the build",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
											MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
											MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value_from": schema.SingleNestedAttribute{
											Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
											MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a ConfigMap.",
													MarkdownDescription: "Selects a key of a ConfigMap.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the ConfigMap or its key must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"field_ref": schema.SingleNestedAttribute{
													Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
													MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
													Attributes: map[string]schema.Attribute{
														"api_version": schema.StringAttribute{
															Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
															MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"field_path": schema.StringAttribute{
															Description:         "Path of the field to select in the specified API version.",
															MarkdownDescription: "Path of the field to select in the specified API version.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resource_field_ref": schema.SingleNestedAttribute{
													Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
													MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
													Attributes: map[string]schema.Attribute{
														"container_name": schema.StringAttribute{
															Description:         "Container name: required for volumes, optional for env vars",
															MarkdownDescription: "Container name: required for volumes, optional for env vars",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"divisor": schema.StringAttribute{
															Description:         "Specifies the output format of the exposed resources, defaults to '1'",
															MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resource": schema.StringAttribute{
															Description:         "Required: resource to select",
															MarkdownDescription: "Required: resource to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a secret in the pod's namespace",
													MarkdownDescription: "Selects a key of a secret in the pod's namespace",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from.  Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the Secret or its key must be defined",
															MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"git_proxy": schema.SingleNestedAttribute{
								Description:         "GitProxy contains the proxy settings for git operations only. If set, this will override any Proxy settings for all git commands, such as git clone.  Values that are not set here will be inherited from DefaultProxy.",
								MarkdownDescription: "GitProxy contains the proxy settings for git operations only. If set, this will override any Proxy settings for all git commands, such as git clone.  Values that are not set here will be inherited from DefaultProxy.",
								Attributes: map[string]schema.Attribute{
									"http_proxy": schema.StringAttribute{
										Description:         "httpProxy is the URL of the proxy for HTTP requests.  Empty means unset and will not result in an env var.",
										MarkdownDescription: "httpProxy is the URL of the proxy for HTTP requests.  Empty means unset and will not result in an env var.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"https_proxy": schema.StringAttribute{
										Description:         "httpsProxy is the URL of the proxy for HTTPS requests.  Empty means unset and will not result in an env var.",
										MarkdownDescription: "httpsProxy is the URL of the proxy for HTTPS requests.  Empty means unset and will not result in an env var.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"no_proxy": schema.StringAttribute{
										Description:         "noProxy is a comma-separated list of hostnames and/or CIDRs and/or IPs for which the proxy should not be used. Empty means unset and will not result in an env var.",
										MarkdownDescription: "noProxy is a comma-separated list of hostnames and/or CIDRs and/or IPs for which the proxy should not be used. Empty means unset and will not result in an env var.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"readiness_endpoints": schema.ListAttribute{
										Description:         "readinessEndpoints is a list of endpoints used to verify readiness of the proxy.",
										MarkdownDescription: "readinessEndpoints is a list of endpoints used to verify readiness of the proxy.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"trusted_ca": schema.SingleNestedAttribute{
										Description:         "trustedCA is a reference to a ConfigMap containing a CA certificate bundle. The trustedCA field should only be consumed by a proxy validator. The validator is responsible for reading the certificate bundle from the required key 'ca-bundle.crt', merging it with the system default trust bundle, and writing the merged trust bundle to a ConfigMap named 'trusted-ca-bundle' in the 'openshift-config-managed' namespace. Clients that expect to make proxy connections must use the trusted-ca-bundle for all HTTPS requests to the proxy, and may use the trusted-ca-bundle for non-proxy HTTPS requests as well.  The namespace for the ConfigMap referenced by trustedCA is 'openshift-config'. Here is an example ConfigMap (in yaml):  apiVersion: v1 kind: ConfigMap metadata: name: user-ca-bundle namespace: openshift-config data: ca-bundle.crt: | -----BEGIN CERTIFICATE----- Custom CA certificate bundle. -----END CERTIFICATE-----",
										MarkdownDescription: "trustedCA is a reference to a ConfigMap containing a CA certificate bundle. The trustedCA field should only be consumed by a proxy validator. The validator is responsible for reading the certificate bundle from the required key 'ca-bundle.crt', merging it with the system default trust bundle, and writing the merged trust bundle to a ConfigMap named 'trusted-ca-bundle' in the 'openshift-config-managed' namespace. Clients that expect to make proxy connections must use the trusted-ca-bundle for all HTTPS requests to the proxy, and may use the trusted-ca-bundle for non-proxy HTTPS requests as well.  The namespace for the ConfigMap referenced by trustedCA is 'openshift-config'. Here is an example ConfigMap (in yaml):  apiVersion: v1 kind: ConfigMap metadata: name: user-ca-bundle namespace: openshift-config data: ca-bundle.crt: | -----BEGIN CERTIFICATE----- Custom CA certificate bundle. -----END CERTIFICATE-----",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is the metadata.name of the referenced config map",
												MarkdownDescription: "name is the metadata.name of the referenced config map",
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

							"image_labels": schema.ListNestedAttribute{
								Description:         "ImageLabels is a list of docker labels that are applied to the resulting image. User can override a default label by providing a label with the same name in their Build/BuildConfig.",
								MarkdownDescription: "ImageLabels is a list of docker labels that are applied to the resulting image. User can override a default label by providing a label with the same name in their Build/BuildConfig.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name defines the name of the label. It must have non-zero length.",
											MarkdownDescription: "Name defines the name of the label. It must have non-zero length.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value defines the literal value of the label.",
											MarkdownDescription: "Value defines the literal value of the label.",
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

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources defines resource requirements to execute the build.",
								MarkdownDescription: "Resources defines resource requirements to execute the build.",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"build_overrides": schema.SingleNestedAttribute{
						Description:         "BuildOverrides controls override settings for builds",
						MarkdownDescription: "BuildOverrides controls override settings for builds",
						Attributes: map[string]schema.Attribute{
							"force_pull": schema.BoolAttribute{
								Description:         "ForcePull overrides, if set, the equivalent value in the builds, i.e. false disables force pull for all builds, true enables force pull for all builds, independently of what each build specifies itself",
								MarkdownDescription: "ForcePull overrides, if set, the equivalent value in the builds, i.e. false disables force pull for all builds, true enables force pull for all builds, independently of what each build specifies itself",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_labels": schema.ListNestedAttribute{
								Description:         "ImageLabels is a list of docker labels that are applied to the resulting image. If user provided a label in their Build/BuildConfig with the same name as one in this list, the user's label will be overwritten.",
								MarkdownDescription: "ImageLabels is a list of docker labels that are applied to the resulting image. If user provided a label in their Build/BuildConfig with the same name as one in this list, the user's label will be overwritten.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name defines the name of the label. It must have non-zero length.",
											MarkdownDescription: "Name defines the name of the label. It must have non-zero length.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value defines the literal value of the label.",
											MarkdownDescription: "Value defines the literal value of the label.",
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

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the build pod to fit on a node",
								MarkdownDescription: "NodeSelector is a selector which must be true for the build pod to fit on a node",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations is a list of Tolerations that will override any existing tolerations set on a build pod.",
								MarkdownDescription: "Tolerations is a list of Tolerations that will override any existing tolerations set on a build pod.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConfigOpenshiftIoBuildV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_build_v1_manifest")

	var model ConfigOpenshiftIoBuildV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("Build")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
