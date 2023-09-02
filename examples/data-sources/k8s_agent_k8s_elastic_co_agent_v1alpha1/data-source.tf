data "k8s_agent_k8s_elastic_co_agent_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
