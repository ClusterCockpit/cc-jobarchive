#!/bin/sh

if [ -d './var' ]; then
    echo 'Directory ./var already exists! Skipping initialization.'
    ./cc-backend --server --dev
else
    make

    cd var
    wget https://hpc-mover.rrze.uni-erlangen.de/HPC-Data/0x7b58aefb/eig7ahyo6fo2bais0ephuf2aitohv1ai/job-archive-dev.tar.xz
    tar xJf job-archive-dev.tar.xz
    rm ./job-archive-dev.tar.xz
    cd ../

    cp ./configs/env-template.txt .env
    cp ./docs/config.json config.json

    ./cc-backend --migrate-db
    ./cc-backend --server --dev --init-db --add-user demo:admin:AdminDev
fi
