data "k8s_elasticache_services_k8s_aws_cache_subnet_group_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
