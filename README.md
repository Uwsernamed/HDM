# HDM
A folder hierarchy mapper.

Example:
  Using the executable:
```
HDF Manager
version: 1.0.0

HDM C:\Users\ADMIN\Desktop\HDM\executables | cd F:\Files\Documents\User\Coding
Successfully changed directory.

HDM F:\Files\Documents\User\Coding | create
File 'F:\Files\Documents\User\Coding/Coding.HDF' created successfully.

HDM F:\Files\Documents\User\Coding |
```
  Coding.HDF   (after easy text removal):
```
[Github]
    [Uwsernamed]
        [HDM]
            [executables]
            [src]
[Go]
    [HDM]
[Python]
[Rust]
```
## Table of contents:
0. **IMPORTANT!**
1. **Why?**
2. **Features.**
3. **How to obtain the executable.**
4. **How to use.**
5. **Thank you message.**

## IMPORTANT!
This program is currently only supported for Windows AMD64. Support for other operating systems may be added in the future. Please note that HDM is still in beta and more updates are planned.

## Why?
The purpose of this app is to provide users with a convenient solution for generating an easy-to-read file that maps out the locations of folders and their corresponding purposes. With the ever-increasing complexity of file systems, it can be challenging to keep track of the various folders and their intended uses. By utilizing this app, users can effortlessly generate a human-readable file that serves as a visual representation of the folder hierarchy and their assigned purposes. This mapping file can be easily updated and edited by the user, ensuring that it remains up-to-date and reflective of the current state of their file system. This app aims to simplify folder organization, enhance productivity, and facilitate better collaboration among users.

## Features.
1. **Generation of maps:** Maps can be generated by using the create command after using it in the right path. These files will take on the parent folder's name + an added '.HDM' (which may or may not be visable. The file is in a standard text format, though each layer from the imidiate children adds indentation and each folder name is incased in square brackets.

## How to obtain the executable.
### Pre-compiled program:
The more simpler method for any supported platform, these can be found in the 'executables' folder in the main project.
### To compile the program:
1. Go to the 'src' folder.
2. Download the Go coding language from 'https://go.dev/dl/' (remember to complete the processes that entail to have 'go' (the program) accessable from any shell. 
3. Download the 'main.go' file from 'https://github.com/Uwsernamed/HDM/tree/main/src'
4. Place the 'main.go' in a folder (preferably where, if a new file is created it can be easily found).
5. Open a shell, make the path of the shell be the path of the folder storing the 'main.go'.
6. run `go build [filename.go]` (in the case the file is called 'main.go' run `go build main.go`.
7. take the new file generated from the command.

## How to use.
### The ways to initialize:
1. **Move & use:** One the simplist ways is to move the program to where ever you want the effects of the commands to take place.
2. **From a shell (more permanent):** Another way is to store the executable in a safe folder, now add the program's path to where the shell can access them; by adding it to the systems path.
3. **From a shell (tedious):** Open a shell and type the program's path.
4. **Turn the program into a profile (like a shell):** Open a terminal that can add profiles, follow the steps to (do it)(, hopefully it won't be too hard).

### How to take help:
1. **Type `help`**

### How to exit:
1. **Type `exit`**

### Thank you message.
Thank you for using HDM. If you need any assistance, feel free to ask someone.
