data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./model",
    "--dialect", "postgres"
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/16/dev"
  url = "postgres://dev:dev@localhost:5432/orion_pay?sslmode=disable"
  migration {
    dir = "file://migration"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}