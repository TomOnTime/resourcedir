
all: beta

velma: *.go
	go build

freebsd/velma_polybeta:
	sed -e 's/CHANGE_TO_INSTANCE/polybeta/g' <freebsd/velma.sh >$@

freebsd/velma_polyprod:
	sed -e 's/CHANGE_TO_INSTANCE/polyprod/g' <freebsd/velma.sh >$@

beta: velma freebsd/velma_polybeta
	sudo install -C -m 0755 -o 0 -g 0 velma /usr/local/bin/velma_polybeta
	sudo install -C -m 0755 -o 0 -g 0 freebsd/velma_polybeta /etc/rc.d/velma_polybeta
	sudo touch /var/log/velma_polybeta.log
	sudo chown velmad:velmad /var/log/velma_polybeta.log

prod: velma freebsd/velma_polyprod
	sudo install -C -m 0755 -o 0 -g 0 velma /usr/local/bin/velma_polyprod
	sudo install -C -m 0755 -o 0 -g 0 freebsd/velma_polyprod /etc/rc.d/velma_polyprod
	sudo touch /var/log/velma_polyprod.log
	sudo chown velmad:velmad /var/log/velma_polyprod.log
