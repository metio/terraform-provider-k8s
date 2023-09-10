data "k8s_virt_virtink_smartx_com_virtual_machine_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
