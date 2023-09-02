output "manifests" {
  value = {
    "example" = data.k8s_keda_sh_cluster_trigger_authentication_v1alpha1_manifest.example.yaml
  }
}
