# summon-file

File provider for [Summon](https://github.com/cyberark/summon).

Download the [latest release](https://github.com/cyberark/summon-file/releases) and extract it to the directory `/usr/local/lib/summon`.

## Usage in isolation

Give summon-file a variable identifier (file path) and it will fetch it for you and print the value to stdout.

```sh-session
$ summon-file path/to/access_key_id
8h9psadf89sdahfp98
```

## Usage as a provider for Summon

[Summon](https://github.com/cyberark/summon/) is a command-line tool that reads a file in secrets.yml format and injects secrets as environment variables into any process. Once the process exits, the secrets are gone.

*Example*

As an example let's use the `env` command: 

Following installation, define your keys in a `secrets.yml` file

```yml
AWS_ACCESS_KEY_ID: !var path/to/aws/iam/user/robot/access_key_id
AWS_SECRET_ACCESS_KEY: !var path/to/aws/iam/user/robot/secret_access_key
```

By default, summon will look for `secrets.yml` in the directory it is called from and export the secret values to the environment of the command it wraps.

Wrap the `env` in summon:

```sh
$ summon --provider summon-file env
...
AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxxxxx
AWS_SECRET_ACCESS_KEY=yyyyyyyyyyyyyyyy
...
```

`summon` resolves the entries in secrets.yml with the file provider and makes the secret values available to the environment of the command `env`.


## Contributing

We welcome contributions of all kinds to this repository. For instructions on how to get started and descriptions of our development workflows, please see our [contributing
guide][contrib].

[contrib]: CONTRIBUTING.md
