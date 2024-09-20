data "k8s_druid_stackable_tech_druid_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    brokers = {
      role_groups = {}
    }
    cluster_config = {
      deep_storage = {}
      metadata_storage_database = {
        conn_string = "postgresql://host:port/name"
        db_type     = "postgresql"
        host        = "some-host"
        port        = 12345
      }
      zookeeper_config_map_name = "some-name"
    }
    coordinators = {
      role_groups = {}
    }
    historicals = {
      role_groups = {}
    }
    image = {}
    middle_managers = {
      role_groups = {}
    }
    routers = {
      role_groups = {}
    }
  }
}
