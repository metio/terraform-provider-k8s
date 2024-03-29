data "k8s_documentdb_services_k8s_aws_db_subnet_group_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
