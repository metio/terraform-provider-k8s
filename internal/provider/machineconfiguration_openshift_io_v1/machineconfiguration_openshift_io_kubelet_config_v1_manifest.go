/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package machineconfiguration_openshift_io_v1

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
	_ datasource.DataSource = &MachineconfigurationOpenshiftIoKubeletConfigV1Manifest{}
)

func NewMachineconfigurationOpenshiftIoKubeletConfigV1Manifest() datasource.DataSource {
	return &MachineconfigurationOpenshiftIoKubeletConfigV1Manifest{}
}

type MachineconfigurationOpenshiftIoKubeletConfigV1Manifest struct{}

type MachineconfigurationOpenshiftIoKubeletConfigV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AutoSizingReserved        *bool              `tfsdk:"auto_sizing_reserved" json:"autoSizingReserved,omitempty"`
		KubeletConfig             *map[string]string `tfsdk:"kubelet_config" json:"kubeletConfig,omitempty"`
		LogLevel                  *int64             `tfsdk:"log_level" json:"logLevel,omitempty"`
		MachineConfigPoolSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"machine_config_pool_selector" json:"machineConfigPoolSelector,omitempty"`
		TlsSecurityProfile *struct {
			Custom *struct {
				Ciphers       *[]string `tfsdk:"ciphers" json:"ciphers,omitempty"`
				MinTLSVersion *string   `tfsdk:"min_tls_version" json:"minTLSVersion,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Intermediate *map[string]string `tfsdk:"intermediate" json:"intermediate,omitempty"`
			Modern       *map[string]string `tfsdk:"modern" json:"modern,omitempty"`
			Old          *map[string]string `tfsdk:"old" json:"old,omitempty"`
			Type         *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"tls_security_profile" json:"tlsSecurityProfile,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MachineconfigurationOpenshiftIoKubeletConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_machineconfiguration_openshift_io_kubelet_config_v1_manifest"
}

func (r *MachineconfigurationOpenshiftIoKubeletConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KubeletConfig describes a customized Kubelet configuration.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "KubeletConfig describes a customized Kubelet configuration.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "KubeletConfigSpec defines the desired state of KubeletConfig",
				MarkdownDescription: "KubeletConfigSpec defines the desired state of KubeletConfig",
				Attributes: map[string]schema.Attribute{
					"auto_sizing_reserved": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kubelet_config": schema.MapAttribute{
						Description:         "kubeletConfig fields are defined in kubernetes upstream. Please refer to the types defined in the version/commit used by OpenShift of the upstream kubernetes. It's important to note that, since the fields of the kubelet configuration are directly fetched from upstream the validation of those values is handled directly by the kubelet. Please refer to the upstream version of the relevant kubernetes for the valid values of these fields. Invalid values of the kubelet configuration fields may render cluster nodes unusable.",
						MarkdownDescription: "kubeletConfig fields are defined in kubernetes upstream. Please refer to the types defined in the version/commit used by OpenShift of the upstream kubernetes. It's important to note that, since the fields of the kubelet configuration are directly fetched from upstream the validation of those values is handled directly by the kubelet. Please refer to the upstream version of the relevant kubernetes for the valid values of these fields. Invalid values of the kubelet configuration fields may render cluster nodes unusable.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_level": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"machine_config_pool_selector": schema.SingleNestedAttribute{
						Description:         "MachineConfigPoolSelector selects which pools the KubeletConfig shoud apply to. A nil selector will result in no pools being selected.",
						MarkdownDescription: "MachineConfigPoolSelector selects which pools the KubeletConfig shoud apply to. A nil selector will result in no pools being selected.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"tls_security_profile": schema.SingleNestedAttribute{
						Description:         "If unset, the default is based on the apiservers.config.openshift.io/cluster resource. Note that only Old and Intermediate profiles are currently supported, and the maximum available minTLSVersion is VersionTLS12.",
						MarkdownDescription: "If unset, the default is based on the apiservers.config.openshift.io/cluster resource. Note that only Old and Intermediate profiles are currently supported, and the maximum available minTLSVersion is VersionTLS12.",
						Attributes: map[string]schema.Attribute{
							"custom": schema.SingleNestedAttribute{
								Description:         "custom is a user-defined TLS security profile. Be extremely careful using a custom profile as invalid configurations can be catastrophic. An example custom profile looks like this:  ciphers:  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  minTLSVersion: VersionTLS11",
								MarkdownDescription: "custom is a user-defined TLS security profile. Be extremely careful using a custom profile as invalid configurations can be catastrophic. An example custom profile looks like this:  ciphers:  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  minTLSVersion: VersionTLS11",
								Attributes: map[string]schema.Attribute{
									"ciphers": schema.ListAttribute{
										Description:         "ciphers is used to specify the cipher algorithms that are negotiated during the TLS handshake.  Operators may remove entries their operands do not support.  For example, to use DES-CBC3-SHA  (yaml):  ciphers: - DES-CBC3-SHA",
										MarkdownDescription: "ciphers is used to specify the cipher algorithms that are negotiated during the TLS handshake.  Operators may remove entries their operands do not support.  For example, to use DES-CBC3-SHA  (yaml):  ciphers: - DES-CBC3-SHA",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_tls_version": schema.StringAttribute{
										Description:         "minTLSVersion is used to specify the minimal version of the TLS protocol that is negotiated during the TLS handshake. For example, to use TLS versions 1.1, 1.2 and 1.3 (yaml):  minTLSVersion: VersionTLS11  NOTE: currently the highest minTLSVersion allowed is VersionTLS12",
										MarkdownDescription: "minTLSVersion is used to specify the minimal version of the TLS protocol that is negotiated during the TLS handshake. For example, to use TLS versions 1.1, 1.2 and 1.3 (yaml):  minTLSVersion: VersionTLS11  NOTE: currently the highest minTLSVersion allowed is VersionTLS12",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("VersionTLS10", "VersionTLS11", "VersionTLS12", "VersionTLS13"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"intermediate": schema.MapAttribute{
								Description:         "intermediate is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  minTLSVersion: VersionTLS12",
								MarkdownDescription: "intermediate is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Intermediate_compatibility_.28recommended.29  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  minTLSVersion: VersionTLS12",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"modern": schema.MapAttribute{
								Description:         "modern is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  minTLSVersion: VersionTLS13",
								MarkdownDescription: "modern is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  minTLSVersion: VersionTLS13",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"old": schema.MapAttribute{
								Description:         "old is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  - DHE-RSA-CHACHA20-POLY1305  - ECDHE-ECDSA-AES128-SHA256  - ECDHE-RSA-AES128-SHA256  - ECDHE-ECDSA-AES128-SHA  - ECDHE-RSA-AES128-SHA  - ECDHE-ECDSA-AES256-SHA384  - ECDHE-RSA-AES256-SHA384  - ECDHE-ECDSA-AES256-SHA  - ECDHE-RSA-AES256-SHA  - DHE-RSA-AES128-SHA256  - DHE-RSA-AES256-SHA256  - AES128-GCM-SHA256  - AES256-GCM-SHA384  - AES128-SHA256  - AES256-SHA256  - AES128-SHA  - AES256-SHA  - DES-CBC3-SHA  minTLSVersion: VersionTLS10",
								MarkdownDescription: "old is a TLS security profile based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Old_backward_compatibility  and looks like this (yaml):  ciphers:  - TLS_AES_128_GCM_SHA256  - TLS_AES_256_GCM_SHA384  - TLS_CHACHA20_POLY1305_SHA256  - ECDHE-ECDSA-AES128-GCM-SHA256  - ECDHE-RSA-AES128-GCM-SHA256  - ECDHE-ECDSA-AES256-GCM-SHA384  - ECDHE-RSA-AES256-GCM-SHA384  - ECDHE-ECDSA-CHACHA20-POLY1305  - ECDHE-RSA-CHACHA20-POLY1305  - DHE-RSA-AES128-GCM-SHA256  - DHE-RSA-AES256-GCM-SHA384  - DHE-RSA-CHACHA20-POLY1305  - ECDHE-ECDSA-AES128-SHA256  - ECDHE-RSA-AES128-SHA256  - ECDHE-ECDSA-AES128-SHA  - ECDHE-RSA-AES128-SHA  - ECDHE-ECDSA-AES256-SHA384  - ECDHE-RSA-AES256-SHA384  - ECDHE-ECDSA-AES256-SHA  - ECDHE-RSA-AES256-SHA  - DHE-RSA-AES128-SHA256  - DHE-RSA-AES256-SHA256  - AES128-GCM-SHA256  - AES256-GCM-SHA384  - AES128-SHA256  - AES256-SHA256  - AES128-SHA  - AES256-SHA  - DES-CBC3-SHA  minTLSVersion: VersionTLS10",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "type is one of Old, Intermediate, Modern or Custom. Custom provides the ability to specify individual TLS security profile parameters. Old, Intermediate and Modern are TLS security profiles based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations  The profiles are intent based, so they may change over time as new ciphers are developed and existing ciphers are found to be insecure.  Depending on precisely which ciphers are available to a process, the list may be reduced.  Note that the Modern profile is currently not supported because it is not yet well adopted by common software libraries.",
								MarkdownDescription: "type is one of Old, Intermediate, Modern or Custom. Custom provides the ability to specify individual TLS security profile parameters. Old, Intermediate and Modern are TLS security profiles based on:  https://wiki.mozilla.org/Security/Server_Side_TLS#Recommended_configurations  The profiles are intent based, so they may change over time as new ciphers are developed and existing ciphers are found to be insecure.  Depending on precisely which ciphers are available to a process, the list may be reduced.  Note that the Modern profile is currently not supported because it is not yet well adopted by common software libraries.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Old", "Intermediate", "Modern", "Custom"),
								},
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

func (r *MachineconfigurationOpenshiftIoKubeletConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_machineconfiguration_openshift_io_kubelet_config_v1_manifest")

	var model MachineconfigurationOpenshiftIoKubeletConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("machineconfiguration.openshift.io/v1")
	model.Kind = pointer.String("KubeletConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
