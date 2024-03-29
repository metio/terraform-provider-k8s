/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package bpfd_dev_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &BpfdDevXdpProgramV1Alpha1Manifest{}
)

func NewBpfdDevXdpProgramV1Alpha1Manifest() datasource.DataSource {
	return &BpfdDevXdpProgramV1Alpha1Manifest{}
}

type BpfdDevXdpProgramV1Alpha1Manifest struct{}

type BpfdDevXdpProgramV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Bpffunctionname *string `tfsdk:"bpffunctionname" json:"bpffunctionname,omitempty"`
		Bytecode        *struct {
			Image *struct {
				Imagepullpolicy *string `tfsdk:"imagepullpolicy" json:"imagepullpolicy,omitempty"`
				Imagepullsecret *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"imagepullsecret" json:"imagepullsecret,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			Path *string `tfsdk:"path" json:"path,omitempty"`
		} `tfsdk:"bytecode" json:"bytecode,omitempty"`
		Globaldata        *map[string]string `tfsdk:"globaldata" json:"globaldata,omitempty"`
		Interfaceselector *struct {
			Interfaces           *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
			Primarynodeinterface *bool     `tfsdk:"primarynodeinterface" json:"primarynodeinterface,omitempty"`
		} `tfsdk:"interfaceselector" json:"interfaceselector,omitempty"`
		Mapownerselector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"mapownerselector" json:"mapownerselector,omitempty"`
		Nodeselector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"nodeselector" json:"nodeselector,omitempty"`
		Priority  *int64    `tfsdk:"priority" json:"priority,omitempty"`
		Proceedon *[]string `tfsdk:"proceedon" json:"proceedon,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BpfdDevXdpProgramV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_bpfd_dev_xdp_program_v1alpha1_manifest"
}

func (r *BpfdDevXdpProgramV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "XdpProgram is the Schema for the XdpPrograms API",
		MarkdownDescription: "XdpProgram is the Schema for the XdpPrograms API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "XdpProgramSpec defines the desired state of XdpProgram",
				MarkdownDescription: "XdpProgramSpec defines the desired state of XdpProgram",
				Attributes: map[string]schema.Attribute{
					"bpffunctionname": schema.StringAttribute{
						Description:         "BpfFunctionName is the name of the function that is the entry point for the BPF program",
						MarkdownDescription: "BpfFunctionName is the name of the function that is the entry point for the BPF program",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"bytecode": schema.SingleNestedAttribute{
						Description:         "Bytecode configures where the bpf program's bytecode should be loaded from.",
						MarkdownDescription: "Bytecode configures where the bpf program's bytecode should be loaded from.",
						Attributes: map[string]schema.Attribute{
							"image": schema.SingleNestedAttribute{
								Description:         "Image used to specify a bytecode container image.",
								MarkdownDescription: "Image used to specify a bytecode container image.",
								Attributes: map[string]schema.Attribute{
									"imagepullpolicy": schema.StringAttribute{
										Description:         "PullPolicy describes a policy for if/when to pull a bytecode image. Defaults to IfNotPresent.",
										MarkdownDescription: "PullPolicy describes a policy for if/when to pull a bytecode image. Defaults to IfNotPresent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
										},
									},

									"imagepullsecret": schema.SingleNestedAttribute{
										Description:         "ImagePullSecret is the name of the secret bpfd should use to get remote image repository secrets.",
										MarkdownDescription: "ImagePullSecret is the name of the secret bpfd should use to get remote image repository secrets.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the secret which contains the credentials to access the image repository.",
												MarkdownDescription: "Name of the secret which contains the credentials to access the image repository.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the secret which contains the credentials to access the image repository.",
												MarkdownDescription: "Namespace of the secret which contains the credentials to access the image repository.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": schema.StringAttribute{
										Description:         "Valid container image URL used to reference a remote bytecode image.",
										MarkdownDescription: "Valid container image URL used to reference a remote bytecode image.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": schema.StringAttribute{
								Description:         "Path is used to specify a bytecode object via filepath.",
								MarkdownDescription: "Path is used to specify a bytecode object via filepath.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"globaldata": schema.MapAttribute{
						Description:         "GlobalData allows the user to to set global variables when the program is loaded with an array of raw bytes. This is a very low level primitive. The caller is responsible for formatting the byte string appropriately considering such things as size, endianness, alignment and packing of data structures.",
						MarkdownDescription: "GlobalData allows the user to to set global variables when the program is loaded with an array of raw bytes. This is a very low level primitive. The caller is responsible for formatting the byte string appropriately considering such things as size, endianness, alignment and packing of data structures.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interfaceselector": schema.SingleNestedAttribute{
						Description:         "Selector to determine the network interface (or interfaces)",
						MarkdownDescription: "Selector to determine the network interface (or interfaces)",
						Attributes: map[string]schema.Attribute{
							"interfaces": schema.ListAttribute{
								Description:         "Interfaces refers to a list of network interfaces to attach the BPF program to.",
								MarkdownDescription: "Interfaces refers to a list of network interfaces to attach the BPF program to.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.List{
									listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("primarynodeinterface")),
								},
							},

							"primarynodeinterface": schema.BoolAttribute{
								Description:         "Attach BPF program to the primary interface on the node. Only 'true' accepted.",
								MarkdownDescription: "Attach BPF program to the primary interface on the node. Only 'true' accepted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Bool{
									boolvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("interfaces")),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"mapownerselector": schema.SingleNestedAttribute{
						Description:         "MapOwnerSelector is used to select the loaded eBPF program this eBPF program will share a map with. The value is a label applied to the BpfProgram to select. The selector must resolve to exactly one instance of a BpfProgram on a given node or the eBPF program will not load.",
						MarkdownDescription: "MapOwnerSelector is used to select the loaded eBPF program this eBPF program will share a map with. The value is a label applied to the BpfProgram to select. The selector must resolve to exactly one instance of a BpfProgram on a given node or the eBPF program will not load.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"nodeselector": schema.SingleNestedAttribute{
						Description:         "NodeSelector allows the user to specify which nodes to deploy the bpf program to.  This field must be specified, to select all nodes use standard metav1.LabelSelector semantics and make it empty.",
						MarkdownDescription: "NodeSelector allows the user to specify which nodes to deploy the bpf program to.  This field must be specified, to select all nodes use standard metav1.LabelSelector semantics and make it empty.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority specifies the priority of the bpf program in relation to other programs of the same type with the same attach point. It is a value from 0 to 1000 where lower values have higher precedence.",
						MarkdownDescription: "Priority specifies the priority of the bpf program in relation to other programs of the same type with the same attach point. It is a value from 0 to 1000 where lower values have higher precedence.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(1000),
						},
					},

					"proceedon": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
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
	}
}

func (r *BpfdDevXdpProgramV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_bpfd_dev_xdp_program_v1alpha1_manifest")

	var model BpfdDevXdpProgramV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("bpfd.dev/v1alpha1")
	model.Kind = pointer.String("XdpProgram")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
