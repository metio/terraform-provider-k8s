data "k8s_grafana_integreatly_org_grafana_datasource_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
