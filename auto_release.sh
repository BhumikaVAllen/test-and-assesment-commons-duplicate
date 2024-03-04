# Publish on github
echo "Publishing new release on Github..."

git fetch --force --tags

# Get the new release tag name
new_tag=$(git describe --tags HEAD~1 | awk 'BEGIN { FS = "." } ; { $3 = $3 + 1; print $1 "." $2 "." $3 }')

# Get the title
name=$(git log --format=%B -n 1 | head -n 1)

# Create a release
release=$(curl -XPOST -H "Authorization:token $1" --data "{\"tag_name\": \"$new_tag\", \"target_commitish\": \"main\", \"name\": \"$name\", \"body\": \"$name\", \"draft\": false, \"prerelease\": false}" https://api.github.com/repos/Allen-Career-Institute/test-and-assessment-commons/releases)
echo $release
