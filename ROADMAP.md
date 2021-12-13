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

## config

`# ~/.config/pj/projects.yaml`

```yaml
---
projects:
    # I want to open the 'pj' project
    # that lives in 'dir'
    # and run the command
    pj:
        dir: "$HOME/repos/pj"
        command: "code ."
    secondproject:
        dir: "$HOME/repos/secondproject"
        navigate: true
        pre_commands:
        - "echo $TMUX"
        - "test -f .tmux"
        command: "vv"
        post_commands:
        - "rm -f tmp/*"
```
