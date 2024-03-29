data "k8s_velero_io_volume_snapshot_location_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
