# save your github access token in .github.token
# use `truncate -s -1 .github.token` to remove last new line
gh_token=$(cat .github.token)

echo $gh_token

curl -H "Authorization: token ${gh_token}" \
  -H 'Accept: application/vnd.github.v3.raw' \
  -o ./endpoints.json \
  -L https://raw.githubusercontent.com/edmodo/aws-endpoints/master/README.md
