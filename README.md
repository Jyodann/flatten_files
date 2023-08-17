# Go Flat - Flatten your directories Easily

## Introduction

Imagine you have a directory like the one in `test_folder`:
```
test_folder
    ├───folder_1
    │       img_1.png
    │       img_2.png
    │       img_3.png
    │
    ├───folder_2
    │       img_1.png
    │       img_2.png
    │
    └───folder_3
        │   img_1.png
        │
        └───folder_3_1
                img_1.png
```

Running this app allows you to essentially "flatten" the file structure, so for instance `test_folder/folder_1/img_1.png` will be a file called `folder_1 img_1.png` when the app is run inside `test_folder` 

Results of running the program are located in the `test_folder` directory in this repository.

This app will be made available for Windows, MacOS and Linix

## Plans

I would like to make it more customizable, so currently I have not made a built version release yet. Here are the following planned changes:

- [x] Customise path the program runs on
- [x] Customise Seperator for File names
- [ ] Allow for Sequential renaming of files
- [ ] Potentially expose a small GUI to make the app more friendly to use
- [ ] Place a 'cache' folder to let program pick up where it left off (incase of sudden crashes)
- [ ] Silent mode when flattening structure
