package lua

import (
    "fmt"
    "strconv"
    "errors"
    "reflect"
log "github.com/sirupsen/logrus"
)

type Decoder struct {
    cursor int
    index int
    data []byte
}

func Unmarshal(data []byte, v interface{}) error {

    par := &Decoder{cursor:0,index:1,data:data}

    return par.Decode(v)
}

func (self *Decoder) Current() (byte,error) {

    if self.cursor >= len(self.data) {
        return 0,errors.New(fmt.Sprintf("Index %d out of range",self.cursor))
    }

    return self.data[self.cursor],nil
}

func (self *Decoder) Decode(v interface{}) error {

    c,err := self.Current()

    if err != nil {
        return err
    }

    if c == '{' {

        return self.DecodeTable(v)
    }

    value := self.DecodePrimitive(string(self.data))


    if reflect.TypeOf(v).Kind() !=  reflect.Ptr {
        return errors.New("destination is not a pointer")
    }

    ptr:=reflect.ValueOf(v)

    if ptr.Elem().Kind() != reflect.ValueOf(value).Kind() {
        return errors.New("destination type does not match")
    }

    ptr.Elem().Set(reflect.ValueOf(value))

    return nil
}

func (self *Decoder) DecodeTable(v interface{}) error {

    log.Debug("Parsing Table")
    var c byte = 0

    for c != '}' {

        log.Debug("Looking for entry")

        key,valuestr,err := self.FindEntry()

        if err != nil {
            return err
        }

        store,ok := v.(map[string]interface{})

        if ! ok {
            return errors.New("Is not an interface map")
        }

        if len(valuestr) == 0 {
            //skip - empty entry
        } else {
            var value interface{}

            if valuestr[0] == '{' {
                substore := make(map[string]interface{})

                err := Unmarshal([]byte(valuestr),substore)

                if err != nil {
                    return err
                }
                value=substore

            } else if valuestr[0] == '"' {
                value = valuestr[1:len(valuestr)-1]
            } else {
                value = self.DecodePrimitive(valuestr)
            }

            if key == "" {
                store[strconv.FormatInt(int64(self.index),10)]=value
                self.index++
            } else if key[0] == '[' {
                if (key[1] == '"') {
                    store[key[2:len(key)-2]]=value
                } else {
                    store[key[1:len(key)-1]]=value
                }
            } else {
                store[key]=value
            }
        }

        c,err = self.Current()

        if err != nil {
            return err
        }

    }

    return nil
}

func (self *Decoder) FindEntry() (string,string,error) {

    var c byte = 0
    var err error

    seppos:=0
    instring:=false
    entry:=""
    start:=self.cursor+1

    for c != ',' && c != '}' {

        self.cursor++

        c,err = self.Current()

        // cursor is out of range
        if err != nil {
            return "","",err
        }

        if c == '=' && seppos == 0 {
        // current byte is a seperator

            seppos=self.cursor-start
        }

        // Cursor is inside a string
        if instring {
            if c == '"' {
                instring=false
                entry+="\""
            } else if c == '\\' {
                self.cursor++

                c,err := self.Current()

                if err != nil {
                    return "","",err
                }
                if c == 'n' {
                    entry+="\n"
                } else {
                    entry+=string(c)
                }
            } else if c == ',' {
                entry+=string(c)
                c=0
            } else {
                entry+=string(c)
            }
        } else if c == '{' {
        // cursor is at the beginning of a table

            c,_ := self.Current()

            s:="{"

            for c != '}' {

                key,value,err := self.FindEntry()

                if err != nil {
                    return "","",err
                }

                if key != "" {
                    s+=key
                    s+="="
                    s+=value

                } else {
                    s+=value
                }
                s+=","

                c,err = self.Current()

                if err != nil {
                    return "","",err
                }
            }

            if len(s) > 1 {
                b:=[]byte(s)
                b[len(b)-1]='}'
                s=string(b)

            } else {
                s+="}"
            }

            entry+=s
            self.cursor++

        } else if c != ',' && c != '}' {
        // cursor reached entry's end

            if c == '"' {
                instring=true
                entry+="\""
            } else {
                entry+=string(c)
            }
        }

        // get byte at cursor position
        c,err = self.Current()

        if err != nil {
            return "","",err
        }
    }

    if seppos == 0 {
        return "",entry,nil

    } else {

        key:= entry[:seppos]
        val:= entry[seppos+1:]

        return key,val,nil
    }
}

func (self *Decoder) DecodePrimitive(s string) interface{} {

    if len(s) > 2{
        if s[0] == '"' && s[len(s)-1] == '"' {
            return s[1:len(s)-1]
        }
    }

    i,err := strconv.ParseInt(s,10,32)

    if err == nil {
        return i
    }

    b,err := strconv.ParseBool(s)

    if err == nil {
        return b
    }

    r,err := strconv.ParseFloat(s,32)

    if err == nil {
        return r
    }

    return s
}

