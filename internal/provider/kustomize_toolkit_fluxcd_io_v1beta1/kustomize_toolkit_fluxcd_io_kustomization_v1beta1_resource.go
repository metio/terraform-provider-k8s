/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kustomize_toolkit_fluxcd_io_v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource{}
	_ resource.ResourceWithConfigure   = &KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource{}
	_ resource.ResourceWithImportState = &KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource{}
)

func NewKustomizeToolkitFluxcdIoKustomizationV1Beta1Resource() resource.Resource {
	return &KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource{}
}

type KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type KustomizeToolkitFluxcdIoKustomizationV1Beta1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
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
		PatchesJson6902 *[]struct {
			Patch *[]struct {
				From  *string            `tfsdk:"from" json:"from,omitempty"`
				Op    *string            `tfsdk:"op" json:"op,omitempty"`
				Path  *string            `tfsdk:"path" json:"path,omitempty"`
				Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"patch" json:"patch,omitempty"`
			Target *struct {
				AnnotationSelector *string `tfsdk:"annotation_selector" json:"annotationSelector,omitempty"`
				Group              *string `tfsdk:"group" json:"group,omitempty"`
				Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
				LabelSelector      *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Version            *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"patches_json6902" json:"patchesJson6902,omitempty"`
		PatchesStrategicMerge *[]string `tfsdk:"patches_strategic_merge" json:"patchesStrategicMerge,omitempty"`
		Path                  *string   `tfsdk:"path" json:"path,omitempty"`
		PostBuild             *struct {
			Substitute     *map[string]string `tfsdk:"substitute" json:"substitute,omitempty"`
			SubstituteFrom *[]struct {
				Kind *string `tfsdk:"kind" json:"kind,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
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
		Validation      *string `tfsdk:"validation" json:"validation,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kustomize_toolkit_fluxcd_io_kustomization_v1beta1"
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "KustomizationSpec defines the desired state of a kustomization.",
				MarkdownDescription: "KustomizationSpec defines the desired state of a kustomization.",
				Attributes: map[string]schema.Attribute{
					"decryption": schema.SingleNestedAttribute{
						Description:         "Decrypt Kubernetes secrets before applying them on the cluster.",
						MarkdownDescription: "Decrypt Kubernetes secrets before applying them on the cluster.",
						Attributes: map[string]schema.Attribute{
							"provider": schema.StringAttribute{
								Description:         "Provider is the name of the decryption engine.",
								MarkdownDescription: "Provider is the name of the decryption engine.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("sops"),
								},
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "The secret name containing the private OpenPGP keys used for decryption.",
								MarkdownDescription: "The secret name containing the private OpenPGP keys used for decryption.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",
										Required:            true,
										Optional:            false,
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

					"depends_on": schema.ListNestedAttribute{
						Description:         "DependsOn may contain a meta.NamespacedObjectReference slice with references to Kustomization resources that must be ready before this Kustomization can be reconciled.",
						MarkdownDescription: "DependsOn may contain a meta.NamespacedObjectReference slice with references to Kustomization resources that must be ready before this Kustomization can be reconciled.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.",
									MarkdownDescription: "Name of the referent.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
									MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",
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

					"force": schema.BoolAttribute{
						Description:         "Force instructs the controller to recreate resources when patching fails due to an immutable field change.",
						MarkdownDescription: "Force instructs the controller to recreate resources when patching fails due to an immutable field change.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the referent.",
									MarkdownDescription: "Kind of the referent.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the referent.",
									MarkdownDescription: "Name of the referent.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent, when not specified it acts as LocalObjectReference.",
									MarkdownDescription: "Namespace of the referent, when not specified it acts as LocalObjectReference.",
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

					"images": schema.ListNestedAttribute{
						Description:         "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
						MarkdownDescription: "Images is a list of (image name, new name, new tag or digest) for changing image names, tags or digests. This can also be achieved with a patch, but this operator is simpler to specify.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"digest": schema.StringAttribute{
									Description:         "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",
									MarkdownDescription: "Digest is the value used to replace the original image tag. If digest is present NewTag value is ignored.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is a tag-less image name.",
									MarkdownDescription: "Name is a tag-less image name.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"new_name": schema.StringAttribute{
									Description:         "NewName is the value used to replace the original name.",
									MarkdownDescription: "NewName is the value used to replace the original name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"new_tag": schema.StringAttribute{
									Description:         "NewTag is the value used to replace the original tag.",
									MarkdownDescription: "NewTag is the value used to replace the original tag.",
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

					"interval": schema.StringAttribute{
						Description:         "The interval at which to reconcile the Kustomization.",
						MarkdownDescription: "The interval at which to reconcile the Kustomization.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"kube_config": schema.SingleNestedAttribute{
						Description:         "The KubeConfig for reconciling the Kustomization on a remote cluster. When specified, KubeConfig takes precedence over ServiceAccountName.",
						MarkdownDescription: "The KubeConfig for reconciling the Kustomization on a remote cluster. When specified, KubeConfig takes precedence over ServiceAccountName.",
						Attributes: map[string]schema.Attribute{
							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef holds the name to a secret that contains a 'value' key with the kubeconfig file as the value. It must be in the same namespace as the Kustomization. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling the Kustomization.",
								MarkdownDescription: "SecretRef holds the name to a secret that contains a 'value' key with the kubeconfig file as the value. It must be in the same namespace as the Kustomization. It is recommended that the kubeconfig is self-contained, and the secret is regularly updated if credentials such as a cloud-access-token expire. Cloud specific 'cmd-path' auth helpers will not function without adding binaries and credentials to the Pod that is responsible for reconciling the Kustomization.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",
										Required:            true,
										Optional:            false,
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

					"patches_json6902": schema.ListNestedAttribute{
						Description:         "JSON 6902 patches, defined as inline YAML objects.",
						MarkdownDescription: "JSON 6902 patches, defined as inline YAML objects.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"patch": schema.ListNestedAttribute{
									Description:         "Patch contains the JSON6902 patch document with an array of operation objects.",
									MarkdownDescription: "Patch contains the JSON6902 patch document with an array of operation objects.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"from": schema.StringAttribute{
												Description:         "From contains a JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
												MarkdownDescription: "From contains a JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"op": schema.StringAttribute{
												Description:         "Op indicates the operation to perform. Its value MUST be one of 'add', 'remove', 'replace', 'move', 'copy', or 'test'. https://datatracker.ietf.org/doc/html/rfc6902#section-4",
												MarkdownDescription: "Op indicates the operation to perform. Its value MUST be one of 'add', 'remove', 'replace', 'move', 'copy', or 'test'. https://datatracker.ietf.org/doc/html/rfc6902#section-4",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("test", "remove", "add", "replace", "move", "copy"),
												},
											},

											"path": schema.StringAttribute{
												Description:         "Path contains the JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op.",
												MarkdownDescription: "Path contains the JSON-pointer value that references a location within the target document where the operation is performed. The meaning of the value depends on the value of Op.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"value": schema.MapAttribute{
												Description:         "Value contains a valid JSON structure. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
												MarkdownDescription: "Value contains a valid JSON structure. The meaning of the value depends on the value of Op, and is NOT taken into account by all operations.",
												ElementType:         types.StringType,
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
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"patches_strategic_merge": schema.ListAttribute{
						Description:         "Strategic merge patches, defined as inline YAML objects.",
						MarkdownDescription: "Strategic merge patches, defined as inline YAML objects.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path to the directory containing the kustomization.yaml file, or the set of plain YAMLs a kustomization.yaml should be generated for. Defaults to 'None', which translates to the root path of the SourceRef.",
						MarkdownDescription: "Path to the directory containing the kustomization.yaml file, or the set of plain YAMLs a kustomization.yaml should be generated for. Defaults to 'None', which translates to the root path of the SourceRef.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
								Optional:            true,
								Computed:            false,
							},

							"substitute_from": schema.ListNestedAttribute{
								Description:         "SubstituteFrom holds references to ConfigMaps and Secrets containing the variables and their values to be substituted in the YAML manifests. The ConfigMap and the Secret data keys represent the var names and they must match the vars declared in the manifests for the substitution to happen.",
								MarkdownDescription: "SubstituteFrom holds references to ConfigMaps and Secrets containing the variables and their values to be substituted in the YAML manifests. The ConfigMap and the Secret data keys represent the var names and they must match the vars declared in the manifests for the substitution to happen.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",
											MarkdownDescription: "Kind of the values referent, valid values are ('Secret', 'ConfigMap').",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Secret", "ConfigMap"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name of the values referent. Should reside in the same namespace as the referring resource.",
											MarkdownDescription: "Name of the values referent. Should reside in the same namespace as the referring resource.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(253),
											},
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

					"prune": schema.BoolAttribute{
						Description:         "Prune enables garbage collection.",
						MarkdownDescription: "Prune enables garbage collection.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"retry_interval": schema.StringAttribute{
						Description:         "The interval at which to retry a previously failed reconciliation. When not specified, the controller uses the KustomizationSpec.Interval value to retry failures.",
						MarkdownDescription: "The interval at which to retry a previously failed reconciliation. When not specified, the controller uses the KustomizationSpec.Interval value to retry failures.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "The name of the Kubernetes service account to impersonate when reconciling this Kustomization.",
						MarkdownDescription: "The name of the Kubernetes service account to impersonate when reconciling this Kustomization.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_ref": schema.SingleNestedAttribute{
						Description:         "Reference of the source where the kustomization file is.",
						MarkdownDescription: "Reference of the source where the kustomization file is.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent",
								MarkdownDescription: "API version of the referent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent",
								MarkdownDescription: "Kind of the referent",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("GitRepository", "Bucket"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent",
								MarkdownDescription: "Name of the referent",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent, defaults to the Kustomization namespace",
								MarkdownDescription: "Namespace of the referent, defaults to the Kustomization namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "This flag tells the controller to suspend subsequent kustomize executions, it does not apply to already started executions. Defaults to false.",
						MarkdownDescription: "This flag tells the controller to suspend subsequent kustomize executions, it does not apply to already started executions. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_namespace": schema.StringAttribute{
						Description:         "TargetNamespace sets or overrides the namespace in the kustomization.yaml file.",
						MarkdownDescription: "TargetNamespace sets or overrides the namespace in the kustomization.yaml file.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(63),
						},
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout for validation, apply and health checking operations. Defaults to 'Interval' duration.",
						MarkdownDescription: "Timeout for validation, apply and health checking operations. Defaults to 'Interval' duration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"validation": schema.StringAttribute{
						Description:         "Validate the Kubernetes objects before applying them on the cluster. The validation strategy can be 'client' (local dry-run), 'server' (APIServer dry-run) or 'none'. When 'Force' is 'true', validation will fallback to 'client' if set to 'server' because server-side validation is not supported in this scenario.",
						MarkdownDescription: "Validate the Kubernetes objects before applying them on the cluster. The validation strategy can be 'client' (local dry-run), 'server' (APIServer dry-run) or 'none'. When 'Force' is 'true', validation will fallback to 'client' if set to 'server' because server-side validation is not supported in this scenario.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "client", "server"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta1")

	var model KustomizeToolkitFluxcdIoKustomizationV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("kustomize.toolkit.fluxcd.io/v1beta1")
	model.Kind = pointer.String("Kustomization")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta1", Resource: "kustomizations"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KustomizeToolkitFluxcdIoKustomizationV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta1")

	var data KustomizeToolkitFluxcdIoKustomizationV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta1", Resource: "kustomizations"}).
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

	var readResponse KustomizeToolkitFluxcdIoKustomizationV1Beta1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta1")

	var model KustomizeToolkitFluxcdIoKustomizationV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kustomize.toolkit.fluxcd.io/v1beta1")
	model.Kind = pointer.String("Kustomization")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta1", Resource: "kustomizations"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KustomizeToolkitFluxcdIoKustomizationV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta1")

	var data KustomizeToolkitFluxcdIoKustomizationV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta1", Resource: "kustomizations"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
