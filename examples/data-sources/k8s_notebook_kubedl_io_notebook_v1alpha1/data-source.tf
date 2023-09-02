data "k8s_notebook_kubedl_io_notebook_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
