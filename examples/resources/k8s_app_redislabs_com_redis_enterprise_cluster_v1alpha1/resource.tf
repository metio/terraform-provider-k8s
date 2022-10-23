resource "k8s_app_redislabs_com_redis_enterprise_cluster_v1alpha1" "minimal" {
  metadata = {
    name = "redis-enterprise"
  }
}

resource "k8s_app_redislabs_com_redis_enterprise_cluster_v1alpha1" "example" {
  metadata = {
    name = "redis-enterprise"
  }
  spec = {
    nodes = 3
    persistent_spec = {
      enabled            = "true"
      storage_class_name = "gp2"
    }
    ui_service_type = "LoadBalancer"
    username        = "admin@acme.com"
    redis_enterprise_node_resources = {
      limits = {
        cpu    = "400m"
        memory = "4Gi"
      }
      requests = {
        cpu    = "400m"
        memory = "4Gi"
      }
    }
    redis_enterprise_image_spec = {
      image_pull_policy = "IfNotPresent"
      repository        = "redislabs/redis"
      version_tag       = "5.4.0-19"
    }
  }
}
