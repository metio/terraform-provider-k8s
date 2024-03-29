data "k8s_dex_gpu_ninja_com_dex_user_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
