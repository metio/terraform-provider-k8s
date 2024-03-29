data "k8s_k8up_io_pre_backup_pod_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
