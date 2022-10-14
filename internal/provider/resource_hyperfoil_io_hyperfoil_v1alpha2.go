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

type HyperfoilIoHyperfoilV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*HyperfoilIoHyperfoilV1Alpha2Resource)(nil)
)

type HyperfoilIoHyperfoilV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HyperfoilIoHyperfoilV1Alpha2GoModel struct {
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
		AdditionalArgs *[]string `tfsdk:"additional_args" yaml:"additionalArgs,omitempty"`

		AgentDeployTimeout *int64 `tfsdk:"agent_deploy_timeout" yaml:"agentDeployTimeout,omitempty"`

		Auth *struct {
			Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
		} `tfsdk:"auth" yaml:"auth,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		Log *string `tfsdk:"log" yaml:"log,omitempty"`

		PersistentVolumeClaim *string `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

		PostHooks *[]string `tfsdk:"post_hooks" yaml:"postHooks,omitempty"`

		PreHooks *[]string `tfsdk:"pre_hooks" yaml:"preHooks,omitempty"`

		Route *struct {
			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Tls *string `tfsdk:"tls" yaml:"tls,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"route" yaml:"route,omitempty"`

		SecretEnvVars *[]string `tfsdk:"secret_env_vars" yaml:"secretEnvVars,omitempty"`

		ServiceType *string `tfsdk:"service_type" yaml:"serviceType,omitempty"`

		TriggerUrl *string `tfsdk:"trigger_url" yaml:"triggerUrl,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHyperfoilIoHyperfoilV1Alpha2Resource() resource.Resource {
	return &HyperfoilIoHyperfoilV1Alpha2Resource{}
}

func (r *HyperfoilIoHyperfoilV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hyperfoil_io_hyperfoil_v1alpha2"
}

func (r *HyperfoilIoHyperfoilV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Hyperfoil is the Schema for the hyperfoils API",
		MarkdownDescription: "Hyperfoil is the Schema for the hyperfoils API",
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
				Description:         "HyperfoilSpec Configures Hyperfoil Controller and related resources.",
				MarkdownDescription: "HyperfoilSpec Configures Hyperfoil Controller and related resources.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"additional_args": {
						Description:         "AdditionalArgs specifies additional arguments to pass to the Hyperfoil controller.",
						MarkdownDescription: "AdditionalArgs specifies additional arguments to pass to the Hyperfoil controller.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"agent_deploy_timeout": {
						Description:         "Deploy timeout for agents, in milliseconds.",
						MarkdownDescription: "Deploy timeout for agents, in milliseconds.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auth": {
						Description:         "Authentication/authorization settings.",
						MarkdownDescription: "Authentication/authorization settings.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"secret": {
								Description:         "Optional; Name of secret used for basic authentication. Must contain key 'password'.",
								MarkdownDescription: "Optional; Name of secret used for basic authentication. Must contain key 'password'.",

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

					"image": {
						Description:         "Controller image. If 'version' is defined, too, the tag is replaced (or appended). Defaults to 'quay.io/hyperfoil/hyperfoil'",
						MarkdownDescription: "Controller image. If 'version' is defined, too, the tag is replaced (or appended). Defaults to 'quay.io/hyperfoil/hyperfoil'",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log": {
						Description:         "Name of the config map and optionally its entry (separated by '/': e.g myconfigmap/log4j2-superverbose.xml) storing Log4j2 configuration file. By default the Controller uses its embedded configuration.",
						MarkdownDescription: "Name of the config map and optionally its entry (separated by '/': e.g myconfigmap/log4j2-superverbose.xml) storing Log4j2 configuration file. By default the Controller uses its embedded configuration.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"persistent_volume_claim": {
						Description:         "Name of the PVC hyperfoil should mount for its workdir.",
						MarkdownDescription: "Name of the PVC hyperfoil should mount for its workdir.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"post_hooks": {
						Description:         "Names of config maps and optionally keys (separated by '/') holding hooks that run after the run finishes.",
						MarkdownDescription: "Names of config maps and optionally keys (separated by '/') holding hooks that run after the run finishes.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pre_hooks": {
						Description:         "Names of config maps and optionally keys (separated by '/') holding hooks that run before the run starts.",
						MarkdownDescription: "Names of config maps and optionally keys (separated by '/') holding hooks that run before the run starts.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route": {
						Description:         "Specification of the exposed route. This setting is ignored when Openshift Routes are not available (on vanilla Kubernetes).",
						MarkdownDescription: "Specification of the exposed route. This setting is ignored when Openshift Routes are not available (on vanilla Kubernetes).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"host": {
								Description:         "Host for the route leading to Controller REST endpoint. Example: hyperfoil.apps.cloud.example.com",
								MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: hyperfoil.apps.cloud.example.com",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": {
								Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
								MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
								MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",

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

					"secret_env_vars": {
						Description:         "List of secrets in this namespace; each entry from those secrets will be mapped as environment variable, using the key as variable name.",
						MarkdownDescription: "List of secrets in this namespace; each entry from those secrets will be mapped as environment variable, using the key as variable name.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_type": {
						Description:         "Type of the service being exposed. By default this is ClusterIP if Openshift Route resource is available (the route will target this service). If Openshift Routes are not available (on vanilla Kubernetes) the default is NodePort.",
						MarkdownDescription: "Type of the service being exposed. By default this is ClusterIP if Openshift Route resource is available (the route will target this service). If Openshift Routes are not available (on vanilla Kubernetes) the default is NodePort.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"trigger_url": {
						Description:         "If this is set the controller does not start benchmark run right away after hitting /benchmark/my-benchmark/start ; instead it responds with status 301 and header Location set to concatenation of this string and 'BENCHMARK=my-benchmark&RUN_ID=xxxx'. CLI interprets that response as a request to hit CI instance on this URL, assuming that CI will trigger a new job that will eventually call /benchmark/my-benchmark/start?runId=xxxx with header 'x-trigger-job'. This is useful if the the CI has to synchronize Hyperfoil to other benchmarks that don't use this controller instance.",
						MarkdownDescription: "If this is set the controller does not start benchmark run right away after hitting /benchmark/my-benchmark/start ; instead it responds with status 301 and header Location set to concatenation of this string and 'BENCHMARK=my-benchmark&RUN_ID=xxxx'. CLI interprets that response as a request to hit CI instance on this URL, assuming that CI will trigger a new job that will eventually call /benchmark/my-benchmark/start?runId=xxxx with header 'x-trigger-job'. This is useful if the the CI has to synchronize Hyperfoil to other benchmarks that don't use this controller instance.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "Tag for controller image. Defaults to version matching the operator version.",
						MarkdownDescription: "Tag for controller image. Defaults to version matching the operator version.",

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
		},
	}, nil
}

func (r *HyperfoilIoHyperfoilV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hyperfoil_io_hyperfoil_v1alpha2")

	var state HyperfoilIoHyperfoilV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HyperfoilIoHyperfoilV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hyperfoil.io/v1alpha2")
	goModel.Kind = utilities.Ptr("Hyperfoil")

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

func (r *HyperfoilIoHyperfoilV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hyperfoil_io_hyperfoil_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *HyperfoilIoHyperfoilV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hyperfoil_io_hyperfoil_v1alpha2")

	var state HyperfoilIoHyperfoilV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HyperfoilIoHyperfoilV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hyperfoil.io/v1alpha2")
	goModel.Kind = utilities.Ptr("Hyperfoil")

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

func (r *HyperfoilIoHyperfoilV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hyperfoil_io_hyperfoil_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}
