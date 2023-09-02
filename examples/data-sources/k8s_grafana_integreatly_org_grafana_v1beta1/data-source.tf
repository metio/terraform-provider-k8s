data "k8s_grafana_integreatly_org_grafana_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
