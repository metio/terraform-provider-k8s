/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package pkg_crossplane_io_v1

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
)

var (
	_ resource.Resource                = &PkgCrossplaneIoConfigurationRevisionV1Resource{}
	_ resource.ResourceWithConfigure   = &PkgCrossplaneIoConfigurationRevisionV1Resource{}
	_ resource.ResourceWithImportState = &PkgCrossplaneIoConfigurationRevisionV1Resource{}
)

func NewPkgCrossplaneIoConfigurationRevisionV1Resource() resource.Resource {
	return &PkgCrossplaneIoConfigurationRevisionV1Resource{}
}

type PkgCrossplaneIoConfigurationRevisionV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type PkgCrossplaneIoConfigurationRevisionV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CommonLabels        *map[string]string `tfsdk:"common_labels" json:"commonLabels,omitempty"`
		ControllerConfigRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"controller_config_ref" json:"controllerConfigRef,omitempty"`
		DesiredState                *string `tfsdk:"desired_state" json:"desiredState,omitempty"`
		EssTLSSecretName            *string `tfsdk:"ess_tls_secret_name" json:"essTLSSecretName,omitempty"`
		IgnoreCrossplaneConstraints *bool   `tfsdk:"ignore_crossplane_constraints" json:"ignoreCrossplaneConstraints,omitempty"`
		Image                       *string `tfsdk:"image" json:"image,omitempty"`
		PackagePullPolicy           *string `tfsdk:"package_pull_policy" json:"packagePullPolicy,omitempty"`
		PackagePullSecrets          *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"package_pull_secrets" json:"packagePullSecrets,omitempty"`
		Revision                 *int64  `tfsdk:"revision" json:"revision,omitempty"`
		SkipDependencyResolution *bool   `tfsdk:"skip_dependency_resolution" json:"skipDependencyResolution,omitempty"`
		TlsClientSecretName      *string `tfsdk:"tls_client_secret_name" json:"tlsClientSecretName,omitempty"`
		TlsServerSecretName      *string `tfsdk:"tls_server_secret_name" json:"tlsServerSecretName,omitempty"`
		WebhookTLSSecretName     *string `tfsdk:"webhook_tls_secret_name" json:"webhookTLSSecretName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_pkg_crossplane_io_configuration_revision_v1"
}

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A ConfigurationRevision that has been added to Crossplane.",
		MarkdownDescription: "A ConfigurationRevision that has been added to Crossplane.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "PackageRevisionSpec specifies the desired state of a PackageRevision.",
				MarkdownDescription: "PackageRevisionSpec specifies the desired state of a PackageRevision.",
				Attributes: map[string]schema.Attribute{
					"common_labels": schema.MapAttribute{
						Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
						MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"controller_config_ref": schema.SingleNestedAttribute{
						Description:         "ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.",
						MarkdownDescription: "ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the ControllerConfig.",
								MarkdownDescription: "Name of the ControllerConfig.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"desired_state": schema.StringAttribute{
						Description:         "DesiredState of the PackageRevision. Can be either Active or Inactive.",
						MarkdownDescription: "DesiredState of the PackageRevision. Can be either Active or Inactive.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ess_tls_secret_name": schema.StringAttribute{
						Description:         "ESSTLSSecretName is the secret name of the TLS certificates that will be used by the provider for External Secret Stores.",
						MarkdownDescription: "ESSTLSSecretName is the secret name of the TLS certificates that will be used by the provider for External Secret Stores.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ignore_crossplane_constraints": schema.BoolAttribute{
						Description:         "IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package. Default is false.",
						MarkdownDescription: "IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package. Default is false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Package image used by install Pod to extract package contents.",
						MarkdownDescription: "Package image used by install Pod to extract package contents.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"package_pull_policy": schema.StringAttribute{
						Description:         "PackagePullPolicy defines the pull policy for the package. It is also applied to any images pulled for the package, such as a provider's controller image. Default is IfNotPresent.",
						MarkdownDescription: "PackagePullPolicy defines the pull policy for the package. It is also applied to any images pulled for the package, such as a provider's controller image. Default is IfNotPresent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"package_pull_secrets": schema.ListNestedAttribute{
						Description:         "PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries. They are also applied to any images pulled for the package, such as a provider's controller image.",
						MarkdownDescription: "PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries. They are also applied to any images pulled for the package, such as a provider's controller image.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

					"revision": schema.Int64Attribute{
						Description:         "Revision number. Indicates when the revision will be garbage collected based on the parent's RevisionHistoryLimit.",
						MarkdownDescription: "Revision number. Indicates when the revision will be garbage collected based on the parent's RevisionHistoryLimit.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"skip_dependency_resolution": schema.BoolAttribute{
						Description:         "SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package. Setting this value to true may have unintended consequences. Default is false.",
						MarkdownDescription: "SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package. Setting this value to true may have unintended consequences. Default is false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls_client_secret_name": schema.StringAttribute{
						Description:         "TLSClientSecretName is the name of the TLS Secret that stores client certificates of the Provider.",
						MarkdownDescription: "TLSClientSecretName is the name of the TLS Secret that stores client certificates of the Provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls_server_secret_name": schema.StringAttribute{
						Description:         "TLSServerSecretName is the name of the TLS Secret that stores server certificates of the Provider.",
						MarkdownDescription: "TLSServerSecretName is the name of the TLS Secret that stores server certificates of the Provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"webhook_tls_secret_name": schema.StringAttribute{
						Description:         "WebhookTLSSecretName is the name of the TLS Secret that will be used by the provider to serve a TLS-enabled webhook server. The certificate will be injected to webhook configurations as well as CRD conversion webhook strategy if needed. If it's not given, provider will not have a certificate mounted to its filesystem, webhook configurations won't be deployed and if there is a CRD with webhook conversion strategy, the installation will fail.",
						MarkdownDescription: "WebhookTLSSecretName is the name of the TLS Secret that will be used by the provider to serve a TLS-enabled webhook server. The certificate will be injected to webhook configurations as well as CRD conversion webhook strategy if needed. If it's not given, provider will not have a certificate mounted to its filesystem, webhook configurations won't be deployed and if there is a CRD with webhook conversion strategy, the installation will fail.",
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

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_pkg_crossplane_io_configuration_revision_v1")

	var model PkgCrossplaneIoConfigurationRevisionV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("pkg.crossplane.io/v1")
	model.Kind = pointer.String("ConfigurationRevision")

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
		Resource(k8sSchema.GroupVersionResource{Group: "pkg.crossplane.io", Version: "v1", Resource: "configurationrevisions"}).
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

	var readResponse PkgCrossplaneIoConfigurationRevisionV1ResourceData
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

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_pkg_crossplane_io_configuration_revision_v1")

	var data PkgCrossplaneIoConfigurationRevisionV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "pkg.crossplane.io", Version: "v1", Resource: "configurationrevisions"}).
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

	var readResponse PkgCrossplaneIoConfigurationRevisionV1ResourceData
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

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_pkg_crossplane_io_configuration_revision_v1")

	var model PkgCrossplaneIoConfigurationRevisionV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("pkg.crossplane.io/v1")
	model.Kind = pointer.String("ConfigurationRevision")

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
		Resource(k8sSchema.GroupVersionResource{Group: "pkg.crossplane.io", Version: "v1", Resource: "configurationrevisions"}).
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

	var readResponse PkgCrossplaneIoConfigurationRevisionV1ResourceData
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

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_pkg_crossplane_io_configuration_revision_v1")

	var data PkgCrossplaneIoConfigurationRevisionV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "pkg.crossplane.io", Version: "v1", Resource: "configurationrevisions"}).
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

func (r *PkgCrossplaneIoConfigurationRevisionV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
