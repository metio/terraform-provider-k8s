data "k8s_troubleshoot_sh_host_preflight_v1beta2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
