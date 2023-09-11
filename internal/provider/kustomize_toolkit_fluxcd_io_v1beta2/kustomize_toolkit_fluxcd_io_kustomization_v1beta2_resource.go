/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kustomize_toolkit_fluxcd_io_v1beta2

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
	"regexp"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource{}
	_ resource.ResourceWithConfigure   = &KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource{}
	_ resource.ResourceWithImportState = &KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource{}
)

func NewKustomizeToolkitFluxcdIoKustomizationV1Beta2Resource() resource.Resource {
	return &KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource{}
}

type KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type KustomizeToolkitFluxcdIoKustomizationV1Beta2ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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
		Validation      *string `tfsdk:"validation" json:"validation,omitempty"`
		Wait            *bool   `tfsdk:"wait" json:"wait,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kustomize_toolkit_fluxcd_io_kustomization_v1beta2"
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"wait_for_upsert": schema.ListNestedAttribute{
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
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
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
						Description:         "Components specifies relative paths to specifications of other Components.",
						MarkdownDescription: "Components specifies relative paths to specifications of other Components.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

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
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
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
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the Secret.",
										MarkdownDescription: "Name of the Secret.",
										Required:            true,
										Optional:            false,
										Computed:            false,
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
						Description:         "JSON 6902 patches, defined as inline YAML objects. Deprecated: Use Patches instead.",
						MarkdownDescription: "JSON 6902 patches, defined as inline YAML objects. Deprecated: Use Patches instead.",
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
						Description:         "Strategic merge patches, defined as inline YAML objects. Deprecated: Use Patches instead.",
						MarkdownDescription: "Strategic merge patches, defined as inline YAML objects. Deprecated: Use Patches instead.",
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

										"optional": schema.BoolAttribute{
											Description:         "Optional indicates whether the referenced resource must exist, or whether to tolerate its absence. If true and the referenced resource is absent, proceed as if the resource was present but empty, without any variables defined.",
											MarkdownDescription: "Optional indicates whether the referenced resource must exist, or whether to tolerate its absence. If true and the referenced resource is absent, proceed as if the resource was present but empty, without any variables defined.",
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
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
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
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
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
								Validators: []validator.String{
									stringvalidator.OneOf("OCIRepository", "GitRepository", "Bucket"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
								MarkdownDescription: "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
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
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"validation": schema.StringAttribute{
						Description:         "Deprecated: Not used in v1beta2.",
						MarkdownDescription: "Deprecated: Not used in v1beta2.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "client", "server"),
						},
					},

					"wait": schema.BoolAttribute{
						Description:         "Wait instructs the controller to check the health of all the reconciled resources. When enabled, the HealthChecks are ignored. Defaults to false.",
						MarkdownDescription: "Wait instructs the controller to check the health of all the reconciled resources. When enabled, the HealthChecks are ignored. Defaults to false.",
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

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2")

	var model KustomizeToolkitFluxcdIoKustomizationV1Beta2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kustomize.toolkit.fluxcd.io/v1beta2")
	model.Kind = pointer.String("Kustomization")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta2", Resource: "kustomizations"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse KustomizeToolkitFluxcdIoKustomizationV1Beta2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2")

	var data KustomizeToolkitFluxcdIoKustomizationV1Beta2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta2", Resource: "kustomizations"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse KustomizeToolkitFluxcdIoKustomizationV1Beta2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2")

	var model KustomizeToolkitFluxcdIoKustomizationV1Beta2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kustomize.toolkit.fluxcd.io/v1beta2")
	model.Kind = pointer.String("Kustomization")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta2", Resource: "kustomizations"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse KustomizeToolkitFluxcdIoKustomizationV1Beta2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2")

	var data KustomizeToolkitFluxcdIoKustomizationV1Beta2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta2", Resource: "kustomizations"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta2", Resource: "kustomizations"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *KustomizeToolkitFluxcdIoKustomizationV1Beta2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
