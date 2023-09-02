data "k8s_org_eclipse_che_che_cluster_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
