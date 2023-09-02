data "k8s_longhorn_io_backup_target_v1beta2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
