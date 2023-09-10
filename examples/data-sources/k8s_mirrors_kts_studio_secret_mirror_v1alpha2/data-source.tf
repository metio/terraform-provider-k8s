data "k8s_mirrors_kts_studio_secret_mirror_v1alpha2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
