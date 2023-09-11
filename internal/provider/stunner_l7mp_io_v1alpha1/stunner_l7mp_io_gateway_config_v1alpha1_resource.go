/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package stunner_l7mp_io_v1alpha1

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
	_ resource.Resource                = &StunnerL7MpIoGatewayConfigV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &StunnerL7MpIoGatewayConfigV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &StunnerL7MpIoGatewayConfigV1Alpha1Resource{}
)

func NewStunnerL7MpIoGatewayConfigV1Alpha1Resource() resource.Resource {
	return &StunnerL7MpIoGatewayConfigV1Alpha1Resource{}
}

type StunnerL7MpIoGatewayConfigV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type StunnerL7MpIoGatewayConfigV1Alpha1ResourceData struct {
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
		AuthLifetime *int64 `tfsdk:"auth_lifetime" json:"authLifetime,omitempty"`
		AuthRef      *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"auth_ref" json:"authRef,omitempty"`
		AuthType                       *string            `tfsdk:"auth_type" json:"authType,omitempty"`
		HealthCheckEndpoint            *string            `tfsdk:"health_check_endpoint" json:"healthCheckEndpoint,omitempty"`
		LoadBalancerServiceAnnotations *map[string]string `tfsdk:"load_balancer_service_annotations" json:"loadBalancerServiceAnnotations,omitempty"`
		LogLevel                       *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		MaxPort                        *int64             `tfsdk:"max_port" json:"maxPort,omitempty"`
		MetricsEndpoint                *string            `tfsdk:"metrics_endpoint" json:"metricsEndpoint,omitempty"`
		MinPort                        *int64             `tfsdk:"min_port" json:"minPort,omitempty"`
		Password                       *string            `tfsdk:"password" json:"password,omitempty"`
		Realm                          *string            `tfsdk:"realm" json:"realm,omitempty"`
		SharedSecret                   *string            `tfsdk:"shared_secret" json:"sharedSecret,omitempty"`
		StunnerConfig                  *string            `tfsdk:"stunner_config" json:"stunnerConfig,omitempty"`
		UserName                       *string            `tfsdk:"user_name" json:"userName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StunnerL7MpIoGatewayConfigV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_stunner_l7mp_io_gateway_config_v1alpha1"
}

func (r *StunnerL7MpIoGatewayConfigV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GatewayConfig is the Schema for the gatewayconfigs API",
		MarkdownDescription: "GatewayConfig is the Schema for the gatewayconfigs API",
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
				Description:         "GatewayConfigSpec defines the desired state of GatewayConfig",
				MarkdownDescription: "GatewayConfigSpec defines the desired state of GatewayConfig",
				Attributes: map[string]schema.Attribute{
					"auth_lifetime": schema.Int64Attribute{
						Description:         "AuthLifetime defines the lifetime of 'longterm' authentication credentials in seconds.",
						MarkdownDescription: "AuthLifetime defines the lifetime of 'longterm' authentication credentials in seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auth_ref": schema.SingleNestedAttribute{
						Description:         "Note that externally set credentials override any inline auth credentials (AuthType, AuthUsername, etc.): if AuthRef is nonempty then it is expected that the referenced Secret exists and *all* authentication credentials are correctly set in the referenced Secret (username/password or shared secret). Mixing of credential sources (inline/external) is not supported.",
						MarkdownDescription: "Note that externally set credentials override any inline auth credentials (AuthType, AuthUsername, etc.): if AuthRef is nonempty then it is expected that the referenced Secret exists and *all* authentication credentials are correctly set in the referenced Secret (username/password or shared secret). Mixing of credential sources (inline/external) is not supported.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
								MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
								MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the referent.",
								MarkdownDescription: "Name is the name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
								MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"auth_type": schema.StringAttribute{
						Description:         "AuthType is the type of the STUN/TURN authentication mechanism.",
						MarkdownDescription: "AuthType is the type of the STUN/TURN authentication mechanism.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^plaintext|static|longterm|ephemeral|timewindowed$`), ""),
						},
					},

					"health_check_endpoint": schema.StringAttribute{
						Description:         "HealthCheckEndpoint is the URI of the form 'http://address:port' exposed for external HTTP health-checking. A liveness probe responder will be exposed on path '/live' and readiness probe on path '/ready'. The scheme ('http://') is mandatory, default is to enable health-checking at 'http://0.0.0.0:8086'.",
						MarkdownDescription: "HealthCheckEndpoint is the URI of the form 'http://address:port' exposed for external HTTP health-checking. A liveness probe responder will be exposed on path '/live' and readiness probe on path '/ready'. The scheme ('http://') is mandatory, default is to enable health-checking at 'http://0.0.0.0:8086'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_balancer_service_annotations": schema.MapAttribute{
						Description:         "LoadBalancerServiceAnnotations is a list of annotations that will go into the LoadBalancer services created automatically by the operator to wrap Gateways.  NOTE: removing annotations from a GatewayConfig will not result in the removal of the corresponding annotations from the LoadBalancer service, in order to prevent the accidental removal of an annotation installed there by Kubernetes or the cloud provider. If you really want to remove an annotation, do this manually or simply remove all Gateways (which will remove the corresponding LoadBalancer services), update the GatewayConfig and then recreate the Gateways, so that the newly created LoadBalancer services will contain the required annotations.",
						MarkdownDescription: "LoadBalancerServiceAnnotations is a list of annotations that will go into the LoadBalancer services created automatically by the operator to wrap Gateways.  NOTE: removing annotations from a GatewayConfig will not result in the removal of the corresponding annotations from the LoadBalancer service, in order to prevent the accidental removal of an annotation installed there by Kubernetes or the cloud provider. If you really want to remove an annotation, do this manually or simply remove all Gateways (which will remove the corresponding LoadBalancer services), update the GatewayConfig and then recreate the Gateways, so that the newly created LoadBalancer services will contain the required annotations.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel specifies the default loglevel for the STUNner daemon.",
						MarkdownDescription: "LogLevel specifies the default loglevel for the STUNner daemon.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_port": schema.Int64Attribute{
						Description:         "MaxRelayPort is the smallest relay port assigned for STUNner relay connections.",
						MarkdownDescription: "MaxRelayPort is the smallest relay port assigned for STUNner relay connections.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_endpoint": schema.StringAttribute{
						Description:         "MetricsEndpoint is the URI in the form 'http://address:port/path' exposed for metric scraping (Prometheus). The scheme ('http://') is mandatory. Default is to expose no metric endpoint.",
						MarkdownDescription: "MetricsEndpoint is the URI in the form 'http://address:port/path' exposed for metric scraping (Prometheus). The scheme ('http://') is mandatory. Default is to expose no metric endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"min_port": schema.Int64Attribute{
						Description:         "MinRelayPort is the smallest relay port assigned for STUNner relay connections.",
						MarkdownDescription: "MinRelayPort is the smallest relay port assigned for STUNner relay connections.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"password": schema.StringAttribute{
						Description:         "Password defines the 'password' credential for 'plaintext' authentication.",
						MarkdownDescription: "Password defines the 'password' credential for 'plaintext' authentication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
						},
					},

					"realm": schema.StringAttribute{
						Description:         "Realm defines the STUN/TURN authentication realm to be used for clients toauthenticate with STUNner.  The realm must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
						MarkdownDescription: "Realm defines the STUN/TURN authentication realm to be used for clients toauthenticate with STUNner.  The realm must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
						},
					},

					"shared_secret": schema.StringAttribute{
						Description:         "SharedSecret defines the shared secret to be used for 'longterm' authentication.",
						MarkdownDescription: "SharedSecret defines the shared secret to be used for 'longterm' authentication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stunner_config": schema.StringAttribute{
						Description:         "StunnerConfig specifies the name of the ConfigMap into which the operator renders the stunnerd configfile.",
						MarkdownDescription: "StunnerConfig specifies the name of the ConfigMap into which the operator renders the stunnerd configfile.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(64),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
						},
					},

					"user_name": schema.StringAttribute{
						Description:         "Username defines the 'username' credential for 'plaintext' authentication.",
						MarkdownDescription: "Username defines the 'username' credential for 'plaintext' authentication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
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

func (r *StunnerL7MpIoGatewayConfigV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *StunnerL7MpIoGatewayConfigV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_stunner_l7mp_io_gateway_config_v1alpha1")

	var model StunnerL7MpIoGatewayConfigV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("stunner.l7mp.io/v1alpha1")
	model.Kind = pointer.String("GatewayConfig")

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
		Resource(k8sSchema.GroupVersionResource{Group: "stunner.l7mp.io", Version: "v1alpha1", Resource: "gatewayconfigs"}).
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

	var readResponse StunnerL7MpIoGatewayConfigV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *StunnerL7MpIoGatewayConfigV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_stunner_l7mp_io_gateway_config_v1alpha1")

	var data StunnerL7MpIoGatewayConfigV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "stunner.l7mp.io", Version: "v1alpha1", Resource: "gatewayconfigs"}).
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

	var readResponse StunnerL7MpIoGatewayConfigV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *StunnerL7MpIoGatewayConfigV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_stunner_l7mp_io_gateway_config_v1alpha1")

	var model StunnerL7MpIoGatewayConfigV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("stunner.l7mp.io/v1alpha1")
	model.Kind = pointer.String("GatewayConfig")

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
		Resource(k8sSchema.GroupVersionResource{Group: "stunner.l7mp.io", Version: "v1alpha1", Resource: "gatewayconfigs"}).
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

	var readResponse StunnerL7MpIoGatewayConfigV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *StunnerL7MpIoGatewayConfigV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_stunner_l7mp_io_gateway_config_v1alpha1")

	var data StunnerL7MpIoGatewayConfigV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "stunner.l7mp.io", Version: "v1alpha1", Resource: "gatewayconfigs"}).
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
				Resource(k8sSchema.GroupVersionResource{Group: "stunner.l7mp.io", Version: "v1alpha1", Resource: "gatewayconfigs"}).
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

func (r *StunnerL7MpIoGatewayConfigV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
