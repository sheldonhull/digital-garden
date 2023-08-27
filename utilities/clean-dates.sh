#!/usr/bin/env bash

shopt -s globstar

for file in notes/blog/posts/**/*.md; do
    # file=notes/blog/posts/2019/2019-04-20t03-48-12-00-00.md
    if [[ ! -f "$file" ]]; then
        echo "No files found"
        return
    fi
    # Extract the date value from the front matter
    date=$(grep -oE 'date: ["'']?[0-9]{4}-[0-9]{2}-[0-9]{2}' "$file" | cut -d' ' -f2)
    datenotcleaned=$(grep -oE 'date: ["'']?[0-9]{4}-[0-9]{2}-[0-9]{2}.*$' "$file")
    echo "date: $date datenotcleaned: $datenotcleaned"
    # regex="^[0-9]{4}-[0-9]{2}-[0-9]{2}.*$"

    # if date != datenotcleaned then run
    if [[ "$date" != "$datenotcleaned" ]]; then
        echo "needs to be cleaned: $file"
        cleanedDate=$(echo "$date" | grep -oE '[0-9]{4}-[0-9]{2}-[0-9]{2}')

        # Replace the date value in the file
        sd "date: .*$" "date: $cleanedDate" "$file"
    else
        echo "✔️ already clean"
    fi
done
