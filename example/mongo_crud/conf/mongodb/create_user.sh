mongo -u "app_root" -p "password" --authenticationDatabase admin \
    --eval "var username = 'app', password = 'password', database = 'accounts';" \
    /home/user.js