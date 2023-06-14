#!/bin/bash
git remote add origin git@github.com:jiajie99/synchronizer.git
git add .
if git checkout -b master; then
  echo "Switched to new branch master"
else
  git checkout master
fi
git commit -m "docs: add $1.md"
git push --set-upstream origin master
