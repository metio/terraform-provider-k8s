data "k8s_reliablesyncs_kubeedge_io_object_sync_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
