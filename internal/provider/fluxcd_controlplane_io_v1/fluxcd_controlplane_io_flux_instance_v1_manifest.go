/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluxcd_controlplane_io_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &FluxcdControlplaneIoFluxInstanceV1Manifest{}
)

func NewFluxcdControlplaneIoFluxInstanceV1Manifest() datasource.DataSource {
	return &FluxcdControlplaneIoFluxInstanceV1Manifest{}
}

type FluxcdControlplaneIoFluxInstanceV1Manifest struct{}

type FluxcdControlplaneIoFluxInstanceV1ManifestData struct {
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
		Cluster *struct {
			Domain                      *string `tfsdk:"domain" json:"domain,omitempty"`
			Multitenant                 *bool   `tfsdk:"multitenant" json:"multitenant,omitempty"`
			NetworkPolicy               *bool   `tfsdk:"network_policy" json:"networkPolicy,omitempty"`
			TenantDefaultServiceAccount *string `tfsdk:"tenant_default_service_account" json:"tenantDefaultServiceAccount,omitempty"`
			Type                        *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"cluster" json:"cluster,omitempty"`
		CommonMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"common_metadata" json:"commonMetadata,omitempty"`
		Components   *[]string `tfsdk:"components" json:"components,omitempty"`
		Distribution *struct {
			Artifact           *string `tfsdk:"artifact" json:"artifact,omitempty"`
			ArtifactPullSecret *string `tfsdk:"artifact_pull_secret" json:"artifactPullSecret,omitempty"`
			ImagePullSecret    *string `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
			Registry           *string `tfsdk:"registry" json:"registry,omitempty"`
			Version            *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"distribution" json:"distribution,omitempty"`
		Kustomize *struct {
			Patches *[]struct {
				Patch  *string `tfsdk:"patch" json:"patch,omitempty"`
				Target *struct {
					AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
					Group              *string `tfsdk:"group" json:"group,omitempty"`
					Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
					LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					Name               *string `tfsdk:"name" json:"name,omitempty"`
					Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Version            *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"patches" json:"patches,omitempty"`
		} `tfsdk:"kustomize" json:"kustomize,omitempty"`
		MigrateResources *bool `tfsdk:"migrate_resources" json:"migrateResources,omitempty"`
		Sharding         *struct {
			Key    *string   `tfsdk:"key" json:"key,omitempty"`
			Shards *[]string `tfsdk:"shards" json:"shards,omitempty"`
		} `tfsdk:"sharding" json:"sharding,omitempty"`
		Storage *struct {
			Class *string `tfsdk:"class" json:"class,omitempty"`
			Size  *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Sync *struct {
			Interval   *string `tfsdk:"interval" json:"interval,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Path       *string `tfsdk:"path" json:"path,omitempty"`
			Provider   *string `tfsdk:"provider" json:"provider,omitempty"`
			PullSecret *string `tfsdk:"pull_secret" json:"pullSecret,omitempty"`
			Ref        *string `tfsdk:"ref" json:"ref,omitempty"`
			Url        *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"sync" json:"sync,omitempty"`
		Wait *bool `tfsdk:"wait" json:"wait,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluxcdControlplaneIoFluxInstanceV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluxcd_controlplane_io_flux_instance_v1_manifest"
}

func (r *FluxcdControlplaneIoFluxInstanceV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FluxInstance is the Schema for the fluxinstances API",
		MarkdownDescription: "FluxInstance is the Schema for the fluxinstances API",
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
				Description:         "FluxInstanceSpec defines the desired state of FluxInstance",
				MarkdownDescription: "FluxInstanceSpec defines the desired state of FluxInstance",
				Attributes: map[string]schema.Attribute{
					"cluster": schema.SingleNestedAttribute{
						Description:         "Cluster holds the specification of the Kubernetes cluster.",
						MarkdownDescription: "Cluster holds the specification of the Kubernetes cluster.",
						Attributes: map[string]schema.Attribute{
							"domain": schema.StringAttribute{
								Description:         "Domain is the cluster domain used for generating the FQDN of services. Defaults to 'cluster.local'.",
								MarkdownDescription: "Domain is the cluster domain used for generating the FQDN of services. Defaults to 'cluster.local'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"multitenant": schema.BoolAttribute{
								Description:         "Multitenant enables the multitenancy lockdown.",
								MarkdownDescription: "Multitenant enables the multitenancy lockdown.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"network_policy": schema.BoolAttribute{
								Description:         "NetworkPolicy restricts network access to the current namespace. Defaults to true.",
								MarkdownDescription: "NetworkPolicy restricts network access to the current namespace. Defaults to true.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tenant_default_service_account": schema.StringAttribute{
								Description:         "TenantDefaultServiceAccount is the name of the service account to use as default when the multitenant lockdown is enabled. Defaults to the 'default' service account from the tenant namespace.",
								MarkdownDescription: "TenantDefaultServiceAccount is the name of the service account to use as default when the multitenant lockdown is enabled. Defaults to the 'default' service account from the tenant namespace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type specifies the distro of the Kubernetes cluster. Defaults to 'kubernetes'.",
								MarkdownDescription: "Type specifies the distro of the Kubernetes cluster. Defaults to 'kubernetes'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("kubernetes", "openshift", "aws", "azure", "gcp"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"common_metadata": schema.SingleNestedAttribute{
						Description:         "CommonMetadata specifies the common labels and annotations that are applied to all resources. Any existing label or annotation will be overridden if its key matches a common one.",
						MarkdownDescription: "CommonMetadata specifies the common labels and annotations that are applied to all resources. Any existing label or annotation will be overridden if its key matches a common one.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to be added to the object's metadata.",
								MarkdownDescription: "Annotations to be added to the object's metadata.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels to be added to the object's metadata.",
								MarkdownDescription: "Labels to be added to the object's metadata.",
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

					"components": schema.ListAttribute{
						Description:         "Components is the list of controllers to install. Defaults to all controllers.",
						MarkdownDescription: "Components is the list of controllers to install. Defaults to all controllers.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"distribution": schema.SingleNestedAttribute{
						Description:         "Distribution specifies the version and container registry to pull images from.",
						MarkdownDescription: "Distribution specifies the version and container registry to pull images from.",
						Attributes: map[string]schema.Attribute{
							"artifact": schema.StringAttribute{
								Description:         "Artifact is the URL to the OCI artifact containing the latest Kubernetes manifests for the distribution, e.g. 'oci://ghcr.io/controlplaneio-fluxcd/flux-operator-manifests:latest'.",
								MarkdownDescription: "Artifact is the URL to the OCI artifact containing the latest Kubernetes manifests for the distribution, e.g. 'oci://ghcr.io/controlplaneio-fluxcd/flux-operator-manifests:latest'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^oci://.*$`), ""),
								},
							},

							"artifact_pull_secret": schema.StringAttribute{
								Description:         "ArtifactPullSecret is the name of the Kubernetes secret to use for pulling the Kubernetes manifests for the distribution specified in the Artifact field.",
								MarkdownDescription: "ArtifactPullSecret is the name of the Kubernetes secret to use for pulling the Kubernetes manifests for the distribution specified in the Artifact field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_secret": schema.StringAttribute{
								Description:         "ImagePullSecret is the name of the Kubernetes secret to use for pulling images.",
								MarkdownDescription: "ImagePullSecret is the name of the Kubernetes secret to use for pulling images.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registry": schema.StringAttribute{
								Description:         "Registry address to pull the distribution images from e.g. 'ghcr.io/fluxcd'.",
								MarkdownDescription: "Registry address to pull the distribution images from e.g. 'ghcr.io/fluxcd'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version semver expression e.g. '2.x', '2.3.x'.",
								MarkdownDescription: "Version semver expression e.g. '2.x', '2.3.x'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"kustomize": schema.SingleNestedAttribute{
						Description:         "Kustomize holds a set of patches that can be applied to the Flux installation, to customize the way Flux operates.",
						MarkdownDescription: "Kustomize holds a set of patches that can be applied to the Flux installation, to customize the way Flux operates.",
						Attributes: map[string]schema.Attribute{
							"patches": schema.ListNestedAttribute{
								Description:         "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",
								MarkdownDescription: "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"patch": schema.StringAttribute{
											Description:         "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",
											MarkdownDescription: "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target": schema.SingleNestedAttribute{
											Description:         "Target points to the resources that the patch document should be applied to.",
											MarkdownDescription: "Target points to the resources that the patch document should be applied to.",
											Attributes: map[string]schema.Attribute{
												"annotation_selector": schema.StringAttribute{
													Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
													MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"group": schema.StringAttribute{
													Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selector": schema.StringAttribute{
													Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
													MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name to match resources with.",
													MarkdownDescription: "Name to match resources with.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace to select resources from.",
													MarkdownDescription: "Namespace to select resources from.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
													MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
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

					"migrate_resources": schema.BoolAttribute{
						Description:         "MigrateResources instructs the controller to migrate the Flux custom resources from the previous version to the latest API version specified in the CRD. Defaults to true.",
						MarkdownDescription: "MigrateResources instructs the controller to migrate the Flux custom resources from the previous version to the latest API version specified in the CRD. Defaults to true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sharding": schema.SingleNestedAttribute{
						Description:         "Sharding holds the specification of the sharding configuration.",
						MarkdownDescription: "Sharding holds the specification of the sharding configuration.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Key is the label key used to shard the resources.",
								MarkdownDescription: "Key is the label key used to shard the resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"shards": schema.ListAttribute{
								Description:         "Shards is the list of shard names.",
								MarkdownDescription: "Shards is the list of shard names.",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage holds the specification of the source-controller persistent volume claim.",
						MarkdownDescription: "Storage holds the specification of the source-controller persistent volume claim.",
						Attributes: map[string]schema.Attribute{
							"class": schema.StringAttribute{
								Description:         "Class is the storage class to use for the PVC.",
								MarkdownDescription: "Class is the storage class to use for the PVC.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size is the size of the PVC.",
								MarkdownDescription: "Size is the size of the PVC.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"sync": schema.SingleNestedAttribute{
						Description:         "Sync specifies the source for the cluster sync operation. When set, a Flux source (GitRepository, OCIRepository or Bucket) and Flux Kustomization are created to sync the cluster state with the source repository.",
						MarkdownDescription: "Sync specifies the source for the cluster sync operation. When set, a Flux source (GitRepository, OCIRepository or Bucket) and Flux Kustomization are created to sync the cluster state with the source repository.",
						Attributes: map[string]schema.Attribute{
							"interval": schema.StringAttribute{
								Description:         "Interval is the time between syncs.",
								MarkdownDescription: "Interval is the time between syncs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is the kind of the source.",
								MarkdownDescription: "Kind is the kind of the source.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("OCIRepository", "GitRepository", "Bucket"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the Flux source and kustomization resources. When not specified, the name is set to the namespace name of the FluxInstance.",
								MarkdownDescription: "Name is the name of the Flux source and kustomization resources. When not specified, the name is set to the namespace name of the FluxInstance.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(63),
								},
							},

							"path": schema.StringAttribute{
								Description:         "Path is the path to the source directory containing the kustomize overlay or plain Kubernetes manifests.",
								MarkdownDescription: "Path is the path to the source directory containing the kustomize overlay or plain Kubernetes manifests.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"provider": schema.StringAttribute{
								Description:         "Provider specifies OIDC provider for source authentication. For OCIRepository and Bucket the provider can be set to 'aws', 'azure' or 'gcp'. for GitRepository the accepted value can be set to 'azure' or 'github'. To disable OIDC authentication the provider can be set to 'generic' or left empty.",
								MarkdownDescription: "Provider specifies OIDC provider for source authentication. For OCIRepository and Bucket the provider can be set to 'aws', 'azure' or 'gcp'. for GitRepository the accepted value can be set to 'azure' or 'github'. To disable OIDC authentication the provider can be set to 'generic' or left empty.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("generic", "aws", "azure", "gcp", "github"),
								},
							},

							"pull_secret": schema.StringAttribute{
								Description:         "PullSecret specifies the Kubernetes Secret containing the authentication credentials for the source. For Git over HTTP/S sources, the secret must contain username and password fields. For Git over SSH sources, the secret must contain known_hosts and identity fields. For OCI sources, the secret must be of type kubernetes.io/dockerconfigjson. For Bucket sources, the secret must contain accesskey and secretkey fields.",
								MarkdownDescription: "PullSecret specifies the Kubernetes Secret containing the authentication credentials for the source. For Git over HTTP/S sources, the secret must contain username and password fields. For Git over SSH sources, the secret must contain known_hosts and identity fields. For OCI sources, the secret must be of type kubernetes.io/dockerconfigjson. For Bucket sources, the secret must contain accesskey and secretkey fields.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ref": schema.StringAttribute{
								Description:         "Ref is the source reference, can be a Git ref name e.g. 'refs/heads/main', an OCI tag e.g. 'latest' or a bucket name e.g. 'flux'.",
								MarkdownDescription: "Ref is the source reference, can be a Git ref name e.g. 'refs/heads/main', an OCI tag e.g. 'latest' or a bucket name e.g. 'flux'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "URL is the source URL, can be a Git repository HTTP/S or SSH address, an OCI repository address or a Bucket endpoint.",
								MarkdownDescription: "URL is the source URL, can be a Git repository HTTP/S or SSH address, an OCI repository address or a Bucket endpoint.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"wait": schema.BoolAttribute{
						Description:         "Wait instructs the controller to check the health of all the reconciled resources. Defaults to true.",
						MarkdownDescription: "Wait instructs the controller to check the health of all the reconciled resources. Defaults to true.",
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

func (r *FluxcdControlplaneIoFluxInstanceV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluxcd_controlplane_io_flux_instance_v1_manifest")

	var model FluxcdControlplaneIoFluxInstanceV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluxcd.controlplane.io/v1")
	model.Kind = pointer.String("FluxInstance")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
