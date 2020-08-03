# gender-pronouns
Simple Go library to convert gender pronouns and nouns to their different forms. Based on [this matrix](https://uwm.edu/lgbtrc/support/gender-pronouns/#:~:text=She%2Fher%2Fhers%20and%20he,gender%2Dneutral%20pronouns%20in%20use.)

> This is not meant to be an exaustive list. I'm not sure what else might be missing but feel free to open up a PR

### Install
```
go get github.com/motivq/gender-pronouns
```

### Usage

```
import  github.com/motivq/gender-pronouns/pkg

func main(){
    subject := pronouns.They
    fmt.Println(subject.String())
    fmt.Println(subject.Object())
    fmt.Println(subject.Possesive())
    fmt.Println(subject.PossesivePronoun())
    fmt.Println(subject.Reflexive())
}
```

### TO DO
Add converting from the something other than the subject