output "resources" {
  value = {
    "minimal" = k8s_apps_3scale_net_ap_icast_v1alpha1.minimal.yaml
    "example" = k8s_apps_3scale_net_ap_icast_v1alpha1.example.yaml
  }
}
