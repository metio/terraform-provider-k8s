resource "k8s_flink_apache_org_flink_session_job_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}