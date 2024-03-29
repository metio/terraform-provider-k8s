output "manifests" {
  value = {
    "example" = data.k8s_cloudwatch_aws_amazon_com_amazon_cloud_watch_agent_v1alpha1_manifest.example.yaml
  }
}
