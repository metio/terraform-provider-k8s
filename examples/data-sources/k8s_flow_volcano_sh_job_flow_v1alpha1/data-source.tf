data "k8s_flow_volcano_sh_job_flow_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}