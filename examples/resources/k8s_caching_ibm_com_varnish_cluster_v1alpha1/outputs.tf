output "resources" {
  value = {
    "minimal" = k8s_caching_ibm_com_varnish_cluster_v1alpha1.minimal.yaml
    "example" = k8s_caching_ibm_com_varnish_cluster_v1alpha1.example.yaml
  }
}
