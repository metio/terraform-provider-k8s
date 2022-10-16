output "resources" {
  value = {
    "minimal" = k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2.minimal.yaml
    "example" = k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2.example.yaml
  }
}
