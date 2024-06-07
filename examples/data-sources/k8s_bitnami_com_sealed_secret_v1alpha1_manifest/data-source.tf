data "k8s_bitnami_com_sealed_secret_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    encrypted_data = {
      "key" = "value"
    }
  }
}
