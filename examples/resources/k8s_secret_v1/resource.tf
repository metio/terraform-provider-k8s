resource "k8s_secret_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_secret_v1" "example" {
  metadata = {
    name = "test"
  }
  data = {
    username = "admin"
    password = "P4ssw0rd"
  }
  type = "kubernetes.io/basic-auth"
}
