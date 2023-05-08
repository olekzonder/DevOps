#!/bin/bash
rm -rf /var/astra-chroot
mkdir /var/astra-chroot
debootstrap --variant=minbase --include locales,lsb-release --components=main,contrib,non-free orel /var/astra-chroot http://dl.astralinux.ru/astra/stable/2.12_x86-64/repository
cat << EOF | chroot /var/astra-chroot
echo "ru_RU.UTF-8 UTF-8" >> /etc/locale.gen
echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen
rm /etc/resolv.conf
locale-gen
update-locale ru_RU.UTF-8
mv /etc/apt/sources.list /etc/apt/sources.list.bak
apt update
mv /etc/apt/sources.list.bak /etc/apt/sources.list
apt-get autoclean
apt clean
exit
EOF
docker rmi astralinux:orel -f
tar -C /var/astra-chroot -cpf - . | docker import - astralinux:orel --change "ENV PATH /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin" --change 'CMD ["/bin/bash"]'  --change "ENV LANG=ru_RU.UTF-8"
