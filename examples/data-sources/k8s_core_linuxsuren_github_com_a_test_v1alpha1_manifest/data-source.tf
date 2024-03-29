data "k8s_core_linuxsuren_github_com_a_test_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
