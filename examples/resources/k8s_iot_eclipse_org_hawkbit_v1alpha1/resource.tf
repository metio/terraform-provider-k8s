resource "k8s_iot_eclipse_org_hawkbit_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_iot_eclipse_org_hawkbit_v1alpha1" "example" {
  metadata = {
    name = "default"
  }
  spec = {
    database = {
      embedded = {}
    }
    rabbit = {
      managed = {}
    }
  }
}
