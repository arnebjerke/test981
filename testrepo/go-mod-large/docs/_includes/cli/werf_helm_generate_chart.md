{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
Generate Werf chart which will contain a valid Helm chart to the specified path.

Werf will generate additional values files, templates Chart.yaml and other files specific to the 
Werf chart. The result is a valid Helm chart

{{ header }} Syntax

```bash
werf helm generate-chart PATH [options]
```

{{ header }} Environments

```bash
  $WERF_SECRET_KEY  Use specified secret key to extract secrets for the deploy. Recommended way to 
                    set secret key in CI-system. 
                    
                    Secret key also can be defined in files:
                    * ~/.werf/global_secret_key (globally),
                    * .werf_secret_key (per project)
```

{{ header }} Options

```bash
      --dir='':
            Change to the specified directory to find werf.yaml config
      --docker-config='':
            Specify docker config directory path. Default $WERF_DOCKER_CONFIG or $DOCKER_CONFIG or 
            ~/.docker (in the order of priority).
            Command needs granted permissions to read and pull images from the specified stages 
            storage and images repo
      --env='':
            Use specified environment (default $WERF_ENV)
  -h, --help=false:
            help for generate-chart
      --home-dir='':
            Use specified dir to store werf cache files and dirs (default $WERF_HOME or ~/.werf)
      --ignore-secret-key=false:
            Disable secrets decryption (default $WERF_IGNORE_SECRET_KEY)
  -i, --images-repo='':
            Docker Repo to store images (default $WERF_IMAGES_REPO)
      --insecure-repo=false:
            Allow usage of insecure docker repos (default $WERF_INSECURE_REPO)
      --namespace='':
            Use specified Kubernetes namespace (default [[ project ]]-[[ env ]] template or 
            deploy.namespace custom template from werf.yaml)
      --secret-values=[]:
            Specify helm secret values in a YAML file (can specify multiple)
      --ssh-key=[]:
            Use only specific ssh keys (Defaults to system ssh-agent or ~/.ssh/{id_rsa|id_dsa}, see 
            https://werf.io/reference/toolbox/ssh.html). Option can be specified multiple times to 
            use multiple keys.
  -s, --stages-storage='':
            Docker Repo to store stages or :local for non-distributed build (only :local is 
            supported for now; default $WERF_STAGES_STORAGE environment).
            More info about stages: https://werf.io/reference/build/stages_and_images.html
      --tag-custom=[]:
            Use custom tagging strategy and tag by the specified arbitrary tags. Option can be used 
            multiple times to produce multiple images with the specified tags.
      --tag-git-branch='':
            Use git-branch tagging strategy and tag by the specified git branch (option can be 
            enabled by specifying git branch in the $WERF_TAG_GIT_BRANCH)
      --tag-git-commit='':
            Use git-commit tagging strategy and tag by the specified git commit hash (option can be 
            enabled by specifying git commit hash in the $WERF_TAG_GIT_COMMIT)
      --tag-git-tag='':
            Use git-tag tagging strategy and tag by the specified git tag (option can be enabled by 
            specifying git tag in the $WERF_TAG_GIT_TAG)
      --tmp-dir='':
            Use specified dir to store tmp files and dirs (default $WERF_TMP_DIR or system tmp dir)
```

