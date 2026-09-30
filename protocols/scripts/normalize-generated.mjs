import { readdir, readFile, writeFile } from 'node:fs/promises'
import { join, extname } from 'node:path'

// OpenAPI Generator leaves trailing whitespace in comments and Markdown.
// Normalize it so regeneration produces clean, reviewable diffs.
async function normalize(directory) {
  for (const entry of await readdir(directory, { withFileTypes: true })) {
    const path = join(directory, entry.name)
    if (entry.isDirectory()) {
      await normalize(path)
    } else if (['.ts', '.md'].includes(extname(path))) {
      const original = await readFile(path, 'utf8')
      const clean = original.split('\n').map(line => line.trimEnd()).join('\n').trimEnd() + '\n'
      if (clean !== original) await writeFile(path, clean)
    }
  }
}

await normalize(new URL('../ts/src', import.meta.url).pathname)
