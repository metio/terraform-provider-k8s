resource "k8s_scheduling_sigs_k8s_io_pod_group_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
