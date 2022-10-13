/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type K8sProvider struct{}

var (
	_ provider.Provider             = (*K8sProvider)(nil)
	_ provider.ProviderWithMetadata = (*K8sProvider)(nil)
)

func New() provider.Provider {
	return &K8sProvider{}
}

func (p *K8sProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "k8s"
}

func (p *K8sProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Provider for custom Kubernetes resources. Requires Terraform 1.0 or later.",
		MarkdownDescription: "Provider for custom [Kubernetes](https://kubernetes.io/) resources. Requires Terraform 1.0 or later.",
	}, nil
}

func (p *K8sProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
	// NO-OP: provider requires no configuration
}

func (p *K8sProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *K8sProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAcidZalanDoOperatorConfigurationV1Resource,
		NewAcidZalanDoPostgresqlV1Resource,
		NewAcidZalanDoPostgresTeamV1Resource,
		NewApicodegenApimaticIoAPIMaticV1Beta1Resource,
		NewApigatewayv2ServicesK8SAwsAPIV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsAuthorizerV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsDeploymentV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsIntegrationV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsRouteV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsStageV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsVPCLinkV1Alpha1Resource,
		NewAppKiegroupOrgKogitoBuildV1Beta1Resource,
		NewAppKiegroupOrgKogitoInfraV1Beta1Resource,
		NewAppKiegroupOrgKogitoRuntimeV1Beta1Resource,
		NewAppKiegroupOrgKogitoSupportingServiceV1Beta1Resource,
		NewAppLightbendComAkkaClusterV1Alpha1Resource,
		NewApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Resource,
		NewApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource,
		NewApps3ScaleNetAPIcastV1Alpha1Resource,
		NewAppsGitlabComRunnerV1Beta2Resource,
		NewAppsM88IIoNexusV1Alpha1Resource,
		NewAquasecurityGithubIoAquaStarboardV1Alpha1Resource,
		NewArgoprojIoApplicationV1Alpha1Resource,
		NewArgoprojIoApplicationSetV1Alpha1Resource,
		NewArgoprojIoAppProjectV1Alpha1Resource,
		NewArgoprojIoArgoCDExportV1Alpha1Resource,
		NewArgoprojIoArgoCDV1Alpha1Resource,
		NewAsdbAerospikeComAerospikeClusterV1Beta1Resource,
		NewCertManagerIoCertificateRequestV1Resource,
		NewCertManagerIoCertificateV1Resource,
		NewAcmeCertManagerIoChallengeV1Resource,
		NewCertManagerIoClusterIssuerV1Resource,
		NewCertManagerIoIssuerV1Resource,
		NewAcmeCertManagerIoOrderV1Resource,
		NewChartsHelmK8SIoSnykMonitorV1Alpha1Resource,
		NewChartsOpdevIoSynapseV1Alpha1Resource,
		NewCheEclipseOrgKubernetesImagePullerV1Alpha1Resource,
		NewConfigGatekeeperShConfigV1Alpha1Resource,
		NewCoreStrimziIoStrimziPodSetV1Beta2Resource,
		NewCouchbaseComCouchbaseAutoscalerV2Resource,
		NewCouchbaseComCouchbaseBackupRestoreV2Resource,
		NewCouchbaseComCouchbaseBackupV2Resource,
		NewCouchbaseComCouchbaseBucketV2Resource,
		NewCouchbaseComCouchbaseClusterV2Resource,
		NewCouchbaseComCouchbaseCollectionGroupV2Resource,
		NewCouchbaseComCouchbaseCollectionV2Resource,
		NewCouchbaseComCouchbaseEphemeralBucketV2Resource,
		NewCouchbaseComCouchbaseGroupV2Resource,
		NewCouchbaseComCouchbaseMemcachedBucketV2Resource,
		NewCouchbaseComCouchbaseMigrationReplicationV2Resource,
		NewCouchbaseComCouchbaseReplicationV2Resource,
		NewCouchbaseComCouchbaseRoleBindingV2Resource,
		NewCouchbaseComCouchbaseScopeGroupV2Resource,
		NewCouchbaseComCouchbaseScopeV2Resource,
		NewCouchbaseComCouchbaseUserV2Resource,
		NewCrdProjectcalicoOrgBGPConfigurationV1Resource,
		NewCrdProjectcalicoOrgBGPPeerV1Resource,
		NewCrdProjectcalicoOrgBlockAffinityV1Resource,
		NewCrdProjectcalicoOrgCalicoNodeStatusV1Resource,
		NewCrdProjectcalicoOrgClusterInformationV1Resource,
		NewCrdProjectcalicoOrgFelixConfigurationV1Resource,
		NewCrdProjectcalicoOrgGlobalNetworkPolicyV1Resource,
		NewCrdProjectcalicoOrgGlobalNetworkSetV1Resource,
		NewCrdProjectcalicoOrgHostEndpointV1Resource,
		NewCrdProjectcalicoOrgIPAMBlockV1Resource,
		NewCrdProjectcalicoOrgIPAMConfigV1Resource,
		NewCrdProjectcalicoOrgIPAMHandleV1Resource,
		NewCrdProjectcalicoOrgIPPoolV1Resource,
		NewCrdProjectcalicoOrgIPReservationV1Resource,
		NewCrdProjectcalicoOrgKubeControllersConfigurationV1Resource,
		NewCrdProjectcalicoOrgNetworkPolicyV1Resource,
		NewCrdProjectcalicoOrgNetworkSetV1Resource,
		NewTelemetryIstioIoTelemetryV1Alpha1Resource,
		NewExtensionsIstioIoWasmPluginV1Alpha1Resource,
		NewExternalSecretsIoClusterSecretStoreV1Alpha1Resource,
		NewExternalSecretsIoExternalSecretV1Alpha1Resource,
		NewExternalSecretsIoSecretStoreV1Alpha1Resource,
		NewExternalSecretsIoClusterExternalSecretV1Beta1Resource,
		NewExternalSecretsIoClusterSecretStoreV1Beta1Resource,
		NewExternalSecretsIoExternalSecretV1Beta1Resource,
		NewExternalSecretsIoSecretStoreV1Beta1Resource,
		NewExternaldataGatekeeperShProviderV1Alpha1Resource,
		NewFlaggerAppAlertProviderV1Beta1Resource,
		NewFlaggerAppCanaryV1Beta1Resource,
		NewFlaggerAppMetricTemplateV1Beta1Resource,
		NewFlinkApacheOrgFlinkDeploymentV1Beta1Resource,
		NewFlinkApacheOrgFlinkSessionJobV1Beta1Resource,
		NewGatewayNetworkingK8SIoGatewayClassV1Alpha2Resource,
		NewGatewayNetworkingK8SIoGatewayV1Alpha2Resource,
		NewGatewayNetworkingK8SIoHTTPRouteV1Alpha2Resource,
		NewGatewayNetworkingK8SIoTCPRouteV1Alpha2Resource,
		NewGatewayNetworkingK8SIoTLSRouteV1Alpha2Resource,
		NewHazelcastComCronHotBackupV1Alpha1Resource,
		NewHazelcastComHazelcastV1Alpha1Resource,
		NewHazelcastComHotBackupV1Alpha1Resource,
		NewHazelcastComManagementCenterV1Alpha1Resource,
		NewHazelcastComMapV1Alpha1Resource,
		NewHazelcastComWanReplicationV1Alpha1Resource,
		NewHelmSigstoreDevRekorV1Alpha1Resource,
		NewHelmToolkitFluxcdIoHelmReleaseV2Beta1Resource,
		NewImageToolkitFluxcdIoImagePolicyV1Alpha1Resource,
		NewImageToolkitFluxcdIoImageRepositoryV1Alpha1Resource,
		NewImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource,
		NewImageToolkitFluxcdIoImagePolicyV1Alpha2Resource,
		NewImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource,
		NewImageToolkitFluxcdIoImageUpdateAutomationV1Alpha2Resource,
		NewImageToolkitFluxcdIoImagePolicyV1Beta1Resource,
		NewImageToolkitFluxcdIoImageRepositoryV1Beta1Resource,
		NewImageToolkitFluxcdIoImageUpdateAutomationV1Beta1Resource,
		NewImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDicomEventDrivenIngestionV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDicomStudyBindingV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDimseIngestionServiceV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource,
		NewInfinispanOrgInfinispanV1Resource,
		NewInfinispanOrgBackupV2Alpha1Resource,
		NewInfinispanOrgBatchV2Alpha1Resource,
		NewInfinispanOrgCacheV2Alpha1Resource,
		NewInfinispanOrgRestoreV2Alpha1Resource,
		NewInstallationMattermostComMattermostV1Beta1Resource,
		NewIntegreatlyOrgGrafanaDashboardV1Alpha1Resource,
		NewIntegreatlyOrgGrafanaDataSourceV1Alpha1Resource,
		NewIntegreatlyOrgGrafanaFolderV1Alpha1Resource,
		NewIntegreatlyOrgGrafanaNotificationChannelV1Alpha1Resource,
		NewIntegreatlyOrgGrafanaV1Alpha1Resource,
		NewJaegertracingIoJaegerV1Resource,
		NewKafkaStrimziIoKafkaBridgeV1Beta2Resource,
		NewKafkaStrimziIoKafkaConnectorV1Beta2Resource,
		NewKafkaStrimziIoKafkaConnectV1Beta2Resource,
		NewKafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource,
		NewKafkaStrimziIoKafkaMirrorMakerV1Beta2Resource,
		NewKafkaStrimziIoKafkaRebalanceV1Beta2Resource,
		NewKafkaStrimziIoKafkaV1Beta2Resource,
		NewKafkaStrimziIoKafkaTopicV1Beta2Resource,
		NewKafkaStrimziIoKafkaUserV1Beta2Resource,
		NewKeycloakOrgKeycloakBackupV1Alpha1Resource,
		NewKeycloakOrgKeycloakClientV1Alpha1Resource,
		NewKeycloakOrgKeycloakRealmV1Alpha1Resource,
		NewKeycloakOrgKeycloakV1Alpha1Resource,
		NewKeycloakOrgKeycloakUserV1Alpha1Resource,
		NewKialiIoKialiV1Alpha1Resource,
		NewKustomizeToolkitFluxcdIoKustomizationV1Beta1Resource,
		NewKustomizeToolkitFluxcdIoKustomizationV1Beta2Resource,
		NewLitmuschaosIoChaosEngineV1Alpha1Resource,
		NewLitmuschaosIoChaosExperimentV1Alpha1Resource,
		NewLitmuschaosIoChaosResultV1Alpha1Resource,
		NewMattermostComClusterInstallationV1Alpha1Resource,
		NewMattermostComMattermostRestoreDBV1Alpha1Resource,
		NewMinioMinIoTenantV1Resource,
		NewMinioMinIoTenantV2Resource,
		NewMonitoringCoreosComAlertmanagerV1Resource,
		NewMonitoringCoreosComPodMonitorV1Resource,
		NewMonitoringCoreosComProbeV1Resource,
		NewMonitoringCoreosComPrometheusV1Resource,
		NewMonitoringCoreosComPrometheusRuleV1Resource,
		NewMonitoringCoreosComServiceMonitorV1Resource,
		NewMonitoringCoreosComThanosRulerV1Resource,
		NewMonitoringCoreosComAlertmanagerConfigV1Alpha1Resource,
		NewMutationsGatekeeperShAssignV1Alpha1Resource,
		NewMutationsGatekeeperShAssignMetadataV1Alpha1Resource,
		NewMutationsGatekeeperShModifySetV1Alpha1Resource,
		NewMutationsGatekeeperShAssignV1Beta1Resource,
		NewMutationsGatekeeperShAssignMetadataV1Beta1Resource,
		NewMutationsGatekeeperShModifySetV1Beta1Resource,
		NewNetworkingIstioIoDestinationRuleV1Alpha3Resource,
		NewNetworkingIstioIoEnvoyFilterV1Alpha3Resource,
		NewNetworkingIstioIoGatewayV1Alpha3Resource,
		NewNetworkingIstioIoServiceEntryV1Alpha3Resource,
		NewNetworkingIstioIoSidecarV1Alpha3Resource,
		NewNetworkingIstioIoVirtualServiceV1Alpha3Resource,
		NewNetworkingIstioIoWorkloadEntryV1Alpha3Resource,
		NewNetworkingIstioIoWorkloadGroupV1Alpha3Resource,
		NewSecurityIstioIoAuthorizationPolicyV1Beta1Resource,
		NewNetworkingIstioIoGatewayV1Beta1Resource,
		NewSecurityIstioIoPeerAuthenticationV1Beta1Resource,
		NewNetworkingIstioIoProxyConfigV1Beta1Resource,
		NewSecurityIstioIoRequestAuthenticationV1Beta1Resource,
		NewNetworkingIstioIoServiceEntryV1Beta1Resource,
		NewNetworkingIstioIoSidecarV1Beta1Resource,
		NewNetworkingIstioIoVirtualServiceV1Beta1Resource,
		NewNetworkingIstioIoWorkloadEntryV1Beta1Resource,
		NewNetworkingIstioIoWorkloadGroupV1Beta1Resource,
		NewNotificationToolkitFluxcdIoAlertV1Beta1Resource,
		NewNotificationToolkitFluxcdIoProviderV1Beta1Resource,
		NewNotificationToolkitFluxcdIoReceiverV1Beta1Resource,
		NewOpentelemetryIoInstrumentationV1Alpha1Resource,
		NewOpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource,
		NewOperatorAquasecComAquaCspV1Alpha1Resource,
		NewOperatorAquasecComAquaDatabaseV1Alpha1Resource,
		NewOperatorAquasecComAquaEnforcerV1Alpha1Resource,
		NewOperatorAquasecComAquaGatewayV1Alpha1Resource,
		NewOperatorAquasecComAquaKubeEnforcerV1Alpha1Resource,
		NewOperatorAquasecComAquaScannerV1Alpha1Resource,
		NewOperatorAquasecComAquaServerV1Alpha1Resource,
		NewOperatorKnativeDevKnativeEventingV1Beta1Resource,
		NewOperatorKnativeDevKnativeServingV1Beta1Resource,
		NewOperatorTektonDevTektonResultV1Alpha1Resource,
		NewOperatorTigeraIoAPIServerV1Resource,
		NewOperatorTigeraIoImageSetV1Resource,
		NewOperatorTigeraIoInstallationV1Resource,
		NewOperatorTigeraIoTigeraStatusV1Resource,
		NewPostgresOperatorCrunchydataComPostgresClusterV1Beta1Resource,
		NewQuayRedhatComQuayRegistryV1Resource,
		NewRedhatcopRedhatIoQuayEcosystemV1Alpha1Resource,
		NewRegistryApicurIoApicurioRegistryV1Resource,
		NewRipsawCloudbulldozerIoBenchmarkV1Alpha1Resource,
		NewRocketmqApacheOrgBrokerV1Alpha1Resource,
		NewRocketmqApacheOrgConsoleV1Alpha1Resource,
		NewRocketmqApacheOrgNameServiceV1Alpha1Resource,
		NewRocketmqApacheOrgTopicTransferV1Alpha1Resource,
		NewSecscanQuayRedhatComImageManifestVulnV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoProfileBindingV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoProfileRecordingV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoSecurityProfileNodeStatusV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoRawSelinuxProfileV1Alpha2Resource,
		NewSecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Resource,
		NewSecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource,
		NewSourceToolkitFluxcdIoBucketV1Beta1Resource,
		NewSourceToolkitFluxcdIoGitRepositoryV1Beta1Resource,
		NewSourceToolkitFluxcdIoHelmChartV1Beta1Resource,
		NewSourceToolkitFluxcdIoHelmRepositoryV1Beta1Resource,
		NewSourceToolkitFluxcdIoBucketV1Beta2Resource,
		NewSourceToolkitFluxcdIoGitRepositoryV1Beta2Resource,
		NewSourceToolkitFluxcdIoHelmChartV1Beta2Resource,
		NewSourceToolkitFluxcdIoHelmRepositoryV1Beta2Resource,
		NewSourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource,
		NewSparkoperatorK8SIoScheduledSparkApplicationV1Beta2Resource,
		NewSparkoperatorK8SIoSparkApplicationV1Beta2Resource,
		NewTemplatesGatekeeperShConstraintTemplateV1Resource,
		NewTemplatesGatekeeperShConstraintTemplateV1Alpha1Resource,
		NewTemplatesGatekeeperShConstraintTemplateV1Beta1Resource,
		NewTraefikContainoUsIngressRouteV1Alpha1Resource,
		NewTraefikContainoUsIngressRouteTCPV1Alpha1Resource,
		NewTraefikContainoUsIngressRouteUDPV1Alpha1Resource,
		NewTraefikContainoUsMiddlewareV1Alpha1Resource,
		NewTraefikContainoUsMiddlewareTCPV1Alpha1Resource,
		NewTraefikContainoUsServersTransportV1Alpha1Resource,
		NewTraefikContainoUsTLSOptionV1Alpha1Resource,
		NewTraefikContainoUsTLSStoreV1Alpha1Resource,
		NewTraefikContainoUsTraefikServiceV1Alpha1Resource,
		NewWildflyOrgWildFlyServerV1Alpha1Resource,
	}
}
