/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type SecretsCrossplaneIoStoreConfigV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SecretsCrossplaneIoStoreConfigV1Alpha1Resource)(nil)
)

type SecretsCrossplaneIoStoreConfigV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SecretsCrossplaneIoStoreConfigV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		DefaultScope *string `tfsdk:"default_scope" yaml:"defaultScope,omitempty"`

		Kubernetes *struct {
			Auth *struct {
				Env *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"env" yaml:"env,omitempty"`

				Fs *struct {
					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"fs" yaml:"fs,omitempty"`

				SecretRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

				Source *string `tfsdk:"source" yaml:"source,omitempty"`
			} `tfsdk:"auth" yaml:"auth,omitempty"`
		} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`

		Vault *struct {
			Auth *struct {
				Method *string `tfsdk:"method" yaml:"method,omitempty"`

				Token *struct {
					Env *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

					Fs *struct {
						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"fs" yaml:"fs,omitempty"`

					SecretRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

					Source *string `tfsdk:"source" yaml:"source,omitempty"`
				} `tfsdk:"token" yaml:"token,omitempty"`
			} `tfsdk:"auth" yaml:"auth,omitempty"`

			CaBundle *struct {
				Env *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"env" yaml:"env,omitempty"`

				Fs *struct {
					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"fs" yaml:"fs,omitempty"`

				SecretRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

				Source *string `tfsdk:"source" yaml:"source,omitempty"`
			} `tfsdk:"ca_bundle" yaml:"caBundle,omitempty"`

			MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

			Server *string `tfsdk:"server" yaml:"server,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"vault" yaml:"vault,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSecretsCrossplaneIoStoreConfigV1Alpha1Resource() resource.Resource {
	return &SecretsCrossplaneIoStoreConfigV1Alpha1Resource{}
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secrets_crossplane_io_store_config_v1alpha1"
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "A StoreConfig configures how Crossplane controllers should store connection details.",
		MarkdownDescription: "A StoreConfig configures how Crossplane controllers should store connection details.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "A StoreConfigSpec defines the desired state of a StoreConfig.",
				MarkdownDescription: "A StoreConfigSpec defines the desired state of a StoreConfig.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"default_scope": {
						Description:         "DefaultScope used for scoping secrets for 'cluster-scoped' resources. If store type is 'Kubernetes', this would mean the default namespace to store connection secrets for cluster scoped resources. In case of 'Vault', this would be used as the default parent path. Typically, should be set as Crossplane installation namespace.",
						MarkdownDescription: "DefaultScope used for scoping secrets for 'cluster-scoped' resources. If store type is 'Kubernetes', this would mean the default namespace to store connection secrets for cluster scoped resources. In case of 'Vault', this would be used as the default parent path. Typically, should be set as Crossplane installation namespace.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"kubernetes": {
						Description:         "Kubernetes configures a Kubernetes secret store. If the 'type' is 'Kubernetes' but no config provided, in cluster config will be used.",
						MarkdownDescription: "Kubernetes configures a Kubernetes secret store. If the 'type' is 'Kubernetes' but no config provided, in cluster config will be used.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"auth": {
								Description:         "Credentials used to connect to the Kubernetes API.",
								MarkdownDescription: "Credentials used to connect to the Kubernetes API.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"env": {
										Description:         "Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.",
										MarkdownDescription: "Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the name of an environment variable.",
												MarkdownDescription: "Name is the name of an environment variable.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fs": {
										Description:         "Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.",
										MarkdownDescription: "Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
												Description:         "Path is a filesystem path.",
												MarkdownDescription: "Path is a filesystem path.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.",
										MarkdownDescription: "A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the secret.",
												MarkdownDescription: "Name of the secret.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace of the secret.",
												MarkdownDescription: "Namespace of the secret.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source": {
										Description:         "Source of the credentials.",
										MarkdownDescription: "Source of the credentials.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("None", "Secret", "Environment", "Filesystem"),
										},
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": {
						Description:         "Type configures which secret store to be used. Only the configuration block for this store will be used and others will be ignored if provided. Default is Kubernetes.",
						MarkdownDescription: "Type configures which secret store to be used. Only the configuration block for this store will be used and others will be ignored if provided. Default is Kubernetes.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vault": {
						Description:         "Vault configures a Vault secret store.",
						MarkdownDescription: "Vault configures a Vault secret store.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"auth": {
								Description:         "Auth configures an authentication method for Vault.",
								MarkdownDescription: "Auth configures an authentication method for Vault.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"method": {
										Description:         "Method configures which auth method will be used.",
										MarkdownDescription: "Method configures which auth method will be used.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"token": {
										Description:         "Token configures Token Auth for Vault.",
										MarkdownDescription: "Token configures Token Auth for Vault.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.",
												MarkdownDescription: "Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of an environment variable.",
														MarkdownDescription: "Name is the name of an environment variable.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fs": {
												Description:         "Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.",
												MarkdownDescription: "Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "Path is a filesystem path.",
														MarkdownDescription: "Path is a filesystem path.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.",
												MarkdownDescription: "A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the secret.",
														MarkdownDescription: "Name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace of the secret.",
														MarkdownDescription: "Namespace of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source": {
												Description:         "Source of the credentials.",
												MarkdownDescription: "Source of the credentials.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("None", "Secret", "Environment", "Filesystem"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"ca_bundle": {
								Description:         "CABundle configures CA bundle for Vault Server.",
								MarkdownDescription: "CABundle configures CA bundle for Vault Server.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"env": {
										Description:         "Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.",
										MarkdownDescription: "Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the name of an environment variable.",
												MarkdownDescription: "Name is the name of an environment variable.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fs": {
										Description:         "Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.",
										MarkdownDescription: "Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
												Description:         "Path is a filesystem path.",
												MarkdownDescription: "Path is a filesystem path.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.",
										MarkdownDescription: "A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the secret.",
												MarkdownDescription: "Name of the secret.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace of the secret.",
												MarkdownDescription: "Namespace of the secret.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source": {
										Description:         "Source of the credentials.",
										MarkdownDescription: "Source of the credentials.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("None", "Secret", "Environment", "Filesystem"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount_path": {
								Description:         "MountPath is the mount path of the KV secrets engine.",
								MarkdownDescription: "MountPath is the mount path of the KV secrets engine.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"server": {
								Description:         "Server is the url of the Vault server, e.g. 'https://vault.acme.org'",
								MarkdownDescription: "Server is the url of the Vault server, e.g. 'https://vault.acme.org'",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"version": {
								Description:         "Version of the KV Secrets engine of Vault. https://www.vaultproject.io/docs/secrets/kv",
								MarkdownDescription: "Version of the KV Secrets engine of Vault. https://www.vaultproject.io/docs/secrets/kv",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_secrets_crossplane_io_store_config_v1alpha1")

	var state SecretsCrossplaneIoStoreConfigV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SecretsCrossplaneIoStoreConfigV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("secrets.crossplane.io/v1alpha1")
	goModel.Kind = utilities.Ptr("StoreConfig")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_crossplane_io_store_config_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_secrets_crossplane_io_store_config_v1alpha1")

	var state SecretsCrossplaneIoStoreConfigV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SecretsCrossplaneIoStoreConfigV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("secrets.crossplane.io/v1alpha1")
	goModel.Kind = utilities.Ptr("StoreConfig")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_secrets_crossplane_io_store_config_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
