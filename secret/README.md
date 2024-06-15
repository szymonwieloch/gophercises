# CLI password manager

This simple project mimics the behaviour of popular password managers such as `LastPass` or `Bitwarden` but in CLI.
It encrypts multiple secrets and stores them in a file - by default the `~/.secrets` file.
You can set or get individual secrets by issuing CLI commands.

## Usage

```
$ ./secret -h
Usage: secret --password PASSWORD [--file FILE] <command> [<args>]

Options:
  --password PASSWORD, -p PASSWORD
                         Password required to encrypt or decrypt the file [env: PASSWORD]
  --file FILE, -f FILE   Path to the secret file where the data is stored, default: '~/.secrets'
  --help, -h             display this help and exit

Commands:
  set                    Sets a secret in the given file
  get                    Gets a secret from a secret file and prints it
```

## Example

```
$ export PASSWORD=mypass
$ secret set mysql_passwd
Provide the secret:
<enter the secret here>
$ secret get mysql_passwd
mysqlpass
$ mysql -h domain.com -u root -p$(secret get msyql_passwd)
```

## Warning

Secrets stored in the history of bash commands or environment variables may be unsafe. As this is a toy project, some compromises were made.
