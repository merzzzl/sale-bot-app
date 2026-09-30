#!/bin/sh
set -eu

if [ "$#" -ne 3 ]; then
  echo "usage: $0 <readme.md> <src> <lang>" >&2
  exit 2
fi

README=$1
SRC=$2
LANG=$3

BEGIN_MARKER="<!-- BEGIN:${SRC} -->"
END_MARKER="<!-- END:${SRC} -->"

[ -f "$README" ] || { echo "readme not found: $README" >&2; exit 1; }
[ -f "$SRC" ] || { echo "src not found: $SRC" >&2; exit 1; }

tmp="${README}.tmp.$$"

awk -v begin="$BEGIN_MARKER" -v end="$END_MARKER" -v src="$SRC" -v lang="$LANG" '
  BEGIN { skip=0; found_begin=0; found_end=0 }
  $0 == begin {
    found_begin=1
    print
    if (lang != "") print "```" lang; else print "```"
    while ((getline line < src) > 0) print line
    close(src)
    print "```"
    skip=1
    next
  }
  $0 == end {
    found_end=1
    skip=0
    print
    next
  }
  skip == 0 { print }
  END {
    if (found_begin == 0 || found_end == 0) exit 3
  }
' "$README" > "$tmp" || {
  code=$?
  rm -f "$tmp"
  if [ "$code" -eq 3 ]; then
    echo "markers not found in $README: $BEGIN_MARKER / $END_MARKER" >&2
  fi
  exit "$code"
}

mv "$tmp" "$README"