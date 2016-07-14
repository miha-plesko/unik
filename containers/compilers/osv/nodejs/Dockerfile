FROM ubuntu:14.04

# Install prerequisites
RUN apt-get update -y
RUN apt-get install -y qemu wget curl build-essential


# Copy files for docker
COPY bin /bin

# Copy files for osv
COPY osv_files /osv_files

# Create mount point (important: must be after copying files for OSv)
VOLUME /project_directory



# Pull basic capstan images and packages
RUN /bin/capstan -u https://mikelangelo-capstan.s3.amazonaws.com/ pull mike/osv-loader
RUN /bin/capstan -u https://mikelangelo-capstan.s3.amazonaws.com/ package pull eu.mikelangelo-project.osv.bootstrap
RUN /bin/capstan -u https://mikelangelo-capstan.s3.amazonaws.com/ package pull eu.mikelangelo-project.app.node-4.4.5


# Compose boot image and copy it to /project_directory (unik expects it there)
ENTRYPOINT cd /project_directory && \
	cp /osv_files/* . -R && \
	perl -pi -e 's/\/\/CALL_NODE_MAIN_HERE/require("\.\/app.js")/g' ./node-wrapper-udp.js && \
	ls -altrh . && \
	/bin/capstan package compose --run "node /node-wrapper-udp.js" --verbose unik/node-image && \		
	qemu-img convert -f qcow2 -O qcow2 -o compat=0.10 /root/.capstan/repository/unik/node-image/node-image.qemu /project_directory/boot.qcow2

# /bin/capstan run -p qemu unik/node-image && \	
# cp /root/.capstan/repository/unik/node-image/node-image.qemu /project_directory/boot.qcow2 && \
# qemu-img convert -f qcow2 -O qcow2 -o compat=0.10 /root/.capstan/instances/qemu/unik-node-image/disk.qcow2 /project_directory/boot.qcow2

#/bin/capstan run -p qemu unik/node-image & sleep 30 && echo Converting && \		
#qemu-img convert -f qcow2 -O qcow2 -o compat=0.10 /root/.capstan/repository/unik/node-image/node-image.qemu /project_directory/boot.qcow2

#run this container with
#docker run --rm -e NODE_MAIN_FILE=app.js -v "/your/project/path/:/project_directory" projectunik/compilers-osv-nodejs:v0.0
