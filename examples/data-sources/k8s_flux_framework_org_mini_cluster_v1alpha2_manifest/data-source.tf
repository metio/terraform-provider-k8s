data "k8s_flux_framework_org_mini_cluster_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
