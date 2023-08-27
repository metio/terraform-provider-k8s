data "k8s_secret_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  data = {
    username = "admin"
    password = "P4ssw0rd"
  }
  type = "kubernetes.io/basic-auth"
}
