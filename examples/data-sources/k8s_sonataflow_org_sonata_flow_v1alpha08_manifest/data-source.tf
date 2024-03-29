data "k8s_sonataflow_org_sonata_flow_v1alpha08_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
