         ___        ______     ____ _                 _  ___  
        / \ \      / / ___|   / ___| | ___  _   _  __| |/ _ \ 
       / _ \ \ /\ / /\___ \  | |   | |/ _ \| | | |/ _` | (_) |
      / ___ \ V  V /  ___) | | |___| | (_) | |_| | (_| |\__, |
     /_/   \_\_/\_/  |____/   \____|_|\___/ \__,_|\__,_|  /_/ 
 ----------------------------------------------------------------- 


Working on Go, Scala and Prolog development

Go - cloud9 has built in Go compiler
    go run hello_world.go
    
Prolog
    sudo yum install gprolog
    
    gplc FILE
    ./FILE

Scala - need to install Scala compiler and jdk 1.8

    Adapted from https://community.c9.io/t/creating-a-scala-app/1605
    
    sudo yum remove java-1.7.0-openjdk
    sudo yum install java-1.8.0-openjdk
    wget https://downloads.lightbend.com/scala/2.12.4/scala-2.12.4.tgz
    tar xvf scala-2.12.4.tgz
    mv scala-2.12.4 ../scala-2.12.4
    nano ../.bash_profile
        Append to end of file:
            SCALA_HOME="$HOME/scala-2.12.4"
            PATH="$PATH:$SCALA_HOME/bin"
        
    Save and open new terminal, scalac then scala to compile and run    
