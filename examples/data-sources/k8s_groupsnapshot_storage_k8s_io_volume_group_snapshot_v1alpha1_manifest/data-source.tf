data "k8s_groupsnapshot_storage_k8s_io_volume_group_snapshot_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    source = {}
  }
}
