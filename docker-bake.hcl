variable "VERSION" {
    default = ""
}

target "vault-plugin-catalog" {
    args = {
      "VERSION" = "${VERSION}"
    }
}

target "default" {
    tags = [ "localhost:5000/soramitsukhmer-lab/vault-plugin-catalog:dev" ]
}

target "release" {
    inherits = [ "vault-plugin-catalog" ]
    platforms = [
        "linux/amd64",
        "linux/arm64",
    ]
    tags = [ "harbor.sorakh.io/soramitsukhmer-lab/vault-plugin-catalog:${VERSION}" ]
}

target "binaries" {
    inherits = [ "vault-plugin-catalog" ]
    output = [ "./binaries" ]
    platforms = [ "local" ]
    target = "binaries"
}
