data "k8s_secretgenerator_mittwald_de_basic_auth_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
