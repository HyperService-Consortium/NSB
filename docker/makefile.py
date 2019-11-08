#!/usr/bin/env python3
import os, shutil
from pymake import require_cls, oqs, entry, pipe

class Makefile:
    current_path = os.path.dirname(os.path.realpath(__file__))
    build_path = os.path.join(current_path, 'build')
    node_image_path = os.path.join(current_path, 'image')
    nsb_file_path = os.path.join(build_path, 'NSB')
    tendermint_file_path = os.path.join(build_path, 'tendermint')
    node_name = 'tendermint-nsb/node'
    count = 4
    compose_file = os.path.join(current_path, 'testnode4.yml')
    compose_run_file = os.path.join(current_path, 'testnode4.run.yml')
    
    @classmethod
    @require_cls('nsb_source', 'tendermint')
    def image(cls, *_):
        pipe('docker build --tag %s node' % Makefile.node_name)

    @classmethod
    def nsb_source(cls, *_):
        shutil.copy(Makefile.nsb_file_path, Makefile.node_image_path)

    @classmethod
    def tendermint(cls, *_):
        os.makedirs(Makefile.build_path, exist_ok=True)
        if not os.path.isfile(Makefile.tendermint_file_path):
            pipe('curl -o %s -L https://github.com/HyperService-Consortium/NSB/releases/download/v0.7.4/tendermint-linux-v0.32.2-8ba8497a' % (Makefile.tendermint_file_path, ))

    @classmethod
    @require_cls('image', 'template')
    def build(cls, *_):
        os.makedirs(Makefile.build_path, exist_ok=True)
        if not os.path.isfile(os.path.join(Makefile.build_path, 'node0/config/genesis.json')):
            pipe('docker run --rm -v %s/:/tendermint:Z %s testnet --v %s --o . --populate-persistent-peers --starting-ip-address 192.167.233.2' %
                (Makefile.build_path, Makefile.node_name, Makefile.count))
        pipe('docker-compose -f %s up' % (Makefile.compose_run_file))

    @classmethod
    def template(cls, *_):
        with open(Makefile.compose_file) as f:
            s = f.read().replace('{{build}}', Makefile.build_path + '/')
            with open(Makefile.compose_run_file, 'w') as o:
                o.write(s)

    @classmethod
    @require_cls('template')
    def down(cls, *_):
	    pipe('docker-compose -f %s down' % (Makefile.compose_run_file)) 
        
    @classmethod
    @require_cls('template')
    def start(cls, *_):
	    pipe('docker-compose -f %s start' % (Makefile.compose_run_file)) 
        
    @classmethod
    @require_cls('template')
    def stop(cls, *_):
	    pipe('docker-compose -f %s stop' % (Makefile.compose_run_file)) 

    @classmethod
    @require_cls('template')
    def restart(cls, *_):
	    pipe('docker-compose -f %s restart' % (Makefile.compose_run_file)) 


    @classmethod
    def clean(cls, *_):
        pipe('rm -rf -r %s/node*' % (Makefile.build_path))
        pipe('rm -rf -r %s/data*' % (Makefile.build_path))
        pipe('rm -rf -r %s/nsbstate.db' % (Makefile.build_path))
        pipe('rm -rf -r %s/trienode.db' % (Makefile.build_path))

    @classmethod
    @require_cls('build')
    def all(cls, *_):
        pass

if __name__ == '__main__':
    entry(Makefile)
