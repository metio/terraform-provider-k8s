//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

var crdv1Sources = []string{
	"https://github.com/zalando/postgres-operator/blob/master/charts/postgres-operator/crds/postgresqls.yaml",
	"https://github.com/zalando/postgres-operator/blob/master/charts/postgres-operator/crds/operatorconfigurations.yaml",
	"https://github.com/zalando/postgres-operator/blob/master/charts/postgres-operator/crds/postgresteams.yaml",

	"https://github.com/elastic/cloud-on-k8s/blob/main/config/crds/v1/bases/agent.k8s.elastic.co_agents.yaml",

	"https://github.com/apimatic/apimatic-kubernetes-operator/blob/main/config/crd/bases/apicodegen.apimatic.io_apimatics.yaml",

	"https://github.com/aws-controllers-k8s/apigatewayv2-controller/blob/main/config/crd/bases/apigatewayv2.services.k8s.aws_apis.yaml",
	"https://github.com/aws-controllers-k8s/apigatewayv2-controller/blob/main/config/crd/bases/apigatewayv2.services.k8s.aws_authorizers.yaml",
	"https://github.com/aws-controllers-k8s/apigatewayv2-controller/blob/main/config/crd/bases/apigatewayv2.services.k8s.aws_deployments.yaml",
	"https://github.com/aws-controllers-k8s/apigatewayv2-controller/blob/main/config/crd/bases/apigatewayv2.services.k8s.aws_integrations.yaml",
	"https://github.com/aws-controllers-k8s/apigatewayv2-controller/blob/main/config/crd/bases/apigatewayv2.services.k8s.aws_routes.yaml",
	"https://github.com/aws-controllers-k8s/apigatewayv2-controller/blob/main/config/crd/bases/apigatewayv2.services.k8s.aws_stages.yaml",
	"https://github.com/aws-controllers-k8s/apigatewayv2-controller/blob/main/config/crd/bases/apigatewayv2.services.k8s.aws_vpclinks.yaml",

	"https://github.com/elastic/cloud-on-k8s/blob/main/config/crds/v1/bases/apm.k8s.elastic.co_apmservers.yaml",

	"https://github.com/kiegroup/kogito-operator/blob/main/config/crd/app/bases/app.kiegroup.org_kogitobuilds.yaml",
	"https://github.com/kiegroup/kogito-operator/blob/main/config/crd/app/bases/app.kiegroup.org_kogitoinfras.yaml",
	"https://github.com/kiegroup/kogito-operator/blob/main/config/crd/app/bases/app.kiegroup.org_kogitoruntimes.yaml",
	"https://github.com/kiegroup/kogito-operator/blob/main/config/crd/app/bases/app.kiegroup.org_kogitosupportingservices.yaml",

	"https://github.com/lightbend/akka-cluster-operator/blob/master/deploy/crds/app_v1alpha1_akkacluster_crd.yaml",

	"https://github.com/RedisLabs/redis-enterprise-k8s-docs/blob/master/crds/rec_crd.yaml",
	"https://github.com/RedisLabs/redis-enterprise-k8s-docs/blob/master/crds/redb_crd.yaml",

	"https://github.com/aws-controllers-k8s/applicationautoscaling-controller/blob/main/config/crd/bases/applicationautoscaling.services.k8s.aws_scalabletargets.yaml",
	"https://github.com/aws-controllers-k8s/applicationautoscaling-controller/blob/main/config/crd/bases/applicationautoscaling.services.k8s.aws_scalingpolicies.yaml",

	"https://github.com/3scale/apicast-operator/blob/master/config/crd/bases/apps.3scale.net_apicasts.yaml",

	"https://gitlab.com/gitlab-org/cloud-native/gitlab-operator/-/blob/master/config/crd/bases/apps.gitlab.com_gitlabs.yaml",
	"https://gitlab.com/gitlab-org/gl-openshift/gitlab-runner-operator/-/blob/master/config/crd_k8s/bases/apps.gitlab.com_runners.yaml",

	"https://github.com/m88i/nexus-operator/blob/main/config/crd/bases/apps.m88i.io_nexus.yaml",

	"https://github.com/redhat-performance/cluster-impairment-operator/blob/main/config/crd/bases/apps.redhat.com_clusterimpairments.yaml",

	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/aquasecurity.github.io_aquastarboards.yaml",
	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/aquasecurity.github.io_clusterconfigauditreports.yaml",
	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/aquasecurity.github.io_configauditreports.yaml",

	"https://github.com/argoproj-labs/argocd-operator/blob/master/config/crd/bases/argoproj.io_applications.yaml",
	"https://github.com/argoproj-labs/argocd-operator/blob/master/config/crd/bases/argoproj.io_applicationsets.yaml",
	"https://github.com/argoproj-labs/argocd-operator/blob/master/config/crd/bases/argoproj.io_appprojects.yaml",
	"https://github.com/argoproj-labs/argocd-operator/blob/master/config/crd/bases/argoproj.io_argocdexports.yaml",
	"https://github.com/argoproj-labs/argocd-operator/blob/master/config/crd/bases/argoproj.io_argocds.yaml",

	"https://github.com/aerospike/aerospike-kubernetes-operator/blob/master/config/crd/bases/asdb.aerospike.com_aerospikeclusters.yaml",

	"https://github.com/elastic/cloud-on-k8s/blob/main/config/crds/v1/bases/autoscaling.k8s.elastic.co_elasticsearchautoscalers.yaml",

	"https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/deploy/vpa-v1-crd-gen.yaml",

	"https://github.com/elastic/cloud-on-k8s/blob/main/config/crds/v1/bases/beat.k8s.elastic.co_beats.yaml",

	"https://github.com/redhat-developer/service-binding-operator/blob/master/config/crd/bases/binding.operators.coreos.com_bindablekinds.yaml",
	"https://github.com/redhat-developer/service-binding-operator/blob/master/config/crd/bases/binding.operators.coreos.com_servicebindings.yaml",

	"https://github.com/IBM/varnish-operator/blob/main/config/crd/bases/caching.ibm.com_varnishclusters.yaml",

	"https://github.com/apache/camel-k/blob/main/config/crd/bases/camel.apache.org_builds.yaml",
	"https://github.com/apache/camel-k/blob/main/config/crd/bases/camel.apache.org_camelcatalogs.yaml",
	"https://github.com/apache/camel-k/blob/main/config/crd/bases/camel.apache.org_integrationkits.yaml",
	"https://github.com/apache/camel-k/blob/main/config/crd/bases/camel.apache.org_integrationplatforms.yaml",
	"https://github.com/apache/camel-k/blob/main/config/crd/bases/camel.apache.org_integrations.yaml",
	"https://github.com/apache/camel-k/blob/main/config/crd/bases/camel.apache.org_kameletbindings.yaml",
	"https://github.com/apache/camel-k/blob/main/config/crd/bases/camel.apache.org_kamelets.yaml",

	"https://github.com/cert-manager/cert-manager/blob/master/deploy/crds/crd-certificaterequests.yaml",
	"https://github.com/cert-manager/cert-manager/blob/master/deploy/crds/crd-certificates.yaml",
	"https://github.com/cert-manager/cert-manager/blob/master/deploy/crds/crd-challenges.yaml",
	"https://github.com/cert-manager/cert-manager/blob/master/deploy/crds/crd-clusterissuers.yaml",
	"https://github.com/cert-manager/cert-manager/blob/master/deploy/crds/crd-issuers.yaml",
	"https://github.com/cert-manager/cert-manager/blob/master/deploy/crds/crd-orders.yaml",

	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_awschaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_azurechaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_blockchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_dnschaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_gcpchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_httpchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_iochaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_jvmchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_kernelchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_networkchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_physicalmachinechaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_physicalmachines.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_podchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_podhttpchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_podiochaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_podnetworkchaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_remoteclusters.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_schedules.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_statuschecks.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_stresschaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_timechaos.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_workflownodes.yaml",
	"https://github.com/chaos-mesh/chaos-mesh/blob/master/config/crd/bases/chaos-mesh.org_workflows.yaml",

	"https://github.com/Flagsmith/flagsmith-operator/blob/master/config/crd/bases/charts.flagsmith.com_flagsmiths.yaml",

	"https://github.com/snyk/kubernetes-monitor/blob/staging/snyk-operator/deploy/olm-catalog/snyk-operator/0.0.0/snykmonitors.charts.helm.k8s.io.crd.yaml",

	"https://github.com/opdev/synapse-helm/blob/master/config/crd/bases/charts.opdev.io_synapses.yaml",

	"https://github.com/dmesser/cockroachdb-operator/blob/main/config/crd/bases/charts.operatorhub.io_cockroachdbs.yaml",

	"https://github.com/che-incubator/kubernetes-image-puller-operator/blob/main/config/crd/bases/che.eclipse.org_kubernetesimagepullers.yaml",

	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumclusterwideenvoyconfigs.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumclusterwidenetworkpolicies.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumegressgatewaypolicies.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumendpoints.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumenvoyconfigs.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumexternalworkloads.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumidentities.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumlocalredirectpolicies.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumnetworkpolicies.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2/ciliumnodes.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2alpha1/ciliumbgploadbalancerippools.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2alpha1/ciliumbgppeeringpolicies.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2alpha1/ciliumegressnatpolicies.yaml",
	"https://github.com/cilium/cilium/blob/master/pkg/k8s/apis/cilium.io/client/crds/v2alpha1/ciliumendpointslices.yaml",

	"https://github.com/schemahero/schemahero/blob/main/config/crds/v1/databases.schemahero.io_databases.yaml",
	"https://github.com/schemahero/schemahero/blob/main/config/crds/v1/schemas.schemahero.io_datatypes.yaml",
	"https://github.com/schemahero/schemahero/blob/main/config/crds/v1/schemas.schemahero.io_migrations.yaml",
	"https://github.com/schemahero/schemahero/blob/main/config/crds/v1/schemas.schemahero.io_tables.yaml",

	"https://github.com/open-policy-agent/gatekeeper/blob/master/deploy/gatekeeper.yaml",

	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/042-Crd-strimzipodset.yaml",

	"https://github.com/couchbase-partners/helm-charts/blob/master/charts/couchbase-operator/crds/couchbase.crds.yaml",

	"https://github.com/projectcalico/calico/blob/master/manifests/crds.yaml",

	"https://github.com/aws-controllers-k8s/dynamodb-controller/blob/main/config/crd/bases/dynamodb.services.k8s.aws_backups.yaml",
	"https://github.com/aws-controllers-k8s/dynamodb-controller/blob/main/config/crd/bases/dynamodb.services.k8s.aws_globaltables.yaml",
	"https://github.com/aws-controllers-k8s/dynamodb-controller/blob/main/config/crd/bases/dynamodb.services.k8s.aws_tables.yaml",

	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_dhcpoptions.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_elasticipaddresses.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_instances.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_internetgateways.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_natgateways.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_routetables.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_securitygroups.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_subnets.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_transitgateways.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_vpcendpoints.yaml",
	"https://github.com/aws-controllers-k8s/ec2-controller/blob/main/config/crd/bases/ec2.services.k8s.aws_vpcs.yaml",

	"https://github.com/aws-controllers-k8s/ecr-controller/blob/main/config/crd/bases/ecr.services.k8s.aws_pullthroughcacherules.yaml",
	"https://github.com/aws-controllers-k8s/ecr-controller/blob/main/config/crd/bases/ecr.services.k8s.aws_repositories.yaml",

	"https://github.com/aws-controllers-k8s/eks-controller/blob/main/config/crd/bases/eks.services.k8s.aws_addons.yaml",
	"https://github.com/aws-controllers-k8s/eks-controller/blob/main/config/crd/bases/eks.services.k8s.aws_clusters.yaml",
	"https://github.com/aws-controllers-k8s/eks-controller/blob/main/config/crd/bases/eks.services.k8s.aws_fargateprofiles.yaml",
	"https://github.com/aws-controllers-k8s/eks-controller/blob/main/config/crd/bases/eks.services.k8s.aws_nodegroups.yaml",

	"https://github.com/aws-controllers-k8s/elasticache-controller/blob/main/config/crd/bases/elasticache.services.k8s.aws_cacheparametergroups.yaml",
	"https://github.com/aws-controllers-k8s/elasticache-controller/blob/main/config/crd/bases/elasticache.services.k8s.aws_cachesubnetgroups.yaml",
	"https://github.com/aws-controllers-k8s/elasticache-controller/blob/main/config/crd/bases/elasticache.services.k8s.aws_replicationgroups.yaml",
	"https://github.com/aws-controllers-k8s/elasticache-controller/blob/main/config/crd/bases/elasticache.services.k8s.aws_snapshots.yaml",
	"https://github.com/aws-controllers-k8s/elasticache-controller/blob/main/config/crd/bases/elasticache.services.k8s.aws_usergroups.yaml",
	"https://github.com/aws-controllers-k8s/elasticache-controller/blob/main/config/crd/bases/elasticache.services.k8s.aws_users.yaml",

	"https://github.com/elastic/cloud-on-k8s/blob/main/config/crds/v1/bases/elasticsearch.k8s.elastic.co_elasticsearches.yaml",

	"https://github.com/aws-controllers-k8s/emrcontainers-controller/blob/main/config/crd/bases/emrcontainers.services.k8s.aws_jobruns.yaml",
	"https://github.com/aws-controllers-k8s/emrcontainers-controller/blob/main/config/crd/bases/emrcontainers.services.k8s.aws_virtualclusters.yaml",

	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/enterprise.gloo.solo.io_v1_AuthConfig.yaml",

	"https://github.com/elastic/cloud-on-k8s/blob/main/config/crds/v1/bases/enterprisesearch.k8s.elastic.co_enterprisesearches.yaml",

	"https://github.com/furiko-io/furiko/blob/main/config/crd/bases/execution.furiko.io_jobconfigs.yaml",
	"https://github.com/furiko-io/furiko/blob/main/config/crd/bases/execution.furiko.io_jobs.yaml",

	"https://github.com/istio/istio/blob/master/manifests/charts/base/crds/crd-all.gen.yaml",

	"https://github.com/external-secrets/external-secrets/blob/main/config/crds/bases/external-secrets.io_clusterexternalsecrets.yaml",
	"https://github.com/external-secrets/external-secrets/blob/main/config/crds/bases/external-secrets.io_clustersecretstores.yaml",
	"https://github.com/external-secrets/external-secrets/blob/main/config/crds/bases/external-secrets.io_externalsecrets.yaml",
	"https://github.com/external-secrets/external-secrets/blob/main/config/crds/bases/external-secrets.io_secretstores.yaml",

	"https://github.com/k8gb-io/k8gb/blob/master/chart/k8gb/crd/dns-endpoint-crd-manifest.yaml",

	"https://github.com/fluxcd/flagger/blob/main/artifacts/flagger/crd.yaml",

	"https://github.com/apache/flink-kubernetes-operator/blob/main/helm/flink-kubernetes-operator/crds/flinkdeployments.flink.apache.org-v1.yml",
	"https://github.com/apache/flink-kubernetes-operator/blob/main/helm/flink-kubernetes-operator/crds/flinksessionjobs.flink.apache.org-v1.yml",

	"https://github.com/fossul/fossul/blob/master/operator/config/crd/bases/fossul.io_backupconfigs.yaml",
	"https://github.com/fossul/fossul/blob/master/operator/config/crd/bases/fossul.io_backups.yaml",
	"https://github.com/fossul/fossul/blob/master/operator/config/crd/bases/fossul.io_backupschedules.yaml",
	"https://github.com/fossul/fossul/blob/master/operator/config/crd/bases/fossul.io_fossuls.yaml",
	"https://github.com/fossul/fossul/blob/master/operator/config/crd/bases/fossul.io_restores.yaml",

	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/gateway.networking.k8s.io_gatewayclasses.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/gateway.networking.k8s.io_gateways.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/gateway.networking.k8s.io_httproutes.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/gateway.networking.k8s.io_tcproutes.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/gateway.networking.k8s.io_tlsroutes.yaml",

	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gateway.solo.io_v1_Gateway.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gateway.solo.io_v1_MatchableHttpGateway.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gateway.solo.io_v1_RouteOption.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gateway.solo.io_v1_RouteTable.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gateway.solo.io_v1_VirtualHostOption.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gateway.solo.io_v1_VirtualService.yaml",

	"https://raw.githubusercontent.com/emissary-ingress/emissary/master/pkg/api/getambassador.io/crds.yaml",

	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gloo.solo.io_v1_Proxy.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gloo.solo.io_v1_Settings.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gloo.solo.io_v1_UpstreamGroup.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/gloo.solo.io_v1_Upstream.yaml",
	"https://github.com/solo-io/gloo/blob/master/install/helm/gloo/crds/graphql.gloo.solo.io_v1beta1_GraphQLApi.yaml",

	"https://github.com/hazelcast/hazelcast-platform-operator/blob/main/config/crd/bases/hazelcast.com_cronhotbackups.yaml",
	"https://github.com/hazelcast/hazelcast-platform-operator/blob/main/config/crd/bases/hazelcast.com_hazelcasts.yaml",
	"https://github.com/hazelcast/hazelcast-platform-operator/blob/main/config/crd/bases/hazelcast.com_hotbackups.yaml",
	"https://github.com/hazelcast/hazelcast-platform-operator/blob/main/config/crd/bases/hazelcast.com_managementcenters.yaml",
	"https://github.com/hazelcast/hazelcast-platform-operator/blob/main/config/crd/bases/hazelcast.com_maps.yaml",
	"https://github.com/hazelcast/hazelcast-platform-operator/blob/main/config/crd/bases/hazelcast.com_wanreplications.yaml",

	"https://github.com/sigstore/sigstore-helm-operator/blob/main/config/crd/bases/helm.sigstore.dev_rekors.yaml",

	"https://github.com/fluxcd/kustomize-controller/blob/main/config/crd/bases/kustomize.toolkit.fluxcd.io_kustomizations.yaml",
	"https://github.com/fluxcd/source-controller/blob/main/config/crd/bases/source.toolkit.fluxcd.io_buckets.yaml",
	"https://github.com/fluxcd/source-controller/blob/main/config/crd/bases/source.toolkit.fluxcd.io_gitrepositories.yaml",
	"https://github.com/fluxcd/source-controller/blob/main/config/crd/bases/source.toolkit.fluxcd.io_helmcharts.yaml",
	"https://github.com/fluxcd/source-controller/blob/main/config/crd/bases/source.toolkit.fluxcd.io_helmrepositories.yaml",
	"https://github.com/fluxcd/source-controller/blob/main/config/crd/bases/source.toolkit.fluxcd.io_ocirepositories.yaml",
	"https://github.com/fluxcd/helm-controller/blob/main/config/crd/bases/helm.toolkit.fluxcd.io_helmreleases.yaml",
	"https://github.com/fluxcd/notification-controller/blob/main/config/crd/bases/notification.toolkit.fluxcd.io_alerts.yaml",
	"https://github.com/fluxcd/notification-controller/blob/main/config/crd/bases/notification.toolkit.fluxcd.io_providers.yaml",
	"https://github.com/fluxcd/notification-controller/blob/main/config/crd/bases/notification.toolkit.fluxcd.io_receivers.yaml",
	"https://github.com/fluxcd/image-reflector-controller/blob/main/config/crd/bases/image.toolkit.fluxcd.io_imagepolicies.yaml",
	"https://github.com/fluxcd/image-reflector-controller/blob/main/config/crd/bases/image.toolkit.fluxcd.io_imagerepositories.yaml",
	"https://github.com/fluxcd/image-automation-controller/blob/main/config/crd/bases/image.toolkit.fluxcd.io_imageupdateautomations.yaml",

	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_checkpoints.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterclaims.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterdeploymentcustomizations.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterdeployments.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterdeprovisions.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterimagesets.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterpools.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterprovisions.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterrelocates.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_clusterstates.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hiveinternal.openshift.io_clustersyncleases.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_dnszones.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_hiveconfigs.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_machinepoolnameleases.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_machinepools.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_selectorsyncidentityproviders.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_selectorsyncsets.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_syncidentityproviders.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hive.openshift.io_syncsets.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hiveinternal.openshift.io_clustersyncs.yaml",
	"https://github.com/openshift/hive/blob/master/config/crds/hiveinternal.openshift.io_fakeclusterinstalls.yaml",

	"https://github.com/Hyperfoil/horreum-operator/blob/master/config/crd/bases/hyperfoil.io_horreums.yaml",
	"https://github.com/Hyperfoil/hyperfoil-operator/blob/master/config/crd/bases/hyperfoil.io_hyperfoils.yaml",

	"https://github.com/aws-controllers-k8s/iam-controller/blob/main/config/crd/bases/iam.services.k8s.aws_groups.yaml",
	"https://github.com/aws-controllers-k8s/iam-controller/blob/main/config/crd/bases/iam.services.k8s.aws_policies.yaml",
	"https://github.com/aws-controllers-k8s/iam-controller/blob/main/config/crd/bases/iam.services.k8s.aws_roles.yaml",

	"https://github.com/composable-operator/composable/blob/main/config/crd/bases/ibmcloud.ibm.com_composables.yaml",

	"https://github.com/Alvearie/imaging-ingestion/blob/main/imaging-ingestion-operator/config/crd/bases/imaging-ingestion.alvearie.org_dicomeventbridges.yaml",
	"https://github.com/Alvearie/imaging-ingestion/blob/main/imaging-ingestion-operator/config/crd/bases/imaging-ingestion.alvearie.org_dicomeventdriveningestions.yaml",
	"https://github.com/Alvearie/imaging-ingestion/blob/main/imaging-ingestion-operator/config/crd/bases/imaging-ingestion.alvearie.org_dicominstancebindings.yaml",
	"https://github.com/Alvearie/imaging-ingestion/blob/main/imaging-ingestion-operator/config/crd/bases/imaging-ingestion.alvearie.org_dicomstudybindings.yaml",
	"https://github.com/Alvearie/imaging-ingestion/blob/main/imaging-ingestion-operator/config/crd/bases/imaging-ingestion.alvearie.org_dicomwebingestionservices.yaml",
	"https://github.com/Alvearie/imaging-ingestion/blob/main/imaging-ingestion-operator/config/crd/bases/imaging-ingestion.alvearie.org_dimseingestionservices.yaml",
	"https://github.com/Alvearie/imaging-ingestion/blob/main/imaging-ingestion-operator/config/crd/bases/imaging-ingestion.alvearie.org_dimseproxies.yaml",

	"https://github.com/infinispan/infinispan-operator/blob/main/config/crd/bases/infinispan.org_infinispans.yaml",
	"https://github.com/infinispan/infinispan-operator/blob/main/config/crd/bases/infinispan.org_backups.yaml",
	"https://github.com/infinispan/infinispan-operator/blob/main/config/crd/bases/infinispan.org_batches.yaml",
	"https://github.com/infinispan/infinispan-operator/blob/main/config/crd/bases/infinispan.org_caches.yaml",
	"https://github.com/infinispan/infinispan-operator/blob/main/config/crd/bases/infinispan.org_restores.yaml",

	"https://github.com/mattermost/mattermost-operator/blob/master/config/crd/bases/installation.mattermost.com_mattermosts.yaml",

	"https://github.com/grafana-operator/grafana-operator/blob/master/config/crd/bases/integreatly.org_grafanadashboards.yaml",
	"https://github.com/grafana-operator/grafana-operator/blob/master/config/crd/bases/integreatly.org_grafanadatasources.yaml",
	"https://github.com/grafana-operator/grafana-operator/blob/master/config/crd/bases/integreatly.org_grafanafolders.yaml",
	"https://github.com/grafana-operator/grafana-operator/blob/master/config/crd/bases/integreatly.org_grafananotificationchannels.yaml",
	"https://github.com/grafana-operator/grafana-operator/blob/master/config/crd/bases/integreatly.org_grafanas.yaml",

	"https://github.com/grafana/loki/blob/main/operator/config/crd/bases/config.grafana.com_projectconfigs.yaml",
	"https://github.com/grafana/loki/blob/main/operator/config/crd/bases/loki.grafana.com_alertingrules.yaml",
	"https://github.com/grafana/loki/blob/main/operator/config/crd/bases/loki.grafana.com_lokistacks.yaml",
	"https://github.com/grafana/loki/blob/main/operator/config/crd/bases/loki.grafana.com_recordingrules.yaml",
	"https://github.com/grafana/loki/blob/main/operator/config/crd/bases/loki.grafana.com_rulerconfigs.yaml",

	"https://github.com/ctron/ditto-operator/blob/main/helm/ditto-operator/crds/ditto.yaml",
	"https://github.com/ctron/hawkbit-operator/blob/main/crds/hawkbit.crd.yaml",

	"https://github.com/jaegertracing/jaeger-operator/blob/main/config/crd/bases/jaegertracing.io_jaegers.yaml",

	"https://github.com/k8gb-io/k8gb/blob/master/chart/k8gb/crd/k8gb.absa.oss_gslbs.yaml",

	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/046-Crd-kafkabridge.yaml",
	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/047-Crd-kafkaconnector.yaml",
	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/041-Crd-kafkaconnect.yaml",
	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/048-Crd-kafkamirrormaker2.yaml",
	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/045-Crd-kafkamirrormaker.yaml",
	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/049-Crd-kafkarebalance.yaml",
	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/040-Crd-kafka.yaml",
	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/043-Crd-kafkatopic.yaml",
	"https://github.com/strimzi/strimzi-kafka-operator/blob/main/helm-charts/helm3/strimzi-kafka-operator/crds/044-Crd-kafkauser.yaml",

	"https://github.com/keycloak/keycloak-operator/blob/main/deploy/crds/keycloak.org_keycloakbackups_crd.yaml",
	"https://github.com/keycloak/keycloak-operator/blob/main/deploy/crds/keycloak.org_keycloakclients_crd.yaml",
	"https://github.com/keycloak/keycloak-operator/blob/main/deploy/crds/keycloak.org_keycloakrealms_crd.yaml",
	"https://github.com/keycloak/keycloak-operator/blob/main/deploy/crds/keycloak.org_keycloaks_crd.yaml",
	"https://github.com/keycloak/keycloak-operator/blob/main/deploy/crds/keycloak.org_keycloakusers_crd.yaml",

	"https://github.com/kiali/kiali-operator/blob/master/crd-docs/crd/kiali.io_kialis.yaml",

	"https://github.com/elastic/cloud-on-k8s/blob/main/config/crds/v1/bases/kibana.k8s.elastic.co_kibanas.yaml",

	"https://github.com/aws-controllers-k8s/kms-controller/blob/main/config/crd/bases/kms.services.k8s.aws_aliases.yaml",
	"https://github.com/aws-controllers-k8s/kms-controller/blob/main/config/crd/bases/kms.services.k8s.aws_grants.yaml",
	"https://github.com/aws-controllers-k8s/kms-controller/blob/main/config/crd/bases/kms.services.k8s.aws_keys.yaml",

	"https://github.com/kyverno/kyverno/blob/main/config/crds/kyverno.io_clusterpolicies.yaml",
	"https://github.com/kyverno/kyverno/blob/main/config/crds/kyverno.io_generaterequests.yaml",
	"https://github.com/kyverno/kyverno/blob/main/config/crds/kyverno.io_policies.yaml",
	"https://github.com/kyverno/kyverno/blob/main/config/crds/kyverno.io_admissionreports.yaml",
	"https://github.com/kyverno/kyverno/blob/main/config/crds/kyverno.io_backgroundscanreports.yaml",
	"https://github.com/kyverno/kyverno/blob/main/config/crds/kyverno.io_clusteradmissionreports.yaml",
	"https://github.com/kyverno/kyverno/blob/main/config/crds/kyverno.io_clusterbackgroundscanreports.yaml",
	"https://github.com/kyverno/kyverno/blob/main/config/crds/kyverno.io_updaterequests.yaml",

	"https://github.com/aws-controllers-k8s/lambda-controller/blob/main/config/crd/bases/lambda.services.k8s.aws_aliases.yaml",
	"https://github.com/aws-controllers-k8s/lambda-controller/blob/main/config/crd/bases/lambda.services.k8s.aws_codesigningconfigs.yaml",
	"https://github.com/aws-controllers-k8s/lambda-controller/blob/main/config/crd/bases/lambda.services.k8s.aws_eventsourcemappings.yaml",
	"https://github.com/aws-controllers-k8s/lambda-controller/blob/main/config/crd/bases/lambda.services.k8s.aws_functions.yaml",
	"https://github.com/aws-controllers-k8s/lambda-controller/blob/main/config/crd/bases/lambda.services.k8s.aws_functionurlconfigs.yaml",

	"https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-crds/templates/serviceprofile.yaml",

	"https://github.com/litmuschaos/chaos-operator/blob/master/deploy/crds/chaosengine_crd.yaml",
	"https://github.com/litmuschaos/chaos-operator/blob/master/deploy/crds/chaosexperiment_crd.yaml",
	"https://github.com/litmuschaos/chaos-operator/blob/master/deploy/crds/chaosresults_crd.yaml",

	"https://github.com/elastic/cloud-on-k8s/blob/main/config/crds/v1/bases/maps.k8s.elastic.co_elasticmapsservers.yaml",

	"https://github.com/mattermost/mattermost-operator/blob/master/config/crd/bases/mattermost.com_clusterinstallations.yaml",
	"https://github.com/mattermost/mattermost-operator/blob/master/config/crd/bases/mattermost.com_mattermostrestoredbs.yaml",

	"https://github.com/microcks/microcks-ansible-operator/blob/master/deploy/crds/microcks_v1alpha1_microcksinstall_crd.yaml",

	"https://github.com/minio/operator/blob/master/resources/base/crds/minio.min.io_tenants.yaml",

	"https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/crds/crd-alertmanagerconfigs.yaml",
	"https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/crds/crd-alertmanagers.yaml",
	"https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/crds/crd-podmonitors.yaml",
	"https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/crds/crd-probes.yaml",
	"https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/crds/crd-prometheuses.yaml",
	"https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/crds/crd-prometheusrules.yaml",
	"https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/crds/crd-servicemonitors.yaml",
	"https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/crds/crd-thanosrulers.yaml",

	"https://github.com/aws-controllers-k8s/mq-controller/blob/main/config/crd/bases/mq.services.k8s.aws_brokers.yaml",

	"https://github.com/aws-controllers-k8s/opensearchservice-controller/blob/main/config/crd/bases/opensearchservice.services.k8s.aws_domains.yaml",

	"https://github.com/open-telemetry/opentelemetry-operator/blob/main/config/crd/bases/opentelemetry.io_instrumentations.yaml",
	"https://github.com/open-telemetry/opentelemetry-operator/blob/main/config/crd/bases/opentelemetry.io_opentelemetrycollectors.yaml",

	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/operator.aquasec.com_aquacsps.yaml",
	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/operator.aquasec.com_aquadatabases.yaml",
	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/operator.aquasec.com_aquaenforcers.yaml",
	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/operator.aquasec.com_aquagateways.yaml",
	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/operator.aquasec.com_aquakubeenforcers.yaml",
	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/operator.aquasec.com_aquascanners.yaml",
	"https://github.com/aquasecurity/aqua-operator/blob/master/config/crd/bases/operator.aquasec.com_aquaservers.yaml",

	"https://github.com/cryostatio/cryostat-operator/blob/main/config/crd/bases/operator.cryostat.io_cryostats.yaml",

	"https://raw.githubusercontent.com/knative/operator/main/config/crd/bases/operator.knative.dev_knativeeventings.yaml",
	"https://raw.githubusercontent.com/knative/operator/main/config/crd/bases/operator.knative.dev_knativeservings.yaml",

	"https://github.com/open-cluster-management-io/registration-operator/blob/main/deploy/cluster-manager/config/crds/0000_01_operator.open-cluster-management.io_clustermanagers.crd.yaml",
	"https://github.com/open-cluster-management-io/registration-operator/blob/main/deploy/klusterlet/config/crds/0000_00_operator.open-cluster-management.io_klusterlets.crd.yaml",

	"https://github.com/tektoncd/operator/blob/main/config/base/300-operator_v1alpha1_chain_crd.yaml",
	"https://github.com/tektoncd/operator/blob/main/config/base/300-operator_v1alpha1_config_crd.yaml",
	"https://github.com/tektoncd/operator/blob/main/config/base/300-operator_v1alpha1_hub_crd.yaml",
	"https://github.com/tektoncd/operator/blob/main/config/base/300-operator_v1alpha1_installer_set_crd.yaml",
	"https://github.com/tektoncd/operator/blob/main/config/base/300-operator_v1alpha1_pipeline_crd.yaml",
	"https://github.com/tektoncd/operator/blob/main/config/base/300-operator_v1alpha1_trigger_crd.yaml",

	"https://github.com/projectcalico/calico/blob/master/charts/tigera-operator/crds/operator.tigera.io_apiservers_crd.yaml",
	"https://github.com/projectcalico/calico/blob/master/charts/tigera-operator/crds/operator.tigera.io_imagesets_crd.yaml",
	"https://github.com/projectcalico/calico/blob/master/charts/tigera-operator/crds/operator.tigera.io_installations_crd.yaml",
	"https://github.com/projectcalico/calico/blob/master/charts/tigera-operator/crds/operator.tigera.io_tigerastatuses_crd.yaml",

	"https://github.com/eclipse-che/che-operator/blob/main/config/crd/bases/org.eclipse.che_checlusters.yaml",

	"https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-crds/templates/policy/authorization-policy.yaml",
	"https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-crds/templates/policy/httproute.yaml",
	"https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-crds/templates/policy/meshtls-authentication.yaml",
	"https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-crds/templates/policy/network-authentication.yaml",
	"https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-crds/templates/policy/server-authorization.yaml",
	"https://github.com/linkerd/linkerd2/blob/main/charts/linkerd-crds/templates/policy/server.yaml",

	"https://github.com/CrunchyData/postgres-operator/blob/master/config/crd/bases/postgres-operator.crunchydata.com_postgresclusters.yaml",

	"https://github.com/aws-controllers-k8s/prometheusservice-controller/blob/main/config/crd/bases/prometheusservice.services.k8s.aws_alertmanagerdefinitions.yaml",
	"https://github.com/aws-controllers-k8s/prometheusservice-controller/blob/main/config/crd/bases/prometheusservice.services.k8s.aws_rulegroupsnamespaces.yaml",
	"https://github.com/aws-controllers-k8s/prometheusservice-controller/blob/main/config/crd/bases/prometheusservice.services.k8s.aws_workspaces.yaml",

	"https://github.com/quay/quay-operator/blob/master/config/crd/bases/quay.redhat.com_quayregistries.yaml",

	"https://github.com/aws-controllers-k8s/rds-controller/blob/main/config/crd/bases/rds.services.k8s.aws_dbclusterparametergroups.yaml",
	"https://github.com/aws-controllers-k8s/rds-controller/blob/main/config/crd/bases/rds.services.k8s.aws_dbclusters.yaml",
	"https://github.com/aws-controllers-k8s/rds-controller/blob/main/config/crd/bases/rds.services.k8s.aws_dbinstances.yaml",
	"https://github.com/aws-controllers-k8s/rds-controller/blob/main/config/crd/bases/rds.services.k8s.aws_dbparametergroups.yaml",
	"https://github.com/aws-controllers-k8s/rds-controller/blob/main/config/crd/bases/rds.services.k8s.aws_dbproxies.yaml",
	"https://github.com/aws-controllers-k8s/rds-controller/blob/main/config/crd/bases/rds.services.k8s.aws_dbsubnetgroups.yaml",
	"https://github.com/aws-controllers-k8s/rds-controller/blob/main/config/crd/bases/rds.services.k8s.aws_globalclusters.yaml",

	"https://github.com/Apicurio/apicurio-registry-operator/blob/main/config/crd/resources/registry.apicur.io_apicurioregistries.yaml",

	"https://github.com/cloud-bulldozer/benchmark-operator/blob/master/config/crd/bases/ripsaw.cloudbulldozer.io_benchmarks.yaml",

	"https://github.com/apache/rocketmq-operator/blob/master/deploy/crds/rocketmq.apache.org_brokers.yaml",
	"https://github.com/apache/rocketmq-operator/blob/master/deploy/crds/rocketmq.apache.org_consoles.yaml",
	"https://github.com/apache/rocketmq-operator/blob/master/deploy/crds/rocketmq.apache.org_nameservices.yaml",
	"https://github.com/apache/rocketmq-operator/blob/master/deploy/crds/rocketmq.apache.org_topictransfers.yaml",

	"https://github.com/rook/rook/blob/master/deploy/examples/crds.yaml",

	"https://github.com/aws-controllers-k8s/s3-controller/blob/main/config/crd/bases/s3.services.k8s.aws_buckets.yaml",

	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_apps.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_dataqualityjobdefinitions.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_domains.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_endpointconfigs.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_endpoints.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_featuregroups.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_hyperparametertuningjobs.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_modelbiasjobdefinitions.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_modelexplainabilityjobdefinitions.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_modelpackagegroups.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_modelpackages.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_modelqualityjobdefinitions.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_models.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_monitoringschedules.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_notebookinstancelifecycleconfigs.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_notebookinstances.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_processingjobs.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_trainingjobs.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_transformjobs.yaml",
	"https://github.com/aws-controllers-k8s/sagemaker-controller/blob/main/config/crd/bases/sagemaker.services.k8s.aws_userprofiles.yaml",

	"https://github.com/scylladb/scylla-operator/blob/master/pkg/api/scylla/v1/scylla.scylladb.com_scyllaclusters.yaml",
	"https://github.com/scylladb/scylla-operator/blob/master/pkg/api/scylla/v1alpha1/scylla.scylladb.com_nodeconfigs.yaml",
	"https://github.com/scylladb/scylla-operator/blob/master/pkg/api/scylla/v1alpha1/scylla.scylladb.com_scyllaoperatorconfigs.yaml",

	"https://github.com/quay/container-security-operator/blob/master/bundle/manifests/imagemanifestvulns.secscan.quay.redhat.com.crd.yaml",

	"https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/deploy/base-crds/crds/apparmorprofile.yaml",
	"https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/deploy/base-crds/crds/profilebinding.yaml",
	"https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/deploy/base-crds/crds/profilerecording.yaml",
	"https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/deploy/base-crds/crds/securityprofilenodestatus.yaml",
	"https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/deploy/base-crds/crds/securityprofilesoperatordaemon.yaml",
	"https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/deploy/base-crds/crds/selinuxpolicy.yaml",
	"https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/deploy/base-crds/crds/selinuxpolicy.yaml",
	"https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/deploy/base-crds/crds/seccompprofile.yaml",

	"https://github.com/sematext/sematext-operator/blob/master/deploy/crds/sematext_v1_sematextagent_crd.yaml",

	"https://github.com/redhat-developer/service-binding-operator/blob/master/config/crd/bases/servicebinding.io_clusterworkloadresourcemappings.yaml",
	"https://github.com/redhat-developer/service-binding-operator/blob/master/config/crd/bases/servicebinding.io_servicebindings.yaml",

	"https://github.com/aws-controllers-k8s/eks-controller/blob/main/config/crd/common/bases/services.k8s.aws_adoptedresources.yaml",
	"https://github.com/aws-controllers-k8s/eks-controller/blob/main/config/crd/common/bases/services.k8s.aws_fieldexports.yaml",

	"https://github.com/aws-controllers-k8s/sfn-controller/blob/main/config/crd/bases/sfn.services.k8s.aws_activities.yaml",
	"https://github.com/aws-controllers-k8s/sfn-controller/blob/main/config/crd/bases/sfn.services.k8s.aws_statemachines.yaml",

	"https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/blob/master/charts/spark-operator-chart/crds/sparkoperator.k8s.io_scheduledsparkapplications.yaml",
	"https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/blob/master/charts/spark-operator-chart/crds/sparkoperator.k8s.io_sparkapplications.yaml",

	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_ingressroutes.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_ingressroutetcps.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_ingressrouteudps.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_middlewares.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_middlewaretcps.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_serverstransports.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_tlsoptions.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_tlsstores.yaml",
	"https://github.com/traefik/traefik/blob/master/docs/content/reference/dynamic-configuration/traefik.containo.us_traefikservices.yaml",

	"https://github.com/kyverno/kyverno/blob/main/config/crds/wgpolicyk8s.io_clusterpolicyreports.yaml",
	"https://github.com/kyverno/kyverno/blob/main/config/crds/wgpolicyk8s.io_policyreports.yaml",

	"https://github.com/wildfly/wildfly-operator/blob/main/config/crd/bases/wildfly.org_wildflyservers.yaml",
}
