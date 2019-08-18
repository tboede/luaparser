package main

import (
    "testing"
    "reflect"
)


func compareMap(m map[string]interface{},k string,v interface{},test *testing.T) {

    v1 := reflect.ValueOf(v)
    v2 := reflect.ValueOf(m[k])

    if v1.Kind() != v2.Kind() {

        test.Errorf("Type at '%s' is '%s' != '%s'",k,v2.Kind(),v1.Kind())

    } else if v1.Interface() != v2.Interface() {

        test.Errorf("Value at '%s' is '%v' != '%v'",k,v2.Interface(),v1.Interface())
    }

}

func compareValue(raw1 interface{},raw2 interface{}, test *testing.T) {

    v1 := reflect.ValueOf(raw1)
    v2 := reflect.ValueOf(raw2)

    if v1.Kind() != v2.Kind() {

        test.Errorf("Type '%s' != '%s'",v2.Kind(),v1.Kind())

    } else if v1.Interface() != v2.Interface() {

        test.Errorf("Value '%v' != '%v'",v2.Interface(),v1.Interface())
    }
}

func TestDecodeInt(t *testing.T) {

    v1 := int64(64)
    var v2 int64

    s := "64"

    err := Unmarshal([]byte(s),&v2)

    if err != nil {
        t.Fatal(err)
    }

    compareValue(v1,v2,t)
}

func TestDecodeBool(t *testing.T) {

    v1 := true
    var v2 bool

    s := "true"

    err := Unmarshal([]byte(s),&v2)

    if err != nil {
        t.Fatal(err)
    }

    compareValue(v1,v2,t)
}

func TestDecodeFloat(t *testing.T) {

//    v1 := float64(64)
    var v2 float64

    s := "64.6464646464"

    err := Unmarshal([]byte(s),&v2)

    if err != nil {
        t.Fatal(err)
    }
}

func TestDecodeString(t *testing.T) {

    v1 := "asdsa"
    var v2 string

    s := "\"asdsa\""

    err := Unmarshal([]byte(s),&v2)

    if err != nil {
        t.Fatal(err)
    }

    compareValue(v1,v2,t)
}

func TestDecodeList(t *testing.T) {

    m := make(map[string]interface{})
    s := `{1,2,3,4,true,false,"aa","bb"}`

    err := Unmarshal([]byte(s),m)

    if err != nil {
        t.Fatal(err)
    }

    compareMap(m,"1",int64(1),t)
    compareMap(m,"2",int64(2),t)
    compareMap(m,"3",int64(3),t)
    compareMap(m,"4",int64(4),t)
    compareMap(m,"5",true,t)
    compareMap(m,"6",false,t)
    compareMap(m,"7","aa",t)
    compareMap(m,"8","bb",t)
}

func TestDecodeMap(t *testing.T) {

    m := make(map[string]interface{})
    s := `{a=1,b=2,c=3,d=4,aa=true,bb=false,["dd-dd-dd"]="aa",cc="aa",dd="bb"}`

    err := Unmarshal([]byte(s),m)

    if err != nil {
        t.Fatal(err)
    }

    compareMap(m,"a",int64(1),t)
    compareMap(m,"b",int64(2),t)
    compareMap(m,"c",int64(3),t)
    compareMap(m,"d",int64(4),t)
    compareMap(m,"aa",true,t)
    compareMap(m,"bb",false,t)
    compareMap(m,"cc","aa",t)
    compareMap(m,"dd","bb",t)
    compareMap(m,"dd-dd-dd","aa",t)

}

func TestDecodeMixed(t *testing.T) {

    m := make(map[string]interface{})
    s := `{1,true,"cc",b=2,[6]=3,[5]=4,bb=false,["ll-aa"]=22,dd="aa"}`

    err := Unmarshal([]byte(s),m)

    if err != nil {
        t.Fatal(err)
    }

    compareMap(m,"1",int64(1),t)
    compareMap(m,"2",true,t)
    compareMap(m,"3","cc",t)
    compareMap(m,"b",int64(2),t)
    compareMap(m,"5",int64(4),t)
    compareMap(m,"6",int64(3),t)
    compareMap(m,"bb",false,t)
    compareMap(m,"dd","aa",t)
    compareMap(m,"ll-aa",int64(22),t)

}

func TestDecodeMixedNested(t *testing.T) {

    m := make(map[string]interface{})
    s := `{1,true,"cc",b=2,[6]=3,n={"a",false,true,gg=2,zz="aa",hh=false,u={"zz",ui=false,[78]=true,lol=22}},[5]=4,bb=false,dd="aa"}`

    err := Unmarshal([]byte(s),m)

    if err != nil {
        t.Fatal(err)
    }

    compareMap(m,"1",int64(1),t)
    compareMap(m,"2",true,t)
    compareMap(m,"3","cc",t)
    compareMap(m,"b",int64(2),t)
    compareMap(m,"5",int64(4),t)
    compareMap(m,"6",int64(3),t)
    compareMap(m,"bb",false,t)
    compareMap(m,"dd","aa",t)

    n,ok := m["n"].(map[string]interface{})

    if ! ok {
        t.Error("n is not a map")
    } else {
        compareMap(n,"1","a",t)
        compareMap(n,"2",false,t)
        compareMap(n,"3",true,t)
        compareMap(n,"gg",int64(2),t)
        compareMap(n,"zz","aa",t)
        compareMap(n,"hh",false,t)

        u,ok := n["u"].(map[string]interface{})

        if ! ok {
            t.Error("u is not a map")
        } else {
            compareMap(u,"1","zz",t)
            compareMap(u,"ui",false,t)
            compareMap(u,"78",true,t)
            compareMap(u,"lol",int64(22),t)
        }
    }

}
