data "k8s_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
