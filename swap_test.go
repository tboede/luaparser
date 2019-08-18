package luaparser

import(
    "testing"
)

func TestSwap (t *testing.T) {

    m := make(map[string]interface{})
    n := make(map[string]interface{})
    u := make(map[string]interface{})

    m["1"]="aa"
    m["zh"]=int64(56)
    m["5"]=true
    m["3"]=false
    m["2"]=int64(34)
    m["tg"]="zz"
    m["n"]=n

    n["2"]=false
    n["1"]=int64(4)
    n["a"]="fff"
    n["u"]=u

    u["1"]=true
    u["2"]="aa"
    u["gg"]=int64(22)

    s,err := Marshal(m)

    if err != nil {
        t.Fatal(err)
    }

    mnew := make(map[string]interface{})

    err = Unmarshal([]byte(s),mnew)

    if err != nil {
        t.Fatal(err)
    }

    for k,_ := range m {
        switch m[k].(type) {
            case map[string]interface{}:
                //pass
            default:
                if m[k] != mnew[k] {
                    t.Errorf("%v != %v",m[k],mnew[k])
                }
        }
    }

    nnew,ok := mnew["n"].(map[string]interface{})

    if ! ok {
        t.Error("mnew[\"n\"] is not a map")
    } else {
        for k,_ := range n {
            switch n[k].(type) {
                case map[string]interface{}:
                    //pass
                default:
                    if n[k] != nnew[k] {
                        t.Errorf("%v != %v",n[k],nnew[k])
                    }
            }
        }

        unew,ok := nnew["u"].(map[string]interface{})

        if ! ok {
            t.Error("mnew[\"n\"] is not a map")
        } else {

            for k,_ := range u {
                switch u[k].(type) {
                    case map[string]interface{}:
                        //pass
                    default:
                        if u[k] != unew[k] {
                            t.Errorf("%v != %v",u[k],unew[k])
                        }
                }
            }
        }
    }

}

