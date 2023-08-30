#!/usr/bin/env bash

SourceImages='/home/sheldonhull/git/github.com/sheldonhull/sheldonhull.hugo/static'
# file=notes/posts/2021/2021-06-28-simplify-aws-developer-security-with-leapp.md
# shopt -s globstar

for file in $(find -path './notes/posts/**/*.md' -type f); do
    if [[ -f $file ]]; then
        # Search for any images/ path using a regular expression
        images=$(grep -oE 'images/[^)]+\.(png|gif|webp|jpe?g)' "$file")
        printf "🔍 images in $file: $images\n"
        # Loop through each image path
        echo $images | while read -r image; do
            # Extract the file name
            filename=$(basename "$image")
            printf "\tbasename: $filename ..."

            # Search for the file in the $SourceImages directory
            sourcefile=$(find "$SourceImages" -name "$filename" -print -quit)
            printf "\nsourcefile: $sourcefile\n"
            for fileToCopy in $sourcefile; do
                printf "\tfileToCopy: $fileToCopy ..."
                # Create the nested directory relative to the path of the file being parsed
                nesteddir="$(dirname "$file")/images"
                mkdir -p "$nesteddir" || true
                # Copy the file to the nested directory
                cp "$sourcefile" "$(realpath $nesteddir)"
                # Output a checkmark and say "copied"
                printf "💾 copied $sourcefile to $(realpath $nesteddir)\n"
            done
        done
    fi
done
