/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package registry_devfile_io_v1alpha1

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
	_ datasource.DataSource = &RegistryDevfileIoDevfileRegistryV1Alpha1Manifest{}
)

func NewRegistryDevfileIoDevfileRegistryV1Alpha1Manifest() datasource.DataSource {
	return &RegistryDevfileIoDevfileRegistryV1Alpha1Manifest{}
}

type RegistryDevfileIoDevfileRegistryV1Alpha1Manifest struct{}

type RegistryDevfileIoDevfileRegistryV1Alpha1ManifestData struct {
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
		DevfileIndex *struct {
			Image           *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			MemoryLimit     *string `tfsdk:"memory_limit" json:"memoryLimit,omitempty"`
		} `tfsdk:"devfile_index" json:"devfileIndex,omitempty"`
		DevfileIndexImage *string `tfsdk:"devfile_index_image" json:"devfileIndexImage,omitempty"`
		Headless          *bool   `tfsdk:"headless" json:"headless,omitempty"`
		K8s               *struct {
			IngressClass  *string `tfsdk:"ingress_class" json:"ingressClass,omitempty"`
			IngressDomain *string `tfsdk:"ingress_domain" json:"ingressDomain,omitempty"`
		} `tfsdk:"k8s" json:"k8s,omitempty"`
		OciRegistry *struct {
			Image           *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			MemoryLimit     *string `tfsdk:"memory_limit" json:"memoryLimit,omitempty"`
		} `tfsdk:"oci_registry" json:"ociRegistry,omitempty"`
		OciRegistryImage *string `tfsdk:"oci_registry_image" json:"ociRegistryImage,omitempty"`
		RegistryViewer   *struct {
			Image           *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			MemoryLimit     *string `tfsdk:"memory_limit" json:"memoryLimit,omitempty"`
		} `tfsdk:"registry_viewer" json:"registryViewer,omitempty"`
		RegistryViewerImage *string `tfsdk:"registry_viewer_image" json:"registryViewerImage,omitempty"`
		Storage             *struct {
			Enabled            *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			RegistryVolumeSize *string `tfsdk:"registry_volume_size" json:"registryVolumeSize,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Telemetry *struct {
			Key                    *string `tfsdk:"key" json:"key,omitempty"`
			RegistryName           *string `tfsdk:"registry_name" json:"registryName,omitempty"`
			RegistryViewerWriteKey *string `tfsdk:"registry_viewer_write_key" json:"registryViewerWriteKey,omitempty"`
		} `tfsdk:"telemetry" json:"telemetry,omitempty"`
		Tls *struct {
			Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RegistryDevfileIoDevfileRegistryV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_registry_devfile_io_devfile_registry_v1alpha1_manifest"
}

func (r *RegistryDevfileIoDevfileRegistryV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DevfileRegistry is a custom resource allows you to create and manage your own index server and registry viewer. In order to be added, the Devfile Registry must be reachable, supports the Devfile v2.0 spec and above, and is not using the default namespace.",
		MarkdownDescription: "DevfileRegistry is a custom resource allows you to create and manage your own index server and registry viewer. In order to be added, the Devfile Registry must be reachable, supports the Devfile v2.0 spec and above, and is not using the default namespace.",
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
				Description:         "DevfileRegistrySpec defines the desired state of DevfileRegistry",
				MarkdownDescription: "DevfileRegistrySpec defines the desired state of DevfileRegistry",
				Attributes: map[string]schema.Attribute{
					"devfile_index": schema.SingleNestedAttribute{
						Description:         "Sets the devfile index container spec to be deployed on the Devfile Registry",
						MarkdownDescription: "Sets the devfile index container spec to be deployed on the Devfile Registry",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "Sets the container image",
								MarkdownDescription: "Sets the container image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "Sets the image pull policy for the container",
								MarkdownDescription: "Sets the image pull policy for the container",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"memory_limit": schema.StringAttribute{
								Description:         "Sets the memory limit for the container",
								MarkdownDescription: "Sets the memory limit for the container",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"devfile_index_image": schema.StringAttribute{
						Description:         "Sets the container image containing devfile stacks to be deployed on the Devfile Registry",
						MarkdownDescription: "Sets the container image containing devfile stacks to be deployed on the Devfile Registry",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"headless": schema.BoolAttribute{
						Description:         "Sets the registry server deployment to run under headless mode",
						MarkdownDescription: "Sets the registry server deployment to run under headless mode",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"k8s": schema.SingleNestedAttribute{
						Description:         "DevfileRegistrySpecK8sOnly defines the desired state of the kubernetes-only fields of the DevfileRegistry",
						MarkdownDescription: "DevfileRegistrySpecK8sOnly defines the desired state of the kubernetes-only fields of the DevfileRegistry",
						Attributes: map[string]schema.Attribute{
							"ingress_class": schema.StringAttribute{
								Description:         "Ingress class for a Kubernetes cluster. Defaults to nginx.",
								MarkdownDescription: "Ingress class for a Kubernetes cluster. Defaults to nginx.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress_domain": schema.StringAttribute{
								Description:         "Ingress domain for a Kubernetes cluster. This MUST be explicitly specified on Kubernetes. There are no defaults",
								MarkdownDescription: "Ingress domain for a Kubernetes cluster. This MUST be explicitly specified on Kubernetes. There are no defaults",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"oci_registry": schema.SingleNestedAttribute{
						Description:         "Sets the OCI registry container spec to be deployed on the Devfile Registry",
						MarkdownDescription: "Sets the OCI registry container spec to be deployed on the Devfile Registry",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "Sets the container image",
								MarkdownDescription: "Sets the container image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "Sets the image pull policy for the container",
								MarkdownDescription: "Sets the image pull policy for the container",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"memory_limit": schema.StringAttribute{
								Description:         "Sets the memory limit for the container",
								MarkdownDescription: "Sets the memory limit for the container",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"oci_registry_image": schema.StringAttribute{
						Description:         "Overrides the container image used for the OCI registry. Recommended to leave blank and default to the image specified by the operator.",
						MarkdownDescription: "Overrides the container image used for the OCI registry. Recommended to leave blank and default to the image specified by the operator.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"registry_viewer": schema.SingleNestedAttribute{
						Description:         "Sets the registry viewer container spec to be deployed on the Devfile Registry",
						MarkdownDescription: "Sets the registry viewer container spec to be deployed on the Devfile Registry",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "Sets the container image",
								MarkdownDescription: "Sets the container image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "Sets the image pull policy for the container",
								MarkdownDescription: "Sets the image pull policy for the container",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"memory_limit": schema.StringAttribute{
								Description:         "Sets the memory limit for the container",
								MarkdownDescription: "Sets the memory limit for the container",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"registry_viewer_image": schema.StringAttribute{
						Description:         "Overrides the container image used for the registry viewer.",
						MarkdownDescription: "Overrides the container image used for the registry viewer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "DevfileRegistrySpecStorage defines the desired state of the storage for the DevfileRegistry",
						MarkdownDescription: "DevfileRegistrySpecStorage defines the desired state of the storage for the DevfileRegistry",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Instructs the operator to deploy the DevfileRegistry with persistent storage Disabled by default.",
								MarkdownDescription: "Instructs the operator to deploy the DevfileRegistry with persistent storage Disabled by default.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registry_volume_size": schema.StringAttribute{
								Description:         "Configures the size of the devfile registry's persistent volume, if enabled. Defaults to 1Gi.",
								MarkdownDescription: "Configures the size of the devfile registry's persistent volume, if enabled. Defaults to 1Gi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"telemetry": schema.SingleNestedAttribute{
						Description:         "Telemetry defines the desired state for telemetry in the DevfileRegistry",
						MarkdownDescription: "Telemetry defines the desired state for telemetry in the DevfileRegistry",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Specify a telemetry key to allow devfile specific data to be sent to a client's own Segment analytics source. If the write key is specified then telemetry will be enabled",
								MarkdownDescription: "Specify a telemetry key to allow devfile specific data to be sent to a client's own Segment analytics source. If the write key is specified then telemetry will be enabled",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registry_name": schema.StringAttribute{
								Description:         "The registry name (can be any string) that is used as identifier for devfile telemetry.",
								MarkdownDescription: "The registry name (can be any string) that is used as identifier for devfile telemetry.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registry_viewer_write_key": schema.StringAttribute{
								Description:         "Specify a telemetry write key for the registry viewer component to allow data to be sent to a client's own Segment analytics source. If the write key is specified then telemetry for the registry viewer component will be enabled",
								MarkdownDescription: "Specify a telemetry write key for the registry viewer component to allow data to be sent to a client's own Segment analytics source. If the write key is specified then telemetry for the registry viewer component will be enabled",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "DevfileRegistrySpecTLS defines the desired state for TLS in the DevfileRegistry",
						MarkdownDescription: "DevfileRegistrySpecTLS defines the desired state for TLS in the DevfileRegistry",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Instructs the operator to deploy the DevfileRegistry with TLS enabled. Enabled by default. Disabling is only recommended for development or test.",
								MarkdownDescription: "Instructs the operator to deploy the DevfileRegistry with TLS enabled. Enabled by default. Disabling is only recommended for development or test.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "Name of an optional, pre-existing TLS secret to use for TLS termination on ingress/route resources.",
								MarkdownDescription: "Name of an optional, pre-existing TLS secret to use for TLS termination on ingress/route resources.",
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
	}
}

func (r *RegistryDevfileIoDevfileRegistryV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_registry_devfile_io_devfile_registry_v1alpha1_manifest")

	var model RegistryDevfileIoDevfileRegistryV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("registry.devfile.io/v1alpha1")
	model.Kind = pointer.String("DevfileRegistry")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
