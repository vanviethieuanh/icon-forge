# Icon Templ Generator

This repository generates Go `templ` components
from SVG icon packs (e.g., Lucide, Feather).
Each icon pack has its own Taskfile configuration and then call
the main Taskfile for cloning, generating, and building Go modules.

## Usage

1. Clone this repository.
2. Pick an icon pack, e.g., Lucide or Feather:

    ```bash
    # Example for Feather icons
    go-task -f feather.Taskfile.yml default
    ````

3. Generated Go module will be in `build/<provider>-templ`. Use it in your projects:

    ```go
    import "github.com/vanviethieuanh/lucide_templ"

    // I created this under above repo.
    templ SomeBiggerComponent(){
        @lucide_templ.ChevronFirst()
        @lucide_templ.X(lucide_templ.Props{Size: 20, StrokeWidth: 2})
    }
    ```

## Add a new icon pack

1. Create a new Taskfile, e.g., `awesome-icons.Taskfile.yml` with appropriate variables:

   * `PROVIDER_NAME`: `lucide` for lucide icon for example.
   * `REPO_URL`: the URL to target icon pack repository.
   * `ICONS_DIR`: path to the target icons directory files from repo root.
   * `USERNAME`: github username, you can use yours or mine.

2. Run the default task to generate the templ module. You can see `lucide.Taskfile.yml`
an example

## Contributing

1. Fork the repository.
2. Add or improve Taskfiles for icon packs.
3. Test to generate repo and test it locally.
4. Submit a pull request.

## Created packages

* [Lucide Icons](https://lucide.dev/): `https://github.com/vanviethieuanh/lucide-templ`
* [Feather Icons](https://feathericons.com/): `https://github.com/vanviethieuanh/feather-templ`
