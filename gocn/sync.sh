git remote add origin git@github.com:jiajie99/synchronizer.git
git add .
git checkout -b "sync-$1"
git commit -m "docs: add $1.md"
git push origin
