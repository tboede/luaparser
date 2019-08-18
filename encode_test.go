package luaparser

import(
    "testing"
    "fmt"
)

func printmap(prefix string,m map[string]interface{}){

    for k,v := range m {

        switch v.(type) {
            case map[string]interface{}:
                mm,_ := v.(map[string]interface{})
                printmap(prefix+"."+k,mm)
            default:
                fmt.Println(prefix,k,":",v)
        }
    }
}

func TestEncodeMixedNested(t *testing.T) {

    m := make(map[string]interface{})
    mn := make(map[string]interface{})

    m["1"]="aa"
    m["zh"]=56
    m["5"]=true
    m["3"]=false
    m["2"]=34
    m["tg"]="zz"
    m["n"]=mn

    mn["2"]=false
    mn["1"]=4
    mn["a"]="fff"
    mn["b"]=true

    s,err := Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    t.Log(s)

    if s[0] != '{' {
        t.Error("String does not begin with '{' ")
    }

    if s[len(s)-1] != '}' {
        t.Error("String does not end with '}' ")
    }
}

func TestEncodeMixed(t *testing.T) {

    m := make(map[string]interface{})

    m["1"]="aa"
    m["zh"]=56
    m["5"]=true
    m["3"]=false
    m["2"]=34
    m["tg"]="zz"

    s,err := Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    t.Log(s)

    if s[0] != '{' {
        t.Error("String does not begin with '{' ")
    }

    if s[len(s)-1] != '}' {
        t.Error("String does not end with '}' ")
    }
}

func TestEncodeInvalidKey(t *testing.T) {

    m := make(map[string]interface{})
    test:=`{["zz-fg#"]="aa"}`

    m["zz-fg#"]="aa"

    s,err := Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    if s != test {
        t.Errorf("%s != %s",s,test)
    }

    m = make(map[string]interface{})
    test =`{["jj+zu+u"]=true}`

    m["jj+zu+u"]=true

    s,err = Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    if s != test {
        t.Errorf("%s != %s",s,test)
    }
}

func TestEncodeBool(t *testing.T) {

    m := true

    s,err := Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    if s != "true" {
        t.Errorf("%s != true",s)
    }
}

func TestEncodeInt(t *testing.T) {

    m := 5

    s,err := Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    if s != "5" {
        t.Errorf("%s != 5",s)
    }
}

func TestEncodeString(t *testing.T) {

    m := "aaaa"

    s,err := Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    if s != `"aaaa"` {
        t.Errorf("%s != true",s)
    }
}

func TestEncodeFloat(t *testing.T) {

    m := 64.6464646464

    s,err := Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    if s != `64.6464646464` {
        t.Errorf("%s != 64.6464646464",s)
    }
}
