/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kustomize_toolkit_fluxcd_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &KustomizeToolkitFluxcdIoKustomizationV1DataSource{}
	_ datasource.DataSourceWithConfigure = &KustomizeToolkitFluxcdIoKustomizationV1DataSource{}
)

func NewKustomizeToolkitFluxcdIoKustomizationV1DataSource() datasource.DataSource {
	return &KustomizeToolkitFluxcdIoKustomizationV1DataSource{}
}

type KustomizeToolkitFluxcdIoKustomizationV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KustomizeToolkitFluxcdIoKustomizationV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CommonMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"common_metadata" json:"commonMetadata,omitempty"`
		Components *[]string `tfsdk:"components" json:"components,omitempty"`
		Decryption *struct {
			Provider  *string `tfsdk:"provider" json:"provider,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"decryption" json:"decryption,omitempty"`
		DependsOn *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"depends_on" json:"dependsOn,omitempty"`
		Force        *bool `tfsdk:"force" json:"force,omitempty"`
		HealthChecks *[]struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"health_checks" json:"healthChecks,omitempty"`
		Images *[]struct {
			Digest  *string `tfsdk:"digest" json:"digest,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			NewName *string `tfsdk:"new_name" json:"newName,omitempty"`
			NewTag  *string `tfsdk:"new_tag" json:"newTag,omitempty"`
		} `tfsdk:"images" json:"images,omitempty"`
		Interval   *string `tfsdk:"interval" json:"interval,omitempty"`
		KubeConfig *struct {
			SecretRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"kube_config" json:"kubeConfig,omitempty"`
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
		Path      *string `tfsdk:"path" json:"path,omitempty"`
		PostBuild *struct {
			Substitute     *map[string]string `tfsdk:"substitute" json:"substitute,omitempty"`
			SubstituteFrom *[]struct {
				Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"substitute_from" json:"substituteFrom,omitempty"`
		} `tfsdk:"post_build" json:"postBuild,omitempty"`
		Prune              *bool   `tfsdk:"prune" json:"prune,omitempty"`
		RetryInterval      *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		SourceRef          *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"source_ref" json:"sourceRef,omitempty"`
		Suspend         *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		TargetNamespace *string `tfsdk:"target_namespace" json:"targetNamespace,omitempty"`
		Timeout         *string `tfsdk:"timeout" json:"timeout,omitempty"`
		Wait            *bool   `tfsdk:"wait" json:"wait,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kustomize_toolkit_fluxcd_io_kustomization_v1"
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Kustomization is the Schema for the kustomizations API.",
		MarkdownDescription: "Kustomization is the Schema for the kustomizations API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "KustomizationSpec defines the configuration to calculate the desired state from a Source using Kustomize.",
				MarkdownDescription: "KustomizationSpec defines the configuration to calculate the desired state from a Source using Kustomize.",
				Attributes: map[string]schema.Attribute{
					"common_metadata": schema.SingleNestedAttribute{
						Description:         "CommonMetadata specifies the common labels and annotations that are applied to all resources. Any existing label or annotation will be overridden if its key matches a common one.",
						MarkdownDescription: "CommonMetadata specifies the common labels and annotations that are applied to all resources. Any existing label or annotation will be overridden if its key matches a common one.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to be added to the object's metadata.",
								MarkdownDescription: "Annotations to be added to the object's metadata.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels to be added to the object's metadata.",
								MarkdownDescription: "Labels to be added to the object's metadata.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"components": schema.ListAttribute{
						Description:         "Components specifies relative paths to specifications of other Components.",
						MarkdownDescription: "Components specifies relative paths to specifications of other Components.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"decryption": schema.SingleNestedAttribute{
						Description:         "Decrypt Kubernetes secrets before applying them on the cluster.",
						MarkdownDescription: "Decrypt Kubernetes secrets before applying them on the cluster.",
						Attributes: map[string]schema.Attribute{
							"provider": schema.StringAttribute{
								Description:         "Provider is the name of the decryption engine.",
								MarkdownDescription: "Provider is the name of the decryption engine.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "The secret name containing the private OpenPGP keys used for decryption.",
								MarkdownDescription: "The secret name containing the private OpenPGP keys used for decryption.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"depends_on": schema.ListNestedAttribute{
						Description:         "DependsOn may contain a meta.NamespacedObjectReference slice with references to Kustomization resources that must be ready before this Kustomization can be reconciled.",
						MarkdownDescription: "DependsOn may contain a meta.NamespacedObjectReference slice with references to Kustomization resources that must be ready before this Kustomization can be reconciled.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.",
									MarkdownDescription: "Name of the referent.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
									MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"force": schema.BoolAttribute{
						Description:         "Force instructs the controller to recreate resources when patching fails due to an immutable field change.",
						MarkdownDescription: "Force instructs the controller to recreate resources when patching fails due to an immutable field change.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"health_checks": schema.ListNestedAttribute{
						Description:         "A list of resources to be included in the health assessment.",
						MarkdownDescription: "A list of resources to be included in the health assessment.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "API version of the referent, if not specified the Kubernetes preferred version will be used.",
									MarkdownDescription: "API version of the referent, if not specified the Kubernetes preferred version will be used.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the referent.",
									MarkdownDescription: "Kind of the referent.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the referent.",
									MarkdownDescription: "Name of the referent.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
									MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"images": schema.ListNestedAttribute{
						Description:         "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
						MarkdownDescription: "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"digest": schema.StringAttribute{
									Description:         "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",
									MarkdownDescription: "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name is a tag-less image name.",
									MarkdownDescription: "Name is a tag-less image name.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"new_name": schema.StringAttribute{
									Description:         "NewName is the value used to replace the original name.",
									MarkdownDescription: "NewName is the value used to replace the original name.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"new_tag": schema.StringAttribute{
									Description:         "NewTag is the value used to replace the original tag.",
									MarkdownDescription: "NewTag is the value used to replace the original tag.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"interval": schema.StringAttribute{
						Description:         "The interval at which to reconcile the Kustomization. This interval is approximate and may be subject to jitter to ensure efficient use of resources.",
						MarkdownDescription: "The interval at which to reconcile the Kustomization. This interval is approximate and may be subject to jitter to ensure efficient use of resources.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"kube_config": schema.SingleNestedAttribute{
						Description:         "The KubeConfig for reconciling the Kustomization on a remote cluster. When used in combination with KustomizationSpec.ServiceAccountName, forces the controller to act on behalf of that Service Account at the target cluster. If the --default-service-account flag is set, its value will be used as a controller level fallback for when KustomizationSpec.ServiceAccountName is empty.",
						MarkdownDescription: "The KubeConfig for reconciling the Kustomization on a remote cluster. When used in combination with KustomizationSpec.ServiceAccountName, forces the controller to act on behalf of that Service Account at the target cluster. If the --default-service-account flag is set, its value will be used as a controller level fallback for when KustomizationSpec.ServiceAccountName is empty.",
						Attributes: map[string]schema.Attribute{
							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef holds the name of a secret that contains a key with the kubeconfig file as the value. If no key is set, the key will default to 'value'. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling Kubernetes resources.",
								MarkdownDescription: "SecretRef holds the name of a secret that contains a key with the kubeconfig file as the value. If no key is set, the key will default to 'value'. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling Kubernetes resources.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "Key in the Secret, when not specified an implementation-specific default key is used.",
										MarkdownDescription: "Key in the Secret, when not specified an implementation-specific default key is used.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the Secret.",
										MarkdownDescription: "Name of the Secret.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"patches": schema.ListNestedAttribute{
						Description:         "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",
						MarkdownDescription: "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"patch": schema.StringAttribute{
									Description:         "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",
									MarkdownDescription: "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"target": schema.SingleNestedAttribute{
									Description:         "Target points to the resources that the patch document should be applied to.",
									MarkdownDescription: "Target points to the resources that the patch document should be applied to.",
									Attributes: map[string]schema.Attribute{
										"annotation_selector": schema.StringAttribute{
											Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
											MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"group": schema.StringAttribute{
											Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
											MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
											MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"label_selector": schema.StringAttribute{
											Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
											MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name to match resources with.",
											MarkdownDescription: "Name to match resources with.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace to select resources from.",
											MarkdownDescription: "Namespace to select resources from.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"version": schema.StringAttribute{
											Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
											MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"path": schema.StringAttribute{
						Description:         "Path to the directory containing the kustomization.yaml file, or the set of plain YAMLs a kustomization.yaml should be generated for. Defaults to 'None', which translates to the root path of the SourceRef.",
						MarkdownDescription: "Path to the directory containing the kustomization.yaml file, or the set of plain YAMLs a kustomization.yaml should be generated for. Defaults to 'None', which translates to the root path of the SourceRef.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"post_build": schema.SingleNestedAttribute{
						Description:         "PostBuild describes which actions to perform on the YAML manifest generated by building the kustomize overlay.",
						MarkdownDescription: "PostBuild describes which actions to perform on the YAML manifest generated by building the kustomize overlay.",
						Attributes: map[string]schema.Attribute{
							"substitute": schema.MapAttribute{
								Description:         "Substitute holds a map of key/value pairs. The variables defined in your YAML manifests that match any of the keys defined in the map will be substituted with the set value. Includes support for bash string replacement functions e.g. ${var:=default}, ${var:position} and ${var/substring/replacement}.",
								MarkdownDescription: "Substitute holds a map of key/value pairs. The variables defined in your YAML manifests that match any of the keys defined in the map will be substituted with the set value. Includes support for bash string replacement functions e.g. ${var:=default}, ${var:position} and ${var/substring/replacement}.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"substitute_from": schema.ListNestedAttribute{
								Description:         "SubstituteFrom holds references to ConfigMaps and Secrets containing the variables and their values to be substituted in the YAML manifests. The ConfigMap and the Secret data keys represent the var names, and they must match the vars declared in the manifests for the substitution to happen.",
								MarkdownDescription: "SubstituteFrom holds references to ConfigMaps and Secrets containing the variables and their values to be substituted in the YAML manifests. The ConfigMap and the Secret data keys represent the var names, and they must match the vars declared in the manifests for the substitution to happen.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",
											MarkdownDescription: "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the values referent. Should reside in the same namespace as the referring resource.",
											MarkdownDescription: "Name of the values referent. Should reside in the same namespace as the referring resource.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"optional": schema.BoolAttribute{
											Description:         "Optional indicates whether the referenced resource must exist, or whether to tolerate its absence. If true and the referenced resource is absent, proceed as if the resource was present but empty, without any variables defined.",
											MarkdownDescription: "Optional indicates whether the referenced resource must exist, or whether to tolerate its absence. If true and the referenced resource is absent, proceed as if the resource was present but empty, without any variables defined.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"prune": schema.BoolAttribute{
						Description:         "Prune enables garbage collection.",
						MarkdownDescription: "Prune enables garbage collection.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"retry_interval": schema.StringAttribute{
						Description:         "The interval at which to retry a previously failed reconciliation. When not specified, the controller uses the KustomizationSpec.Interval value to retry failures.",
						MarkdownDescription: "The interval at which to retry a previously failed reconciliation. When not specified, the controller uses the KustomizationSpec.Interval value to retry failures.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "The name of the Kubernetes service account to impersonate when reconciling this Kustomization.",
						MarkdownDescription: "The name of the Kubernetes service account to impersonate when reconciling this Kustomization.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"source_ref": schema.SingleNestedAttribute{
						Description:         "Reference of the source where the kustomization file is.",
						MarkdownDescription: "Reference of the source where the kustomization file is.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
								MarkdownDescription: "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"suspend": schema.BoolAttribute{
						Description:         "This flag tells the controller to suspend subsequent kustomize executions, it does not apply to already started executions. Defaults to false.",
						MarkdownDescription: "This flag tells the controller to suspend subsequent kustomize executions, it does not apply to already started executions. Defaults to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"target_namespace": schema.StringAttribute{
						Description:         "TargetNamespace sets or overrides the namespace in the kustomization.yaml file.",
						MarkdownDescription: "TargetNamespace sets or overrides the namespace in the kustomization.yaml file.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout for validation, apply and health checking operations. Defaults to 'Interval' duration.",
						MarkdownDescription: "Timeout for validation, apply and health checking operations. Defaults to 'Interval' duration.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"wait": schema.BoolAttribute{
						Description:         "Wait instructs the controller to check the health of all the reconciled resources. When enabled, the HealthChecks are ignored. Defaults to false.",
						MarkdownDescription: "Wait instructs the controller to check the health of all the reconciled resources. When enabled, the HealthChecks are ignored. Defaults to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kustomize_toolkit_fluxcd_io_kustomization_v1")

	var data KustomizeToolkitFluxcdIoKustomizationV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1", Resource: "Kustomization"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KustomizeToolkitFluxcdIoKustomizationV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("kustomize.toolkit.fluxcd.io/v1")
	data.Kind = pointer.String("Kustomization")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
