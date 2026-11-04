data "k8s_core_kubeadmiral_io_collected_status_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }

  clusters         = []
  last_update_time = "2026-01-01T00:00:00Z"
}
