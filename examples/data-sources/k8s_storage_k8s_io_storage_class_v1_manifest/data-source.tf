data "k8s_storage_k8s_io_storage_class_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  k8s_provisioner = "kubernetes.io/gce-pd"
}
