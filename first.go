package main
import (
//"strconv"
"fmt"
"bufio"
"io"
"regexp"
"strings"
"sync"
)
type Results map[string]string
type Count map[string]int

type  Fields struct
{
result Results
}

type Entry struct
{
count Count
}

func(input *Fields) Result(name string) (value string , err error){
value,ok := input.result[name]
if !ok{
fmt.Println("field '%v' does not found in record %+v", name, value)
	}
	return


}
func (input *Fields) SetField(name string,value string){
input.result[name] = value
}
func (input1 *Entry) SetField(name string , value int){
input1.count[name] = value
}

type Parser struct{
format String
regexp *regexp.Regexp
}

func (p *Parser) pattern() *Parser{
format1 = p.format
re := regexp.MustCompile(`\\\$([a-z_]+)(\\?(.))`)
result := re.ReplaceAllString(regexp.QuoteMeta(format1+" "),"(?P<$1>[^$3]*)$2")
return &Parser{format:format1,regexp:regexp.MustCompile(fmt.Sprintf("^%v$", strings.Trim(re, " ")))}
}


func (parser *Parser) Parselog(line string) result *Results{
regex := parser.regexp
fields := re.FindStringSubmatch(line)
if fields == nil {
		err = fmt.Errorf("access log line '%v' does not match given format '%v'", line, re)
		return
	}
for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		result.SetField(name, fields[i])
	}
	return
}
func readLine(reader *bufio.Reader) (string, error) {
	line, isPrefix, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	if !isPrefix {
		return string(line), nil
	}
	var buffer bytes.Buffer
	_, err = buffer.Write(line)
	for isPrefix && err == nil {
		line, isPrefix, err = reader.ReadLine()
		if err == nil {
			_, err = buffer.Write(line)
		}
	}
	return buffer.String(), err
}
func handleError(err error) {
	fmt.Fprintln(os.Stderr, err)
}

var format string
var logFile string

func init() {
	flag.StringVar(&format, "format", "$remote_addr [$time_local] \"$request\"", "Log format")
	flag.StringVar(&logFile, "log", "dummy", "Log file name to read. Read from STDIN if file name is '-'")
}
func main() {
	flag.Parse()
        var logReader io.Reader

	if logFile == "dummy" {
		logReader = strings.NewReader(`89.234.89.123 [08/Nov/2013:13:39:18 +0000] "GET /api/foo/bar HTTP/1.1"`)
	} else if logFile == "-" {
		logReader = os.Stdin
	} else {
		file, err := os.Open(logFile)
		if err != nil {
			panic(err)
		}
		logReader = file
		defer file.Close()
	}

       var lines = make(chan string)
       var entries = make(chan *Results, 10)
       go func(topload int){
             var sem = make(chan bool,topload)
             for i := 0 ; i < 10; i++{
             sem <- true
}
var wg sync.Waitgroup
for {
      if !<-sem{
             break
}
wg.Add(1)
go func(){
defer wg.Done()
line, ok := <-lines
if !ok {
sem <- false
return 
}
entry, err := parser.ParseString(line)
if err==nil{
entries <- entry
}
else{
handleError(err)
}
sem <- true
}()
}
wg.Wait()
close(enties)
}(cap(entries))
go func() {
		reader := bufio.NewReader(file)
		line, err := readLine(reader)
		for err == nil {
			// Read next line from the file and feed mapper routines.
			lines <- line
			line, err = readLine(reader)
		}
		close(lines)

		if err != nil && err != io.EOF {
			handleError(err)
		}
	}()

}


        
       



