# LXC Tumbler

LXC Tumbler is an LXC orchestrator that allows you to start or stop containers based on their dependencies.

## Usage

```sh
lxc-tmblr [commands]
```

### Available Commands

```sh
-start    Start the container with all the dependencies

-stop     Stop the container with all the dependencies
```

### Flags

```sh
-n, -no-dependencies    Ignore dependencies and work only with specified containers
```

### Examples

Start `foo`, `bar` and stop `baz`

```sh
lxc-tmblr -start foo -start bar -stop baz
```

## Defining dependencies

LXC Tumbler supports configuration files in YAML format. You can place config in one of the following files:

- `/etc/lxc-tmblr/config.yml`
- `$HOME/.config/lxc-tmblr/config.yml`
- `./config.yml`

### Config example

```yaml
socket: /var/snap/lxd/common/lxd/unix.socket

# A list of containers with dependencies.
# root is the actual name of the container. If root is not specified,
# the actual name of the container is equal to the name of the partition.
#
# dependencies are the names of containers or sections that the current
# container depends on.
containers:

   foo:
      dependencies:
         - bar
         - baz-with-subdependencies

   baz-with-subdependencies:
      root: baz
      dependencies:
         - bar
         - quux
```
