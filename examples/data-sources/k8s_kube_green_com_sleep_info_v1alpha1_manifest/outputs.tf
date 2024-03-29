output "manifests" {
  value = {
    "example" = data.k8s_kube_green_com_sleep_info_v1alpha1_manifest.example.yaml
  }
}
