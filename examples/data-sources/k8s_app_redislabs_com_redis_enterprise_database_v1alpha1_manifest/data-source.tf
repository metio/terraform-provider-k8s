data "k8s_app_redislabs_com_redis_enterprise_database_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
