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

  (0..3).each do |n|
	config.vm.define "n#{n}" do |node|
		node.vm.hostname = "n#{n}"
		node.vm.network "private_network", ip: "172.20.20.#{n+10}"

		if n == 0 then
			node.vm.network "forwarded_port", guest: 8500, guest_ip: "172.20.20.10", host: 8500
			node.vm.network "forwarded_port", guest: 9200, guest_ip: "172.20.20.10", host: 9200
			node.vm.network "forwarded_port", guest: 5601, guest_ip: "172.20.20.10", host: 5601
		end
	
		node.vm.synced_folder "./n#{n}", "/n#{n}"
		if File.exist?("./n#{n}/provision.sh") then
			node.vm.provision "shell", path: "./n#{n}/provision.sh"
		end


		node.vm.provision "docker" do |d|
			if File.exist?("./n#{n}/Dockerfile") then
				node.vm.provision "shell",
					inline: <<-EOS
						echo 'DOCKER_OPTS="--dns 172.20.20.#{n+10}"' \
							>> /etc/default/docker && \
							service docker restart
					EOS

				d.build_image "/n#{n}", args: "-t twinfield/n#{n}"
				d.run "n#{n}", image: "twinfield/n#{n}", args: "-p 9090:9090 -p 9091:9091 -d"
				
			end
			if n == 0 then
				d.run "consul", image: "consul", 
					args: "--net=host -e 'CONSUL_LOCAL_CONFIG={\"skip_leave_on_interrupt\": true}'",
					cmd: "agent -server -bind=172.20.20.10 -client=172.20.20.10 -ui -bootstrap-expect=1"
				d.run "elasticsearch", image: "elasticsearch",
					args: "--net=host"
				d.run "kibana", image: "kibana",
					args: "--net=host -e ELASTICSEARCH_URL=http://localhost:9200"
				d.run "logstash", image: "logstash",
					args: "--net=host -v /n0:/config-dir",
					cmd: "-f /config-dir/logstash.conf"
			else
				d.run "consul", image: "consul",
					args: "-d --net=host -e 'CONSUL_LOCAL_CONFIG={\"leave_on_terminate\": true}' -v /consul/config:/consul/config",
					cmd: "agent -bind=172.20.20.#{n+10} -join=172.20.20.10"
			end
		end
		if File.exist?("./n#{n}/post-provision.sh") then
			node.vm.provision "shell", path: "./n#{n}/post-provision.sh"
		end

	end
  end

end
