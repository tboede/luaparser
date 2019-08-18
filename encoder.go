package luaparser

import(
    "fmt"
    "strconv"
    "errors"
    "regexp"
)

type Encoder struct {

}


func Marshal(v interface{}) (string,error) {

    enc := &Encoder{}

    return enc.Encode(v)
}

func (self *Encoder) Encode(v interface{}) (string,error) {

    s:=""

    switch v.(type) {
        case map[string]interface{}:

            val,_ := v.(map[string]interface{})

            sub,err := self.EncodeMap(val)

            if err != nil {
                return "",err
            }
            s=sub
        default:
            sub,err := self.EncodePrimitive(v)
            if err != nil {
                return "",err
            }
            s=sub
    }
    return s,nil
}

func (self *Encoder) EncodeMap(m map[string]interface{}) (string,error) {

    s:="{"

    allowed,err := regexp.Compile("^[A-Z|a-z|0-9|_]*$")

    if err != nil {
        return "",errors.New("Invalid check regex")
    }

    for k,v := range m {

        _,err := strconv.ParseInt(k,10,32)

        var key string

        if err == nil {
            key=fmt.Sprintf("[%s]",k)
        } else if ! allowed.Match([]byte(k)) {
            key=fmt.Sprintf("[\"%s\"]",k)
        } else {
            key=k
        }

        s+=key+"="

        sub,err := self.Encode(v)

        if err != nil {
            return "",err
        }

        s+=sub
        s+=","
    }

    if len(s) > 1 {
        b:=[]byte(s)
        b[len(b)-1]='}'
        s=string(b)
    } else {
        s+="}"
    }

    return s,nil
}


func (self *Encoder) EncodePrimitive(v interface{}) (string,error) {

    s:=""

    switch v.(type) {
        case int:
            val,_:=v.(int)
            s=strconv.FormatInt(int64(val),10)
        case int32:
            val,_:=v.(int32)
            s=strconv.FormatInt(int64(val),10)
        case int64:
            val,_:=v.(int64)
            s=strconv.FormatInt(val,10)
        case uint:
            val,_:=v.(uint)
            s=strconv.FormatUint(uint64(val),10)
        case uint32:
            val,_:=v.(uint32)
            s=strconv.FormatUint(uint64(val),10)
        case uint64:
            val,_:=v.(uint64)
            s=strconv.FormatUint(val,10)
        case float32:
            val,_:=v.(float32)
            s=strconv.FormatFloat(float64(val),'f',10,32)
        case float64:
            val,_:=v.(float64)
            s=strconv.FormatFloat(val,'f',10,64)
        case bool:
            val,_:=v.(bool)
            s=strconv.FormatBool(val)
        case string:
            val,_:=v.(string)
            s="\""+val+"\""
        default:
            return "",errors.New(fmt.Sprintf("Unknown type %T",v))
    }

    return s,nil
}
