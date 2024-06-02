# multirename

CLI tool renaming multiple files at once.

## Usage

```bash
$ ./multirename -h
Usage: multirename --name NAME [--recursive] [--filter FILTER] [--dry] [DIRS [DIRS ...]]

Positional arguments:
  DIRS                   List of directories to process. By default uses the current directory.

Options:
  --name NAME, -n NAME   New nae for the files
  --recursive, -r        Recursive lookup.
  --filter FILTER, -f FILTER
                         Filters files using the file name. Uses common bash syntax.
  --dry, -d              Dry run.
  --help, -h             display this help and exit
```
## Example

Find recursively all `*.jpg` files in the `sample` directory and rename them to `Party.jpg` with an added number (per directory)

```
$ ./multirename -n Party  -r -f '*.jpg' sample
Renaming 'sample/nested/IMG_ABC.jpg' into 'sample/nested/Party (1 of 3).jpg'
Renaming 'sample/nested/IMG_DEF.jpg' into 'sample/nested/Party (2 of 3).jpg'
Renaming 'sample/nested/IMG_GHI.jpg' into 'sample/nested/Party (3 of 3).jpg'
```