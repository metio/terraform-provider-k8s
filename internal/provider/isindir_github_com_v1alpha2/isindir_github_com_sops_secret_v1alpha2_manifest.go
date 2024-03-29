/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package isindir_github_com_v1alpha2

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
	_ datasource.DataSource = &IsindirGithubComSopsSecretV1Alpha2Manifest{}
)

func NewIsindirGithubComSopsSecretV1Alpha2Manifest() datasource.DataSource {
	return &IsindirGithubComSopsSecretV1Alpha2Manifest{}
}

type IsindirGithubComSopsSecretV1Alpha2Manifest struct{}

type IsindirGithubComSopsSecretV1Alpha2ManifestData struct {
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

	Sops *struct {
		Age *[]struct {
			Enc       *string `tfsdk:"enc" json:"enc,omitempty"`
			Recipient *string `tfsdk:"recipient" json:"recipient,omitempty"`
		} `tfsdk:"age" json:"age,omitempty"`
		Azure_kv *[]struct {
			Created_at *string `tfsdk:"created_at" json:"created_at,omitempty"`
			Enc        *string `tfsdk:"enc" json:"enc,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Vault_url  *string `tfsdk:"vault_url" json:"vault_url,omitempty"`
			Version    *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"azure_kv" json:"azure_kv,omitempty"`
		Encrypted_regex  *string `tfsdk:"encrypted_regex" json:"encrypted_regex,omitempty"`
		Encrypted_suffix *string `tfsdk:"encrypted_suffix" json:"encrypted_suffix,omitempty"`
		Gcp_kms          *[]struct {
			Created_at  *string `tfsdk:"created_at" json:"created_at,omitempty"`
			Enc         *string `tfsdk:"enc" json:"enc,omitempty"`
			Resource_id *string `tfsdk:"resource_id" json:"resource_id,omitempty"`
		} `tfsdk:"gcp_kms" json:"gcp_kms,omitempty"`
		Hc_vault *[]struct {
			Created_at    *string `tfsdk:"created_at" json:"created_at,omitempty"`
			Enc           *string `tfsdk:"enc" json:"enc,omitempty"`
			Engine_path   *string `tfsdk:"engine_path" json:"engine_path,omitempty"`
			Key_name      *string `tfsdk:"key_name" json:"key_name,omitempty"`
			Vault_address *string `tfsdk:"vault_address" json:"vault_address,omitempty"`
		} `tfsdk:"hc_vault" json:"hc_vault,omitempty"`
		Kms *[]struct {
			Arn         *string `tfsdk:"arn" json:"arn,omitempty"`
			Aws_profile *string `tfsdk:"aws_profile" json:"aws_profile,omitempty"`
			Created_at  *string `tfsdk:"created_at" json:"created_at,omitempty"`
			Enc         *string `tfsdk:"enc" json:"enc,omitempty"`
			Role        *string `tfsdk:"role" json:"role,omitempty"`
		} `tfsdk:"kms" json:"kms,omitempty"`
		Lastmodified *string `tfsdk:"lastmodified" json:"lastmodified,omitempty"`
		Mac          *string `tfsdk:"mac" json:"mac,omitempty"`
		Pgp          *[]struct {
			Created_at *string `tfsdk:"created_at" json:"created_at,omitempty"`
			Enc        *string `tfsdk:"enc" json:"enc,omitempty"`
			Fp         *string `tfsdk:"fp" json:"fp,omitempty"`
		} `tfsdk:"pgp" json:"pgp,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"sops" json:"sops,omitempty"`
	Spec *struct {
		SecretTemplates *[]struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Data        *map[string]string `tfsdk:"data" json:"data,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
			Type        *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"secret_templates" json:"secretTemplates,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *IsindirGithubComSopsSecretV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_isindir_github_com_sops_secret_v1alpha2_manifest"
}

func (r *IsindirGithubComSopsSecretV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SopsSecret is the Schema for the sopssecrets API",
		MarkdownDescription: "SopsSecret is the Schema for the sopssecrets API",
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

			"sops": schema.SingleNestedAttribute{
				Description:         "SopsSecret metadata",
				MarkdownDescription: "SopsSecret metadata",
				Attributes: map[string]schema.Attribute{
					"age": schema.ListNestedAttribute{
						Description:         "Age configuration",
						MarkdownDescription: "Age configuration",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"enc": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"recipient": schema.StringAttribute{
									Description:         "Recipient which private key can be used for decription",
									MarkdownDescription: "Recipient which private key can be used for decription",
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

					"azure_kv": schema.ListNestedAttribute{
						Description:         "Azure KMS configuration",
						MarkdownDescription: "Azure KMS configuration",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created_at": schema.StringAttribute{
									Description:         "Object creation date",
									MarkdownDescription: "Object creation date",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enc": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"vault_url": schema.StringAttribute{
									Description:         "Azure KMS vault URL",
									MarkdownDescription: "Azure KMS vault URL",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"version": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"encrypted_regex": schema.StringAttribute{
						Description:         "Regex used to encrypt SopsSecret resourceThis opstion should be used with more care, as it can make resource unapplicable to the cluster.",
						MarkdownDescription: "Regex used to encrypt SopsSecret resourceThis opstion should be used with more care, as it can make resource unapplicable to the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encrypted_suffix": schema.StringAttribute{
						Description:         "Suffix used to encrypt SopsSecret resource",
						MarkdownDescription: "Suffix used to encrypt SopsSecret resource",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"gcp_kms": schema.ListNestedAttribute{
						Description:         "Gcp KMS configuration",
						MarkdownDescription: "Gcp KMS configuration",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created_at": schema.StringAttribute{
									Description:         "Object creation date",
									MarkdownDescription: "Object creation date",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enc": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resource_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"hc_vault": schema.ListNestedAttribute{
						Description:         "Hashicorp Vault KMS configurarion",
						MarkdownDescription: "Hashicorp Vault KMS configurarion",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created_at": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enc": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"engine_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"vault_address": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"kms": schema.ListNestedAttribute{
						Description:         "Aws KMS configuration",
						MarkdownDescription: "Aws KMS configuration",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"arn": schema.StringAttribute{
									Description:         "Arn - KMS key ARN to use",
									MarkdownDescription: "Arn - KMS key ARN to use",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"aws_profile": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"created_at": schema.StringAttribute{
									Description:         "Object creation date",
									MarkdownDescription: "Object creation date",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enc": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"role": schema.StringAttribute{
									Description:         "AWS Iam Role",
									MarkdownDescription: "AWS Iam Role",
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

					"lastmodified": schema.StringAttribute{
						Description:         "LastModified date when SopsSecret was last modified",
						MarkdownDescription: "LastModified date when SopsSecret was last modified",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mac": schema.StringAttribute{
						Description:         "Mac - sops setting",
						MarkdownDescription: "Mac - sops setting",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pgp": schema.ListNestedAttribute{
						Description:         "PGP configuration",
						MarkdownDescription: "PGP configuration",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"created_at": schema.StringAttribute{
									Description:         "Object creation date",
									MarkdownDescription: "Object creation date",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enc": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"fp": schema.StringAttribute{
									Description:         "PGP FingerPrint of the key which can be used for decryption",
									MarkdownDescription: "PGP FingerPrint of the key which can be used for decryption",
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

					"version": schema.StringAttribute{
						Description:         "Version of the sops tool used to encrypt SopsSecret",
						MarkdownDescription: "Version of the sops tool used to encrypt SopsSecret",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "SopsSecret Spec definition",
				MarkdownDescription: "SopsSecret Spec definition",
				Attributes: map[string]schema.Attribute{
					"secret_templates": schema.ListNestedAttribute{
						Description:         "Secrets template is a list of definitions to create Kubernetes Secrets",
						MarkdownDescription: "Secrets template is a list of definitions to create Kubernetes Secrets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"annotations": schema.MapAttribute{
									Description:         "Annotations to apply to Kubernetes secret",
									MarkdownDescription: "Annotations to apply to Kubernetes secret",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"data": schema.MapAttribute{
									Description:         "Data map to use in Kubernetes secret (equivalent to Kubernetes Secret object stringData, please see for moreinformation: https://kubernetes.io/docs/concepts/configuration/secret/#overview-of-secrets)",
									MarkdownDescription: "Data map to use in Kubernetes secret (equivalent to Kubernetes Secret object stringData, please see for moreinformation: https://kubernetes.io/docs/concepts/configuration/secret/#overview-of-secrets)",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels to apply to Kubernetes secret",
									MarkdownDescription: "Labels to apply to Kubernetes secret",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the Kubernetes secret to create",
									MarkdownDescription: "Name of the Kubernetes secret to create",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Kubernetes secret type. Default: Opauqe. Possible values: Opauqe,kubernetes.io/service-account-token, kubernetes.io/dockercfg,kubernetes.io/dockerconfigjson, kubernetes.io/basic-auth,kubernetes.io/ssh-auth, kubernetes.io/tls, bootstrap.kubernetes.io/token",
									MarkdownDescription: "Kubernetes secret type. Default: Opauqe. Possible values: Opauqe,kubernetes.io/service-account-token, kubernetes.io/dockercfg,kubernetes.io/dockerconfigjson, kubernetes.io/basic-auth,kubernetes.io/ssh-auth, kubernetes.io/tls, bootstrap.kubernetes.io/token",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
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

func (r *IsindirGithubComSopsSecretV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_isindir_github_com_sops_secret_v1alpha2_manifest")

	var model IsindirGithubComSopsSecretV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("isindir.github.com/v1alpha2")
	model.Kind = pointer.String("SopsSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
