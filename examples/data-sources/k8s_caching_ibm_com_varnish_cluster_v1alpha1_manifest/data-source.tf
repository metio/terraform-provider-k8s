data "k8s_caching_ibm_com_varnish_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
