/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_networking_k8s_io_v1beta1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &GatewayNetworkingK8SIoReferenceGrantV1Beta1Manifest{}
)

func NewGatewayNetworkingK8SIoReferenceGrantV1Beta1Manifest() datasource.DataSource {
	return &GatewayNetworkingK8SIoReferenceGrantV1Beta1Manifest{}
}

type GatewayNetworkingK8SIoReferenceGrantV1Beta1Manifest struct{}

type GatewayNetworkingK8SIoReferenceGrantV1Beta1ManifestData struct {
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
		From *[]struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"from" json:"from,omitempty"`
		To *[]struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"to" json:"to,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingK8SIoReferenceGrantV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_k8s_io_reference_grant_v1beta1_manifest"
}

func (r *GatewayNetworkingK8SIoReferenceGrantV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ReferenceGrant identifies kinds of resources in other namespaces that aretrusted to reference the specified kinds of resources in the same namespaceas the policy.Each ReferenceGrant can be used to represent a unique trust relationship.Additional Reference Grants can be used to add to the set of trustedsources of inbound references for the namespace they are defined within.All cross-namespace references in Gateway API (with the exception of cross-namespaceGateway-route attachment) require a ReferenceGrant.ReferenceGrant is a form of runtime verification allowing users to assertwhich cross-namespace object references are permitted. Implementations thatsupport ReferenceGrant MUST NOT permit cross-namespace references which haveno grant, and MUST respond to the removal of a grant by revoking the accessthat the grant allowed.",
		MarkdownDescription: "ReferenceGrant identifies kinds of resources in other namespaces that aretrusted to reference the specified kinds of resources in the same namespaceas the policy.Each ReferenceGrant can be used to represent a unique trust relationship.Additional Reference Grants can be used to add to the set of trustedsources of inbound references for the namespace they are defined within.All cross-namespace references in Gateway API (with the exception of cross-namespaceGateway-route attachment) require a ReferenceGrant.ReferenceGrant is a form of runtime verification allowing users to assertwhich cross-namespace object references are permitted. Implementations thatsupport ReferenceGrant MUST NOT permit cross-namespace references which haveno grant, and MUST respond to the removal of a grant by revoking the accessthat the grant allowed.",
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
				Description:         "Spec defines the desired state of ReferenceGrant.",
				MarkdownDescription: "Spec defines the desired state of ReferenceGrant.",
				Attributes: map[string]schema.Attribute{
					"from": schema.ListNestedAttribute{
						Description:         "From describes the trusted namespaces and kinds that can reference theresources described in 'To'. Each entry in this list MUST be consideredto be an additional place that references can be valid from, or to putthis another way, entries MUST be combined using OR.Support: Core",
						MarkdownDescription: "From describes the trusted namespaces and kinds that can reference theresources described in 'To'. Each entry in this list MUST be consideredto be an additional place that references can be valid from, or to putthis another way, entries MUST be combined using OR.Support: Core",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group is the group of the referent.When empty, the Kubernetes core API group is inferred.Support: Core",
									MarkdownDescription: "Group is the group of the referent.When empty, the Kubernetes core API group is inferred.Support: Core",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is the kind of the referent. Although implementations may supportadditional resources, the following types are part of the 'Core'support level for this field.When used to permit a SecretObjectReference:* GatewayWhen used to permit a BackendObjectReference:* GRPCRoute* HTTPRoute* TCPRoute* TLSRoute* UDPRoute",
									MarkdownDescription: "Kind is the kind of the referent. Although implementations may supportadditional resources, the following types are part of the 'Core'support level for this field.When used to permit a SecretObjectReference:* GatewayWhen used to permit a BackendObjectReference:* GRPCRoute* HTTPRoute* TCPRoute* TLSRoute* UDPRoute",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
									},
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace is the namespace of the referent.Support: Core",
									MarkdownDescription: "Namespace is the namespace of the referent.Support: Core",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"to": schema.ListNestedAttribute{
						Description:         "To describes the resources that may be referenced by the resourcesdescribed in 'From'. Each entry in this list MUST be considered to be anadditional place that references can be valid to, or to put this anotherway, entries MUST be combined using OR.Support: Core",
						MarkdownDescription: "To describes the resources that may be referenced by the resourcesdescribed in 'From'. Each entry in this list MUST be considered to be anadditional place that references can be valid to, or to put this anotherway, entries MUST be combined using OR.Support: Core",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group is the group of the referent.When empty, the Kubernetes core API group is inferred.Support: Core",
									MarkdownDescription: "Group is the group of the referent.When empty, the Kubernetes core API group is inferred.Support: Core",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is the kind of the referent. Although implementations may supportadditional resources, the following types are part of the 'Core'support level for this field:* Secret when used to permit a SecretObjectReference* Service when used to permit a BackendObjectReference",
									MarkdownDescription: "Kind is the kind of the referent. Although implementations may supportadditional resources, the following types are part of the 'Core'support level for this field:* Secret when used to permit a SecretObjectReference* Service when used to permit a BackendObjectReference",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the referent. When unspecified, this policyrefers to all resources of the specified Group and Kind in the localnamespace.",
									MarkdownDescription: "Name is the name of the referent. When unspecified, this policyrefers to all resources of the specified Group and Kind in the localnamespace.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
									},
								},
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
		},
	}
}

func (r *GatewayNetworkingK8SIoReferenceGrantV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_k8s_io_reference_grant_v1beta1_manifest")

	var model GatewayNetworkingK8SIoReferenceGrantV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("gateway.networking.k8s.io/v1beta1")
	model.Kind = pointer.String("ReferenceGrant")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
