#!/bin/bash

# Display usage information
usage() {
    echo "Usage: $0 -i <license_file> -l <line_number> -d <directory>"
    exit 1
}

# Check if the correct number of arguments is provided
if [ "$#" -ne 6 ]; then
    usage
fi

# Parse the input arguments
while getopts ":i:l:d:" opt; do
    case ${opt} in
        i)
            license_file=${OPTARG}
            ;;
        l)
            line_number=${OPTARG}
            ;;
        d)
            target_directory=${OPTARG}
            ;;
        \?)
            echo "Invalid option: $OPTARG" 1>&2
            usage
            ;;
        :)
            echo "Option -$OPTARG requires an argument." 1>&2
            usage
            ;;
    esac
done

# Check if the provided license file exists
if [ ! -f "$license_file" ]; then
    echo "License file $license_file not found!"
    exit 1
fi

# Create a backup directory within the specified directory
backup_dir="$target_directory/backup"
mkdir -p "$backup_dir"

# Find all .go files in the target directory and its subdirectories
find "$target_directory" -type f -name "*.go" | while read -r go_file; do
    # Create a backup of the original file
    cp "$go_file" "$backup_dir/$(basename "$go_file").bak"

    # Insert the content of license.txt at the specified line number
    awk -v n="$line_number" -v license_file="$license_file" 'NR==n{while((getline line < license_file) > 0) print line}1' "$go_file" > "$go_file.tmp" && mv "$go_file.tmp" "$go_file"
done

echo "License inserted and backup created in $backup_dir."
