---
title: Publish
sidebar: reference
permalink: reference/registry/publish.html
author: Timofey Kirillov <timofey.kirillov@flant.com>
---

<!--Docker images should be pushed into the docker registry for further usage in most cases. The usage includes these demands:-->

<!--1. Using an image to run an application (for example in kubernetes). These images will be referred to as **images for running**.-->
<!--2. Using an existing old image version from a docker registry as a cache to build a new image version. Usually, it is default behavior. However, some additional actions may be required to organize a build environment with multiple build hosts or build hosts with no persistent local storage. These images will be referred to as **distributed images cache**.-->

<!--## What can be published-->

<!--The result of werf [build commands]({{ site.baseurl }}/cli/build/build.html) is a _stages_ in _stages storage_ related to images defined in the `werf.yaml` config. -->
<!--Werf can be used to publish either:-->

<!--* Images. These can only be used as _images for running_. -->
<!--These images are not suitable for _distributed images cache_, because werf build algorithm implies creating separate images for _stages_. -->
<!--When you pull a image from a docker registry, you do not receive _stages_ for this image.-->
<!--* Images with a stages cache images. These images can be used as _images for running_ and also as a _distributed images cache_.-->

<!--Werf pushes image into a docker registry with a so-called [**image publish procedure**](#image-publish-procedure). Also, werf pushes stages cache of all images from config with a so-called [**stages publish procedure**](#stages-publish-procedure).-->

<!--Before digging into these algorithms, it is helpful to see how to publish images using Docker.-->

<!--### Standard publish procedure-->

Normally in the Docker world to publish an already built arbitrary docker image, the following steps are required in general:

 1. Get a name or an id of the local built image.
 2. Create new temporary image name alias for this image, which consists of two parts:
     - [docker repository name](https://docs.docker.com/glossary/?term=repository) with embedded docker registry address;
     - [docker tag name](https://docs.docker.com/glossary/?term=tag).
 3. Push image by newly created alias into docker registry.
 4. Delete temporary image name alias.

This process will be referred to as **standard publish procedure**. There is a docker command for each of these steps, and usually, they are performed by calling corresponding docker commands.

## Image publish procedure

To publish a image from the config werf implements the **image publish procedure**. It consists of the following steps:

1. Perform [**werf tag procedure**]({{ site.baseurl }}/reference/registry/image_naming.html#werf-tag-procedure) for built image. The result of werf tag is an image with a name that is compatible with the step 2 of _standard publish procedure_. I.e., this image is ready to be published.
2. Push newly created image into docker registry.
3. Delete temporary image created in the 1'st step.

All of these steps are performed with a single werf publish command, which will be described below.

The result of this procedure is a image named by the [image naming]({{ site.baseurl }}/reference/registry/image_naming.html) rules pushed into the docker registry.

<!--### Stages publish procedure-->

<!--To publish stages cache of a image from the config werf implements the **stages publish procedure**. It consists of the following steps:-->

<!-- 1. Create temporary image names aliases for all docker images in stages cache, so that:-->
<!--     - [docker repository name](https://docs.docker.com/glossary/?term=repository) is a `REPO` parameter specified by the user without changes ([details about `REPO`]({{ site.baseurl }}/reference/registry/image_naming.html#repo-parameter)).-->
<!--     - [docker tag name](https://docs.docker.com/glossary/?term=tag) constructed as a signature prefixed with a word `image-stage-` (for example `image-stage-41772c141b158349804ad27b354247df8984ead077a5dd601f3940536ebe9a11`).-->
<!-- 2. Push images by newly created aliases into docker registry.-->
<!-- 3. Delete temporary image names aliases.-->

<!--All of these steps are also performed with a single werf command, which will be described below.-->

<!--The result of this procedure is multiple images from stages cache of image pushed into the docker registry.-->

## Publish command

{% include /cli/werf_publish.md %}

## Build and publish command

{% include /cli/werf_build_and_publish.md %}
