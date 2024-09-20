data "k8s_spark_stackable_tech_spark_history_server_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    image              = {}
    log_file_directory = {}
    nodes = {
      role_groups = {}
    }
  }
}
