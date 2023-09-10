/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_istio_io_v1beta1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SecurityIstioIoRequestAuthenticationV1Beta1Manifest{}
)

func NewSecurityIstioIoRequestAuthenticationV1Beta1Manifest() datasource.DataSource {
	return &SecurityIstioIoRequestAuthenticationV1Beta1Manifest{}
}

type SecurityIstioIoRequestAuthenticationV1Beta1Manifest struct{}

type SecurityIstioIoRequestAuthenticationV1Beta1ManifestData struct {
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
			OutputClaimToHeaders *[]struct {
				Claim  *string `tfsdk:"claim" json:"claim,omitempty"`
				Header *string `tfsdk:"header" json:"header,omitempty"`
			} `tfsdk:"output_claim_to_headers" json:"outputClaimToHeaders,omitempty"`
			OutputPayloadToHeader *string `tfsdk:"output_payload_to_header" json:"outputPayloadToHeader,omitempty"`
		} `tfsdk:"jwt_rules" json:"jwtRules,omitempty"`
		Selector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecurityIstioIoRequestAuthenticationV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_istio_io_request_authentication_v1beta1_manifest"
}

func (r *SecurityIstioIoRequestAuthenticationV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
									Optional:            true,
									Computed:            false,
								},

								"forward_original_token": schema.BoolAttribute{
									Description:         "If set to true, the original token will be kept for the upstream request.",
									MarkdownDescription: "If set to true, the original token will be kept for the upstream request.",
									Required:            false,
									Optional:            true,
									Computed:            false,
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
												Optional:            true,
												Computed:            false,
											},

											"prefix": schema.StringAttribute{
												Description:         "The prefix that should be stripped before decoding the token.",
												MarkdownDescription: "The prefix that should be stripped before decoding the token.",
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

								"from_params": schema.ListAttribute{
									Description:         "List of query parameters from which JWT is expected.",
									MarkdownDescription: "List of query parameters from which JWT is expected.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"issuer": schema.StringAttribute{
									Description:         "Identifies the issuer that issued the JWT.",
									MarkdownDescription: "Identifies the issuer that issued the JWT.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"jwks": schema.StringAttribute{
									Description:         "JSON Web Key Set of public keys to validate signature of the JWT.",
									MarkdownDescription: "JSON Web Key Set of public keys to validate signature of the JWT.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"jwks_uri": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
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
												Optional:            true,
												Computed:            false,
											},

											"header": schema.StringAttribute{
												Description:         "The name of the header to be created.",
												MarkdownDescription: "The name of the header to be created.",
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

								"output_payload_to_header": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"selector": schema.SingleNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "group is the group of the target resource.",
								MarkdownDescription: "group is the group of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is kind of the target resource.",
								MarkdownDescription: "kind is kind of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the target resource.",
								MarkdownDescription: "name is the name of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace is the namespace of the referent.",
								MarkdownDescription: "namespace is the namespace of the referent.",
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

func (r *SecurityIstioIoRequestAuthenticationV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_istio_io_request_authentication_v1beta1_manifest")

	var model SecurityIstioIoRequestAuthenticationV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("security.istio.io/v1beta1")
	model.Kind = pointer.String("RequestAuthentication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
