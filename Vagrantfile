# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

ENV['VAGRANT_DEFAULT_PROVIDER'] = 'virtualbox'

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.provider "virtualbox" do |vb|
     vb.memory = "1024"
  end


  config.vm.synced_folder ".", "/vagrant", disabled: true, type: "rsync"

  #config.vm.provision "shell", path: "./scripts/install-consul.sh"

  (0..2).each do |n|
	config.vm.define "n#{n}" do |node|
		node.vm.hostname = "n#{n}"
		node.vm.network "private_network", ip: "172.20.20.#{n+10}", virtualbox__intnet: true

		#if n > 0 then
		#	node.vm.network "forwarded_port", guest: 9090+n-1, host: 9090+n-1
		#else
		#	node.vm.network "forwarded_port", guest: 8500, guest_ip: "172.20.20.10", host: 8500
		#end
	
		node.vm.synced_folder "./n#{n}", "/n#{n}"
		if File.exist?("./n#{n}/provision.sh") then
			node.vm.provision "shell", path: "./n#{n}/provision.sh"
		end

		node.vm.provision "docker" do |d|
			if File.exist?("./n#{n}/Dockerfile") then
				d.build_image "/n#{n}", args: "-t twinfield/n#{n}"
				d.run "n#{n}", image: "twinfield/n#{n}", args: "-p 909#{n-1}:909#{n-1} -d"
			end
			if n == 0 then
				d.run "consul", image: "consul", 
					args: "--net=host -e 'CONSUL_LOCAL_CONFIG={\"skip_leave_on_interrupt\": true}' -p 8500:8500",
					cmd: "agent -server -bind=172.20.20.10 -client=172.20.20.10 -ui -bootstrap-expect=1"
			else
				d.run "consul", image: "consul",
					args: "-d --net=host -e 'CONSUL_LOCAL_CONFIG={\"leave_on_terminate\": true}'",
					cmd: "agent -bind=172.20.20.#{n+10} -join=172.20.20.10"
			end
		end
	end
  end

end
