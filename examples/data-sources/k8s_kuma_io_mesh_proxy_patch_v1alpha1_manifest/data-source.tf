data "k8s_kuma_io_mesh_proxy_patch_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
