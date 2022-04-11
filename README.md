# `eac`

[![Go Report Card](https://goreportcard.com/badge/github.com/thetillhoff/eac)](https://goreportcard.com/report/github.com/thetillhoff/eac)

This tool helps managing apps on the local machine. Should work platform agnostic, but is only test with linux for now.

## The story

Somehow I kept finding myself installing, updating and configuring the same apps several times over a lot of (virtual) machines.
And the worst part of it was to visit each app-homepage again and again to find out whether it is installed via apt, snap or just a simple binary I had to put into my path.

While searching for the installation instructions for an app once again I wondered if this could be automated.
My initial draft used ansible (and sadly, I invested a lot of time and effort to finally make installations reusable and idempotent), but in the end I decided it was not the right tool for the task because:

- It's overkill. I don't want to use remote execution, I don't want to setup my hostname (at least not initially), ...
- It's complicated. I don't want to lookup every damn command in the docs, when most quick-start instructions look so simple!
- It's time-consuming. I better not think about how much time I've invested in writing ansible-playbooks, just to install some apps...
- It requires a linux host-system. I do have to use windows for some tasks (or I'm just too lazy to switch the system), and that is a major pain. Thankfully chocolatey/scoop already sooth the pain a bit, but I still rely on too many powershell-scripts... <rant> If you ever develop a windows app, please make sure to at least provide instructions on how to set settings via commandline - or ~please~ don't rely on the windows-registry for configs & settings. </rant>

> That being said, I do like the things you can do with ansible but you cannot with `eac`. IMO they have different use-cases.

So when I was once again searching for some stupid installation instructions, I bore the idea to create `eac`.

## The vision

`eac`'s goal is to
- be a simple app management tool - simple as in simple to understand and simple to use. If it's not simple, it's broken.
- make it easy to install a specific version of a software - including latest.
- make it easy to switch between app versions.
- make the same apps available on all (or at least the most) platforms in the same way.
- make adding new apps _fun_. If maintaining it becomes work, it's broken.
- be script-friendly. All "interactive" parts must be configurable via flags.
- support dotfiles. Desired apps, their desired versions and all flags must be configurable with a dotfile in the users `$HOME`.

Out of scope is dependency management of any kind.

## The name

`eac` stands for `environment-as-code`.

> The idea was to have a short and unique name. It is about "iac", but that name would not only be confusing but also not really fitting - since `eac` is only for managing my local _environment_.

# The commands

- `eac list available`: List all apps that can be managed with eac.
- `eac list downloaded`: List all apps that were downloaded and therefore cached by `eac`.
- `eac list downloaded versions`: List all apps with all cached versions.
- `eac install`: Install all apps as specified in config file.
- `eac install <appname>[@<version>]*`: Install specified apps. If version is specified, install in specified version.
<!-- - `eac upgrade` / `eac check`: Check if newer versions are available. -->
<!-- - `eac upgrade <appnames> [-y]`: Check specified apps for updates and apply them. -->
<!-- - `eac uninstall <appnames`: Uninstall specified apps. -->
- `eac clean`: Removes all caches files.
- `eac clean <appnames>`: Remove cached files for app.
- `eac version`: Print `eac`'s version.

# The config file
The config file is located at `~/.eac` and YAML-formatted.

You can define flags like `verbose`, `dry-run` and `latest` in there, as well as a list of apps (each optionally with a desired version).

Here is an example of its contents:
```
verbose: true
apps:
- kubectl
- terraform@1.1.7

```

# The roadmap

- `eac install [<appname>[==<version>]]*` installs the specified apps with versions.
  - [ ] `--latest` flag skips checking config file and directly retrieves the latest version.
- (Automatic) tests
  - [ ] Write unit tests.
  - [ ] Create github actions for unit tests
  - [ ] Write an automated test that verifies each app (`eac list available`?).
  - [ ] Add that automated test to github actions
- [ ] `eac status` compares the installedVersion and wantedVersion for each app and prints them.
      Example:
      ```
      [OK] kubectl
      [1.2.3] terraform@1.3.2
      [OK] eac@1.2.3
      ```
  - [ ] Find a solution on how to detect version of app. (`/tmp/eac/versions.yaml` which is updated every time something is installed? Or introduce version command for each app)
- `eac uninstall <appname>[ <appname>]*` uninstalls the apps with the provided names.
  - [ ] Removes the app.Destination file/folder
  - [ ] Removes the version from the `~/.eac` (if exists && if listed there).
  - [ ] `uninstall` should verify whether files/folders exist before attempting to delete them. (If not, print something like "<appname> is not installed")
- Allow further configuration after the installation
  - [ ] Allow addition of environment variable via .profile
  - [ ] Allow addition of paths to PATH environment variable via .profile
  - [ ] Automatic setup of autocompletion for all installed shells.
- Add snapshot/restore feature
  - [ ] `eac snapshot versions` stores the currently active version in `~/.eac`.
  - [ ] add `settingsLocation` field to apps (multiple paths possible, so dotfiles can be added), so the following works:
  - [ ] `eac snapshot apps` zips/tars the result together with the `.eac.yaml`/`versions.yaml`. Eases initial setup and migrations between machines (limited to the same platform, so switching between linux and windows won't be covered here - sorry).
        Only snapshots one platform obviously... maybe add a feature for combining them?
  - [ ] automatically detect if settings for an app exist (== if it is installed?) - if not, don't save the settings instead of failing
  - [ ] `eac restore` uses a snapshot file (zip/tar) and installs all the apps according to the contained `.eac.yaml`/`versions.yaml` and (if they are part of the zip/tar) restores the settings for each.
- [ ] some packages in some installer variants, f.e. golang/docker (not sure) apt-get don't support installing older versions. How should they be treated?
- 12 factor app todos
  - [ ] concurrency of installations
  - [ ] ensure that no matter on which part the app stops/crashes, the state is as valid as possible (not possible if killed during copying files)
- Add version-profiles. Instead of having one set of app-versions in the config file, support multiple sets of versions that can be grouped and enable switching between those groups

# The apps

You can get all available apps via `eac list available`, but here are some examples:
- `eac`
- `golang`
- `helm`
- `kubectl`
- `terraform`

## How can I add app _xyz_ to `eac`?

You can always open an GitHub issue with a request.

Apart from that, the apps are defined in [`/pkg/eac/internal/apps/<appname>.go`](./pkg/eac/internal/apps/).

Feel free to open a Pull-Request where you added your app. Don't forget to mention it in the section for the next version in the [`CHANGELOG.md`](./CHANGELOG.md).

For further information, please read [`DEVELOPMENT.md`](./DEVELOPMENT.md).
