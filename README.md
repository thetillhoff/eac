# `eac`

This tool helps managing apps on the local machine. Should work platform agnostic, but is only test with linux for now.


## The story

Somehow I kept finding myself installing, updating and configuring the same apps several times over a lot of (virtual) machines.
And the worst part of it was to visit each app-homepage again and again to find out whether it is installed via apt, snap or just a simple binary I had to put into my path.

While searching for the installation instructions for an app once again I wondered if this could be automated.
My initial draft used ansible (and sadly, I invested so much time and effort that it finally made installations reusable), but it was not the right tool for the task because:

- It's overkill. I don't want to use remote execution, I don't want to setup my hostname (at least not initially), ...
- It's complicated. I don't want to lookup every damn command in the docs, when all the quick-starts already give me shell-commands to execute.
- It's time-consuming. I better not think about how much time I've invested in writing ansible-playbooks, just to install some apps (the proper way)...
- It requires a linux host-system. I do have to use windows for some tasks (or I'm just too lazy to switch the system), and that is a major pain. Thankfully chocolatey already soothes the pain a bit, but I still rely on too many powershell-scripts... <rant> If you ever develop a windows app, please make sure to at least provide instructions on how to set settings via commandline - and ~please~ don't rely too much on the windows-registry for configs & settings. </rant>

> That being said, I do like the things you can do with ansible but you cannot with eac. IMO they have different use-cases.

So when I was once again searching for some stupid installation instructions, I bore the idea to create `eac`.

## The vision

`eac`'s goal is to
- be a simple app management tool - simple as in simple to understand and simple to use. If it's not simple, it's broken.
- make it easy to install a specific version of a software - as well as latest.
- make the same apps available on all (or at least the most) platforms in the same way.
- enable the user to use 'native' scripts for each installation, configuration, uninstallation, ... - The installation instructions are already given by the app developers in that way, why would any tool require you to migrate those commands to another format (and then probably format them back before the actual execution)?
- enable users to use it in scripts. All interactive parts must be able to be configured via flags.

Out of scope is dependency management of any kind.

## The name

`eac` stands for `environment-as-code`.

> I (the author) wanted a short name and it should be unique, too. It is about "iac", but that name would not only be confusing but also not really fitting - eac is only for managing my local _environment_.


## The features

### Everyday features

- [x] `eac init` creates the folder structure (`~/.eac`) and adds the first app: `eac` itself. This includes creating the `versions.yaml` at `~/.apps/versions.yaml`.
- [x] `eac list` prints all apps that are managed via `eac` (== contained in `versions.yaml`). Per default one app per line in the format `<app>[==<installedVersion>]`. The seperator can be edited with a flag. Created by trying to run the getInstalledVersion script.
- [x] `eac status` compares the installedVersion and wantedVersion for each app and prints them.
  - [ ] improve description, and what exactly is output - yellow text means update is available, white means latest already.
- [x] `eac install <appname>[==<version>][ <appname>[==<version>]]*` installs the apps with the provided names.
  - [x] If no version for the app is specified in `versions.yaml`, add the newest.
  - [x] If not locally available, check whether the repository contains the app. If yes download it automatically and install then. If not, recommend the user to create it with `eac create`.
  - [ ] If a version is specified in the `<appname>==<version>` format, saves the version to the `versions.yaml`.
  - [ ] `eac install` (no arguments) checks whether all apps are installed as described in `versions.yaml`. If not (or the getInstalledVersion script fails), the app is installed.
  - [ ] `--latest` flag skips checking the `versions.yaml` and directly retrieves the latest version. //TODO And saves that version to the `versions.yaml`.
- [x] `eac uninstall <appname>[ <appname>]*` uninstalls the apps with the provided names.
  - [ ] Removes the version from the `versions.yaml` (if exists).
  - [ ] ~downloads appfiles automatically as well~
- [x] `eac update[ <appname>]*` checks whether updates for the provided apps are available. If yes, only the version is updated, not the app.
  - [x] `eac update <app>` should not fail if app is not installed. Instead get the latest version and store it to the `versions.yaml`
    -> Only interact with versions.yaml, no getInstalledVersion.
    - [ ] add flags `--minor` (implicit patch) and `--patch` to `upgrade` command (no flag: implicit minor and implicit patch)
  - [ ] If no name/argument is provided, all apps are checked.
  - [ ] `--quiet / -q` updates the version (if necessary) without asking the user.
- [ ] some packages in some installer variants, f.e. golang/docker (not sure) apt-get don't support installing older versions. How should they be treated?
- [ ] move to per-installer-type installation scripts instead of per-app installation scripts. That hopefully removes redundancy.
- [ ] `eac version` print the current version of `eac`.

### App maintainer features

- [x] `eac create <appname>` creates the default files and folder structure for the new app. Without additional flags only default files for the current OS are created.
  - ~`--force` overwrites existing files~ Not needed if the scripts are embedded in the executable.
  - ~`--no-default-files` disables creation of default files completely (-> only the folders are created).~ Doesn't really make sense. Why would you want that?
  - [x] `--platform [linux,darwin,windows,all]` creates the folders and default files for the specified platform(s). Multiple occurances of this flag are possible.
  - [x] `--githubUser <githubUser>` adjusts the default files so they fit for github releases. The githubUser is the owner of the repository.
  - Add helper-scripts for common installation, getLatest variants: *Check with roadmap before working on it*
    - [ ] `--apt`: "common/apt-install.sh <package-name> <repo-url> <repo-key-url>"
    - [ ] `--github`: "common/github-getLatestVersion.sh <repo-owner> <repo-name>"
    - [ ] Create helpers/scripts/sources folder, where for each type of generic tool scripts can be placed. F.e. github binary release, github tar.gz release, github zip release, apt
    - [x] Add helper-scripts for common tasks as well: Add to path, remove from path
- [ ] `eac validate <appname>[ <appname>*]` checks whether the app configurations are set up in a valid way. TODO what exactly is validated here?
- [x] `eac delete <appname>` deletes the folder structure and all contents for the specified app.
  [x] `--platform [linux,darwin,windows,all]` deletes only the folders and files for the specified platform(s). Multiple occurances of this flag are possible.
<!-- [ ] autoupdate scripts; requires checksum for each iteration of app-scripts `eac flag APPNAME=<version>` -->


## The roadmap

> "My" in the following context means the owner(s)/author(s)/member(s) of this tool. Currently this is only me and I think its sometimes easier to write from my point of view.

- Create/add some example apps for each type, good examples so far are:
  - [ ] docker
  - [ ] golang
  - [ ] git
  - [ ] eac
- ~Add Dockercontainer to registry for eac usage.~ Not possible, since too many folders would have to be shared with the host.
- Resolve `//TODO`s. There are quite a lot.
- [ ] embed the scripts in the executable
      This ensures eac has a strict set of scripts that might be executed and noone can create an attack surface out of them and
      enables to use different scripts per branch -> easier development workflow
- [ ] move from scripts per app to scripts per installer variant (apt, github-binary, github-tar-gz...) *Check with `eac create` before working on it*
      This removes the hassle of filemanaged hassle on each machine, as well as
      simplifying how to update the scripts when anything changes in them on the repo side.
- More tests
  - [ ] Write more (unit) tests.
  - [ ] Create github actions for tests
- [ ] Check & test how `eac` behaves with settings files. (And maybe use it as replacement for the custom config struct?)
- Add snapshot/restore feature
  - [ ] `eac snapshot` uses a script `saveSettings` or similar for all apps and zips the result together with the `versions.yaml`. Eases initial setup and migrations between machines (limited to the same platform, so switching between linux and windows won't be covered here - sorry).
        Only snapshots one platform obviously... maybe add a feature for combining them?
  - [ ] `--no-settings` so only versions are snapshotted.
  - [ ] automatically detect if settings for an app exist (== if it is installed?) - if not, don't save the settings instead of failing
  - [ ] `eac restore` uses a snapshot file and installs all the apps and restores the settings (`restoreSettings` script)
- Add guideline on how to write scripts for new apps.
  - [ ] use _shared scripts_
  - [ ] message before sudo commands
  - [ ] custom uninstall commands can/should be added before installation. Sometimes required, for example for golang. (Don't only overwrite, but clean before.)
        Can include older variants, as done in app `docker`. Done by each application, so settings are not affected.
  - [ ] `uninstall` should verify whether files/folders exist before attempting to delete them. Only if they exist, anything should be printed.

- Find a solution for integrating dotfiles and implement it accordingly. -> Actual featureset tbd (or is there a tool for this? maybe an idea for my next programming project?)
- Future problems:
  - Different versions of the same app might install completely differently. How to treat backwards compatability?
  - How to treat authentication? Integrate with github? What about (geolocation) distribution mirrors?
  - How to proceed when there are hundreds of apps managed with `eac` and many changes come in on a daily basis?
    - Completely different repos for all apps has to potential to introduce the "which one was it again for this app" bullshit `eac` is trying to solve. It _could_ be solved with some standardization like f.e. the repo for app 'xyz' has to be at github.com/xyz/xyz, but what happens if this is already taken / containing something else?
    - Another solution would be to have still completely different repos, which have to register at the main repo of `eac`. This would leave some basic trust on my side, and move the actual implementation/review out of my responsiblities. Others are using something like _verified_ labels for this, but I'd like to only have verified apps managable by `eac`.
    - Either way, some way of signing will be required, to ensure my trust into the repo-owner (or better phrased: my trust into the believe that the repo owner is still the repo owner after a change).
    - Wouldn't this basically be the same as copying/forking the whole external repo into the main repo? Apart from the storage, that is.
    - An optimal solution would be to have one (my) main repo, where people can "upload" to (==autocreate PRs for their app), and some form of automation takes place, which will test if it works, and if yes, automatically merge the PR.
    - OR we move to scripts per installer instead of the current scripts per app setup and only manage some metadata per app, like github-repo location, version format and the likes. -> *preferred solution for now*


## The bugs
- `~` can't be resolved in the scripts. Use `$HOME` instead.
- The default script (at least on wsl, untested on others) for script executions is `dash`. So it seems the `shell` param doesn't work properly... -> removed functionality


## Other todos

- [ ] Add autocompletions for bash, zsh (and fish?), depending on whether a .zshrc and a .bashrc exist. (_shared scripts_)
- ~`uninstall` should download appfiles automatically as well~
