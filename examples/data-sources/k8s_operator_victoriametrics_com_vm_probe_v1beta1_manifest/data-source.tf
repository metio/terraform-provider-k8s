data "k8s_operator_victoriametrics_com_vm_probe_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    vm_prober_spec = {
      url = "some-url"
    }
  }
}
