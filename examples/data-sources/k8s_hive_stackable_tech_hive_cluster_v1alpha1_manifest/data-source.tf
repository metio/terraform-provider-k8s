data "k8s_hive_stackable_tech_hive_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      database = {
        conn_string        = "postgresql://host:port/name"
        credentials_secret = "some-secret"
        db_type            = "postgres"
      }
    }
    image = {}
  }
}
