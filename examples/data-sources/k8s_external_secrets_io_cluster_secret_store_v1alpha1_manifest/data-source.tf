data "k8s_external_secrets_io_cluster_secret_store_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    
  }
}
