data "k8s_app_redislabs_com_redis_enterprise_cluster_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
