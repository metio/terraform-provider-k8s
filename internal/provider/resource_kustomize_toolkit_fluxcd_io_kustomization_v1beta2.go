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

type KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource struct{}

var (
	_ resource.Resource = (*KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource)(nil)
)

type KustomizeToolkitFluxcdIoKustomizationV1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KustomizeToolkitFluxcdIoKustomizationV1Beta2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Decryption *struct {
			Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
		} `tfsdk:"decryption" yaml:"decryption,omitempty"`

		Force *bool `tfsdk:"force" yaml:"force,omitempty"`

		PatchesJson6902 *[]struct {
			Patch *[]struct {
				From *string `tfsdk:"from" yaml:"from,omitempty"`

				Op *string `tfsdk:"op" yaml:"op,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Value *map[string]string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"patch" yaml:"patch,omitempty"`

			Target *struct {
				AnnotationSelector *string `tfsdk:"annotation_selector" yaml:"annotationSelector,omitempty"`

				Group *string `tfsdk:"group" yaml:"group,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				LabelSelector *string `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"target" yaml:"target,omitempty"`
		} `tfsdk:"patches_json6902" yaml:"patchesJson6902,omitempty"`

		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		Wait *bool `tfsdk:"wait" yaml:"wait,omitempty"`

		Images *[]struct {
			Digest *string `tfsdk:"digest" yaml:"digest,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			NewName *string `tfsdk:"new_name" yaml:"newName,omitempty"`

			NewTag *string `tfsdk:"new_tag" yaml:"newTag,omitempty"`
		} `tfsdk:"images" yaml:"images,omitempty"`

		KubeConfig *struct {
			SecretRef *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
		} `tfsdk:"kube_config" yaml:"kubeConfig,omitempty"`

		SourceRef *struct {
			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`
		} `tfsdk:"source_ref" yaml:"sourceRef,omitempty"`

		Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

		HealthChecks *[]struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"health_checks" yaml:"healthChecks,omitempty"`

		Patches *[]struct {
			Patch *string `tfsdk:"patch" yaml:"patch,omitempty"`

			Target *struct {
				LabelSelector *string `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`

				AnnotationSelector *string `tfsdk:"annotation_selector" yaml:"annotationSelector,omitempty"`

				Group *string `tfsdk:"group" yaml:"group,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
			} `tfsdk:"target" yaml:"target,omitempty"`
		} `tfsdk:"patches" yaml:"patches,omitempty"`

		Path *string `tfsdk:"path" yaml:"path,omitempty"`

		PostBuild *struct {
			Substitute *map[string]string `tfsdk:"substitute" yaml:"substitute,omitempty"`

			SubstituteFrom *[]struct {
				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"substitute_from" yaml:"substituteFrom,omitempty"`
		} `tfsdk:"post_build" yaml:"postBuild,omitempty"`

		Prune *bool `tfsdk:"prune" yaml:"prune,omitempty"`

		TargetNamespace *string `tfsdk:"target_namespace" yaml:"targetNamespace,omitempty"`

		Validation *string `tfsdk:"validation" yaml:"validation,omitempty"`

		DependsOn *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"depends_on" yaml:"dependsOn,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		PatchesStrategicMerge *[]string `tfsdk:"patches_strategic_merge" yaml:"patchesStrategicMerge,omitempty"`

		RetryInterval *string `tfsdk:"retry_interval" yaml:"retryInterval,omitempty"`

		ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKustomizeToolkitFluxcdIoKustomizationV1Beta2Resource() resource.Resource {
	return &KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource{}
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kustomize_toolkit_fluxcd_io_kustomization_v1beta2"
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Kustomization is the Schema for the kustomizations API.",
		MarkdownDescription: "Kustomization is the Schema for the kustomizations API.",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "KustomizationSpec defines the configuration to calculate the desired state from a Source using Kustomize.",
				MarkdownDescription: "KustomizationSpec defines the configuration to calculate the desired state from a Source using Kustomize.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"decryption": {
						Description:         "Decrypt Kubernetes secrets before applying them on the cluster.",
						MarkdownDescription: "Decrypt Kubernetes secrets before applying them on the cluster.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"provider": {
								Description:         "Provider is the name of the decryption engine.",
								MarkdownDescription: "Provider is the name of the decryption engine.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"secret_ref": {
								Description:         "The secret name containing the private OpenPGP keys used for decryption.",
								MarkdownDescription: "The secret name containing the private OpenPGP keys used for decryption.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"force": {
						Description:         "Force instructs the controller to recreate resources when patching fails due to an immutable field change.",
						MarkdownDescription: "Force instructs the controller to recreate resources when patching fails due to an immutable field change.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"patches_json6902": {
						Description:         "JSON 6902 patches, defined as inline YAML objects. Deprecated: Use Patches instead.",
						MarkdownDescription: "JSON 6902 patches, defined as inline YAML objects. Deprecated: Use Patches instead.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"patch": {
								Description:         "Patch contains the JSON6902 patch document with an array of operation objects.",
								MarkdownDescription: "Patch contains the JSON6902 patch document with an array of operation objects.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"from": {
										Description:         "From contains a JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
										MarkdownDescription: "From contains a JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"op": {
										Description:         "Op indicates the operation to perform. Its value MUST be one of 'add', 'remove', 'replace', 'move', 'copy', or 'test'. https://datatracker.ietf.org/doc/html/rfc6902#section-4",
										MarkdownDescription: "Op indicates the operation to perform. Its value MUST be one of 'add', 'remove', 'replace', 'move', 'copy', or 'test'. https://datatracker.ietf.org/doc/html/rfc6902#section-4",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"path": {
										Description:         "Path contains the JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op.",
										MarkdownDescription: "Path contains the JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Value contains a valid JSON structure. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
										MarkdownDescription: "Value contains a valid JSON structure. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"target": {
								Description:         "Target points to the resources that the patch document should be applied to.",
								MarkdownDescription: "Target points to the resources that the patch document should be applied to.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotation_selector": {
										Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
										MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"group": {
										Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
										MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
										MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"label_selector": {
										Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
										MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name to match resources with.",
										MarkdownDescription: "Name to match resources with.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace to select resources from.",
										MarkdownDescription: "Namespace to select resources from.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
										MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
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

					"suspend": {
						Description:         "This flag tells the controller to suspend subsequent kustomize executions, it does not apply to already started executions. Defaults to false.",
						MarkdownDescription: "This flag tells the controller to suspend subsequent kustomize executions, it does not apply to already started executions. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wait": {
						Description:         "Wait instructs the controller to check the health of all the reconciled resources. When enabled, the HealthChecks are ignored. Defaults to false.",
						MarkdownDescription: "Wait instructs the controller to check the health of all the reconciled resources. When enabled, the HealthChecks are ignored. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"images": {
						Description:         "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
						MarkdownDescription: "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"digest": {
								Description:         "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",
								MarkdownDescription: "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name is a tag-less image name.",
								MarkdownDescription: "Name is a tag-less image name.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"new_name": {
								Description:         "NewName is the value used to replace the original name.",
								MarkdownDescription: "NewName is the value used to replace the original name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"new_tag": {
								Description:         "NewTag is the value used to replace the original tag.",
								MarkdownDescription: "NewTag is the value used to replace the original tag.",

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

					"kube_config": {
						Description:         "The KubeConfig for reconciling the Kustomization on a remote cluster. When used in combination with KustomizationSpec.ServiceAccountName, forces the controller to act on behalf of that Service Account at the target cluster. If the --default-service-account flag is set, its value will be used as a controller level fallback for when KustomizationSpec.ServiceAccountName is empty.",
						MarkdownDescription: "The KubeConfig for reconciling the Kustomization on a remote cluster. When used in combination with KustomizationSpec.ServiceAccountName, forces the controller to act on behalf of that Service Account at the target cluster. If the --default-service-account flag is set, its value will be used as a controller level fallback for when KustomizationSpec.ServiceAccountName is empty.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"secret_ref": {
								Description:         "SecretRef holds the name of a secret that contains a key with the kubeconfig file as the value. If no key is set, the key will default to 'value'. The secret must be in the same namespace as the Kustomization. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling the Kustomization.",
								MarkdownDescription: "SecretRef holds the name of a secret that contains a key with the kubeconfig file as the value. If no key is set, the key will default to 'value'. The secret must be in the same namespace as the Kustomization. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling the Kustomization.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "Key in the Secret, when not specified an implementation-specific default key is used.",
										MarkdownDescription: "Key in the Secret, when not specified an implementation-specific default key is used.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the Secret.",
										MarkdownDescription: "Name of the Secret.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_ref": {
						Description:         "Reference of the source where the kustomization file is.",
						MarkdownDescription: "Reference of the source where the kustomization file is.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"kind": {
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
								MarkdownDescription: "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"api_version": {
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"timeout": {
						Description:         "Timeout for validation, apply and health checking operations. Defaults to 'Interval' duration.",
						MarkdownDescription: "Timeout for validation, apply and health checking operations. Defaults to 'Interval' duration.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"health_checks": {
						Description:         "A list of resources to be included in the health assessment.",
						MarkdownDescription: "A list of resources to be included in the health assessment.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "API version of the referent, if not specified the Kubernetes preferred version will be used.",
								MarkdownDescription: "API version of the referent, if not specified the Kubernetes preferred version will be used.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
								MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",

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

					"patches": {
						Description:         "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",
						MarkdownDescription: "Strategic merge and JSON patches, defined as inline YAML objects, capable of targeting objects based on kind, label and annotation selectors.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"patch": {
								Description:         "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",
								MarkdownDescription: "Patch contains an inline StrategicMerge patch or an inline JSON6902 patch with an array of operation objects.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target": {
								Description:         "Target points to the resources that the patch document should be applied to.",
								MarkdownDescription: "Target points to the resources that the patch document should be applied to.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"label_selector": {
										Description:         "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",
										MarkdownDescription: "LabelSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource labels.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name to match resources with.",
										MarkdownDescription: "Name to match resources with.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace to select resources from.",
										MarkdownDescription: "Namespace to select resources from.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
										MarkdownDescription: "Version of the API Group to select resources from. Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"annotation_selector": {
										Description:         "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",
										MarkdownDescription: "AnnotationSelector is a string that follows the label selection expression https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api It matches with the resource annotations.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"group": {
										Description:         "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
										MarkdownDescription: "Group is the API group to select resources from. Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",
										MarkdownDescription: "Kind of the API Group to select resources from. Together with Group and Version it is capable of unambiguously identifying and/or selecting resources. https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md",

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

						Required: false,
						Optional: true,
						Computed: false,
					},

					"path": {
						Description:         "Path to the directory containing the kustomization.yaml file, or the set of plain YAMLs a kustomization.yaml should be generated for. Defaults to 'None', which translates to the root path of the SourceRef.",
						MarkdownDescription: "Path to the directory containing the kustomization.yaml file, or the set of plain YAMLs a kustomization.yaml should be generated for. Defaults to 'None', which translates to the root path of the SourceRef.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"post_build": {
						Description:         "PostBuild describes which actions to perform on the YAML manifest generated by building the kustomize overlay.",
						MarkdownDescription: "PostBuild describes which actions to perform on the YAML manifest generated by building the kustomize overlay.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"substitute": {
								Description:         "Substitute holds a map of key/value pairs. The variables defined in your YAML manifests that match any of the keys defined in the map will be substituted with the set value. Includes support for bash string replacement functions e.g. ${var:=default}, ${var:position} and ${var/substring/replacement}.",
								MarkdownDescription: "Substitute holds a map of key/value pairs. The variables defined in your YAML manifests that match any of the keys defined in the map will be substituted with the set value. Includes support for bash string replacement functions e.g. ${var:=default}, ${var:position} and ${var/substring/replacement}.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"substitute_from": {
								Description:         "SubstituteFrom holds references to ConfigMaps and Secrets containing the variables and their values to be substituted in the YAML manifests. The ConfigMap and the Secret data keys represent the var names and they must match the vars declared in the manifests for the substitution to happen.",
								MarkdownDescription: "SubstituteFrom holds references to ConfigMaps and Secrets containing the variables and their values to be substituted in the YAML manifests. The ConfigMap and the Secret data keys represent the var names and they must match the vars declared in the manifests for the substitution to happen.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"kind": {
										Description:         "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",
										MarkdownDescription: "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the values referent. Should reside in the same namespace as the referring resource.",
										MarkdownDescription: "Name of the values referent. Should reside in the same namespace as the referring resource.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"optional": {
										Description:         "Optional indicates whether the referenced resource must exist, or whether to tolerate its absence. If true and the referenced resource is absent, proceed as if the resource was present but empty, without any variables defined.",
										MarkdownDescription: "Optional indicates whether the referenced resource must exist, or whether to tolerate its absence. If true and the referenced resource is absent, proceed as if the resource was present but empty, without any variables defined.",

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

					"prune": {
						Description:         "Prune enables garbage collection.",
						MarkdownDescription: "Prune enables garbage collection.",

						Type: types.BoolType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"target_namespace": {
						Description:         "TargetNamespace sets or overrides the namespace in the kustomization.yaml file.",
						MarkdownDescription: "TargetNamespace sets or overrides the namespace in the kustomization.yaml file.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"validation": {
						Description:         "Deprecated: Not used in v1beta2.",
						MarkdownDescription: "Deprecated: Not used in v1beta2.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"depends_on": {
						Description:         "DependsOn may contain a meta.NamespacedObjectReference slice with references to Kustomization resources that must be ready before this Kustomization can be reconciled.",
						MarkdownDescription: "DependsOn may contain a meta.NamespacedObjectReference slice with references to Kustomization resources that must be ready before this Kustomization can be reconciled.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
								MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",

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

					"interval": {
						Description:         "The interval at which to reconcile the Kustomization.",
						MarkdownDescription: "The interval at which to reconcile the Kustomization.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"patches_strategic_merge": {
						Description:         "Strategic merge patches, defined as inline YAML objects. Deprecated: Use Patches instead.",
						MarkdownDescription: "Strategic merge patches, defined as inline YAML objects. Deprecated: Use Patches instead.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"retry_interval": {
						Description:         "The interval at which to retry a previously failed reconciliation. When not specified, the controller uses the KustomizationSpec.Interval value to retry failures.",
						MarkdownDescription: "The interval at which to retry a previously failed reconciliation. When not specified, the controller uses the KustomizationSpec.Interval value to retry failures.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_name": {
						Description:         "The name of the Kubernetes service account to impersonate when reconciling this Kustomization.",
						MarkdownDescription: "The name of the Kubernetes service account to impersonate when reconciling this Kustomization.",

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

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2")

	var state KustomizeToolkitFluxcdIoKustomizationV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KustomizeToolkitFluxcdIoKustomizationV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kustomize.toolkit.fluxcd.io/v1beta2")
	goModel.Kind = utilities.Ptr("Kustomization")

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

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2")

	var state KustomizeToolkitFluxcdIoKustomizationV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KustomizeToolkitFluxcdIoKustomizationV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kustomize.toolkit.fluxcd.io/v1beta2")
	goModel.Kind = utilities.Ptr("Kustomization")

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

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
