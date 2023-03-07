#!/bin/bash
# Traverse current working directory and add files to git

# Check for files that need to be added to git
for f in *; do
    if [[ $(git status --porcelain | grep "$f") ]]; then
        # Add file to git
        git add "$f"

        # Prompt for commit message
        echo "Enter commit message for file $f"
        read commit_message

        # Commit changes
        git commit -m "$commit_message"
    fi
done
