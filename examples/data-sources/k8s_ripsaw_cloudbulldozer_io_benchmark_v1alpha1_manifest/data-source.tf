data "k8s_ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
