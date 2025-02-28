# Binary Installation From Release

This is the binary installation guide for GoToSocial. It is assumed that you already have a [properly configured VPS running in the cloud, or a suitable homeserver that is accessible with port forwarding](index.md).

## 1: Prepare VPS

In a terminal on the VPS or your homeserver, make the directory that GoToSocial will run from, the directory it will use as storage, and the directory it will store LetsEncrypt certificates in:

```bash
mkdir /gotosocial && mkdir /gotosocial/storage && mkdir /gotosocial/storage/certs
```

If you don't have root permissions on the machine, use something like `~/gotosocial` instead.

## 2: Download Release

In a terminal on the VPS or your homeserver, cd into the base directory for GoToSocial you just created above:

```bash
cd /gotosocial
```

Now, download the latest GoToSocial release archive corresponding to the operating system and architecture you're running on.

For example, for version 0.1.0 running on 64-bit Linux:

```bash
wget https://github.com/superseriousbusiness/gotosocial/releases/download/v0.1.0/gotosocial_0.1.0_linux_amd64.tar.gz
```

Then extract it:

```bash
tar -xzf gotosocial_0.1.0_linux_amd64.tar.gz
```

This will put the `gotosocial` binary in your current directory, in addition to the `web` folder, which contains assets for the web frontend, and an `example` folder, which contains a sample configuration file.

## 3. Edit Configuration File

Copy the configuration file from the example folder into your current directory:

```bash
cp ./example/config.yaml .
```

Now open the file in your text editor of choice so that you can set some important configuration values. Change the following settings:

* Set `host` to whatever hostname you're going to be running the server on (eg., `example.org`).
* Set `port` to `443`.
* Set `db-type` to `sqlite`.
* Set `db-address` to `sqlite.db`.
* Set `storage-local-base-path` to the storage directory you created above (eg., `/gotosocial/storage`).
* Set `letsencrypt-cert-dir` to the certificate storage directory you created above (eg., `/gotosocial/storage/certs`).

The above options assume you're using SQLite as your database. If you want to use Postgres instead, see [here](../configuration/database.md) for the config options.

## 4: Run the Binary

You can now run the binary.

Start the GoToSocial server with the following command:

```bash
./gotosocial server start --config-path ./config.yaml
```

The server should now start up and you should be able to access the splash page by navigating to your domain in the browser. Note that it might take up to a minute or so for your LetsEncrypt certificates to be created for the first time, so refresh a few times if necessary.

Note that for this example we're assuming that we're allowed to run on port 443 (standard https port), and that nothing else is running on this port.

## 5: Create and confirm your user

You can use the GoToSocial binary to also create, confirm, and promote your user account.

Run the following command to create a new account:

```bash
./gotosocial admin account create --config-path ./config.yaml --username some_username --email some_email@whatever.org --password SOME_PASSWORD
```

In the above command, replace `some_username` with your desired username, `some_email@whatever.org` with the email address you want to associate with your account, and `SOME_PASSWORD` with a secure password.

Run the following command to confirm the account you just created:

```bash
./gotosocial admin account confirm --config-path ./config.yaml --username some_username
```

Replace `some_username` with the username of the account you just created.

If you want your user to have admin rights, you can promote them using a similar command:

```bash
./gotosocial admin account promote --config-path ./config.yaml --username some_username
```

Replace `some_username` with the username of the account you just created.

## 6. Login

You should now be able to log in to your instance using the email address and password of the account you just created. We recommend using [Pinafore](https://pinafore.social) or [Tusky](https://tusky.app) for this.

## 7. Install the Admin Control Panel (optional)

At some point you'll likely want to do things like change instance information, and block domains you don't want to interact with. See the [admin panel](../admin/admin_panel.md) instructions for this.
