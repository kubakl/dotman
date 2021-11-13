# Dotman
## Description
A command line tool for managing dotfiles. It allows you to easily control all of your dotfiles using only few simple commands.
## Usage
Adding a new dotfile:
```shell
dotman add ~/.config/nvim/init.vim NVIM
```
Showing all your dotfiles:
```shell
dotman show 

[
  {
    Dotfile's name: NVIM
    Original path: /home/kuba/.config/nvim/init.vim
    Dotfile's path: /home/kuba/.dotfiles/init.vim
    Creation date: 13/11/2021 19:59
  }
]
```
Removing a dotfile:
```shell
dotman remove NVIM
```
