/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package image_toolkit_fluxcd_io_v1beta2

import (
	"context"
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
	_ datasource.DataSource = &ImageToolkitFluxcdIoImageRepositoryV1Beta2Manifest{}
)

func NewImageToolkitFluxcdIoImageRepositoryV1Beta2Manifest() datasource.DataSource {
	return &ImageToolkitFluxcdIoImageRepositoryV1Beta2Manifest{}
}

type ImageToolkitFluxcdIoImageRepositoryV1Beta2Manifest struct{}

type ImageToolkitFluxcdIoImageRepositoryV1Beta2ManifestData struct {
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
		AccessFrom *struct {
			NamespaceSelectors *[]struct {
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selectors" json:"namespaceSelectors,omitempty"`
		} `tfsdk:"access_from" json:"accessFrom,omitempty"`
		CertSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
		ExclusionList  *[]string `tfsdk:"exclusion_list" json:"exclusionList,omitempty"`
		Image          *string   `tfsdk:"image" json:"image,omitempty"`
		Insecure       *bool     `tfsdk:"insecure" json:"insecure,omitempty"`
		Interval       *string   `tfsdk:"interval" json:"interval,omitempty"`
		Provider       *string   `tfsdk:"provider" json:"provider,omitempty"`
		ProxySecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"proxy_secret_ref" json:"proxySecretRef,omitempty"`
		SecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		Suspend            *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		Timeout            *string `tfsdk:"timeout" json:"timeout,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ImageToolkitFluxcdIoImageRepositoryV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_image_toolkit_fluxcd_io_image_repository_v1beta2_manifest"
}

func (r *ImageToolkitFluxcdIoImageRepositoryV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ImageRepository is the Schema for the imagerepositories API",
		MarkdownDescription: "ImageRepository is the Schema for the imagerepositories API",
		Attributes: map[string]schema.Attribute{
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
				Description:         "ImageRepositorySpec defines the parameters for scanning an imagerepository, e.g., 'fluxcd/flux'.",
				MarkdownDescription: "ImageRepositorySpec defines the parameters for scanning an imagerepository, e.g., 'fluxcd/flux'.",
				Attributes: map[string]schema.Attribute{
					"access_from": schema.SingleNestedAttribute{
						Description:         "AccessFrom defines an ACL for allowing cross-namespace referencesto the ImageRepository object based on the caller's namespace labels.",
						MarkdownDescription: "AccessFrom defines an ACL for allowing cross-namespace referencesto the ImageRepository object based on the caller's namespace labels.",
						Attributes: map[string]schema.Attribute{
							"namespace_selectors": schema.ListNestedAttribute{
								Description:         "NamespaceSelectors is the list of namespace selectors to which this ACL applies.Items in this list are evaluated using a logical OR operation.",
								MarkdownDescription: "NamespaceSelectors is the list of namespace selectors to which this ACL applies.Items in this list are evaluated using a logical OR operation.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match_labels": schema.MapAttribute{
											Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cert_secret_ref": schema.SingleNestedAttribute{
						Description:         "CertSecretRef can be given the name of a Secret containingeither or both of- a PEM-encoded client certificate ('tls.crt') and privatekey ('tls.key');- a PEM-encoded CA certificate ('ca.crt')and whichever are supplied, will be used for connecting to theregistry. The client cert and key are useful if you areauthenticating with a certificate; the CA cert is useful ifyou are using a self-signed server certificate. The Secret mustbe of type 'Opaque' or 'kubernetes.io/tls'.Note: Support for the 'caFile', 'certFile' and 'keyFile' keys hasbeen deprecated.",
						MarkdownDescription: "CertSecretRef can be given the name of a Secret containingeither or both of- a PEM-encoded client certificate ('tls.crt') and privatekey ('tls.key');- a PEM-encoded CA certificate ('ca.crt')and whichever are supplied, will be used for connecting to theregistry. The client cert and key are useful if you areauthenticating with a certificate; the CA cert is useful ifyou are using a self-signed server certificate. The Secret mustbe of type 'Opaque' or 'kubernetes.io/tls'.Note: Support for the 'caFile', 'certFile' and 'keyFile' keys hasbeen deprecated.",
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

					"exclusion_list": schema.ListAttribute{
						Description:         "ExclusionList is a list of regex strings used to exclude certain tagsfrom being stored in the database.",
						MarkdownDescription: "ExclusionList is a list of regex strings used to exclude certain tagsfrom being stored in the database.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Image is the name of the image repository",
						MarkdownDescription: "Image is the name of the image repository",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"insecure": schema.BoolAttribute{
						Description:         "Insecure allows connecting to a non-TLS HTTP container registry.",
						MarkdownDescription: "Insecure allows connecting to a non-TLS HTTP container registry.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interval": schema.StringAttribute{
						Description:         "Interval is the length of time to wait betweenscans of the image repository.",
						MarkdownDescription: "Interval is the length of time to wait betweenscans of the image repository.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"provider": schema.StringAttribute{
						Description:         "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'.When not specified, defaults to 'generic'.",
						MarkdownDescription: "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'.When not specified, defaults to 'generic'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("generic", "aws", "azure", "gcp"),
						},
					},

					"proxy_secret_ref": schema.SingleNestedAttribute{
						Description:         "ProxySecretRef specifies the Secret containing the proxy configurationto use while communicating with the container registry.",
						MarkdownDescription: "ProxySecretRef specifies the Secret containing the proxy configurationto use while communicating with the container registry.",
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

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretRef can be given the name of a secret containingcredentials to use for the image registry. The secret should becreated with 'kubectl create secret docker-registry', or theequivalent.",
						MarkdownDescription: "SecretRef can be given the name of a secret containingcredentials to use for the image registry. The secret should becreated with 'kubectl create secret docker-registry', or theequivalent.",
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

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticatethe image pull if the service account has attached pull secrets.",
						MarkdownDescription: "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticatethe image pull if the service account has attached pull secrets.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
						},
					},

					"suspend": schema.BoolAttribute{
						Description:         "This flag tells the controller to suspend subsequent image scans.It does not apply to already started scans. Defaults to false.",
						MarkdownDescription: "This flag tells the controller to suspend subsequent image scans.It does not apply to already started scans. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout for image scanning.Defaults to 'Interval' duration.",
						MarkdownDescription: "Timeout for image scanning.Defaults to 'Interval' duration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
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

func (r *ImageToolkitFluxcdIoImageRepositoryV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_image_toolkit_fluxcd_io_image_repository_v1beta2_manifest")

	var model ImageToolkitFluxcdIoImageRepositoryV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("image.toolkit.fluxcd.io/v1beta2")
	model.Kind = pointer.String("ImageRepository")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
