data "k8s_akri_sh_instance_v0_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
