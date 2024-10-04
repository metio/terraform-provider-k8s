/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package trino_stackable_tech_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &TrinoStackableTechTrinoCatalogV1Alpha1Manifest{}
)

func NewTrinoStackableTechTrinoCatalogV1Alpha1Manifest() datasource.DataSource {
	return &TrinoStackableTechTrinoCatalogV1Alpha1Manifest{}
}

type TrinoStackableTechTrinoCatalogV1Alpha1Manifest struct{}

type TrinoStackableTechTrinoCatalogV1Alpha1ManifestData struct {
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
		ConfigOverrides *map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
		Connector       *struct {
			BlackHole *map[string]string `tfsdk:"black_hole" json:"blackHole,omitempty"`
			DeltaLake *struct {
				Hdfs *struct {
					ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
				} `tfsdk:"hdfs" json:"hdfs,omitempty"`
				Metastore *struct {
					ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
				} `tfsdk:"metastore" json:"metastore,omitempty"`
				S3 *struct {
					Inline *struct {
						AccessStyle *string `tfsdk:"access_style" json:"accessStyle,omitempty"`
						Credentials *struct {
							Scope *struct {
								ListenerVolumes *[]string `tfsdk:"listener_volumes" json:"listenerVolumes,omitempty"`
								Node            *bool     `tfsdk:"node" json:"node,omitempty"`
								Pod             *bool     `tfsdk:"pod" json:"pod,omitempty"`
								Services        *[]string `tfsdk:"services" json:"services,omitempty"`
							} `tfsdk:"scope" json:"scope,omitempty"`
							SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
						} `tfsdk:"credentials" json:"credentials,omitempty"`
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *int64  `tfsdk:"port" json:"port,omitempty"`
						Tls  *struct {
							Verification *struct {
								None   *map[string]string `tfsdk:"none" json:"none,omitempty"`
								Server *struct {
									CaCert *struct {
										SecretClass *string            `tfsdk:"secret_class" json:"secretClass,omitempty"`
										WebPki      *map[string]string `tfsdk:"web_pki" json:"webPki,omitempty"`
									} `tfsdk:"ca_cert" json:"caCert,omitempty"`
								} `tfsdk:"server" json:"server,omitempty"`
							} `tfsdk:"verification" json:"verification,omitempty"`
						} `tfsdk:"tls" json:"tls,omitempty"`
					} `tfsdk:"inline" json:"inline,omitempty"`
					Reference *string `tfsdk:"reference" json:"reference,omitempty"`
				} `tfsdk:"s3" json:"s3,omitempty"`
			} `tfsdk:"delta_lake" json:"deltaLake,omitempty"`
			Generic *struct {
				ConnectorName *string `tfsdk:"connector_name" json:"connectorName,omitempty"`
				Properties    *struct {
					Value              *string `tfsdk:"value" json:"value,omitempty"`
					ValueFromConfigMap *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"value_from_config_map" json:"valueFromConfigMap,omitempty"`
					ValueFromSecret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"value_from_secret" json:"valueFromSecret,omitempty"`
				} `tfsdk:"properties" json:"properties,omitempty"`
			} `tfsdk:"generic" json:"generic,omitempty"`
			GoogleSheet *struct {
				Cache *struct {
					SheetsDataExpireAfterWrite *string `tfsdk:"sheets_data_expire_after_write" json:"sheetsDataExpireAfterWrite,omitempty"`
					SheetsDataMaxCacheSize     *string `tfsdk:"sheets_data_max_cache_size" json:"sheetsDataMaxCacheSize,omitempty"`
				} `tfsdk:"cache" json:"cache,omitempty"`
				CredentialsSecret *string `tfsdk:"credentials_secret" json:"credentialsSecret,omitempty"`
				MetadataSheetId   *string `tfsdk:"metadata_sheet_id" json:"metadataSheetId,omitempty"`
			} `tfsdk:"google_sheet" json:"googleSheet,omitempty"`
			Hive *struct {
				Hdfs *struct {
					ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
				} `tfsdk:"hdfs" json:"hdfs,omitempty"`
				Metastore *struct {
					ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
				} `tfsdk:"metastore" json:"metastore,omitempty"`
				S3 *struct {
					Inline *struct {
						AccessStyle *string `tfsdk:"access_style" json:"accessStyle,omitempty"`
						Credentials *struct {
							Scope *struct {
								ListenerVolumes *[]string `tfsdk:"listener_volumes" json:"listenerVolumes,omitempty"`
								Node            *bool     `tfsdk:"node" json:"node,omitempty"`
								Pod             *bool     `tfsdk:"pod" json:"pod,omitempty"`
								Services        *[]string `tfsdk:"services" json:"services,omitempty"`
							} `tfsdk:"scope" json:"scope,omitempty"`
							SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
						} `tfsdk:"credentials" json:"credentials,omitempty"`
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *int64  `tfsdk:"port" json:"port,omitempty"`
						Tls  *struct {
							Verification *struct {
								None   *map[string]string `tfsdk:"none" json:"none,omitempty"`
								Server *struct {
									CaCert *struct {
										SecretClass *string            `tfsdk:"secret_class" json:"secretClass,omitempty"`
										WebPki      *map[string]string `tfsdk:"web_pki" json:"webPki,omitempty"`
									} `tfsdk:"ca_cert" json:"caCert,omitempty"`
								} `tfsdk:"server" json:"server,omitempty"`
							} `tfsdk:"verification" json:"verification,omitempty"`
						} `tfsdk:"tls" json:"tls,omitempty"`
					} `tfsdk:"inline" json:"inline,omitempty"`
					Reference *string `tfsdk:"reference" json:"reference,omitempty"`
				} `tfsdk:"s3" json:"s3,omitempty"`
			} `tfsdk:"hive" json:"hive,omitempty"`
			Iceberg *struct {
				Hdfs *struct {
					ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
				} `tfsdk:"hdfs" json:"hdfs,omitempty"`
				Metastore *struct {
					ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
				} `tfsdk:"metastore" json:"metastore,omitempty"`
				S3 *struct {
					Inline *struct {
						AccessStyle *string `tfsdk:"access_style" json:"accessStyle,omitempty"`
						Credentials *struct {
							Scope *struct {
								ListenerVolumes *[]string `tfsdk:"listener_volumes" json:"listenerVolumes,omitempty"`
								Node            *bool     `tfsdk:"node" json:"node,omitempty"`
								Pod             *bool     `tfsdk:"pod" json:"pod,omitempty"`
								Services        *[]string `tfsdk:"services" json:"services,omitempty"`
							} `tfsdk:"scope" json:"scope,omitempty"`
							SecretClass *string `tfsdk:"secret_class" json:"secretClass,omitempty"`
						} `tfsdk:"credentials" json:"credentials,omitempty"`
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *int64  `tfsdk:"port" json:"port,omitempty"`
						Tls  *struct {
							Verification *struct {
								None   *map[string]string `tfsdk:"none" json:"none,omitempty"`
								Server *struct {
									CaCert *struct {
										SecretClass *string            `tfsdk:"secret_class" json:"secretClass,omitempty"`
										WebPki      *map[string]string `tfsdk:"web_pki" json:"webPki,omitempty"`
									} `tfsdk:"ca_cert" json:"caCert,omitempty"`
								} `tfsdk:"server" json:"server,omitempty"`
							} `tfsdk:"verification" json:"verification,omitempty"`
						} `tfsdk:"tls" json:"tls,omitempty"`
					} `tfsdk:"inline" json:"inline,omitempty"`
					Reference *string `tfsdk:"reference" json:"reference,omitempty"`
				} `tfsdk:"s3" json:"s3,omitempty"`
			} `tfsdk:"iceberg" json:"iceberg,omitempty"`
			Tpcds *map[string]string `tfsdk:"tpcds" json:"tpcds,omitempty"`
			Tpch  *map[string]string `tfsdk:"tpch" json:"tpch,omitempty"`
		} `tfsdk:"connector" json:"connector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TrinoStackableTechTrinoCatalogV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_trino_stackable_tech_trino_catalog_v1alpha1_manifest"
}

func (r *TrinoStackableTechTrinoCatalogV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for TrinoCatalogSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for TrinoCatalogSpec via 'CustomResource'",
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
				Description:         "The TrinoCatalog resource can be used to define catalogs in Kubernetes objects. Read more about it in the [Trino operator concept docs](https://docs.stackable.tech/home/nightly/trino/concepts) and the [Trino operator usage guide](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/). The documentation also contains a list of all the supported backends.",
				MarkdownDescription: "The TrinoCatalog resource can be used to define catalogs in Kubernetes objects. Read more about it in the [Trino operator concept docs](https://docs.stackable.tech/home/nightly/trino/concepts) and the [Trino operator usage guide](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/). The documentation also contains a list of all the supported backends.",
				Attributes: map[string]schema.Attribute{
					"config_overrides": schema.MapAttribute{
						Description:         "The 'configOverrides' allow overriding arbitrary Trino settings. For example, for Hive you could add 'hive.metastore.username: trino'.",
						MarkdownDescription: "The 'configOverrides' allow overriding arbitrary Trino settings. For example, for Hive you could add 'hive.metastore.username: trino'.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"connector": schema.SingleNestedAttribute{
						Description:         "The 'connector' defines which connector is used.",
						MarkdownDescription: "The 'connector' defines which connector is used.",
						Attributes: map[string]schema.Attribute{
							"black_hole": schema.MapAttribute{
								Description:         "A [Black Hole](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/black-hole) connector.",
								MarkdownDescription: "A [Black Hole](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/black-hole) connector.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delta_lake": schema.SingleNestedAttribute{
								Description:         "An [Delta Lake](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/delta-lake) connector.",
								MarkdownDescription: "An [Delta Lake](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/delta-lake) connector.",
								Attributes: map[string]schema.Attribute{
									"hdfs": schema.SingleNestedAttribute{
										Description:         "Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS.",
										MarkdownDescription: "Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS.",
										Attributes: map[string]schema.Attribute{
											"config_map": schema.StringAttribute{
												Description:         "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
												MarkdownDescription: "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"metastore": schema.SingleNestedAttribute{
										Description:         "Mandatory connection to a Hive Metastore, which will be used as a storage for metadata.",
										MarkdownDescription: "Mandatory connection to a Hive Metastore, which will be used as a storage for metadata.",
										Attributes: map[string]schema.Attribute{
											"config_map": schema.StringAttribute{
												Description:         "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.",
												MarkdownDescription: "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"s3": schema.SingleNestedAttribute{
										Description:         "Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
										MarkdownDescription: "Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
										Attributes: map[string]schema.Attribute{
											"inline": schema.SingleNestedAttribute{
												Description:         "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
												MarkdownDescription: "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
												Attributes: map[string]schema.Attribute{
													"access_style": schema.StringAttribute{
														Description:         "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
														MarkdownDescription: "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Path", "VirtualHosted"),
														},
													},

													"credentials": schema.SingleNestedAttribute{
														Description:         "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
														MarkdownDescription: "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
														Attributes: map[string]schema.Attribute{
															"scope": schema.SingleNestedAttribute{
																Description:         "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																MarkdownDescription: "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																Attributes: map[string]schema.Attribute{
																	"listener_volumes": schema.ListAttribute{
																		Description:         "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																		MarkdownDescription: "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"node": schema.BoolAttribute{
																		Description:         "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																		MarkdownDescription: "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pod": schema.BoolAttribute{
																		Description:         "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																		MarkdownDescription: "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"services": schema.ListAttribute{
																		Description:         "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
																		MarkdownDescription: "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
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

															"secret_class": schema.StringAttribute{
																Description:         "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																MarkdownDescription: "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"host": schema.StringAttribute{
														Description:         "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
														MarkdownDescription: "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.Int64Attribute{
														Description:         "Port the S3 server listens on. If not specified the product will determine the port to use.",
														MarkdownDescription: "Port the S3 server listens on. If not specified the product will determine the port to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"tls": schema.SingleNestedAttribute{
														Description:         "Use a TLS connection. If not specified no TLS will be used.",
														MarkdownDescription: "Use a TLS connection. If not specified no TLS will be used.",
														Attributes: map[string]schema.Attribute{
															"verification": schema.SingleNestedAttribute{
																Description:         "The verification method used to verify the certificates of the server and/or the client.",
																MarkdownDescription: "The verification method used to verify the certificates of the server and/or the client.",
																Attributes: map[string]schema.Attribute{
																	"none": schema.MapAttribute{
																		Description:         "Use TLS but don't verify certificates.",
																		MarkdownDescription: "Use TLS but don't verify certificates.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"server": schema.SingleNestedAttribute{
																		Description:         "Use TLS and a CA certificate to verify the server.",
																		MarkdownDescription: "Use TLS and a CA certificate to verify the server.",
																		Attributes: map[string]schema.Attribute{
																			"ca_cert": schema.SingleNestedAttribute{
																				Description:         "CA cert to verify the server.",
																				MarkdownDescription: "CA cert to verify the server.",
																				Attributes: map[string]schema.Attribute{
																					"secret_class": schema.StringAttribute{
																						Description:         "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																						MarkdownDescription: "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"web_pki": schema.MapAttribute{
																						Description:         "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																						MarkdownDescription: "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
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
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"reference": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"generic": schema.SingleNestedAttribute{
								Description:         "A [generic](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/generic) connector.",
								MarkdownDescription: "A [generic](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/generic) connector.",
								Attributes: map[string]schema.Attribute{
									"connector_name": schema.StringAttribute{
										Description:         "Name of the Trino connector. Will be passed to 'connector.name'.",
										MarkdownDescription: "Name of the Trino connector. Will be passed to 'connector.name'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"properties": schema.SingleNestedAttribute{
										Description:         "A map of properties to put in the connector configuration file. They can be specified either as a raw value or be read from a Secret or ConfigMap.",
										MarkdownDescription: "A map of properties to put in the connector configuration file. They can be specified either as a raw value or be read from a Secret or ConfigMap.",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from_config_map": schema.SingleNestedAttribute{
												Description:         "Selects a key from a ConfigMap.",
												MarkdownDescription: "Selects a key from a ConfigMap.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from_secret": schema.SingleNestedAttribute{
												Description:         "SecretKeySelector selects a key of a Secret.",
												MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from. Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"google_sheet": schema.SingleNestedAttribute{
								Description:         "A [Google sheets](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/google-sheets) connector.",
								MarkdownDescription: "A [Google sheets](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/google-sheets) connector.",
								Attributes: map[string]schema.Attribute{
									"cache": schema.SingleNestedAttribute{
										Description:         "Cache the contents of sheets. This is used to reduce Google Sheets API usage and latency.",
										MarkdownDescription: "Cache the contents of sheets. This is used to reduce Google Sheets API usage and latency.",
										Attributes: map[string]schema.Attribute{
											"sheets_data_expire_after_write": schema.StringAttribute{
												Description:         "How long to cache spreadsheet data or metadata, defaults to '5m'.",
												MarkdownDescription: "How long to cache spreadsheet data or metadata, defaults to '5m'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sheets_data_max_cache_size": schema.StringAttribute{
												Description:         "Maximum number of spreadsheets to cache, defaults to 1000.",
												MarkdownDescription: "Maximum number of spreadsheets to cache, defaults to 1000.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"credentials_secret": schema.StringAttribute{
										Description:         "The Secret containing the Google API JSON key file. The key used from the Secret is 'credentials'.",
										MarkdownDescription: "The Secret containing the Google API JSON key file. The key used from the Secret is 'credentials'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"metadata_sheet_id": schema.StringAttribute{
										Description:         "Sheet ID of the spreadsheet, that contains the table mapping.",
										MarkdownDescription: "Sheet ID of the spreadsheet, that contains the table mapping.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"hive": schema.SingleNestedAttribute{
								Description:         "An [Apache Hive](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/hive) connector.",
								MarkdownDescription: "An [Apache Hive](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/hive) connector.",
								Attributes: map[string]schema.Attribute{
									"hdfs": schema.SingleNestedAttribute{
										Description:         "Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS.",
										MarkdownDescription: "Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS.",
										Attributes: map[string]schema.Attribute{
											"config_map": schema.StringAttribute{
												Description:         "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
												MarkdownDescription: "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"metastore": schema.SingleNestedAttribute{
										Description:         "Mandatory connection to a Hive Metastore, which will be used as a storage for metadata.",
										MarkdownDescription: "Mandatory connection to a Hive Metastore, which will be used as a storage for metadata.",
										Attributes: map[string]schema.Attribute{
											"config_map": schema.StringAttribute{
												Description:         "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.",
												MarkdownDescription: "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"s3": schema.SingleNestedAttribute{
										Description:         "Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
										MarkdownDescription: "Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
										Attributes: map[string]schema.Attribute{
											"inline": schema.SingleNestedAttribute{
												Description:         "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
												MarkdownDescription: "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
												Attributes: map[string]schema.Attribute{
													"access_style": schema.StringAttribute{
														Description:         "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
														MarkdownDescription: "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Path", "VirtualHosted"),
														},
													},

													"credentials": schema.SingleNestedAttribute{
														Description:         "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
														MarkdownDescription: "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
														Attributes: map[string]schema.Attribute{
															"scope": schema.SingleNestedAttribute{
																Description:         "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																MarkdownDescription: "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																Attributes: map[string]schema.Attribute{
																	"listener_volumes": schema.ListAttribute{
																		Description:         "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																		MarkdownDescription: "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"node": schema.BoolAttribute{
																		Description:         "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																		MarkdownDescription: "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pod": schema.BoolAttribute{
																		Description:         "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																		MarkdownDescription: "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"services": schema.ListAttribute{
																		Description:         "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
																		MarkdownDescription: "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
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

															"secret_class": schema.StringAttribute{
																Description:         "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																MarkdownDescription: "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"host": schema.StringAttribute{
														Description:         "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
														MarkdownDescription: "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.Int64Attribute{
														Description:         "Port the S3 server listens on. If not specified the product will determine the port to use.",
														MarkdownDescription: "Port the S3 server listens on. If not specified the product will determine the port to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"tls": schema.SingleNestedAttribute{
														Description:         "Use a TLS connection. If not specified no TLS will be used.",
														MarkdownDescription: "Use a TLS connection. If not specified no TLS will be used.",
														Attributes: map[string]schema.Attribute{
															"verification": schema.SingleNestedAttribute{
																Description:         "The verification method used to verify the certificates of the server and/or the client.",
																MarkdownDescription: "The verification method used to verify the certificates of the server and/or the client.",
																Attributes: map[string]schema.Attribute{
																	"none": schema.MapAttribute{
																		Description:         "Use TLS but don't verify certificates.",
																		MarkdownDescription: "Use TLS but don't verify certificates.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"server": schema.SingleNestedAttribute{
																		Description:         "Use TLS and a CA certificate to verify the server.",
																		MarkdownDescription: "Use TLS and a CA certificate to verify the server.",
																		Attributes: map[string]schema.Attribute{
																			"ca_cert": schema.SingleNestedAttribute{
																				Description:         "CA cert to verify the server.",
																				MarkdownDescription: "CA cert to verify the server.",
																				Attributes: map[string]schema.Attribute{
																					"secret_class": schema.StringAttribute{
																						Description:         "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																						MarkdownDescription: "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"web_pki": schema.MapAttribute{
																						Description:         "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																						MarkdownDescription: "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
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
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"reference": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"iceberg": schema.SingleNestedAttribute{
								Description:         "An [Apache Iceberg](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/iceberg) connector.",
								MarkdownDescription: "An [Apache Iceberg](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/iceberg) connector.",
								Attributes: map[string]schema.Attribute{
									"hdfs": schema.SingleNestedAttribute{
										Description:         "Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS.",
										MarkdownDescription: "Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS.",
										Attributes: map[string]schema.Attribute{
											"config_map": schema.StringAttribute{
												Description:         "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
												MarkdownDescription: "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"metastore": schema.SingleNestedAttribute{
										Description:         "Mandatory connection to a Hive Metastore, which will be used as a storage for metadata.",
										MarkdownDescription: "Mandatory connection to a Hive Metastore, which will be used as a storage for metadata.",
										Attributes: map[string]schema.Attribute{
											"config_map": schema.StringAttribute{
												Description:         "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.",
												MarkdownDescription: "Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"s3": schema.SingleNestedAttribute{
										Description:         "Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
										MarkdownDescription: "Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3).",
										Attributes: map[string]schema.Attribute{
											"inline": schema.SingleNestedAttribute{
												Description:         "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
												MarkdownDescription: "S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3).",
												Attributes: map[string]schema.Attribute{
													"access_style": schema.StringAttribute{
														Description:         "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
														MarkdownDescription: "Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Path", "VirtualHosted"),
														},
													},

													"credentials": schema.SingleNestedAttribute{
														Description:         "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
														MarkdownDescription: "If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient.",
														Attributes: map[string]schema.Attribute{
															"scope": schema.SingleNestedAttribute{
																Description:         "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																MarkdownDescription: "[Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass).",
																Attributes: map[string]schema.Attribute{
																	"listener_volumes": schema.ListAttribute{
																		Description:         "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																		MarkdownDescription: "The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"node": schema.BoolAttribute{
																		Description:         "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																		MarkdownDescription: "The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pod": schema.BoolAttribute{
																		Description:         "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																		MarkdownDescription: "The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"services": schema.ListAttribute{
																		Description:         "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
																		MarkdownDescription: "The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.",
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

															"secret_class": schema.StringAttribute{
																Description:         "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																MarkdownDescription: "[SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"host": schema.StringAttribute{
														Description:         "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
														MarkdownDescription: "Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.Int64Attribute{
														Description:         "Port the S3 server listens on. If not specified the product will determine the port to use.",
														MarkdownDescription: "Port the S3 server listens on. If not specified the product will determine the port to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"tls": schema.SingleNestedAttribute{
														Description:         "Use a TLS connection. If not specified no TLS will be used.",
														MarkdownDescription: "Use a TLS connection. If not specified no TLS will be used.",
														Attributes: map[string]schema.Attribute{
															"verification": schema.SingleNestedAttribute{
																Description:         "The verification method used to verify the certificates of the server and/or the client.",
																MarkdownDescription: "The verification method used to verify the certificates of the server and/or the client.",
																Attributes: map[string]schema.Attribute{
																	"none": schema.MapAttribute{
																		Description:         "Use TLS but don't verify certificates.",
																		MarkdownDescription: "Use TLS but don't verify certificates.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"server": schema.SingleNestedAttribute{
																		Description:         "Use TLS and a CA certificate to verify the server.",
																		MarkdownDescription: "Use TLS and a CA certificate to verify the server.",
																		Attributes: map[string]schema.Attribute{
																			"ca_cert": schema.SingleNestedAttribute{
																				Description:         "CA cert to verify the server.",
																				MarkdownDescription: "CA cert to verify the server.",
																				Attributes: map[string]schema.Attribute{
																					"secret_class": schema.StringAttribute{
																						Description:         "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																						MarkdownDescription: "Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"web_pki": schema.MapAttribute{
																						Description:         "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
																						MarkdownDescription: "Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.",
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
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"reference": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tpcds": schema.MapAttribute{
								Description:         "A [TPC-DS](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/tpcds) connector.",
								MarkdownDescription: "A [TPC-DS](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/tpcds) connector.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tpch": schema.MapAttribute{
								Description:         "A [TPC-H](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/tpch) connector.",
								MarkdownDescription: "A [TPC-H](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/tpch) connector.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *TrinoStackableTechTrinoCatalogV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_trino_stackable_tech_trino_catalog_v1alpha1_manifest")

	var model TrinoStackableTechTrinoCatalogV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("trino.stackable.tech/v1alpha1")
	model.Kind = pointer.String("TrinoCatalog")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
