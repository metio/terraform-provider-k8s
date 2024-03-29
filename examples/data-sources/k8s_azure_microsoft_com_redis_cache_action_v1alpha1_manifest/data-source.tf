data "k8s_azure_microsoft_com_redis_cache_action_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
