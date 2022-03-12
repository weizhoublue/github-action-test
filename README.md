[![Build Image Release](https://github.com/spidernet-io/spiderpool/actions/workflows/build-release-image.yaml/badge.svg)](https://github.com/spidernet-io/spiderpool/actions/workflows/build-release-image.yaml)

[![Lint Go-lint checks](https://github.com/spidernet-io/spiderpool/actions/workflows/lint-golang.yaml/badge.svg)](https://github.com/spidernet-io/spiderpool/actions/workflows/lint-golang.yaml)


feature

(1) go-lint for source code

(2) codeql check

(3) codeowners , who reivew PR

(4) auto package base image

(5) auto package release image by tag

(6) build CI image for each PR or PUSH to main , to check your code is right , 

and then auto gc the image at intarval ( at now only image of orgs is supported , personal image failed)