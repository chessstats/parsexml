package main

import(
	"fmt"
	"io/ioutil"
	"os"
)

const WAIT_TAG=0
const READ_TAG=1
const READ_VALUE=2
const WAIT_END=3

var outf *os.File

func print(what string) {
	if _,err:=outf.WriteString(what); err!=nil {
		panic(err)
	}
}

var numplayers int
func parse_xml(buff []byte, header bool) {
	numplayers=0
	var tagpuff [100]byte
	var valuepuff [500]byte
	tagi:=0
	valuei:=0
	status:=WAIT_TAG
	first:=true
	for i:=0; i<len(buff); i++ {
		if status==WAIT_TAG {
			if buff[i]=='<' {
				status=READ_TAG
			}
		} else if status==READ_TAG {
			if buff[i]=='>' {
				status=READ_VALUE
				tag:=string(tagpuff[0:tagi])
				if tag=="player" {
					status=WAIT_TAG
					numplayers++
					if (numplayers%50000)==0 {
						fmt.Printf("processed %d\n",numplayers)
					}
					tagi=0
					valuei=0
					first=true
				} else if tag=="/player" {
					if header {
						print("\r\n")
						return
					}
					status=WAIT_TAG
					tagi=0
					valuei=0
					print("\r\n")
				} else if tag=="playerslist" {
					status=WAIT_TAG
					tagi=0
					valuei=0
				} else {
					if header {
						if first {
							first=false
							print(fmt.Sprintf("\"%s\"",tag))
						} else {
							print(fmt.Sprintf(" \"%s\"",tag))
						}
					}
				}
			} else {
				tagpuff[tagi]=buff[i]
				tagi++
			}
		} else if status==READ_VALUE {
			if buff[i]=='<' {
				status=WAIT_END
				value:=string(valuepuff[0:valuei])
				if !header {
					if first {
						first=false
						print(fmt.Sprintf("\"%d\" \"%s\"",numplayers,value))
					} else {
						print(fmt.Sprintf(" \"%s\"",value))
					}
				}
				tagi=0
				valuei=0
			} else {
				valuepuff[valuei]=buff[i]
				valuei++
			}
		} else if status==WAIT_END {
			if buff[i]=='>' {
				status=WAIT_TAG
			}
		}
	}
}

func main() {
	fmt.Printf("\nparse XML @ chessstats\n\n")
	fmt.Printf("usage:\n\nparsexml path/to/xml/ name.xml\n")
	fmt.Printf("  -> parses path/to/xml/name.xml\n")
	fmt.Printf("parsexml path/to/xml/\n")
	fmt.Printf("  -> parses path/to/xml/players.xml\n")
	fmt.Printf("parsexml\n")
	fmt.Printf("  -> parses ./players.xml\n\n")
	path:=""
	name:="players.xml"
	if len(os.Args)>1 {
		path=os.Args[1]
	}
	if len(os.Args)>2 {
		name=os.Args[2]
	}
	fullpath:=path+name
	fmt.Printf("reading file %s\n\n",fullpath)
	buff,err:=ioutil.ReadFile(fullpath)
	if err!=nil {
		fmt.Printf("io error, parse failed")
		return
	} else {
		f,err:=os.OpenFile(path+"players.txt",os.O_CREATE|os.O_WRONLY,0666)
		if err!=nil {
		    panic(err)
		}
		outf=f

		fmt.Printf("file read, size %d\n\n",len(buff))
		parse_xml(buff,true)
		parse_xml(buff,false)
		fmt.Printf("\ndone, num players %d\n",numplayers)
	}
}