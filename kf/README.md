## File commands

### Usage:

-   Example:

```bash
# FILE COMMANDS
$ kf somefile.txt # creates new file named somefile.txt
$ kf create somefile.txt # creates new file named somefile.txt (alias: c)
$ kf create somefile1.txt somefile2.txt # creates multiple files (alias: c)
$ kf create ./somefolder/somefile.txt # creates new file named somefile.txt in somefolder (creates folder if it doesn't exist)
$ kf somefile.txt somefile2.txt # renames/moves somefile.txt to somefile2.txt
$ kf somefile.txt ./somefolder/somefile.txt # moves somefile.txt to somefolder (creates folder if it doesn't exist)
$ kf rename somefile.txt somefile2.txt # renames somefile.txt to somefile2.txt (alias: r)
$ kf delete somefile.txt # deletes somefile.txt (alias: d)
$ kf delete somefile1.txt somefile2.txt # deletes multiple files (alias: d)

# FOLDER COMMANDS
$ kf somefolder # creates a new folder named somefolder

```
