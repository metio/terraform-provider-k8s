resource "k8s_bitnami_com_sealed_secret_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
