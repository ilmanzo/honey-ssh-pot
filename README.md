# honey-ssh-pot
a fake ssh server that works as a honeypot, logging credentials of the login attempts

Curious about who and how attempts ssh login to your home server ? Me too... 
So I wrote a very simple ssh honeypot, just to collect interesting info about the kind guys who knocks my door :) 

warning: this is safe, but don't run the service (well, ANY service) as root user. Even better if you can run it as a dedicate unprivileged user.

This program is only for didactic use and not intended for deployment in a production network environment.

If you want to have it exposed on the public internet, so you must remap port 22 of your wan router to the internal server port ( 2222 by default)... 
**Do it at your risk!**


Download and compilation:

    $ git clone https://github.com/ilmanzo/honey-ssh-pot.git
    $ go get github.com/gliderlabs/ssh
    $ go build -ldflags="-s -w"
    $ ./honey-ssh-pot 

... now wait and keep an eye on your syslog ;)

    $ tail -f /var/log/messages
    ./honey-ssh-pot[24125]: 2018/06/26 21:50:09 starting ssh server on port 2222...
    ./honey-ssh-pot[24125]: 2018/06/26 21:50:40 User: evilhacker connecting from 127.0.0.1:51182 with password: secretpassword

have fun...

credits to https://github.com/gliderlabs for the awesome ssh wrapper package. Thank you guys !

