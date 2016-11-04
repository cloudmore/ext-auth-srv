# -*- mode: ruby -*-
# vi: set ft=ruby :
VAGRANTFILE_API_VERSION='2'
ENV['VAGRANT_DEFAULT_PROVIDER']='docker'
require 'yaml'
boxes=YAML.load_file('box.yml')

Vagrant.configure(VAGRANTFILE_API_VERSION)do|config|
  boxes.each do|box|
    config.vm.define box['name']do|srv|
      srv.vm.provider "docker" do|docker|
        puts "ADDING #{box['name']}"
        docker.name=box['name']
        docker.env=box['env']if box['env']
        docker.cmd=box['cmd']if box['cmd']
        docker.image=box['image']if box['image']
        docker.ports=box['ports']if box['ports']
        docker.create_args=box['args']if box['args']
        docker.build_dir=box['build']if box['build']
        docker.volumes=box['volumes'].each{|v|v.gsub!('`pwd`',Dir.pwd)}if box['volumes']
        box['link'].each{|link|docker.link(link)}if box['link']
      end
    end
  end
end

