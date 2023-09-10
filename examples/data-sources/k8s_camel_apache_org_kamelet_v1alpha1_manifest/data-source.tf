data "k8s_camel_apache_org_kamelet_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    definition = {
      title       = "Telegram Text Source"
      description = <<-TEXT
Receive all text messages that people send to your telegram bot.

# Instructions
Description can include Markdown and guide the final user to configure the Kamelet parameters.
TEXT
      required    = ["botToken"]
      properties = {
        botToken = {
          title         = "Token"
          description   = "The token to access your bot on Telegram"
          type          = "string"
          x-descriptors = ["urn:alm:descriptor:com.tectonic.ui:password"]
        }
      }
    }
    types = {
      out = {
        media_type = "text/plain"
      }
    }
    template = {
      from = {
        uri = "telegram:bots"
        parameters = {
          authorization_token = "#property:botToken"
        }
        steps = [
          {
            convert-body-to = {
              type       = "java.lang.String"
              type-class = "java.lang.String"
              charset    = "UTF8"
            }
          },
          {
            filter = {
              simple = "$${body} != null"
            }
          },
          {
            log = "$${body}"
          },
          {
            to = "kamelet:sink"
          }
        ]
      }
    }
  }
}
