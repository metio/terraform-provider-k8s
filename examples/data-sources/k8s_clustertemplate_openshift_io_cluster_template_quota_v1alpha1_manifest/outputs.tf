output "manifests" {
  value = {
    "example" = data.k8s_clustertemplate_openshift_io_cluster_template_quota_v1alpha1_manifest.example.yaml
  }
}
