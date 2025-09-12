data "k8s_spark_stackable_tech_spark_connect_server_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
