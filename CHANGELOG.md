# CHANGELOG

## v0.2.0 on 2023-10-14
- Update release pipeline
- Add check pipeline for PRs
- Adjust version command slightly

## v0.1.0 on 2022-04-04
- complete rework of `eac`: Many things were only patchwork-level before, since the vision changed over time. While the new `eac` only supports a handful of apps right now, its architecture is now more stable, safer to use and the app maintenance is way easier.
- The currently supported apps are:
  - eac
  - github-cli
  - golang
  - helm
  - kind
  - kubectl
  - terraform

## v0.0.5 on 2022-02-13
- bugfixes
- log improvements
- `install` now verifies whether the app is available online before trying and failing
- Added at least one success messag for actions, so commands are never completely silent

## v0.0.4 on 2022-01-11
- `install` now properly downloads app manifests, instead of only a subset

## v0.0.3 on 2022-01-11
- `install` now downloads app manifests automatically
- under the hood there was a truckload of changes and minor adjustments

## v0.0.2 on 2021-09-27
- instead of previously using a relative path at `./*` for the versionsFile and the appsDir, eac now uses `~/.eac/*`
- reworked default files structure
- several reworks on how several cmds like `create` work under the hood
- new feature structure in `README.md`

## v0.0.1 on 2021-09-23
- initial release
- added github actions release workflow
