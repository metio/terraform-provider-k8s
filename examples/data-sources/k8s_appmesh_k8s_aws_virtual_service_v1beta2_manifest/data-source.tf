data "k8s_appmesh_k8s_aws_virtual_service_v1beta2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
