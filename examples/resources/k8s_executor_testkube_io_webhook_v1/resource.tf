resource "k8s_executor_testkube_io_webhook_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
