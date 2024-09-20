data "k8s_airflow_stackable_tech_airflow_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      credentials_secret = "some-secret"
    }
    image = {}
  }
}
