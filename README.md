# Contributions farmer

## Step-by-step

**First of all you need to generate two github tokens.**

- [Fine-grained access token](https://github.com/settings/tokens?type=beta)
- [Classic token](https://github.com/settings/tokens)

**Create and configure the `app.env` file in the config folder.** </br>
**Farmer will not work if you don't configure these fields:**

```
ACCESS_TOKEN=
CLASSIC_TOKEN=
USER_NAME=
USER_EMAIL=

START_DATE=
END_DATE=
```

Date format: "2022-01-01T00:00:00Z". <br />
If you don't fill in ```START_DATE```, the default will be current day.

It's not neccessarry to configure other parameters, but it's on your own.

```
REPOSITORIES_PATH=
FILE_NAME=
TARGET_REPOSITORY=
REPOSITORY_PREFIX=
```


## How-to-run

- Windows

```bash
make all-w
```

- Linux

```bash
make all-l
```
