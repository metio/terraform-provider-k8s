output "manifests" {
  value = {
    "example" = data.k8s_edp_epam_com_nexus_blob_store_v1alpha1_manifest.example.yaml
  }
}
