#!/usr/bin/env bash

# Parameters for commit message and branch to commit to.
commitmessage=$1
branch=$2

# Commit and push to X branch.
git add -A
git commit -m '$commitmessage'
git push origin $branch

