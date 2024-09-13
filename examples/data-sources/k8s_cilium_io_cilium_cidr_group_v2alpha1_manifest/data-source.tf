data "k8s_cilium_io_cilium_cidr_group_v2alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    external_cidrs = []
  }
}
