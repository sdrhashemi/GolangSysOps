<h1>Sadra Hashemi Mansoor Graph Task</h1>
for gmail saying my zip file contained virus i can't send the binary of my applications.
you can build it on your own via:

    cd cmd/file_writer
    go build -o file_writer
    and
    cd cmd/dir_monitor
    go build -o dir monitor
    and also you can run the app via:
     go run . <params> in each module respectively

<h3>1- you can run the file_writer application by executing the binary (built for linux ofc) and providing the following arguments:</h3>

    
    ./file_writer.bin {direcotry name} {number of concurrent file writers}
    or
    go run . {direcotry name} {number of concurrent file writers}
    - Path to Directory: relative or full path
    - Number of Writers
    

Example:
    <code>./file_writer.bin ../direcotry 4</code>
    <code>go run . ../direcotry 4</code>

<h3>2- then you can run the directory monitor tool binary (built for linux ofc)  to watch for any file creation, you should provide the full path to the directory where the file are being written:</h3>
    
    ./dir_monitor.bin /home/{user}/graph-task/cmd/{directory name}
    or
    go run . /home/{user}/graph-task/cmd/{directory name}

Example:
    <code>./dir_monitor.bin /home/cedrik/graph-task/cmd/directory</code>
    <code>go run . /home/cedrik/graph-task/cmd/directory</code>

<br>
the file writer is working concurrently and the monitor tool is working sequentially.
I implemented graceful shutdown to shutdown the apps properly 
and also I wrote some test for the file_writer