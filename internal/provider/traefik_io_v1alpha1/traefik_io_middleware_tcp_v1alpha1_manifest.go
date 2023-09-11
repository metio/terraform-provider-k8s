/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &TraefikIoMiddlewareTcpV1Alpha1Manifest{}
)

func NewTraefikIoMiddlewareTcpV1Alpha1Manifest() datasource.DataSource {
	return &TraefikIoMiddlewareTcpV1Alpha1Manifest{}
}

type TraefikIoMiddlewareTcpV1Alpha1Manifest struct{}

type TraefikIoMiddlewareTcpV1Alpha1ManifestData struct {
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
		InFlightConn *struct {
			Amount *int64 `tfsdk:"amount" json:"amount,omitempty"`
		} `tfsdk:"in_flight_conn" json:"inFlightConn,omitempty"`
		IpAllowList *struct {
			SourceRange *[]string `tfsdk:"source_range" json:"sourceRange,omitempty"`
		} `tfsdk:"ip_allow_list" json:"ipAllowList,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoMiddlewareTcpV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_middleware_tcp_v1alpha1_manifest"
}

func (r *TraefikIoMiddlewareTcpV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MiddlewareTCP is the CRD implementation of a Traefik TCP middleware. More info: https://doc.traefik.io/traefik/v3.0/middlewares/overview/",
		MarkdownDescription: "MiddlewareTCP is the CRD implementation of a Traefik TCP middleware. More info: https://doc.traefik.io/traefik/v3.0/middlewares/overview/",
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
				Description:         "MiddlewareTCPSpec defines the desired state of a MiddlewareTCP.",
				MarkdownDescription: "MiddlewareTCPSpec defines the desired state of a MiddlewareTCP.",
				Attributes: map[string]schema.Attribute{
					"in_flight_conn": schema.SingleNestedAttribute{
						Description:         "InFlightConn defines the InFlightConn middleware configuration.",
						MarkdownDescription: "InFlightConn defines the InFlightConn middleware configuration.",
						Attributes: map[string]schema.Attribute{
							"amount": schema.Int64Attribute{
								Description:         "Amount defines the maximum amount of allowed simultaneous connections. The middleware closes the connection if there are already amount connections opened.",
								MarkdownDescription: "Amount defines the maximum amount of allowed simultaneous connections. The middleware closes the connection if there are already amount connections opened.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ip_allow_list": schema.SingleNestedAttribute{
						Description:         "IPAllowList defines the IPAllowList middleware configuration.",
						MarkdownDescription: "IPAllowList defines the IPAllowList middleware configuration.",
						Attributes: map[string]schema.Attribute{
							"source_range": schema.ListAttribute{
								Description:         "SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).",
								MarkdownDescription: "SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *TraefikIoMiddlewareTcpV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_io_middleware_tcp_v1alpha1_manifest")

	var model TraefikIoMiddlewareTcpV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("traefik.io/v1alpha1")
	model.Kind = pointer.String("MiddlewareTCP")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
