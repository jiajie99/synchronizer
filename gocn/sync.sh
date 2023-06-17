#!/bin/bash
git init
git add .
git remote add origin git@github.com:jiajie99/synchronizer.git
if git checkout -b master; then
  echo "Switched to new branch master"
else
  git checkout master
fi
git pull
git commit -m "docs: add $1.md"
git push --set-upstream origin master
