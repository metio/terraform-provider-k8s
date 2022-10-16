output "resources" {
  value = {
    "minimal" = k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3.minimal.yaml
    "example" = k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3.example.yaml
  }
}
