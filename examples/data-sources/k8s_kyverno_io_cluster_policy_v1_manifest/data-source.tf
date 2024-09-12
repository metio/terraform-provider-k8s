data "k8s_kyverno_io_cluster_policy_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    rules = [
      {
        name = "some-rule"
        context = [
          {
            name = "response"
            api_call = {
              method = "POST"
              data = [
                {
                  key   = "images"
                  value = jsonencode("some-string")
                },
              ]
            }
          }
        ]
      }
    ]
  }
}

data "k8s_kyverno_io_cluster_policy_v1_manifest" "int_value" {
  metadata = {
    name = "some-name"
  }
  spec = {
    rules = [
      {
        name = "some-rule"
        context = [
          {
            name = "response"
            api_call = {
              method = "POST"
              data = [
                {
                  key   = "images"
                  value = jsonencode(123)
                },
              ]
            }
          }
        ]
      }
    ]
  }
}

data "k8s_kyverno_io_cluster_policy_v1_manifest" "bool_value" {
  metadata = {
    name = "some-name"
  }
  spec = {
    rules = [
      {
        name = "some-rule"
        context = [
          {
            name = "response"
            api_call = {
              method = "POST"
              data = [
                {
                  key   = "images"
                  value = jsonencode(true)
                },
              ]
            }
          }
        ]
      }
    ]
  }
}

data "k8s_kyverno_io_cluster_policy_v1_manifest" "array_value" {
  metadata = {
    name = "some-name"
  }
  spec = {
    rules = [
      {
        name = "some-rule"
        context = [
          {
            name = "response"
            api_call = {
              method = "POST"
              data = [
                {
                  key   = "images"
                  value = jsonencode([123, 456, 789])
                },
              ]
            }
          }
        ]
      }
    ]
  }
}

data "k8s_kyverno_io_cluster_policy_v1_manifest" "map_value" {
  metadata = {
    name = "some-name"
  }
  spec = {
    rules = [
      {
        name = "some-rule"
        context = [
          {
            name = "response"
            api_call = {
              method = "POST"
              data = [
                {
                  key   = "images"
                  value = jsonencode({"a": "b", "c": "d"})
                },
              ]
            }
          }
        ]
      }
    ]
  }
}

data "k8s_kyverno_io_cluster_policy_v1_manifest" "mixed_value" {
  metadata = {
    name = "some-name"
  }
  spec = {
    rules = [
      {
        name = "some-rule"
        context = [
          {
            name = "response"
            api_call = {
              method = "POST"
              data = [
                {
                  key   = "images"
                  value = jsonencode({"a": [123, 456], "c": {"d": true, "e": "f", "g": 789}})
                },
              ]
            }
          }
        ]
      }
    ]
  }
}
