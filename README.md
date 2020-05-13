# jk-community project

Design doc: https://docs.google.com/document/d/1kwL6_2Fm8F3uYPlskM2Z5z_H1hqpf-OecLjr2eo2eoc/edit

Database diagram: https://jamboard.google.com/d/1Y1Km3cSgrav2-rrgNMkcscFA_hD4aisiww2vNDWQ9B8/viewer?f=0

Trello: https://trello.com/jkcommunity 

### How to build on Mac:

1. Install home-brew if you haven't
2. Install `npm` and `nodejs`: `brew install node`
3. Install `mysql`: `brew install mysql@5.6`
1. Run `brew services start mysql@5.6`
1. Roboot your Mac and then run `mysql_secure_installation`.

```shell
# login to mysql to create sample database
mysql -u root -p
# (your password created in previous step)
> create database qmplus;
> use qmplus;
> source server/db/qmplus.sql;
> \q
```

<!--
```shell
brew services start mysql@5.6
mysqladmin -u root password Aa@6447985
# login to mysql to create sample database
mysql -u root -p
Aa@6447985
> create database qmplus;
> use qmplus;
> source server/db/qmplus.sql;
> \q
```
-->



1. Install `Golang` for MacOS: https://golang.org/dl/

2. Put to .bashrc or .zshrc

   `export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin`

3. Install visual studio code with plugins (Optional)
   `Go` and `Vue`

Open a new terminal:

```shell
cd server
go build
./gin-vue-admin
```

Open another terminal:

```shell
cd web
npm install
npm audit fix
npm run serve
```

Done.

