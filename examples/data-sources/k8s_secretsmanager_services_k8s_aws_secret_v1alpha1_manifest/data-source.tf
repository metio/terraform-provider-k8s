data "k8s_secretsmanager_services_k8s_aws_secret_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
