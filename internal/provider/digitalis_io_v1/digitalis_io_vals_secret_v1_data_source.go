/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package digitalis_io_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &DigitalisIoValsSecretV1DataSource{}
	_ datasource.DataSourceWithConfigure = &DigitalisIoValsSecretV1DataSource{}
)

func NewDigitalisIoValsSecretV1DataSource() datasource.DataSource {
	return &DigitalisIoValsSecretV1DataSource{}
}

type DigitalisIoValsSecretV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type DigitalisIoValsSecretV1DataSourceData struct {
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
		Data *struct {
			Encoding *string `tfsdk:"encoding" json:"encoding,omitempty"`
			Ref      *string `tfsdk:"ref" json:"ref,omitempty"`
		} `tfsdk:"data" json:"data,omitempty"`
		Databases *[]struct {
			Driver           *string   `tfsdk:"driver" json:"driver,omitempty"`
			Hosts            *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			LoginCredentials *struct {
				Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"login_credentials" json:"loginCredentials,omitempty"`
			PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
			UserHost    *string `tfsdk:"user_host" json:"userHost,omitempty"`
			UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
		} `tfsdk:"databases" json:"databases,omitempty"`
		Name     *string            `tfsdk:"name" json:"name,omitempty"`
		Template *map[string]string `tfsdk:"template" json:"template,omitempty"`
		Ttl      *int64             `tfsdk:"ttl" json:"ttl,omitempty"`
		Type     *string            `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DigitalisIoValsSecretV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_digitalis_io_vals_secret_v1"
}

func (r *DigitalisIoValsSecretV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ValsSecret is the Schema for the valssecrets API",
		MarkdownDescription: "ValsSecret is the Schema for the valssecrets API",
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
				Description:         "ValsSecretSpec defines the desired state of ValsSecret",
				MarkdownDescription: "ValsSecretSpec defines the desired state of ValsSecret",
				Attributes: map[string]schema.Attribute{
					"data": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"encoding": schema.StringAttribute{
								Description:         "Encoding type for the secret. Only base64 supported. Optional",
								MarkdownDescription: "Encoding type for the secret. Only base64 supported. Optional",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ref": schema.StringAttribute{
								Description:         "Ref value to the secret in the format ref+backend://path https://github.com/helmfile/vals",
								MarkdownDescription: "Ref value to the secret in the format ref+backend://path https://github.com/helmfile/vals",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"databases": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"driver": schema.StringAttribute{
									Description:         "Defines the database type",
									MarkdownDescription: "Defines the database type",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"hosts": schema.ListAttribute{
									Description:         "List of hosts to connect to, they'll be tried in sequence until one succeeds",
									MarkdownDescription: "List of hosts to connect to, they'll be tried in sequence until one succeeds",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"login_credentials": schema.SingleNestedAttribute{
									Description:         "Credentials to access the database",
									MarkdownDescription: "Credentials to access the database",
									Attributes: map[string]schema.Attribute{
										"namespace": schema.StringAttribute{
											Description:         "Optional namespace of the secret, default current namespace",
											MarkdownDescription: "Optional namespace of the secret, default current namespace",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"password_key": schema.StringAttribute{
											Description:         "Key in the secret containing the database username",
											MarkdownDescription: "Key in the secret containing the database username",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Name of the secret containing the credentials to be able to log in to the database",
											MarkdownDescription: "Name of the secret containing the credentials to be able to log in to the database",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"username_key": schema.StringAttribute{
											Description:         "Key in the secret containing the database username",
											MarkdownDescription: "Key in the secret containing the database username",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"password_key": schema.StringAttribute{
									Description:         "Key in the secret containing the database username",
									MarkdownDescription: "Key in the secret containing the database username",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"port": schema.Int64Attribute{
									Description:         "Database port number",
									MarkdownDescription: "Database port number",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"user_host": schema.StringAttribute{
									Description:         "Used for MySQL only, the host part for the username",
									MarkdownDescription: "Used for MySQL only, the host part for the username",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"username_key": schema.StringAttribute{
									Description:         "Key in the secret containing the database username",
									MarkdownDescription: "Key in the secret containing the database username",
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

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"template": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ttl": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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
	}
}

func (r *DigitalisIoValsSecretV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *DigitalisIoValsSecretV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_digitalis_io_vals_secret_v1")

	var data DigitalisIoValsSecretV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "digitalis.io", Version: "v1", Resource: "valssecrets"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse DigitalisIoValsSecretV1DataSourceData
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
	data.ApiVersion = pointer.String("digitalis.io/v1")
	data.Kind = pointer.String("ValsSecret")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
