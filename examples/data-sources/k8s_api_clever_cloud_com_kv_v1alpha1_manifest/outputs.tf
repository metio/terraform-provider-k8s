output "manifests" {
  value = {
    "example" = data.k8s_api_clever_cloud_com_kv_v1alpha1_manifest.example.yaml
  }
}
