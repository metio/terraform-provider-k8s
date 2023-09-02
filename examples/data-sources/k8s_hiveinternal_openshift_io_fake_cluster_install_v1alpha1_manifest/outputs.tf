output "manifests" {
  value = {
    "example" = data.k8s_hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest.example.yaml
  }
}
