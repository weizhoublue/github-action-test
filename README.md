[![Build Image Release](https://github.com/spidernet-io/spiderpool/actions/workflows/build-release-image.yaml/badge.svg)](https://github.com/spidernet-io/spiderpool/actions/workflows/build-release-image.yaml)

[![Lint Go-lint checks](https://github.com/spidernet-io/spiderpool/actions/workflows/lint-golang.yaml/badge.svg)](https://github.com/spidernet-io/spiderpool/actions/workflows/lint-golang.yaml)

#requirement

secret.PAT 

## ========== workflow ==================

# feature

## (1) go-lint for source code

## (2) codeql check

## (3) codeowners , who reivew PR

## (4) auto package base image

## (5) auto package release image by tag

## (6) CI build

build CI image for each PR or PUSH to main , to check your code is right , 

and then auto gc the image at intarval ( at now only image of orgs is supported , personal image failed)

the go pkg cache could accelerate the build , 

the cache is cleaned auto at interval , or by manual

## (7) build beta image by manual

## (8) issue manage

issue template , auto assign label and assignees

auto stale inactive issue after 60 and auto close stale issue after 14

## (8) PR manage

CODEOWNERS auto assign reviewer ( the one who should be the repo member and not me , then you can see it )

auto close stale PR after 60 and auto close stale issue after 14

label check. only when the pr labled with changelog-related label, the pr could be approved. 
changelog-related label could be used to auto generating changelog when releasing

## (9) check chart under charts dierctory 

## (10) auto generating changelog to /changelogs/***

when tag or dispatch by manual , auto generate changelog by the related-label history PR between tags,
then commit the pr to main branch , then auto approve it

## (11) pr label and auto changelog

### all pr should label one of bellowing , so could be merged to changelog

pr/release/bug, pr/release/feature-new, pr/release/feature-changed,
pr/release/doc, pr/release/robot_changelog, pr/release/changelog,
pr/release/none-required

### other label for pr 



### below label of pr will be the changelog

pr/release/bug for Fixes

pr/release/feature-new for New Features

pr/release/feature-changed for Changed Features

## (12) label syncer

## auto add lable 'pr/approved,pr/need-release-label' to reviewed PR

we can get all issiue labeld with "pr/need-release-label" 
and label them with release-related lables , and merge it

## check license missing in go file

## auto publish chart

## use "chart" branch as github page and provide chart repo


## ========== makefile ==================

#### auto generate license of vendor

#### auto globally modify go version in files

#### 


## ============================

# manage flow

### for issue

check these labeled with "issue/not-assign" , and assign them

auto close stale one after 60 and auto close stale one after 14


### for PR

2 reviewer approve , and auto label with "pr/approved" , 

PR must be labeled with "pr/release/**" for merge , which used to generate changelog auto

auto close stale PR after 60 and auto close stale issue after 14

### for Release

if push tag with v*.*.* , 

(1) auto trigger building release image
(2) auto commit changelog PR with label "pr/release/robot_changelog"
(3) auto commit chart PR with label "pr/release/robot_chart", to /docs/charts/* of "chart" branch , and generate /docs/index.yaml , please set /docs of "chart" branch as github page
(4) auto create the release

pr who labeled with following label, will exist in changelog of release:

label "pr/release/feature-new" will generate "New Features" category

label "pr/release/bug" will generate "Fixes" category

label "pr/release/feature-changed" will generate "Changed Features" category

