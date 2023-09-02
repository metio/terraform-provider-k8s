/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_istio_io_v1

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
	_ datasource.DataSource              = &SecurityIstioIoRequestAuthenticationV1DataSource{}
	_ datasource.DataSourceWithConfigure = &SecurityIstioIoRequestAuthenticationV1DataSource{}
)

func NewSecurityIstioIoRequestAuthenticationV1DataSource() datasource.DataSource {
	return &SecurityIstioIoRequestAuthenticationV1DataSource{}
}

type SecurityIstioIoRequestAuthenticationV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SecurityIstioIoRequestAuthenticationV1DataSourceData struct {
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
		JwtRules *[]struct {
			Audiences            *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
			ForwardOriginalToken *bool     `tfsdk:"forward_original_token" json:"forwardOriginalToken,omitempty"`
			FromHeaders          *[]struct {
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
			} `tfsdk:"from_headers" json:"fromHeaders,omitempty"`
			FromParams           *[]string `tfsdk:"from_params" json:"fromParams,omitempty"`
			Issuer               *string   `tfsdk:"issuer" json:"issuer,omitempty"`
			Jwks                 *string   `tfsdk:"jwks" json:"jwks,omitempty"`
			JwksUri              *string   `tfsdk:"jwks_uri" json:"jwksUri,omitempty"`
			Jwks_uri             *string   `tfsdk:"jwks_uri" json:"jwks_uri,omitempty"`
			OutputClaimToHeaders *[]struct {
				Claim  *string `tfsdk:"claim" json:"claim,omitempty"`
				Header *string `tfsdk:"header" json:"header,omitempty"`
			} `tfsdk:"output_claim_to_headers" json:"outputClaimToHeaders,omitempty"`
			OutputPayloadToHeader *string `tfsdk:"output_payload_to_header" json:"outputPayloadToHeader,omitempty"`
		} `tfsdk:"jwt_rules" json:"jwtRules,omitempty"`
		Selector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecurityIstioIoRequestAuthenticationV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_istio_io_request_authentication_v1"
}

func (r *SecurityIstioIoRequestAuthenticationV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "RequestAuthentication defines what request authentication methods are supported by a workload.",
				MarkdownDescription: "RequestAuthentication defines what request authentication methods are supported by a workload.",
				Attributes: map[string]schema.Attribute{
					"jwt_rules": schema.ListNestedAttribute{
						Description:         "Define the list of JWTs that can be validated at the selected workloads' proxy.",
						MarkdownDescription: "Define the list of JWTs that can be validated at the selected workloads' proxy.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"audiences": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"forward_original_token": schema.BoolAttribute{
									Description:         "If set to true, the original token will be kept for the upstream request.",
									MarkdownDescription: "If set to true, the original token will be kept for the upstream request.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"from_headers": schema.ListNestedAttribute{
									Description:         "List of header locations from which JWT is expected.",
									MarkdownDescription: "List of header locations from which JWT is expected.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "The HTTP header name.",
												MarkdownDescription: "The HTTP header name.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"prefix": schema.StringAttribute{
												Description:         "The prefix that should be stripped before decoding the token.",
												MarkdownDescription: "The prefix that should be stripped before decoding the token.",
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

								"from_params": schema.ListAttribute{
									Description:         "List of query parameters from which JWT is expected.",
									MarkdownDescription: "List of query parameters from which JWT is expected.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"issuer": schema.StringAttribute{
									Description:         "Identifies the issuer that issued the JWT.",
									MarkdownDescription: "Identifies the issuer that issued the JWT.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"jwks": schema.StringAttribute{
									Description:         "JSON Web Key Set of public keys to validate signature of the JWT.",
									MarkdownDescription: "JSON Web Key Set of public keys to validate signature of the JWT.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"jwks_uri": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"jwks_uri": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"output_claim_to_headers": schema.ListNestedAttribute{
									Description:         "This field specifies a list of operations to copy the claim to HTTP headers on a successfully verified token.",
									MarkdownDescription: "This field specifies a list of operations to copy the claim to HTTP headers on a successfully verified token.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"claim": schema.StringAttribute{
												Description:         "The name of the claim to be copied from.",
												MarkdownDescription: "The name of the claim to be copied from.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"header": schema.StringAttribute{
												Description:         "The name of the header to be created.",
												MarkdownDescription: "The name of the header to be created.",
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

								"output_payload_to_header": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"selector": schema.SingleNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SecurityIstioIoRequestAuthenticationV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SecurityIstioIoRequestAuthenticationV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_security_istio_io_request_authentication_v1")

	var data SecurityIstioIoRequestAuthenticationV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "security.istio.io", Version: "v1", Resource: "RequestAuthentication"}).
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

	var readResponse SecurityIstioIoRequestAuthenticationV1DataSourceData
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
	data.ApiVersion = pointer.String("security.istio.io/v1")
	data.Kind = pointer.String("RequestAuthentication")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
