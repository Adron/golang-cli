

commitmessage=$1
branch=$2

git add -A
git commit -m '$commitmessage'
git push origin $branch
