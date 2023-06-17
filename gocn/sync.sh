#!/bin/bash

# Set Git username and email
git config --global user.email "your-email@example.com"
git config --global user.name "Your Name"

# Initialize Git repository and add all files
git init
git add .

# Set Git remote URL and fetch changes
git remote add origin git@github.com:jiajie99/synchronizer.git
git fetch

# Switch to the master branch, or create it if it doesn't exist
if git checkout -b master; then
  echo "Switched to new branch master"
else
  git checkout master
fi

# Pull any changes from remote master branch
git pull

# Commit changes with message and push to remote repository
git commit -m "docs: add $1.md"
git push --set-upstream origin master
