data "k8s_cilium_io_cilium_pod_ip_pool_v2alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
