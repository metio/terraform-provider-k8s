resource "k8s_rds_services_k8s_aws_db_cluster_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
