data "k8s_hive_openshift_io_sync_identity_provider_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
