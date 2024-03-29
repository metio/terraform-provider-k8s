output "manifests" {
  value = {
    "example" = data.k8s_pipes_services_k8s_aws_pipe_v1alpha1_manifest.example.yaml
  }
}
