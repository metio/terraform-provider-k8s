resource "k8s_secretgenerator_mittwald_de_string_secret_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
