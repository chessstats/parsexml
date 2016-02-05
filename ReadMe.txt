parses the Fide players list in XML format and converts it to an R data table

#######################################

installation

#######################################

under Windows 64 bit

1) dowload the repository as a zip
2) unzip
3) parsexml.exe is in the root directory of the unzipped repository ready for use

on any platform:

1) install the Go language
2) create a workspace directory for go
3) set the PATH environment variable to the full path of the workspace directory
4) install git and make sure it is in your path
5) open a console ( command prompt ) window and type:

go get github.com/chessstats/parsexml

6) the executable will be created in the bin directory of the workspace

#######################################

usage:

#######################################

1) open a console ( command prompt ) window

2) run parsexml in one of the following ways:

a) parsexml path/to/xml name.xml

this will parse path/to/xml/name.xml and create path/to/xml/players.txt

b) parsexml path/to/xml

this will parse path/to/xml/players.xml and create path/to/xml/players.txt

c) parsexml

this will parse players.xml in the working directory of the program
and create players.txt in the working directory of the program

------------------

hint: simplest way is to copy the xml file into the program's directory,
rename it players.xml and then run the program by clicking on it