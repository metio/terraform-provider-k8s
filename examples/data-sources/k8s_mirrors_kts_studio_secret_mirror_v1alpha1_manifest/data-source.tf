data "k8s_mirrors_kts_studio_secret_mirror_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
