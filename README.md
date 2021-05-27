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
- be a simple app management tool,
- make the same apps available on all (or at least the most important) platforms in the same way,
- enable me to use shell-scripts for each installation, configuration, uninstallation, ... - I mean the installation-instructions are already given by the app developers, why would you require users to migrate those commands to another format (and then probably format them back to the initial commands)?


## The name

`eac` stands for `environment-as-code`.

I wanted a short name and it should be unique, too. The meaning of "iac" is basically it, but I guess that would be not only confusing but also not properly fitting - eac is only for managing my local environment.


## The roadmap

> "My" in the following context means the owner(s)/author(s)/member(s) of this tool. Currently this is only me and I think its easier to write from my point of view anyway.

- Create/add some example apps
- Resolve `//TODO`s. There are quite a lot.
- Write Tests.
- Check & test how eac behaves with settings files.
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
