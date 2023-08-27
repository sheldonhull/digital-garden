#!/usr/bin/env bash

SourceImages='myrepopath/hugo/static/images'
shopt -s globstar

for file in notes/posts/**/*.md; do
    if [[ -f "$file" ]]; then
        # Search for any images/ path using a regular expression
        images=$(grep -oE 'images/[^)]+' "$file")

        # Loop through each image path
        for image in $images; do
            # Extract the file name
            filename=$(basename "$image")

            # Search for the file in the $SourceImages directory
            sourcefile=$(find "$SourceImages" -name "$filename" -print -quit)

            # If the file is found, copy it to the directory containing the markdown file
            if [[ -n "$sourcefile" ]]; then
                # Create the nested directory relative to the path of the file being parsed
                nesteddir=$(dirname "$file")/$(dirname "$image")
                mkdir -p "$nesteddir"

                # Copy the file to the nested directory
                cp "$sourcefile" "$nesteddir"

                # Output a checkmark and say "copied"
                echo -e "💾 copied $filename to $nesteddir"
            fi
        done
    fi
done
