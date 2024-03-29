data "k8s_charts_flagsmith_com_flagsmith_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
