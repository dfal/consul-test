sudo mkdir /consul/config -p
sudo cp /n2/*.json /consul/config/

sudo apt-get install -y dnsmasq
echo "conf-dir=/etc/dnsmasq.d" | sudo tee -a /etc/dnsmasq.conf
echo "server=/consul/127.0.0.1#8600" | sudo tee /etc/dnsmasq.d/10-consul
sudo service dnsmasq restart
