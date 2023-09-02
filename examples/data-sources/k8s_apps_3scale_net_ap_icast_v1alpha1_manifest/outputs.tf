output "manifests" {
  value = {
    "example" = data.k8s_apps_3scale_net_ap_icast_v1alpha1_manifest.example.yaml
  }
}
