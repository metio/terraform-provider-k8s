resource "k8s_hive_openshift_io_checkpoint_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
