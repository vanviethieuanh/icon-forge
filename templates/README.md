# {{.Env.ICON_NAME}} Templ Icons

[![release](https://img.shields.io/github/release/{{ .Env.ICON_REPO }})](https://github.com/{{ .Env.ICON_REPO }}/releases)
[![LICENSE](https://img.shields.io/badge/LICENCE-GPL3-red?style=flat)](./LICENSE)

This module provides all **{{.Env.ICON_NAME}}** icons as [templ](https://templ.guide) components.
Icons are generated automatically using [icon-forge](https://github.com/vanviethieuanh/icon-forge).

## Versioning

This commit using:

![Templ Version](https://img.shields.io/badge/templ-{{ .Env.TEMPL_VERSION }}-blue)
![Icons](https://img.shields.io/badge/{{ .Env.ICON_NAME }}-{{ .Env.ICON_PACK_VERSION }}-orange)

## Installation

Tags format as: `v<{{.Env.ICON_NAME}}-version>-templ<templ-version>`.

So if you want to install
{{.Env.ICON_NAME}} version `0.0.1` which generated using templ version `0.0.2`,
then you likely want to install with tag `v0.0.1-templ0.0.2`.

Install the main branch, which is `latest` release of both templ and {{.Env.ICON_NAME}} using:

```bash
go get github.com/{{.Env.ICON_REPO}}
```

## Usage Example

```go
import "github.com/{{.Env.ICON_REPO}}"

templ SomeBiggerComponent(){
    @{{.Env.ICON_PACK}}.ChevronFirst()
    @{{.Env.ICON_PACK}}.X(lucide_templ.Props{Size: 20, StrokeWidth: 2})
}
```
