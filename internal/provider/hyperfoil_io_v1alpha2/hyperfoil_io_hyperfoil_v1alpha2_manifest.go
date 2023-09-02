/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hyperfoil_io_v1alpha2

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
	_ datasource.DataSource = &HyperfoilIoHyperfoilV1Alpha2Manifest{}
)

func NewHyperfoilIoHyperfoilV1Alpha2Manifest() datasource.DataSource {
	return &HyperfoilIoHyperfoilV1Alpha2Manifest{}
}

type HyperfoilIoHyperfoilV1Alpha2Manifest struct{}

type HyperfoilIoHyperfoilV1Alpha2ManifestData struct {
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
		AdditionalArgs     *[]string `tfsdk:"additional_args" json:"additionalArgs,omitempty"`
		AgentDeployTimeout *int64    `tfsdk:"agent_deploy_timeout" json:"agentDeployTimeout,omitempty"`
		Auth               *struct {
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"auth" json:"auth,omitempty"`
		Image                 *string   `tfsdk:"image" json:"image,omitempty"`
		Log                   *string   `tfsdk:"log" json:"log,omitempty"`
		PersistentVolumeClaim *string   `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
		PostHooks             *[]string `tfsdk:"post_hooks" json:"postHooks,omitempty"`
		PreHooks              *[]string `tfsdk:"pre_hooks" json:"preHooks,omitempty"`
		Route                 *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Tls  *string `tfsdk:"tls" json:"tls,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"route" json:"route,omitempty"`
		SecretEnvVars *[]string `tfsdk:"secret_env_vars" json:"secretEnvVars,omitempty"`
		ServiceType   *string   `tfsdk:"service_type" json:"serviceType,omitempty"`
		TriggerUrl    *string   `tfsdk:"trigger_url" json:"triggerUrl,omitempty"`
		Version       *string   `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HyperfoilIoHyperfoilV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hyperfoil_io_hyperfoil_v1alpha2_manifest"
}

func (r *HyperfoilIoHyperfoilV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Hyperfoil is the Schema for the hyperfoils API",
		MarkdownDescription: "Hyperfoil is the Schema for the hyperfoils API",
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
				Description:         "HyperfoilSpec Configures Hyperfoil Controller and related resources.",
				MarkdownDescription: "HyperfoilSpec Configures Hyperfoil Controller and related resources.",
				Attributes: map[string]schema.Attribute{
					"additional_args": schema.ListAttribute{
						Description:         "AdditionalArgs specifies additional arguments to pass to the Hyperfoil controller.",
						MarkdownDescription: "AdditionalArgs specifies additional arguments to pass to the Hyperfoil controller.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_deploy_timeout": schema.Int64Attribute{
						Description:         "Deploy timeout for agents, in milliseconds.",
						MarkdownDescription: "Deploy timeout for agents, in milliseconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auth": schema.SingleNestedAttribute{
						Description:         "Authentication/authorization settings.",
						MarkdownDescription: "Authentication/authorization settings.",
						Attributes: map[string]schema.Attribute{
							"secret": schema.StringAttribute{
								Description:         "Optional; Name of secret used for basic authentication. Must contain key 'password'.",
								MarkdownDescription: "Optional; Name of secret used for basic authentication. Must contain key 'password'.",
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
						Description:         "Controller image. If 'version' is defined, too, the tag is replaced (or appended). Defaults to 'quay.io/hyperfoil/hyperfoil'",
						MarkdownDescription: "Controller image. If 'version' is defined, too, the tag is replaced (or appended). Defaults to 'quay.io/hyperfoil/hyperfoil'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log": schema.StringAttribute{
						Description:         "Name of the config map and optionally its entry (separated by '/': e.g myconfigmap/log4j2-superverbose.xml) storing Log4j2 configuration file. By default the Controller uses its embedded configuration.",
						MarkdownDescription: "Name of the config map and optionally its entry (separated by '/': e.g myconfigmap/log4j2-superverbose.xml) storing Log4j2 configuration file. By default the Controller uses its embedded configuration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"persistent_volume_claim": schema.StringAttribute{
						Description:         "Name of the PVC hyperfoil should mount for its workdir.",
						MarkdownDescription: "Name of the PVC hyperfoil should mount for its workdir.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"post_hooks": schema.ListAttribute{
						Description:         "Names of config maps and optionally keys (separated by '/') holding hooks that run after the run finishes.",
						MarkdownDescription: "Names of config maps and optionally keys (separated by '/') holding hooks that run after the run finishes.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pre_hooks": schema.ListAttribute{
						Description:         "Names of config maps and optionally keys (separated by '/') holding hooks that run before the run starts.",
						MarkdownDescription: "Names of config maps and optionally keys (separated by '/') holding hooks that run before the run starts.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route": schema.SingleNestedAttribute{
						Description:         "Specification of the exposed route. This setting is ignored when Openshift Routes are not available (on vanilla Kubernetes).",
						MarkdownDescription: "Specification of the exposed route. This setting is ignored when Openshift Routes are not available (on vanilla Kubernetes).",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "Host for the route leading to Controller REST endpoint. Example: hyperfoil.apps.cloud.example.com",
								MarkdownDescription: "Host for the route leading to Controller REST endpoint. Example: hyperfoil.apps.cloud.example.com",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.StringAttribute{
								Description:         "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
								MarkdownDescription: "Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
								MarkdownDescription: "Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_env_vars": schema.ListAttribute{
						Description:         "List of secrets in this namespace; each entry from those secrets will be mapped as environment variable, using the key as variable name.",
						MarkdownDescription: "List of secrets in this namespace; each entry from those secrets will be mapped as environment variable, using the key as variable name.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_type": schema.StringAttribute{
						Description:         "Type of the service being exposed. By default this is ClusterIP if Openshift Route resource is available (the route will target this service). If Openshift Routes are not available (on vanilla Kubernetes) the default is NodePort.",
						MarkdownDescription: "Type of the service being exposed. By default this is ClusterIP if Openshift Route resource is available (the route will target this service). If Openshift Routes are not available (on vanilla Kubernetes) the default is NodePort.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"trigger_url": schema.StringAttribute{
						Description:         "If this is set the controller does not start benchmark run right away after hitting /benchmark/my-benchmark/start ; instead it responds with status 301 and header Location set to concatenation of this string and 'BENCHMARK=my-benchmark&RUN_ID=xxxx'. CLI interprets that response as a request to hit CI instance on this URL, assuming that CI will trigger a new job that will eventually call /benchmark/my-benchmark/start?runId=xxxx with header 'x-trigger-job'. This is useful if the the CI has to synchronize Hyperfoil to other benchmarks that don't use this controller instance.",
						MarkdownDescription: "If this is set the controller does not start benchmark run right away after hitting /benchmark/my-benchmark/start ; instead it responds with status 301 and header Location set to concatenation of this string and 'BENCHMARK=my-benchmark&RUN_ID=xxxx'. CLI interprets that response as a request to hit CI instance on this URL, assuming that CI will trigger a new job that will eventually call /benchmark/my-benchmark/start?runId=xxxx with header 'x-trigger-job'. This is useful if the the CI has to synchronize Hyperfoil to other benchmarks that don't use this controller instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "Tag for controller image. Defaults to version matching the operator version.",
						MarkdownDescription: "Tag for controller image. Defaults to version matching the operator version.",
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
	}
}

func (r *HyperfoilIoHyperfoilV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hyperfoil_io_hyperfoil_v1alpha2_manifest")

	var model HyperfoilIoHyperfoilV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("hyperfoil.io/v1alpha2")
	model.Kind = pointer.String("Hyperfoil")

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
