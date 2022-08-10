---
title: "devspace cleanup images --help"
sidebar_label: devspace cleanup images
---


Deletes all locally created images from docker

## Synopsis

 
```
devspace cleanup images [flags]
```

```
#######################################################
############# devspace cleanup images #################
#######################################################
Deletes all locally created docker images from docker
#######################################################
```


## Flags

```
  -h, --help   help for images
```


## Global & Inherited Flags

```
      --debug                        Prints the stack trace if an error occurs
      --disable-profile-activation   If true will ignore all profile activations
      --inactivity-timeout int       Minutes the current user is inactive (no mouse or keyboard interaction) until DevSpace will exit automatically. 0 to disable. Only supported on windows and mac operating systems
      --kube-context string          The kubernetes context to use
  -n, --namespace string             The kubernetes namespace to use
      --no-warn                      If true does not show any warning when deploying into a different namespace or kube-context than before
      --override-name string         If specified will override the devspace.yaml name
  -p, --profile strings              The DevSpace profiles to apply. Multiple profiles are applied in the order they are specified
      --silent                       Run in silent mode and prevents any devspace log output except panics & fatals
  -s, --switch-context               Switches and uses the last kube context and namespace that was used to deploy the DevSpace project
      --var strings                  Variables to override during execution (e.g. --var=MYVAR=MYVALUE)
```

