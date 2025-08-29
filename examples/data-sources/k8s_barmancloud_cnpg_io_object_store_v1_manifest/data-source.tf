data "k8s_barmancloud_cnpg_io_object_store_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
