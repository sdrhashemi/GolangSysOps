<h1>Sadra Hashemi Mansoor Graph Task</h1>
<h3>1- you can run the file_writer application by executing the binary and providing the following arguments:</h3>

    
    ./file_writer {direcotry name} {number of concurrent file writers}
    - Path to Directory: relative or full path
    - Number of Writers
    

Example:
    <code>./file_writer ../direcotry 4</code>

<h3>2- then you can run the directory monitor tool to watch for any file creation, you should provide the full path to the directory where the file are being written:</h3>
    
    ./dir_monitor /home/{user}/graph-task/cmd/{directory name}

Example:
    <code>./dir_monitor /home/cedrik/graph-task/cmd/directory</code>

<br>
the file writer is working concurrently and the monitor tool is working sequentially.
I implemented graceful shutdown to shutdown the apps properly 
and also I wrote some test for the file_writer