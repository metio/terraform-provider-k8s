output "manifests" {
  value = {
    "example" = data.k8s_kinesis_services_k8s_aws_stream_v1alpha1_manifest.example.yaml
  }
}
