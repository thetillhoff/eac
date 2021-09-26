# `eac`

This tool helps managing apps on the local machine. Should work platform agnostic, but is only test with linux for now.


## The story

Somehow I kept finding myself installing, updating and configuring the same apps several times over a lot of (virtual) machines.
And the worst part of it was to visit each app-homepage again and again to find out whether it is installed via apt, snap or just a simple binary I had to put into my path.

While searching for the installation instructions for an app once again I wondered if this could be automated.
My initial draft used ansible (and sadly, I invested so much time and effort that it finally made installations reusable), but it was not the right tool for the task because:

- It's overkill. I don't want to use remoting, I don't want to setup my hostname (at least not initially), ...
- It's complicated. I don't want to lookup every damn command in the docs, when all the quick-starts already give me shell-commands to execute.
- It's time-consuming. I better not think about how much time I've invested in writing ansible-playbooks, just to install some apps (the proper way)...
- It requires a linux host-system. I do have to use windows for some tasks (or I'm just to lazy to switch the system), and that is a major main. Thankfully chocolatey already heals up half of the pain, but I still rely on too many powershell-scripts... <rant> If you ever develop a windows app, please make sure to at least provide instructions on how to set settings via commandline - and ~please~ don't rely too much on the windows-registry for configs & settings </rant>

> That being said, I do like the things you can do with ansible but you cannot with eac. IMO they have different use-cases.

So when I again searched for some stupid installation commands, I bore the idea to create eac.

## The vision

eac's goal is to
- be a simple app management tool - simple as in simple to understand and simple to use. When the input is ambiguous, ask the user.
- make it easy to install a specific version of a software.
- make the same apps available on all (or at least the most) platforms in the same way.
- enable the user to use shell commands (or even better: shell-scripts) for each installation, configuration, uninstallation, ... - The installation-instructions are already given by the app developers, why would you require users to migrate those commands to another format (and then probably format them back for the actual execution)?
- enable the user to script its usage. It should be easy to use eac in scripts and other tools. All interactive parts must be able to be configured via flags.

Out of scope is dependency management of any kind.

## The name

`eac` stands for `environment-as-code`.

I (the author) wanted a short name and it should be unique, too. It is about "iac", but that name would not only be confusing but also not properly fitting - eac is only for managing my local _environment_.


## The features

### Everyday features

- [x] `eac init` creates the folder structure (`~/.apps`) and adds the first app: eac itself. This includes creating the `versions.yaml` at `~/.apps/versions.yaml`.
- [ ] `eac list` prints all apps that are managed via `eac` (== contained in `versions.yaml`). Per default one app per line in the format `<app>[==<installedVersion>]`. The seperator can be edited with a flag.
  [ ] `eac list local` lists all apps that are available locally.
  [ ] `eac list online` lists all apps that are available online.
  [ ] `eac list all` lists all apps that are available locally and/or online.
  [ ] `eac list installed` lists only the installed apps. Created by trying to run the getLocalVersion script.
- [ ] `eac install <appname>[==<version>][ <appname>[==<version>]*]` installs the apps with the provided names. If no version for the app is specified in `versions.yaml`, add it.
  [ ] If not locally available, check whether the repository contains the app. If yes, ask the user to download it automatically and install then. If not, recommend the user to create it with `eac create`. The former can be disabled with `--offline`
  [ ] `eac install` checks whether all apps are installed as described in `versions.yaml`. If not (or the getLocalVersion script fails), the app is installed.
- [ ] `eac uninstall <appname>[ <appname>*]` uninstalls the apps with the provided names. Removes the version from the `versions.yaml` (if exists).
- [ ] `eac update[ <appname>*]` checks whether updates for the provided apps are available. If yes, only the version is updated, not the app. If no name is provided, all apps are checked.
  [ ] `eac update[ <appname>*]` checks whether updates for the provided apps are available. If yes, the user is asked whether only the version should be updated or the app should be upgraded as well. If no app is provided, all apps are checked.
  [ ] `--versions` only updates the version without asking the user.
  [ ] `--upgrade` updates the version AND installs the app in the new version.
- [ ] `eac upgrade[ <appname>*]` checks whether the desired version and the installed version of the provided apps are equal. For each app where this is not the case, install the desired version.

### App maintainer features

- [x] `eac create <appname>` creates the folder structure and default files for the new app under `apps`. Without additional flags only default files for the current OS are created.
  [ ] `--no-default-files` disables creation of default files completely (-> only the folders are created).
  [x] `--platform [linux,darwin,windows,all]` creates the folders and default files for the specified platform(s). Multiple occurances of this flag are possible.
  [x] `--githubUser <githubUser>` adjusts the default files so they fit for github releases. The githubUser is the owner of the repository.
- [ ] `eac validate <appname>[ <appname>*]` checks whether the app configurations are set up in a valid way. TODO what exactly is validated here?
- [x] `eac delete <appname>` deletes the folder structure and all contents for the specified app.
  [x] `--platform [linux,darwin,windows,all]` deletes only the folders and files for the specified platform(s). Multiple occurances of this flag are possible.


## The roadmap

> "My" in the following context means the owner(s)/author(s)/member(s) of this tool. Currently this is only me and I think its easier to write from my point of view anyway.

- `eac install` and `eac upgrade` should uninstall the old version first. At least for golang this is required. On the other hand, this would delete all settings...
  Maybe this is better added to the install script of the app itself.
- Add helper-scripts for common installation options, f.e. "common/apt-install.sh <package-name> <repo-url> <repo-key-url>", "common/github-getLatestVersion.sh <repo-owner> <repo-name>". The same could be done for getLatestVersion. This could alternatively  (or additionally) be added as param for create like `eac create <app> --github` or `eac create <app> --apt`.
  Create helpers/scripts/sources folder, where for each type of generic tool scripts can be placed. F.e. github binary release, github tar.gz release, github zip release, apt
  Add helper-scripts for common tasks as well: Add to path, remove from path
- Write down each possible command with description what is does and what it does not.
- `eac update <app>` should not fail if app is not installed. Instead get the latest version and store it to the versions.yaml
  Make it possible to upgrade currently not installed apps, without trying to get a local version to compare to. -> compare only against versionsFile
- Create/add some example apps
- support minor & major releases - if both is possible, ask the user '(and add potentially a param for that)
- Add `install --self` and `uninstall --self` commands for eac self-management. (not updating, just plain install & ununstall)
- check folder `apps` into git, and store all "configured" apps there. Later on, they might be downloaded from there as well.
- Add Dockercontainer to registry for eac usage.
- Make it possible to install apps without adding their version to the versionsFile - and with it. (Currently the versionsFile is never edited.) -> --latest
- Resolve `//TODO`s. There are quite a lot.
- add bash/zsh autocompletion
- `install --tmpFolder` parameter for specifying specific folder. Add checking for already existing files there.
- Should it be 'appsDirPath/platform/appName' or 'appsDirPath/appName/platform'? Currently it is the latter. // -> might be better, as like this the apps are selfcontained
- Write (unit) tests.
- Check & test how eac behaves with settings files.
- Create github actions for tests, release
- Create initial release (f.e. 0.0.1).
- Write proper documentation - as of now, the tool is "self-documenting" as every command has its own help-message and example. Probably not sufficient though.
- Add `snapshot` command, which tries to gather all installed apps and their respective settings, so the initial setup will be easier and it will simplify migration between machines (limited to the same platform, so switching between linux and windows won't be covered here - sorry). Improvement: exclude an app, only list(==print) found apps.
- Find a solution for integrating dotfiles and implement it accordingly. -> Actual featureset tbd
- public repository:
  - Add `download` command, which downloads the required files and (sane?) default settings for an app you don't already have. Further improvement: Make the download repository flexible via flag. And one more: only download for ~my~ platform (to save some space).
  - Add `upload` command, which creates a PR at the repository for an previously unexisting app. Again, an additional improvement is to make the repositoy flexible via flag. And also again, add an option to upload only for ~my~ platform because f.e. I might only have this single machine (at hand).
  - Usefulness:
    - This will be useful for custom apps, mirroring the repository and potential enterprise blockades like restricted internet access.
  - Problems:
    - Different versions of the same app might install completely differently. How to treat backwards compatability?
    - How to treat authentication? Integrate with github? What about the mirrors?
    - Who "pulls the trigger" and checks all those PRs? There are way too many apps out there to do this myself.
      - Codeowners might introduce problems with trust (at least IMO - I know you already have to trust ~me~ - but I'm trustworthy, am I not? ;))
      - Completely different repos for all apps has to potential to introduce the "which one was it again for this app" bullshit I'm basically trying to solve with this whole tool. It ~could~ be solved with some standardization like f.e. the repo for app 'xyz' has to be at github.com/xyz/xyz, but what happens if this is already taken / containing something else?
      - Another solution would be to have still completely different repos, which have to register at a/some main repo. This would leave the "base"-trust on my side, and move the actual implementation/review out of my todo.
      - Either way, some way of signing will be required, to ensure my trust into the repo-owner (or better phrased: my trust into the believe that the repo owner is still the repo owner after a change).
      - Wouldn't this basically be the same as copying/forking the whole external repo into the main repo? Apart from the storage, that is.
      - An optimal solution would be to have one (my) main repo, where people can "upload" to (==autocreate PRs for their app), and some form of automation takes place, which will test if it works, and if yes, automatically merge the PR.
    - How to handle updates?
    - How does everyone know how to handle settings? Wouldn't this require an additional command `snapshot-settings` or something like that?
  - Summary: This is a large step and will require quite some time and even more thoughts. For now, I'll leave it as is. This means without any `download` or `upload` commands and without a single source of truth for installing apps. I'll leave it here though, so my future self will know where to continue.

## The bugs
- `~` can't be resolved in the scripts. Use `$HOME` instead.
- The default script (at least on wsl, untested on others) for script executions is `dash`. So it seems the `shell` param doesn't work properly...
