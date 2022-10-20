/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type CamelApacheOrgIntegrationKitV1Resource struct{}

var (
	_ resource.Resource = (*CamelApacheOrgIntegrationKitV1Resource)(nil)
)

type CamelApacheOrgIntegrationKitV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CamelApacheOrgIntegrationKitV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Configuration *[]struct {
			ResourceKey *string `tfsdk:"resource_key" yaml:"resourceKey,omitempty"`

			ResourceMountPoint *string `tfsdk:"resource_mount_point" yaml:"resourceMountPoint,omitempty"`

			ResourceType *string `tfsdk:"resource_type" yaml:"resourceType,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"configuration" yaml:"configuration,omitempty"`

		Dependencies *[]string `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		Profile *string `tfsdk:"profile" yaml:"profile,omitempty"`

		Repositories *[]string `tfsdk:"repositories" yaml:"repositories,omitempty"`

		Traits *struct {
			Addons utilities.Dynamic `tfsdk:"addons" yaml:"addons,omitempty"`

			Builder *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Properties *[]string `tfsdk:"properties" yaml:"properties,omitempty"`

				Verbose *bool `tfsdk:"verbose" yaml:"verbose,omitempty"`
			} `tfsdk:"builder" yaml:"builder,omitempty"`

			Quarkus *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				PackageTypes *[]string `tfsdk:"package_types" yaml:"packageTypes,omitempty"`
			} `tfsdk:"quarkus" yaml:"quarkus,omitempty"`

			Registry *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"registry" yaml:"registry,omitempty"`
		} `tfsdk:"traits" yaml:"traits,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCamelApacheOrgIntegrationKitV1Resource() resource.Resource {
	return &CamelApacheOrgIntegrationKitV1Resource{}
}

func (r *CamelApacheOrgIntegrationKitV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camel_apache_org_integration_kit_v1"
}

func (r *CamelApacheOrgIntegrationKitV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "IntegrationKit defines a container image and additional configuration needed to run an 'Integration'. An 'IntegrationKit' is a generic image generally built from the requirements of an 'Integration', but agnostic to it, in order to be reused by any other 'Integration' which has the same required set of capabilities. An 'IntegrationKit' may be used for other kits as a base container layer, when the 'incremental' build option is enabled.",
		MarkdownDescription: "IntegrationKit defines a container image and additional configuration needed to run an 'Integration'. An 'IntegrationKit' is a generic image generally built from the requirements of an 'Integration', but agnostic to it, in order to be reused by any other 'Integration' which has the same required set of capabilities. An 'IntegrationKit' may be used for other kits as a base container layer, when the 'incremental' build option is enabled.",
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

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "the desired configuration",
				MarkdownDescription: "the desired configuration",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"configuration": {
						Description:         "configuration used by the kit TODO: we should deprecate in future releases in favour of mount, openapi or camel traits",
						MarkdownDescription: "configuration used by the kit TODO: we should deprecate in future releases in favour of mount, openapi or camel traits",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"resource_key": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_mount_point": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_type": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "represents the type of configuration, ie: property, configmap, secret, ...",
								MarkdownDescription: "represents the type of configuration, ie: property, configmap, secret, ...",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"value": {
								Description:         "the value to assign to the configuration (syntax may vary depending on the 'Type')",
								MarkdownDescription: "the value to assign to the configuration (syntax may vary depending on the 'Type')",

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

					"dependencies": {
						Description:         "a list of Camel dependecies used by this kit",
						MarkdownDescription: "a list of Camel dependecies used by this kit",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "the container image as identified in the container registry",
						MarkdownDescription: "the container image as identified in the container registry",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"profile": {
						Description:         "the profile which is expected by this kit",
						MarkdownDescription: "the profile which is expected by this kit",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"repositories": {
						Description:         "Maven repositories that can be used by the kit",
						MarkdownDescription: "Maven repositories that can be used by the kit",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"traits": {
						Description:         "traits that the kit will execute",
						MarkdownDescription: "traits that the kit will execute",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"addons": {
								Description:         "The collection of addon trait configurations",
								MarkdownDescription: "The collection of addon trait configurations",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"builder": {
								Description:         "The builder trait is internally used to determine the best strategy to build and configure IntegrationKits.",
								MarkdownDescription: "The builder trait is internally used to determine the best strategy to build and configure IntegrationKits.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"properties": {
										Description:         "A list of properties to be provided to the build task",
										MarkdownDescription: "A list of properties to be provided to the build task",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verbose": {
										Description:         "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										MarkdownDescription: "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"quarkus": {
								Description:         "The Quarkus trait configures the Quarkus runtime. It's enabled by default. NOTE: Compiling to a native executable, i.e. when using 'package-type=native', is only supported for kamelets, as well as YAML and XML integrations. It also requires at least 4GiB of memory, so the Pod running the native build, that is either the operator Pod, or the build Pod (depending on the build strategy configured for the platform), must have enough memory available.",
								MarkdownDescription: "The Quarkus trait configures the Quarkus runtime. It's enabled by default. NOTE: Compiling to a native executable, i.e. when using 'package-type=native', is only supported for kamelets, as well as YAML and XML integrations. It also requires at least 4GiB of memory, so the Pod running the native build, that is either the operator Pod, or the build Pod (depending on the build strategy configured for the platform), must have enough memory available.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"package_types": {
										Description:         "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
										MarkdownDescription: "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"registry": {
								Description:         "The Registry trait sets up Maven to use the Image registry as a Maven repository.",
								MarkdownDescription: "The Registry trait sets up Maven to use the Image registry as a Maven repository.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

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

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CamelApacheOrgIntegrationKitV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_camel_apache_org_integration_kit_v1")

	var state CamelApacheOrgIntegrationKitV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgIntegrationKitV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("IntegrationKit")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CamelApacheOrgIntegrationKitV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_integration_kit_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CamelApacheOrgIntegrationKitV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_camel_apache_org_integration_kit_v1")

	var state CamelApacheOrgIntegrationKitV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgIntegrationKitV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("IntegrationKit")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CamelApacheOrgIntegrationKitV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_camel_apache_org_integration_kit_v1")
	// NO-OP: Terraform removes the state automatically for us
}
