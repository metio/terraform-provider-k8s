data "k8s_dynamodb_services_k8s_aws_global_table_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
