data "k8s_scheduling_sigs_k8s_io_elastic_quota_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
