data "k8s_flink_apache_org_flink_deployment_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
