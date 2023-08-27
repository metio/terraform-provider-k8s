output "manifests" {
  value = {
    "example" = data.k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2_manifest.example.yaml
  }
}
