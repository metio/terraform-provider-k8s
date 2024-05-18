data "k8s_groupsnapshot_storage_k8s_io_volume_group_snapshot_content_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    deletion_policy           = "Retain"
    driver                    = "some-driver"
    source                    = {}
    volume_group_snapshot_ref = {}
  }
}
