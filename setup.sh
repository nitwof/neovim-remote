#!/bin/sh
set -e

cd /tmp
WORK_PATH=/neovim
NVIM_PATH=${WORK_PATH}/nvim
PYTHON_PATH=${WORK_PATH}/python
mkdir ${WORK_PATH}

curl -L "https://github.com/niess/python-appimage/releases/download/python3.9/python3.9.6-cp39-cp39-manylinux2014_x86_64.AppImage" -o /tmp/python.appimage
chmod u+x ./python.appimage
./python.appimage --appimage-extract
mv ./squashfs-root ${PYTHON_PATH}
${PYTHON_PATH}/AppRun -m pip install --no-cache --upgrade pynvim

curl -L "https://github.com/neovim/neovim/releases/download/v0.5.0/nvim.appimage" -o /tmp/nvim.appimage
chmod u+x ./nvim.appimage
./nvim.appimage --appimage-extract
mv ./squashfs-root ${NVIM_PATH}
ln -s ${NVIM_PATH}/AppRun /usr/bin/nvim

XDG_CONFIG_HOME=${WORK_DIR}/config
XDG_DATA_HOME=${WORK_DIR}/data
NVIM_LOG_FILE=${WORK_DIR}/nvim.log

### init.vim
# let g:loaded_python_provider = 0
# let g:python3_host_prog = '/neovim/python/AppRun'
