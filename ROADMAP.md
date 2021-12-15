# pj Roadmap

Commands:

pj [--list]
    - main listing of registered projects

pj --edit
    - edit the yaml config

pj projectname
    - set up the project and run the entrypoint


## Details

pj should disown all subprocesses

pj should have no output unless wanted or verbose logging

pj should respect all environment variables

pj should respect all aliases

## Implementation

user commands get put into tmp .sh files with #! and get executed by exec.Command.
this preserves environment vars

## config

`# ~/.config/pj/projects.yaml`

```yaml
---
projects:

- name: "pj"
  directory: "$HOME/repos/pj"
  command: "code ."

- name: "uotdx"
  env:
  - TEST_USERNAME: "some_username"
  - TEST_PASSWORD: "secret1"
  command: "code $HOME/repos/uotdx"

- name: "secondproject"
  directory: "$HOME/repos/secondproject"
  pre_commands:
  - "echo $TMUX"
  - "test -f .tmux"
  command: "vv"
  post_commands:
  - "rm -f tmp/*"
```
