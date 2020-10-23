# Week 7

Goal: To complete the 6th chapter which is about `methods`.

Result: 

I was slightly confused with the example program intset. 
There we were wrote a type called `IntSet` which is a public 
struct (or in OOP terms a class GO doesnt have objects only custom types)

in the implementation there is the following 
```go
type IntSet struct {
	words []uint64
}
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

```

which defines that the IntSet has a words subparameter 
and it has the following methods `Has` `Add` `UnionWith` and `String`
All those functions are capitalized to let go know that we can access 
these functions outside of the package. 

What I was stumped with was that each function is creating `word, bit := x/64, uint(x%64)`
 I know that it is moding the x by 64 but to do what? and the Has function is super confusing to me because of all the boolean arithmatic.