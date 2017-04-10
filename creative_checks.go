package atomx

type CreativeChecks []CreativeCheck

func (cc *CreativeChecks) Add(c CreativeCheck) {
	*cc = append(*cc, c)
}
