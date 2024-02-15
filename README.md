From the week 2 track you have grasped lots of important low level and core functionalities
in Go. To put all these to the test we have decided to give you a task of designing a simple
REPL Database - what can be more low level than that.
The database only operates on numbers and has limited functionality. It is similar to SQLite,
it uses a single file to store everything, the goal is to have persistent data. The database can
store integer arrays and perform some operations on them.


### Executing program
git pull https://github.com/shkkr01/WeekTwoAssignment.git

Commands that you should run 
1. ./wkn new ---> will give you a error if the .wkn file is already present 
Error: Db file already present
2. ./wkn ----> it will start the REPL database where you can perform the operations.
3. Functionality ;
   a. # Create an empty array
        wkn> new array_name
        CREATED
   b. # Create an array with data, in the output, 8 represents the length of created array
        wkn> new array_name 1,2,3,4,7,0,-1,4
        CREATED (8)
   c. # Print an existing array - empty one
        wkn> show array_1
        [ ]
   d. # Print a non existent array
        wkn> show array_3
        Error: “array_3” does not exist
   e. # Delete an array
        wkn> del array_1
        DELETED
   f. # Merge Two arrays - the data from 2nd array gets merged to first array, without removing
        data from 2nd array.
        wkn> merge array_1 array_2
        MERGED
   g. # To exit
        wkn> exit
        Bye!
   h. # Unknown command
        wkn> sort array_1
        Error: “sort” is not a supported operation
