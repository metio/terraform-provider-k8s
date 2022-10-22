output "resources" {
  value = {
    "minimal" = k8s_gloo_solo_io_settings_v1.minimal.yaml
  }
}
