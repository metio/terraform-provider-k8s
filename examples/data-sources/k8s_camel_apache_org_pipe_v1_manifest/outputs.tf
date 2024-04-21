output "manifests" {
  value = {
    "example" = data.k8s_camel_apache_org_pipe_v1_manifest.example.yaml
  }
}
