output "manifests" {
  value = {
    "example" = data.k8s_caching_ibm_com_varnish_cluster_v1alpha1_manifest.example.yaml
  }
}
