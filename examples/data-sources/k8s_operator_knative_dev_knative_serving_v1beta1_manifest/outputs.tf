output "manifests" {
  value = {
    "example" = data.k8s_operator_knative_dev_knative_serving_v1beta1_manifest.example.yaml
  }
}
