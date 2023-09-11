/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v2

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
	"time"
)

var (
	_ resource.Resource                = &GetambassadorIoDevPortalV2Resource{}
	_ resource.ResourceWithConfigure   = &GetambassadorIoDevPortalV2Resource{}
	_ resource.ResourceWithImportState = &GetambassadorIoDevPortalV2Resource{}
)

func NewGetambassadorIoDevPortalV2Resource() resource.Resource {
	return &GetambassadorIoDevPortalV2Resource{}
}

type GetambassadorIoDevPortalV2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type GetambassadorIoDevPortalV2ResourceData struct {
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
		Ambassador_id *[]string `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Content       *struct {
			Branch *string `tfsdk:"branch" json:"branch,omitempty"`
			Dir    *string `tfsdk:"dir" json:"dir,omitempty"`
			Url    *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"content" json:"content,omitempty"`
		Default *bool `tfsdk:"default" json:"default,omitempty"`
		Docs    *[]struct {
			Service    *string `tfsdk:"service" json:"service,omitempty"`
			Timeout_ms *int64  `tfsdk:"timeout_ms" json:"timeout_ms,omitempty"`
			Url        *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"docs" json:"docs,omitempty"`
		Naming_scheme    *string `tfsdk:"naming_scheme" json:"naming_scheme,omitempty"`
		Preserve_servers *bool   `tfsdk:"preserve_servers" json:"preserve_servers,omitempty"`
		Search           *struct {
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"search" json:"search,omitempty"`
		Selector *struct {
			MatchLabels     *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			MatchNamespaces *[]string          `tfsdk:"match_namespaces" json:"matchNamespaces,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoDevPortalV2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_dev_portal_v2"
}

func (r *GetambassadorIoDevPortalV2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DevPortal is the Schema for the DevPortals API  DevPortal resources specify the 'what' and 'how' is shown in a DevPortal:  1. 'what' is in a DevPortal can be controlled with  - a 'selector', that can be used for filtering 'Mappings'.  - a 'docs' listing of (services, url)  2. 'how' is a pointer to some 'contents' (a checkout of a Git repository with go-templates/markdown/css).  Multiple 'DevPortal's can exist in the cluster, and the Dev Portal server will show them at different endpoints. A 'DevPortal' resource with a special name, 'ambassador', will be used for configuring the default Dev Portal (served at '/docs/' by default).",
		MarkdownDescription: "DevPortal is the Schema for the DevPortals API  DevPortal resources specify the 'what' and 'how' is shown in a DevPortal:  1. 'what' is in a DevPortal can be controlled with  - a 'selector', that can be used for filtering 'Mappings'.  - a 'docs' listing of (services, url)  2. 'how' is a pointer to some 'contents' (a checkout of a Git repository with go-templates/markdown/css).  Multiple 'DevPortal's can exist in the cluster, and the Dev Portal server will show them at different endpoints. A 'DevPortal' resource with a special name, 'ambassador', will be used for configuring the default Dev Portal (served at '/docs/' by default).",
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
				Description:         "DevPortalSpec defines the desired state of DevPortal",
				MarkdownDescription: "DevPortalSpec defines the desired state of DevPortal",
				Attributes: map[string]schema.Attribute{
					"ambassador_id": schema.ListAttribute{
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"content": schema.SingleNestedAttribute{
						Description:         "Content specifies where the content shown in the DevPortal come from",
						MarkdownDescription: "Content specifies where the content shown in the DevPortal come from",
						Attributes: map[string]schema.Attribute{
							"branch": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dir": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"default": schema.BoolAttribute{
						Description:         "Default must be true when this is the default DevPortal",
						MarkdownDescription: "Default must be true when this is the default DevPortal",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"docs": schema.ListNestedAttribute{
						Description:         "Docs is a static docs definition",
						MarkdownDescription: "Docs is a static docs definition",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"service": schema.StringAttribute{
									Description:         "Service is the service being documented",
									MarkdownDescription: "Service is the service being documented",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"timeout_ms": schema.Int64Attribute{
									Description:         "Timeout specifies the amount of time devportal will wait for the downstream service to report an openapi spec back",
									MarkdownDescription: "Timeout specifies the amount of time devportal will wait for the downstream service to report an openapi spec back",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "URL is the URL used for obtaining docs",
									MarkdownDescription: "URL is the URL used for obtaining docs",
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

					"naming_scheme": schema.StringAttribute{
						Description:         "Describes how to display 'services' in the DevPortal. Default namespace.name",
						MarkdownDescription: "Describes how to display 'services' in the DevPortal. Default namespace.name",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("namespace.name", "name.prefix"),
						},
					},

					"preserve_servers": schema.BoolAttribute{
						Description:         "Configures this DevPortal to use server definitions from the openAPI doc instead of rewriting them based on the url used for the connection.",
						MarkdownDescription: "Configures this DevPortal to use server definitions from the openAPI doc instead of rewriting them based on the url used for the connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"search": schema.SingleNestedAttribute{
						Description:         "DevPortalSearchSpec allows configuration over search functionality for the DevPortal",
						MarkdownDescription: "DevPortalSearchSpec allows configuration over search functionality for the DevPortal",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type of search. 'title-only' does a fuzzy search over openapi and page titles 'all-content' will fuzzy search over all openapi and page content. 'title-only' is the default. warning:  using all-content may incur a larger memory footprint",
								MarkdownDescription: "Type of search. 'title-only' does a fuzzy search over openapi and page titles 'all-content' will fuzzy search over all openapi and page content. 'title-only' is the default. warning:  using all-content may incur a larger memory footprint",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("title-only", "all-content"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector is used for choosing what is shown in the DevPortal",
						MarkdownDescription: "Selector is used for choosing what is shown in the DevPortal",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "MatchLabels specifies the list of labels that must be present in Mappings for being present in this DevPortal.",
								MarkdownDescription: "MatchLabels specifies the list of labels that must be present in Mappings for being present in this DevPortal.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"match_namespaces": schema.ListAttribute{
								Description:         "MatchNamespaces is a list of namespaces that will be included in this DevPortal.",
								MarkdownDescription: "MatchNamespaces is a list of namespaces that will be included in this DevPortal.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *GetambassadorIoDevPortalV2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *GetambassadorIoDevPortalV2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_getambassador_io_dev_portal_v2")

	var model GetambassadorIoDevPortalV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("getambassador.io/v2")
	model.Kind = pointer.String("DevPortal")

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
		Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v2", Resource: "devportals"}).
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

	var readResponse GetambassadorIoDevPortalV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *GetambassadorIoDevPortalV2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_dev_portal_v2")

	var data GetambassadorIoDevPortalV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v2", Resource: "devportals"}).
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

	var readResponse GetambassadorIoDevPortalV2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *GetambassadorIoDevPortalV2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_getambassador_io_dev_portal_v2")

	var model GetambassadorIoDevPortalV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("getambassador.io/v2")
	model.Kind = pointer.String("DevPortal")

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
		Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v2", Resource: "devportals"}).
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

	var readResponse GetambassadorIoDevPortalV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *GetambassadorIoDevPortalV2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_getambassador_io_dev_portal_v2")

	var data GetambassadorIoDevPortalV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v2", Resource: "devportals"}).
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
				Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v2", Resource: "devportals"}).
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

func (r *GetambassadorIoDevPortalV2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
