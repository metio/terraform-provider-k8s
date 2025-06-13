data "k8s_troubleshoot_sh_support_bundle_v1beta2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
