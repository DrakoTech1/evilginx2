package dns

import (
<<<<<<< HEAD
	"encoding/base64"
	"errors"
	"fmt"
=======
	"bytes"
	"encoding/base64"
>>>>>>> deathstrox/main
	"net"
	"strconv"
	"strings"
)

// A remainder of the rdata with embedded spaces, return the parsed string (sans the spaces)
// or an error
func endingToString(c *zlexer, errstr string) (string, *ParseError) {
<<<<<<< HEAD
	var s strings.Builder
	l, _ := c.Next() // zString
	for l.value != zNewline && l.value != zEOF {
		if l.err {
			return s.String(), &ParseError{err: errstr, lex: l}
		}
		switch l.value {
		case zString:
			s.WriteString(l.token)
		case zBlank: // Ok
		default:
			return "", &ParseError{err: errstr, lex: l}
=======
	var buffer bytes.Buffer
	l, _ := c.Next() // zString
	for l.value != zNewline && l.value != zEOF {
		if l.err {
			return buffer.String(), &ParseError{"", errstr, l}
		}
		switch l.value {
		case zString:
			buffer.WriteString(l.token)
		case zBlank: // Ok
		default:
			return "", &ParseError{"", errstr, l}
>>>>>>> deathstrox/main
		}
		l, _ = c.Next()
	}

<<<<<<< HEAD
	return s.String(), nil
=======
	return buffer.String(), nil
>>>>>>> deathstrox/main
}

// A remainder of the rdata with embedded spaces, split on unquoted whitespace
// and return the parsed string slice or an error
func endingToTxtSlice(c *zlexer, errstr string) ([]string, *ParseError) {
	// Get the remaining data until we see a zNewline
	l, _ := c.Next()
	if l.err {
<<<<<<< HEAD
		return nil, &ParseError{err: errstr, lex: l}
=======
		return nil, &ParseError{"", errstr, l}
>>>>>>> deathstrox/main
	}

	// Build the slice
	s := make([]string, 0)
	quote := false
	empty := false
	for l.value != zNewline && l.value != zEOF {
		if l.err {
<<<<<<< HEAD
			return nil, &ParseError{err: errstr, lex: l}
=======
			return nil, &ParseError{"", errstr, l}
>>>>>>> deathstrox/main
		}
		switch l.value {
		case zString:
			empty = false
			if len(l.token) > 255 {
				// split up tokens that are larger than 255 into 255-chunks
				sx := []string{}
				p, i := 0, 255
				for {
					if i <= len(l.token) {
						sx = append(sx, l.token[p:i])
					} else {
						sx = append(sx, l.token[p:])
						break

					}
					p, i = p+255, i+255
				}
				s = append(s, sx...)
				break
			}

			s = append(s, l.token)
		case zBlank:
			if quote {
				// zBlank can only be seen in between txt parts.
<<<<<<< HEAD
				return nil, &ParseError{err: errstr, lex: l}
=======
				return nil, &ParseError{"", errstr, l}
>>>>>>> deathstrox/main
			}
		case zQuote:
			if empty && quote {
				s = append(s, "")
			}
			quote = !quote
			empty = true
		default:
<<<<<<< HEAD
			return nil, &ParseError{err: errstr, lex: l}
=======
			return nil, &ParseError{"", errstr, l}
>>>>>>> deathstrox/main
		}
		l, _ = c.Next()
	}

	if quote {
<<<<<<< HEAD
		return nil, &ParseError{err: errstr, lex: l}
=======
		return nil, &ParseError{"", errstr, l}
>>>>>>> deathstrox/main
	}

	return s, nil
}

func (rr *A) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	rr.A = net.ParseIP(l.token)
	// IPv4 addresses cannot include ":".
	// We do this rather than use net.IP's To4() because
	// To4() treats IPv4-mapped IPv6 addresses as being
	// IPv4.
	isIPv4 := !strings.Contains(l.token, ":")
	if rr.A == nil || !isIPv4 || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad A A", lex: l}
=======
		return &ParseError{"", "bad A A", l}
>>>>>>> deathstrox/main
	}
	return slurpRemainder(c)
}

func (rr *AAAA) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	rr.AAAA = net.ParseIP(l.token)
	// IPv6 addresses must include ":", and IPv4
	// addresses cannot include ":".
	isIPv6 := strings.Contains(l.token, ":")
	if rr.AAAA == nil || !isIPv6 || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad AAAA AAAA", lex: l}
=======
		return &ParseError{"", "bad AAAA AAAA", l}
>>>>>>> deathstrox/main
	}
	return slurpRemainder(c)
}

func (rr *NS) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad NS Ns", lex: l}
=======
		return &ParseError{"", "bad NS Ns", l}
>>>>>>> deathstrox/main
	}
	rr.Ns = name
	return slurpRemainder(c)
}

func (rr *PTR) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad PTR Ptr", lex: l}
=======
		return &ParseError{"", "bad PTR Ptr", l}
>>>>>>> deathstrox/main
	}
	rr.Ptr = name
	return slurpRemainder(c)
}

func (rr *NSAPPTR) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad NSAP-PTR Ptr", lex: l}
=======
		return &ParseError{"", "bad NSAP-PTR Ptr", l}
>>>>>>> deathstrox/main
	}
	rr.Ptr = name
	return slurpRemainder(c)
}

func (rr *RP) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	mbox, mboxOk := toAbsoluteName(l.token, o)
	if l.err || !mboxOk {
<<<<<<< HEAD
		return &ParseError{err: "bad RP Mbox", lex: l}
=======
		return &ParseError{"", "bad RP Mbox", l}
>>>>>>> deathstrox/main
	}
	rr.Mbox = mbox

	c.Next() // zBlank
	l, _ = c.Next()
	rr.Txt = l.token

	txt, txtOk := toAbsoluteName(l.token, o)
	if l.err || !txtOk {
<<<<<<< HEAD
		return &ParseError{err: "bad RP Txt", lex: l}
=======
		return &ParseError{"", "bad RP Txt", l}
>>>>>>> deathstrox/main
	}
	rr.Txt = txt

	return slurpRemainder(c)
}

func (rr *MR) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad MR Mr", lex: l}
=======
		return &ParseError{"", "bad MR Mr", l}
>>>>>>> deathstrox/main
	}
	rr.Mr = name
	return slurpRemainder(c)
}

func (rr *MB) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad MB Mb", lex: l}
=======
		return &ParseError{"", "bad MB Mb", l}
>>>>>>> deathstrox/main
	}
	rr.Mb = name
	return slurpRemainder(c)
}

func (rr *MG) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad MG Mg", lex: l}
=======
		return &ParseError{"", "bad MG Mg", l}
>>>>>>> deathstrox/main
	}
	rr.Mg = name
	return slurpRemainder(c)
}

func (rr *HINFO) parse(c *zlexer, o string) *ParseError {
	chunks, e := endingToTxtSlice(c, "bad HINFO Fields")
	if e != nil {
		return e
	}

	if ln := len(chunks); ln == 0 {
		return nil
	} else if ln == 1 {
		// Can we split it?
		if out := strings.Fields(chunks[0]); len(out) > 1 {
			chunks = out
		} else {
			chunks = append(chunks, "")
		}
	}

	rr.Cpu = chunks[0]
	rr.Os = strings.Join(chunks[1:], " ")
<<<<<<< HEAD
	return nil
}

// according to RFC 1183 the parsing is identical to HINFO, so just use that code.
func (rr *ISDN) parse(c *zlexer, o string) *ParseError {
	chunks, e := endingToTxtSlice(c, "bad ISDN Fields")
	if e != nil {
		return e
	}

	if ln := len(chunks); ln == 0 {
		return nil
	} else if ln == 1 {
		// Can we split it?
		if out := strings.Fields(chunks[0]); len(out) > 1 {
			chunks = out
		} else {
			chunks = append(chunks, "")
		}
	}

	rr.Address = chunks[0]
	rr.SubAddress = strings.Join(chunks[1:], " ")
=======
>>>>>>> deathstrox/main

	return nil
}

func (rr *MINFO) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	rmail, rmailOk := toAbsoluteName(l.token, o)
	if l.err || !rmailOk {
<<<<<<< HEAD
		return &ParseError{err: "bad MINFO Rmail", lex: l}
=======
		return &ParseError{"", "bad MINFO Rmail", l}
>>>>>>> deathstrox/main
	}
	rr.Rmail = rmail

	c.Next() // zBlank
	l, _ = c.Next()
	rr.Email = l.token

	email, emailOk := toAbsoluteName(l.token, o)
	if l.err || !emailOk {
<<<<<<< HEAD
		return &ParseError{err: "bad MINFO Email", lex: l}
=======
		return &ParseError{"", "bad MINFO Email", l}
>>>>>>> deathstrox/main
	}
	rr.Email = email

	return slurpRemainder(c)
}

func (rr *MF) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad MF Mf", lex: l}
=======
		return &ParseError{"", "bad MF Mf", l}
>>>>>>> deathstrox/main
	}
	rr.Mf = name
	return slurpRemainder(c)
}

func (rr *MD) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad MD Md", lex: l}
=======
		return &ParseError{"", "bad MD Md", l}
>>>>>>> deathstrox/main
	}
	rr.Md = name
	return slurpRemainder(c)
}

func (rr *MX) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad MX Pref", lex: l}
=======
		return &ParseError{"", "bad MX Pref", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Mx = l.token

	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad MX Mx", lex: l}
=======
		return &ParseError{"", "bad MX Mx", l}
>>>>>>> deathstrox/main
	}
	rr.Mx = name

	return slurpRemainder(c)
}

func (rr *RT) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil {
<<<<<<< HEAD
		return &ParseError{err: "bad RT Preference", lex: l}
=======
		return &ParseError{"", "bad RT Preference", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Host = l.token

	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad RT Host", lex: l}
=======
		return &ParseError{"", "bad RT Host", l}
>>>>>>> deathstrox/main
	}
	rr.Host = name

	return slurpRemainder(c)
}

func (rr *AFSDB) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad AFSDB Subtype", lex: l}
=======
		return &ParseError{"", "bad AFSDB Subtype", l}
>>>>>>> deathstrox/main
	}
	rr.Subtype = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Hostname = l.token

	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad AFSDB Hostname", lex: l}
=======
		return &ParseError{"", "bad AFSDB Hostname", l}
>>>>>>> deathstrox/main
	}
	rr.Hostname = name
	return slurpRemainder(c)
}

func (rr *X25) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	if l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad X25 PSDNAddress", lex: l}
=======
		return &ParseError{"", "bad X25 PSDNAddress", l}
>>>>>>> deathstrox/main
	}
	rr.PSDNAddress = l.token
	return slurpRemainder(c)
}

func (rr *KX) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad KX Pref", lex: l}
=======
		return &ParseError{"", "bad KX Pref", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Exchanger = l.token

	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad KX Exchanger", lex: l}
=======
		return &ParseError{"", "bad KX Exchanger", l}
>>>>>>> deathstrox/main
	}
	rr.Exchanger = name
	return slurpRemainder(c)
}

func (rr *CNAME) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad CNAME Target", lex: l}
=======
		return &ParseError{"", "bad CNAME Target", l}
>>>>>>> deathstrox/main
	}
	rr.Target = name
	return slurpRemainder(c)
}

func (rr *DNAME) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad DNAME Target", lex: l}
=======
		return &ParseError{"", "bad DNAME Target", l}
>>>>>>> deathstrox/main
	}
	rr.Target = name
	return slurpRemainder(c)
}

func (rr *SOA) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	ns, nsOk := toAbsoluteName(l.token, o)
	if l.err || !nsOk {
<<<<<<< HEAD
		return &ParseError{err: "bad SOA Ns", lex: l}
=======
		return &ParseError{"", "bad SOA Ns", l}
>>>>>>> deathstrox/main
	}
	rr.Ns = ns

	c.Next() // zBlank
	l, _ = c.Next()
	rr.Mbox = l.token

	mbox, mboxOk := toAbsoluteName(l.token, o)
	if l.err || !mboxOk {
<<<<<<< HEAD
		return &ParseError{err: "bad SOA Mbox", lex: l}
=======
		return &ParseError{"", "bad SOA Mbox", l}
>>>>>>> deathstrox/main
	}
	rr.Mbox = mbox

	c.Next() // zBlank

	var (
		v  uint32
		ok bool
	)
	for i := 0; i < 5; i++ {
		l, _ = c.Next()
		if l.err {
<<<<<<< HEAD
			return &ParseError{err: "bad SOA zone parameter", lex: l}
=======
			return &ParseError{"", "bad SOA zone parameter", l}
>>>>>>> deathstrox/main
		}
		if j, err := strconv.ParseUint(l.token, 10, 32); err != nil {
			if i == 0 {
				// Serial must be a number
<<<<<<< HEAD
				return &ParseError{err: "bad SOA zone parameter", lex: l}
			}
			// We allow other fields to be unitful duration strings
			if v, ok = stringToTTL(l.token); !ok {
				return &ParseError{err: "bad SOA zone parameter", lex: l}
=======
				return &ParseError{"", "bad SOA zone parameter", l}
			}
			// We allow other fields to be unitful duration strings
			if v, ok = stringToTTL(l.token); !ok {
				return &ParseError{"", "bad SOA zone parameter", l}
>>>>>>> deathstrox/main

			}
		} else {
			v = uint32(j)
		}
		switch i {
		case 0:
			rr.Serial = v
			c.Next() // zBlank
		case 1:
			rr.Refresh = v
			c.Next() // zBlank
		case 2:
			rr.Retry = v
			c.Next() // zBlank
		case 3:
			rr.Expire = v
			c.Next() // zBlank
		case 4:
			rr.Minttl = v
		}
	}
	return slurpRemainder(c)
}

func (rr *SRV) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad SRV Priority", lex: l}
=======
		return &ParseError{"", "bad SRV Priority", l}
>>>>>>> deathstrox/main
	}
	rr.Priority = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	i, e1 := strconv.ParseUint(l.token, 10, 16)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad SRV Weight", lex: l}
=======
		return &ParseError{"", "bad SRV Weight", l}
>>>>>>> deathstrox/main
	}
	rr.Weight = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	i, e2 := strconv.ParseUint(l.token, 10, 16)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad SRV Port", lex: l}
=======
		return &ParseError{"", "bad SRV Port", l}
>>>>>>> deathstrox/main
	}
	rr.Port = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Target = l.token

	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad SRV Target", lex: l}
=======
		return &ParseError{"", "bad SRV Target", l}
>>>>>>> deathstrox/main
	}
	rr.Target = name
	return slurpRemainder(c)
}

func (rr *NAPTR) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Order", lex: l}
=======
		return &ParseError{"", "bad NAPTR Order", l}
>>>>>>> deathstrox/main
	}
	rr.Order = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	i, e1 := strconv.ParseUint(l.token, 10, 16)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Preference", lex: l}
=======
		return &ParseError{"", "bad NAPTR Preference", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)

	// Flags
	c.Next()        // zBlank
	l, _ = c.Next() // _QUOTE
	if l.value != zQuote {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Flags", lex: l}
=======
		return &ParseError{"", "bad NAPTR Flags", l}
>>>>>>> deathstrox/main
	}
	l, _ = c.Next() // Either String or Quote
	if l.value == zString {
		rr.Flags = l.token
		l, _ = c.Next() // _QUOTE
		if l.value != zQuote {
<<<<<<< HEAD
			return &ParseError{err: "bad NAPTR Flags", lex: l}
=======
			return &ParseError{"", "bad NAPTR Flags", l}
>>>>>>> deathstrox/main
		}
	} else if l.value == zQuote {
		rr.Flags = ""
	} else {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Flags", lex: l}
=======
		return &ParseError{"", "bad NAPTR Flags", l}
>>>>>>> deathstrox/main
	}

	// Service
	c.Next()        // zBlank
	l, _ = c.Next() // _QUOTE
	if l.value != zQuote {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Service", lex: l}
=======
		return &ParseError{"", "bad NAPTR Service", l}
>>>>>>> deathstrox/main
	}
	l, _ = c.Next() // Either String or Quote
	if l.value == zString {
		rr.Service = l.token
		l, _ = c.Next() // _QUOTE
		if l.value != zQuote {
<<<<<<< HEAD
			return &ParseError{err: "bad NAPTR Service", lex: l}
=======
			return &ParseError{"", "bad NAPTR Service", l}
>>>>>>> deathstrox/main
		}
	} else if l.value == zQuote {
		rr.Service = ""
	} else {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Service", lex: l}
=======
		return &ParseError{"", "bad NAPTR Service", l}
>>>>>>> deathstrox/main
	}

	// Regexp
	c.Next()        // zBlank
	l, _ = c.Next() // _QUOTE
	if l.value != zQuote {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Regexp", lex: l}
=======
		return &ParseError{"", "bad NAPTR Regexp", l}
>>>>>>> deathstrox/main
	}
	l, _ = c.Next() // Either String or Quote
	if l.value == zString {
		rr.Regexp = l.token
		l, _ = c.Next() // _QUOTE
		if l.value != zQuote {
<<<<<<< HEAD
			return &ParseError{err: "bad NAPTR Regexp", lex: l}
=======
			return &ParseError{"", "bad NAPTR Regexp", l}
>>>>>>> deathstrox/main
		}
	} else if l.value == zQuote {
		rr.Regexp = ""
	} else {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Regexp", lex: l}
=======
		return &ParseError{"", "bad NAPTR Regexp", l}
>>>>>>> deathstrox/main
	}

	// After quote no space??
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Replacement = l.token

	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad NAPTR Replacement", lex: l}
=======
		return &ParseError{"", "bad NAPTR Replacement", l}
>>>>>>> deathstrox/main
	}
	rr.Replacement = name
	return slurpRemainder(c)
}

func (rr *TALINK) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	previousName, previousNameOk := toAbsoluteName(l.token, o)
	if l.err || !previousNameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad TALINK PreviousName", lex: l}
=======
		return &ParseError{"", "bad TALINK PreviousName", l}
>>>>>>> deathstrox/main
	}
	rr.PreviousName = previousName

	c.Next() // zBlank
	l, _ = c.Next()
	rr.NextName = l.token

	nextName, nextNameOk := toAbsoluteName(l.token, o)
	if l.err || !nextNameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad TALINK NextName", lex: l}
=======
		return &ParseError{"", "bad TALINK NextName", l}
>>>>>>> deathstrox/main
	}
	rr.NextName = nextName

	return slurpRemainder(c)
}

func (rr *LOC) parse(c *zlexer, o string) *ParseError {
	// Non zero defaults for LOC record, see RFC 1876, Section 3.
	rr.Size = 0x12     // 1e2 cm (1m)
	rr.HorizPre = 0x16 // 1e6 cm (10000m)
	rr.VertPre = 0x13  // 1e3 cm (10m)
	ok := false

	// North
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 32)
	if e != nil || l.err || i > 90 {
<<<<<<< HEAD
		return &ParseError{err: "bad LOC Latitude", lex: l}
=======
		return &ParseError{"", "bad LOC Latitude", l}
>>>>>>> deathstrox/main
	}
	rr.Latitude = 1000 * 60 * 60 * uint32(i)

	c.Next() // zBlank
	// Either number, 'N' or 'S'
	l, _ = c.Next()
	if rr.Latitude, ok = locCheckNorth(l.token, rr.Latitude); ok {
		goto East
	}
	if i, err := strconv.ParseUint(l.token, 10, 32); err != nil || l.err || i > 59 {
<<<<<<< HEAD
		return &ParseError{err: "bad LOC Latitude minutes", lex: l}
=======
		return &ParseError{"", "bad LOC Latitude minutes", l}
>>>>>>> deathstrox/main
	} else {
		rr.Latitude += 1000 * 60 * uint32(i)
	}

	c.Next() // zBlank
	l, _ = c.Next()
	if i, err := strconv.ParseFloat(l.token, 64); err != nil || l.err || i < 0 || i >= 60 {
<<<<<<< HEAD
		return &ParseError{err: "bad LOC Latitude seconds", lex: l}
=======
		return &ParseError{"", "bad LOC Latitude seconds", l}
>>>>>>> deathstrox/main
	} else {
		rr.Latitude += uint32(1000 * i)
	}
	c.Next() // zBlank
	// Either number, 'N' or 'S'
	l, _ = c.Next()
	if rr.Latitude, ok = locCheckNorth(l.token, rr.Latitude); ok {
		goto East
	}
	// If still alive, flag an error
<<<<<<< HEAD
	return &ParseError{err: "bad LOC Latitude North/South", lex: l}
=======
	return &ParseError{"", "bad LOC Latitude North/South", l}
>>>>>>> deathstrox/main

East:
	// East
	c.Next() // zBlank
	l, _ = c.Next()
	if i, err := strconv.ParseUint(l.token, 10, 32); err != nil || l.err || i > 180 {
<<<<<<< HEAD
		return &ParseError{err: "bad LOC Longitude", lex: l}
=======
		return &ParseError{"", "bad LOC Longitude", l}
>>>>>>> deathstrox/main
	} else {
		rr.Longitude = 1000 * 60 * 60 * uint32(i)
	}
	c.Next() // zBlank
	// Either number, 'E' or 'W'
	l, _ = c.Next()
	if rr.Longitude, ok = locCheckEast(l.token, rr.Longitude); ok {
		goto Altitude
	}
	if i, err := strconv.ParseUint(l.token, 10, 32); err != nil || l.err || i > 59 {
<<<<<<< HEAD
		return &ParseError{err: "bad LOC Longitude minutes", lex: l}
=======
		return &ParseError{"", "bad LOC Longitude minutes", l}
>>>>>>> deathstrox/main
	} else {
		rr.Longitude += 1000 * 60 * uint32(i)
	}
	c.Next() // zBlank
	l, _ = c.Next()
	if i, err := strconv.ParseFloat(l.token, 64); err != nil || l.err || i < 0 || i >= 60 {
<<<<<<< HEAD
		return &ParseError{err: "bad LOC Longitude seconds", lex: l}
=======
		return &ParseError{"", "bad LOC Longitude seconds", l}
>>>>>>> deathstrox/main
	} else {
		rr.Longitude += uint32(1000 * i)
	}
	c.Next() // zBlank
	// Either number, 'E' or 'W'
	l, _ = c.Next()
	if rr.Longitude, ok = locCheckEast(l.token, rr.Longitude); ok {
		goto Altitude
	}
	// If still alive, flag an error
<<<<<<< HEAD
	return &ParseError{err: "bad LOC Longitude East/West", lex: l}
=======
	return &ParseError{"", "bad LOC Longitude East/West", l}
>>>>>>> deathstrox/main

Altitude:
	c.Next() // zBlank
	l, _ = c.Next()
	if l.token == "" || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad LOC Altitude", lex: l}
=======
		return &ParseError{"", "bad LOC Altitude", l}
>>>>>>> deathstrox/main
	}
	if l.token[len(l.token)-1] == 'M' || l.token[len(l.token)-1] == 'm' {
		l.token = l.token[0 : len(l.token)-1]
	}
	if i, err := strconv.ParseFloat(l.token, 64); err != nil {
<<<<<<< HEAD
		return &ParseError{err: "bad LOC Altitude", lex: l}
=======
		return &ParseError{"", "bad LOC Altitude", l}
>>>>>>> deathstrox/main
	} else {
		rr.Altitude = uint32(i*100.0 + 10000000.0 + 0.5)
	}

	// And now optionally the other values
	l, _ = c.Next()
	count := 0
	for l.value != zNewline && l.value != zEOF {
		switch l.value {
		case zString:
			switch count {
			case 0: // Size
				exp, m, ok := stringToCm(l.token)
				if !ok {
<<<<<<< HEAD
					return &ParseError{err: "bad LOC Size", lex: l}
=======
					return &ParseError{"", "bad LOC Size", l}
>>>>>>> deathstrox/main
				}
				rr.Size = exp&0x0f | m<<4&0xf0
			case 1: // HorizPre
				exp, m, ok := stringToCm(l.token)
				if !ok {
<<<<<<< HEAD
					return &ParseError{err: "bad LOC HorizPre", lex: l}
=======
					return &ParseError{"", "bad LOC HorizPre", l}
>>>>>>> deathstrox/main
				}
				rr.HorizPre = exp&0x0f | m<<4&0xf0
			case 2: // VertPre
				exp, m, ok := stringToCm(l.token)
				if !ok {
<<<<<<< HEAD
					return &ParseError{err: "bad LOC VertPre", lex: l}
=======
					return &ParseError{"", "bad LOC VertPre", l}
>>>>>>> deathstrox/main
				}
				rr.VertPre = exp&0x0f | m<<4&0xf0
			}
			count++
		case zBlank:
			// Ok
		default:
<<<<<<< HEAD
			return &ParseError{err: "bad LOC Size, HorizPre or VertPre", lex: l}
=======
			return &ParseError{"", "bad LOC Size, HorizPre or VertPre", l}
>>>>>>> deathstrox/main
		}
		l, _ = c.Next()
	}
	return nil
}

func (rr *HIP) parse(c *zlexer, o string) *ParseError {
	// HitLength is not represented
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad HIP PublicKeyAlgorithm", lex: l}
=======
		return &ParseError{"", "bad HIP PublicKeyAlgorithm", l}
>>>>>>> deathstrox/main
	}
	rr.PublicKeyAlgorithm = uint8(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	if l.token == "" || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad HIP Hit", lex: l}
=======
		return &ParseError{"", "bad HIP Hit", l}
>>>>>>> deathstrox/main
	}
	rr.Hit = l.token // This can not contain spaces, see RFC 5205 Section 6.
	rr.HitLength = uint8(len(rr.Hit)) / 2

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	if l.token == "" || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad HIP PublicKey", lex: l}
=======
		return &ParseError{"", "bad HIP PublicKey", l}
>>>>>>> deathstrox/main
	}
	rr.PublicKey = l.token // This cannot contain spaces
	decodedPK, decodedPKerr := base64.StdEncoding.DecodeString(rr.PublicKey)
	if decodedPKerr != nil {
<<<<<<< HEAD
		return &ParseError{err: "bad HIP PublicKey", lex: l}
=======
		return &ParseError{"", "bad HIP PublicKey", l}
>>>>>>> deathstrox/main
	}
	rr.PublicKeyLength = uint16(len(decodedPK))

	// RendezvousServers (if any)
	l, _ = c.Next()
	var xs []string
	for l.value != zNewline && l.value != zEOF {
		switch l.value {
		case zString:
			name, nameOk := toAbsoluteName(l.token, o)
			if l.err || !nameOk {
<<<<<<< HEAD
				return &ParseError{err: "bad HIP RendezvousServers", lex: l}
=======
				return &ParseError{"", "bad HIP RendezvousServers", l}
>>>>>>> deathstrox/main
			}
			xs = append(xs, name)
		case zBlank:
			// Ok
		default:
<<<<<<< HEAD
			return &ParseError{err: "bad HIP RendezvousServers", lex: l}
=======
			return &ParseError{"", "bad HIP RendezvousServers", l}
>>>>>>> deathstrox/main
		}
		l, _ = c.Next()
	}

	rr.RendezvousServers = xs
	return nil
}

func (rr *CERT) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	if v, ok := StringToCertType[l.token]; ok {
		rr.Type = v
	} else if i, err := strconv.ParseUint(l.token, 10, 16); err != nil {
<<<<<<< HEAD
		return &ParseError{err: "bad CERT Type", lex: l}
=======
		return &ParseError{"", "bad CERT Type", l}
>>>>>>> deathstrox/main
	} else {
		rr.Type = uint16(i)
	}
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad CERT KeyTag", lex: l}
=======
		return &ParseError{"", "bad CERT KeyTag", l}
>>>>>>> deathstrox/main
	}
	rr.KeyTag = uint16(i)
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	if v, ok := StringToAlgorithm[l.token]; ok {
		rr.Algorithm = v
	} else if i, err := strconv.ParseUint(l.token, 10, 8); err != nil {
<<<<<<< HEAD
		return &ParseError{err: "bad CERT Algorithm", lex: l}
=======
		return &ParseError{"", "bad CERT Algorithm", l}
>>>>>>> deathstrox/main
	} else {
		rr.Algorithm = uint8(i)
	}
	s, e1 := endingToString(c, "bad CERT Certificate")
	if e1 != nil {
		return e1
	}
	rr.Certificate = s
	return nil
}

func (rr *OPENPGPKEY) parse(c *zlexer, o string) *ParseError {
	s, e := endingToString(c, "bad OPENPGPKEY PublicKey")
	if e != nil {
		return e
	}
	rr.PublicKey = s
	return nil
}

func (rr *CSYNC) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	j, e := strconv.ParseUint(l.token, 10, 32)
	if e != nil {
		// Serial must be a number
<<<<<<< HEAD
		return &ParseError{err: "bad CSYNC serial", lex: l}
=======
		return &ParseError{"", "bad CSYNC serial", l}
>>>>>>> deathstrox/main
	}
	rr.Serial = uint32(j)

	c.Next() // zBlank

	l, _ = c.Next()
	j, e1 := strconv.ParseUint(l.token, 10, 16)
	if e1 != nil {
		// Serial must be a number
<<<<<<< HEAD
		return &ParseError{err: "bad CSYNC flags", lex: l}
=======
		return &ParseError{"", "bad CSYNC flags", l}
>>>>>>> deathstrox/main
	}
	rr.Flags = uint16(j)

	rr.TypeBitMap = make([]uint16, 0)
	var (
		k  uint16
		ok bool
	)
	l, _ = c.Next()
	for l.value != zNewline && l.value != zEOF {
		switch l.value {
		case zBlank:
			// Ok
		case zString:
			tokenUpper := strings.ToUpper(l.token)
			if k, ok = StringToType[tokenUpper]; !ok {
				if k, ok = typeToInt(l.token); !ok {
<<<<<<< HEAD
					return &ParseError{err: "bad CSYNC TypeBitMap", lex: l}
=======
					return &ParseError{"", "bad CSYNC TypeBitMap", l}
>>>>>>> deathstrox/main
				}
			}
			rr.TypeBitMap = append(rr.TypeBitMap, k)
		default:
<<<<<<< HEAD
			return &ParseError{err: "bad CSYNC TypeBitMap", lex: l}
=======
			return &ParseError{"", "bad CSYNC TypeBitMap", l}
>>>>>>> deathstrox/main
		}
		l, _ = c.Next()
	}
	return nil
}

func (rr *ZONEMD) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 32)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad ZONEMD Serial", lex: l}
=======
		return &ParseError{"", "bad ZONEMD Serial", l}
>>>>>>> deathstrox/main
	}
	rr.Serial = uint32(i)

	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad ZONEMD Scheme", lex: l}
=======
		return &ParseError{"", "bad ZONEMD Scheme", l}
>>>>>>> deathstrox/main
	}
	rr.Scheme = uint8(i)

	c.Next() // zBlank
	l, _ = c.Next()
	i, err := strconv.ParseUint(l.token, 10, 8)
	if err != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad ZONEMD Hash Algorithm", lex: l}
=======
		return &ParseError{"", "bad ZONEMD Hash Algorithm", l}
>>>>>>> deathstrox/main
	}
	rr.Hash = uint8(i)

	s, e2 := endingToString(c, "bad ZONEMD Digest")
	if e2 != nil {
		return e2
	}
	rr.Digest = s
	return nil
}

func (rr *SIG) parse(c *zlexer, o string) *ParseError { return rr.RRSIG.parse(c, o) }

func (rr *RRSIG) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	tokenUpper := strings.ToUpper(l.token)
	if t, ok := StringToType[tokenUpper]; !ok {
		if strings.HasPrefix(tokenUpper, "TYPE") {
			t, ok = typeToInt(l.token)
			if !ok {
<<<<<<< HEAD
				return &ParseError{err: "bad RRSIG Typecovered", lex: l}
			}
			rr.TypeCovered = t
		} else {
			return &ParseError{err: "bad RRSIG Typecovered", lex: l}
=======
				return &ParseError{"", "bad RRSIG Typecovered", l}
			}
			rr.TypeCovered = t
		} else {
			return &ParseError{"", "bad RRSIG Typecovered", l}
>>>>>>> deathstrox/main
		}
	} else {
		rr.TypeCovered = t
	}

	c.Next() // zBlank
	l, _ = c.Next()
<<<<<<< HEAD
	if l.err {
		return &ParseError{err: "bad RRSIG Algorithm", lex: l}
	}
	i, e := strconv.ParseUint(l.token, 10, 8)
	rr.Algorithm = uint8(i) // if 0 we'll check the mnemonic in the if
	if e != nil {
		v, ok := StringToAlgorithm[l.token]
		if !ok {
			return &ParseError{err: "bad RRSIG Algorithm", lex: l}
		}
		rr.Algorithm = v
	}
=======
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
		return &ParseError{"", "bad RRSIG Algorithm", l}
	}
	rr.Algorithm = uint8(i)
>>>>>>> deathstrox/main

	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad RRSIG Labels", lex: l}
=======
		return &ParseError{"", "bad RRSIG Labels", l}
>>>>>>> deathstrox/main
	}
	rr.Labels = uint8(i)

	c.Next() // zBlank
	l, _ = c.Next()
	i, e2 := strconv.ParseUint(l.token, 10, 32)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad RRSIG OrigTtl", lex: l}
=======
		return &ParseError{"", "bad RRSIG OrigTtl", l}
>>>>>>> deathstrox/main
	}
	rr.OrigTtl = uint32(i)

	c.Next() // zBlank
	l, _ = c.Next()
	if i, err := StringToTime(l.token); err != nil {
		// Try to see if all numeric and use it as epoch
		if i, err := strconv.ParseUint(l.token, 10, 32); err == nil {
			rr.Expiration = uint32(i)
		} else {
<<<<<<< HEAD
			return &ParseError{err: "bad RRSIG Expiration", lex: l}
=======
			return &ParseError{"", "bad RRSIG Expiration", l}
>>>>>>> deathstrox/main
		}
	} else {
		rr.Expiration = i
	}

	c.Next() // zBlank
	l, _ = c.Next()
	if i, err := StringToTime(l.token); err != nil {
		if i, err := strconv.ParseUint(l.token, 10, 32); err == nil {
			rr.Inception = uint32(i)
		} else {
<<<<<<< HEAD
			return &ParseError{err: "bad RRSIG Inception", lex: l}
=======
			return &ParseError{"", "bad RRSIG Inception", l}
>>>>>>> deathstrox/main
		}
	} else {
		rr.Inception = i
	}

	c.Next() // zBlank
	l, _ = c.Next()
	i, e3 := strconv.ParseUint(l.token, 10, 16)
	if e3 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad RRSIG KeyTag", lex: l}
=======
		return &ParseError{"", "bad RRSIG KeyTag", l}
>>>>>>> deathstrox/main
	}
	rr.KeyTag = uint16(i)

	c.Next() // zBlank
	l, _ = c.Next()
	rr.SignerName = l.token
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad RRSIG SignerName", lex: l}
=======
		return &ParseError{"", "bad RRSIG SignerName", l}
>>>>>>> deathstrox/main
	}
	rr.SignerName = name

	s, e4 := endingToString(c, "bad RRSIG Signature")
	if e4 != nil {
		return e4
	}
	rr.Signature = s

	return nil
}

<<<<<<< HEAD
func (rr *NXT) parse(c *zlexer, o string) *ParseError { return rr.NSEC.parse(c, o) }

=======
>>>>>>> deathstrox/main
func (rr *NSEC) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC NextDomain", lex: l}
=======
		return &ParseError{"", "bad NSEC NextDomain", l}
>>>>>>> deathstrox/main
	}
	rr.NextDomain = name

	rr.TypeBitMap = make([]uint16, 0)
	var (
		k  uint16
		ok bool
	)
	l, _ = c.Next()
	for l.value != zNewline && l.value != zEOF {
		switch l.value {
		case zBlank:
			// Ok
		case zString:
			tokenUpper := strings.ToUpper(l.token)
			if k, ok = StringToType[tokenUpper]; !ok {
				if k, ok = typeToInt(l.token); !ok {
<<<<<<< HEAD
					return &ParseError{err: "bad NSEC TypeBitMap", lex: l}
=======
					return &ParseError{"", "bad NSEC TypeBitMap", l}
>>>>>>> deathstrox/main
				}
			}
			rr.TypeBitMap = append(rr.TypeBitMap, k)
		default:
<<<<<<< HEAD
			return &ParseError{err: "bad NSEC TypeBitMap", lex: l}
=======
			return &ParseError{"", "bad NSEC TypeBitMap", l}
>>>>>>> deathstrox/main
		}
		l, _ = c.Next()
	}
	return nil
}

func (rr *NSEC3) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC3 Hash", lex: l}
=======
		return &ParseError{"", "bad NSEC3 Hash", l}
>>>>>>> deathstrox/main
	}
	rr.Hash = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC3 Flags", lex: l}
=======
		return &ParseError{"", "bad NSEC3 Flags", l}
>>>>>>> deathstrox/main
	}
	rr.Flags = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e2 := strconv.ParseUint(l.token, 10, 16)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC3 Iterations", lex: l}
=======
		return &ParseError{"", "bad NSEC3 Iterations", l}
>>>>>>> deathstrox/main
	}
	rr.Iterations = uint16(i)
	c.Next()
	l, _ = c.Next()
	if l.token == "" || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC3 Salt", lex: l}
=======
		return &ParseError{"", "bad NSEC3 Salt", l}
>>>>>>> deathstrox/main
	}
	if l.token != "-" {
		rr.SaltLength = uint8(len(l.token)) / 2
		rr.Salt = l.token
	}

	c.Next()
	l, _ = c.Next()
	if l.token == "" || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC3 NextDomain", lex: l}
=======
		return &ParseError{"", "bad NSEC3 NextDomain", l}
>>>>>>> deathstrox/main
	}
	rr.HashLength = 20 // Fix for NSEC3 (sha1 160 bits)
	rr.NextDomain = l.token

	rr.TypeBitMap = make([]uint16, 0)
	var (
		k  uint16
		ok bool
	)
	l, _ = c.Next()
	for l.value != zNewline && l.value != zEOF {
		switch l.value {
		case zBlank:
			// Ok
		case zString:
			tokenUpper := strings.ToUpper(l.token)
			if k, ok = StringToType[tokenUpper]; !ok {
				if k, ok = typeToInt(l.token); !ok {
<<<<<<< HEAD
					return &ParseError{err: "bad NSEC3 TypeBitMap", lex: l}
=======
					return &ParseError{"", "bad NSEC3 TypeBitMap", l}
>>>>>>> deathstrox/main
				}
			}
			rr.TypeBitMap = append(rr.TypeBitMap, k)
		default:
<<<<<<< HEAD
			return &ParseError{err: "bad NSEC3 TypeBitMap", lex: l}
=======
			return &ParseError{"", "bad NSEC3 TypeBitMap", l}
>>>>>>> deathstrox/main
		}
		l, _ = c.Next()
	}
	return nil
}

func (rr *NSEC3PARAM) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC3PARAM Hash", lex: l}
=======
		return &ParseError{"", "bad NSEC3PARAM Hash", l}
>>>>>>> deathstrox/main
	}
	rr.Hash = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC3PARAM Flags", lex: l}
=======
		return &ParseError{"", "bad NSEC3PARAM Flags", l}
>>>>>>> deathstrox/main
	}
	rr.Flags = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e2 := strconv.ParseUint(l.token, 10, 16)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NSEC3PARAM Iterations", lex: l}
=======
		return &ParseError{"", "bad NSEC3PARAM Iterations", l}
>>>>>>> deathstrox/main
	}
	rr.Iterations = uint16(i)
	c.Next()
	l, _ = c.Next()
	if l.token != "-" {
		rr.SaltLength = uint8(len(l.token) / 2)
		rr.Salt = l.token
	}
	return slurpRemainder(c)
}

func (rr *EUI48) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	if len(l.token) != 17 || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad EUI48 Address", lex: l}
=======
		return &ParseError{"", "bad EUI48 Address", l}
>>>>>>> deathstrox/main
	}
	addr := make([]byte, 12)
	dash := 0
	for i := 0; i < 10; i += 2 {
		addr[i] = l.token[i+dash]
		addr[i+1] = l.token[i+1+dash]
		dash++
		if l.token[i+1+dash] != '-' {
<<<<<<< HEAD
			return &ParseError{err: "bad EUI48 Address", lex: l}
=======
			return &ParseError{"", "bad EUI48 Address", l}
>>>>>>> deathstrox/main
		}
	}
	addr[10] = l.token[15]
	addr[11] = l.token[16]

	i, e := strconv.ParseUint(string(addr), 16, 48)
	if e != nil {
<<<<<<< HEAD
		return &ParseError{err: "bad EUI48 Address", lex: l}
=======
		return &ParseError{"", "bad EUI48 Address", l}
>>>>>>> deathstrox/main
	}
	rr.Address = i
	return slurpRemainder(c)
}

func (rr *EUI64) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	if len(l.token) != 23 || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad EUI64 Address", lex: l}
=======
		return &ParseError{"", "bad EUI64 Address", l}
>>>>>>> deathstrox/main
	}
	addr := make([]byte, 16)
	dash := 0
	for i := 0; i < 14; i += 2 {
		addr[i] = l.token[i+dash]
		addr[i+1] = l.token[i+1+dash]
		dash++
		if l.token[i+1+dash] != '-' {
<<<<<<< HEAD
			return &ParseError{err: "bad EUI64 Address", lex: l}
=======
			return &ParseError{"", "bad EUI64 Address", l}
>>>>>>> deathstrox/main
		}
	}
	addr[14] = l.token[21]
	addr[15] = l.token[22]

	i, e := strconv.ParseUint(string(addr), 16, 64)
	if e != nil {
<<<<<<< HEAD
		return &ParseError{err: "bad EUI68 Address", lex: l}
=======
		return &ParseError{"", "bad EUI68 Address", l}
>>>>>>> deathstrox/main
	}
	rr.Address = i
	return slurpRemainder(c)
}

func (rr *SSHFP) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad SSHFP Algorithm", lex: l}
=======
		return &ParseError{"", "bad SSHFP Algorithm", l}
>>>>>>> deathstrox/main
	}
	rr.Algorithm = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad SSHFP Type", lex: l}
=======
		return &ParseError{"", "bad SSHFP Type", l}
>>>>>>> deathstrox/main
	}
	rr.Type = uint8(i)
	c.Next() // zBlank
	s, e2 := endingToString(c, "bad SSHFP Fingerprint")
	if e2 != nil {
		return e2
	}
	rr.FingerPrint = s
	return nil
}

func (rr *DNSKEY) parseDNSKEY(c *zlexer, o, typ string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad " + typ + " Flags", lex: l}
=======
		return &ParseError{"", "bad " + typ + " Flags", l}
>>>>>>> deathstrox/main
	}
	rr.Flags = uint16(i)
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad " + typ + " Protocol", lex: l}
=======
		return &ParseError{"", "bad " + typ + " Protocol", l}
>>>>>>> deathstrox/main
	}
	rr.Protocol = uint8(i)
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	i, e2 := strconv.ParseUint(l.token, 10, 8)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad " + typ + " Algorithm", lex: l}
=======
		return &ParseError{"", "bad " + typ + " Algorithm", l}
>>>>>>> deathstrox/main
	}
	rr.Algorithm = uint8(i)
	s, e3 := endingToString(c, "bad "+typ+" PublicKey")
	if e3 != nil {
		return e3
	}
	rr.PublicKey = s
	return nil
}

func (rr *DNSKEY) parse(c *zlexer, o string) *ParseError  { return rr.parseDNSKEY(c, o, "DNSKEY") }
func (rr *KEY) parse(c *zlexer, o string) *ParseError     { return rr.parseDNSKEY(c, o, "KEY") }
func (rr *CDNSKEY) parse(c *zlexer, o string) *ParseError { return rr.parseDNSKEY(c, o, "CDNSKEY") }
func (rr *DS) parse(c *zlexer, o string) *ParseError      { return rr.parseDS(c, o, "DS") }
func (rr *DLV) parse(c *zlexer, o string) *ParseError     { return rr.parseDS(c, o, "DLV") }
func (rr *CDS) parse(c *zlexer, o string) *ParseError     { return rr.parseDS(c, o, "CDS") }

<<<<<<< HEAD
func (rr *IPSECKEY) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	num, err := strconv.ParseUint(l.token, 10, 8)
	if err != nil || l.err {
		return &ParseError{err: "bad IPSECKEY value", lex: l}
	}
	rr.Precedence = uint8(num)
	c.Next() // zBlank

	l, _ = c.Next()
	num, err = strconv.ParseUint(l.token, 10, 8)
	if err != nil || l.err {
		return &ParseError{err: "bad IPSECKEY value", lex: l}
	}
	rr.GatewayType = uint8(num)
	c.Next() // zBlank

	l, _ = c.Next()
	num, err = strconv.ParseUint(l.token, 10, 8)
	if err != nil || l.err {
		return &ParseError{err: "bad IPSECKEY value", lex: l}
	}
	rr.Algorithm = uint8(num)
	c.Next() // zBlank

	l, _ = c.Next()
	if l.err {
		return &ParseError{err: "bad IPSECKEY gateway", lex: l}
	}

	rr.GatewayAddr, rr.GatewayHost, err = parseAddrHostUnion(l.token, o, rr.GatewayType)
	if err != nil {
		return &ParseError{wrappedErr: fmt.Errorf("IPSECKEY %w", err), lex: l}
	}

	c.Next() // zBlank

	s, pErr := endingToString(c, "bad IPSECKEY PublicKey")
	if pErr != nil {
		return pErr
	}
	rr.PublicKey = s
	return slurpRemainder(c)
}

func (rr *AMTRELAY) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	num, err := strconv.ParseUint(l.token, 10, 8)
	if err != nil || l.err {
		return &ParseError{err: "bad AMTRELAY value", lex: l}
	}
	rr.Precedence = uint8(num)
	c.Next() // zBlank

	l, _ = c.Next()
	if l.err || !(l.token == "0" || l.token == "1") {
		return &ParseError{err: "bad discovery value", lex: l}
	}
	if l.token == "1" {
		rr.GatewayType = 0x80
	}

	c.Next() // zBlank

	l, _ = c.Next()
	num, err = strconv.ParseUint(l.token, 10, 8)
	if err != nil || l.err {
		return &ParseError{err: "bad AMTRELAY value", lex: l}
	}
	rr.GatewayType |= uint8(num)
	c.Next() // zBlank

	l, _ = c.Next()
	if l.err {
		return &ParseError{err: "bad AMTRELAY gateway", lex: l}
	}

	rr.GatewayAddr, rr.GatewayHost, err = parseAddrHostUnion(l.token, o, rr.GatewayType&0x7f)
	if err != nil {
		return &ParseError{wrappedErr: fmt.Errorf("AMTRELAY %w", err), lex: l}
	}

	return slurpRemainder(c)
}

// same constants and parsing between IPSECKEY and AMTRELAY
func parseAddrHostUnion(token, o string, gatewayType uint8) (addr net.IP, host string, err error) {
	switch gatewayType {
	case IPSECGatewayNone:
		if token != "." {
			return addr, host, errors.New("gateway type none with gateway set")
		}
	case IPSECGatewayIPv4, IPSECGatewayIPv6:
		addr = net.ParseIP(token)
		if addr == nil {
			return addr, host, errors.New("gateway IP invalid")
		}
		if (addr.To4() == nil) == (gatewayType == IPSECGatewayIPv4) {
			return addr, host, errors.New("gateway IP family mismatch")
		}
	case IPSECGatewayHost:
		var ok bool
		host, ok = toAbsoluteName(token, o)
		if !ok {
			return addr, host, errors.New("invalid gateway host")
		}
	}

	return addr, host, nil
}

=======
>>>>>>> deathstrox/main
func (rr *RKEY) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad RKEY Flags", lex: l}
=======
		return &ParseError{"", "bad RKEY Flags", l}
>>>>>>> deathstrox/main
	}
	rr.Flags = uint16(i)
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad RKEY Protocol", lex: l}
=======
		return &ParseError{"", "bad RKEY Protocol", l}
>>>>>>> deathstrox/main
	}
	rr.Protocol = uint8(i)
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	i, e2 := strconv.ParseUint(l.token, 10, 8)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad RKEY Algorithm", lex: l}
=======
		return &ParseError{"", "bad RKEY Algorithm", l}
>>>>>>> deathstrox/main
	}
	rr.Algorithm = uint8(i)
	s, e3 := endingToString(c, "bad RKEY PublicKey")
	if e3 != nil {
		return e3
	}
	rr.PublicKey = s
	return nil
}

func (rr *EID) parse(c *zlexer, o string) *ParseError {
	s, e := endingToString(c, "bad EID Endpoint")
	if e != nil {
		return e
	}
	rr.Endpoint = s
	return nil
}

func (rr *NIMLOC) parse(c *zlexer, o string) *ParseError {
	s, e := endingToString(c, "bad NIMLOC Locator")
	if e != nil {
		return e
	}
	rr.Locator = s
	return nil
}

func (rr *GPOS) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	_, e := strconv.ParseFloat(l.token, 64)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad GPOS Longitude", lex: l}
=======
		return &ParseError{"", "bad GPOS Longitude", l}
>>>>>>> deathstrox/main
	}
	rr.Longitude = l.token
	c.Next() // zBlank
	l, _ = c.Next()
	_, e1 := strconv.ParseFloat(l.token, 64)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad GPOS Latitude", lex: l}
=======
		return &ParseError{"", "bad GPOS Latitude", l}
>>>>>>> deathstrox/main
	}
	rr.Latitude = l.token
	c.Next() // zBlank
	l, _ = c.Next()
	_, e2 := strconv.ParseFloat(l.token, 64)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad GPOS Altitude", lex: l}
=======
		return &ParseError{"", "bad GPOS Altitude", l}
>>>>>>> deathstrox/main
	}
	rr.Altitude = l.token
	return slurpRemainder(c)
}

func (rr *DS) parseDS(c *zlexer, o, typ string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad " + typ + " KeyTag", lex: l}
=======
		return &ParseError{"", "bad " + typ + " KeyTag", l}
>>>>>>> deathstrox/main
	}
	rr.KeyTag = uint16(i)
	c.Next() // zBlank
	l, _ = c.Next()
	if i, err := strconv.ParseUint(l.token, 10, 8); err != nil {
		tokenUpper := strings.ToUpper(l.token)
		i, ok := StringToAlgorithm[tokenUpper]
		if !ok || l.err {
<<<<<<< HEAD
			return &ParseError{err: "bad " + typ + " Algorithm", lex: l}
=======
			return &ParseError{"", "bad " + typ + " Algorithm", l}
>>>>>>> deathstrox/main
		}
		rr.Algorithm = i
	} else {
		rr.Algorithm = uint8(i)
	}
	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad " + typ + " DigestType", lex: l}
=======
		return &ParseError{"", "bad " + typ + " DigestType", l}
>>>>>>> deathstrox/main
	}
	rr.DigestType = uint8(i)
	s, e2 := endingToString(c, "bad "+typ+" Digest")
	if e2 != nil {
		return e2
	}
	rr.Digest = s
	return nil
}

func (rr *TA) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad TA KeyTag", lex: l}
=======
		return &ParseError{"", "bad TA KeyTag", l}
>>>>>>> deathstrox/main
	}
	rr.KeyTag = uint16(i)
	c.Next() // zBlank
	l, _ = c.Next()
	if i, err := strconv.ParseUint(l.token, 10, 8); err != nil {
		tokenUpper := strings.ToUpper(l.token)
		i, ok := StringToAlgorithm[tokenUpper]
		if !ok || l.err {
<<<<<<< HEAD
			return &ParseError{err: "bad TA Algorithm", lex: l}
=======
			return &ParseError{"", "bad TA Algorithm", l}
>>>>>>> deathstrox/main
		}
		rr.Algorithm = i
	} else {
		rr.Algorithm = uint8(i)
	}
	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad TA DigestType", lex: l}
=======
		return &ParseError{"", "bad TA DigestType", l}
>>>>>>> deathstrox/main
	}
	rr.DigestType = uint8(i)
	s, e2 := endingToString(c, "bad TA Digest")
	if e2 != nil {
		return e2
	}
	rr.Digest = s
	return nil
}

func (rr *TLSA) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad TLSA Usage", lex: l}
=======
		return &ParseError{"", "bad TLSA Usage", l}
>>>>>>> deathstrox/main
	}
	rr.Usage = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad TLSA Selector", lex: l}
=======
		return &ParseError{"", "bad TLSA Selector", l}
>>>>>>> deathstrox/main
	}
	rr.Selector = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e2 := strconv.ParseUint(l.token, 10, 8)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad TLSA MatchingType", lex: l}
=======
		return &ParseError{"", "bad TLSA MatchingType", l}
>>>>>>> deathstrox/main
	}
	rr.MatchingType = uint8(i)
	// So this needs be e2 (i.e. different than e), because...??t
	s, e3 := endingToString(c, "bad TLSA Certificate")
	if e3 != nil {
		return e3
	}
	rr.Certificate = s
	return nil
}

func (rr *SMIMEA) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad SMIMEA Usage", lex: l}
=======
		return &ParseError{"", "bad SMIMEA Usage", l}
>>>>>>> deathstrox/main
	}
	rr.Usage = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad SMIMEA Selector", lex: l}
=======
		return &ParseError{"", "bad SMIMEA Selector", l}
>>>>>>> deathstrox/main
	}
	rr.Selector = uint8(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e2 := strconv.ParseUint(l.token, 10, 8)
	if e2 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad SMIMEA MatchingType", lex: l}
=======
		return &ParseError{"", "bad SMIMEA MatchingType", l}
>>>>>>> deathstrox/main
	}
	rr.MatchingType = uint8(i)
	// So this needs be e2 (i.e. different than e), because...??t
	s, e3 := endingToString(c, "bad SMIMEA Certificate")
	if e3 != nil {
		return e3
	}
	rr.Certificate = s
	return nil
}

func (rr *RFC3597) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	if l.token != "\\#" {
<<<<<<< HEAD
		return &ParseError{err: "bad RFC3597 Rdata", lex: l}
=======
		return &ParseError{"", "bad RFC3597 Rdata", l}
>>>>>>> deathstrox/main
	}

	c.Next() // zBlank
	l, _ = c.Next()
	rdlength, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad RFC3597 Rdata ", lex: l}
=======
		return &ParseError{"", "bad RFC3597 Rdata ", l}
>>>>>>> deathstrox/main
	}

	s, e1 := endingToString(c, "bad RFC3597 Rdata")
	if e1 != nil {
		return e1
	}
	if int(rdlength)*2 != len(s) {
<<<<<<< HEAD
		return &ParseError{err: "bad RFC3597 Rdata", lex: l}
=======
		return &ParseError{"", "bad RFC3597 Rdata", l}
>>>>>>> deathstrox/main
	}
	rr.Rdata = s
	return nil
}

func (rr *SPF) parse(c *zlexer, o string) *ParseError {
	s, e := endingToTxtSlice(c, "bad SPF Txt")
	if e != nil {
		return e
	}
	rr.Txt = s
	return nil
}

func (rr *AVC) parse(c *zlexer, o string) *ParseError {
	s, e := endingToTxtSlice(c, "bad AVC Txt")
	if e != nil {
		return e
	}
	rr.Txt = s
	return nil
}

func (rr *TXT) parse(c *zlexer, o string) *ParseError {
	// no zBlank reading here, because all this rdata is TXT
	s, e := endingToTxtSlice(c, "bad TXT Txt")
	if e != nil {
		return e
	}
	rr.Txt = s
	return nil
}

// identical to setTXT
func (rr *NINFO) parse(c *zlexer, o string) *ParseError {
	s, e := endingToTxtSlice(c, "bad NINFO ZSData")
	if e != nil {
		return e
	}
	rr.ZSData = s
	return nil
}

func (rr *URI) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad URI Priority", lex: l}
=======
		return &ParseError{"", "bad URI Priority", l}
>>>>>>> deathstrox/main
	}
	rr.Priority = uint16(i)
	c.Next() // zBlank
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 16)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad URI Weight", lex: l}
=======
		return &ParseError{"", "bad URI Weight", l}
>>>>>>> deathstrox/main
	}
	rr.Weight = uint16(i)

	c.Next() // zBlank
	s, e2 := endingToTxtSlice(c, "bad URI Target")
	if e2 != nil {
		return e2
	}
	if len(s) != 1 {
<<<<<<< HEAD
		return &ParseError{err: "bad URI Target", lex: l}
=======
		return &ParseError{"", "bad URI Target", l}
>>>>>>> deathstrox/main
	}
	rr.Target = s[0]
	return nil
}

func (rr *DHCID) parse(c *zlexer, o string) *ParseError {
	// awesome record to parse!
	s, e := endingToString(c, "bad DHCID Digest")
	if e != nil {
		return e
	}
	rr.Digest = s
	return nil
}

func (rr *NID) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad NID Preference", lex: l}
=======
		return &ParseError{"", "bad NID Preference", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	u, e1 := stringToNodeID(l)
	if e1 != nil || l.err {
		return e1
	}
	rr.NodeID = u
	return slurpRemainder(c)
}

func (rr *L32) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad L32 Preference", lex: l}
=======
		return &ParseError{"", "bad L32 Preference", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Locator32 = net.ParseIP(l.token)
	if rr.Locator32 == nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad L32 Locator", lex: l}
=======
		return &ParseError{"", "bad L32 Locator", l}
>>>>>>> deathstrox/main
	}
	return slurpRemainder(c)
}

func (rr *LP) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad LP Preference", lex: l}
=======
		return &ParseError{"", "bad LP Preference", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Fqdn = l.token
	name, nameOk := toAbsoluteName(l.token, o)
	if l.err || !nameOk {
<<<<<<< HEAD
		return &ParseError{err: "bad LP Fqdn", lex: l}
=======
		return &ParseError{"", "bad LP Fqdn", l}
>>>>>>> deathstrox/main
	}
	rr.Fqdn = name
	return slurpRemainder(c)
}

func (rr *L64) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad L64 Preference", lex: l}
=======
		return &ParseError{"", "bad L64 Preference", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)
	c.Next()        // zBlank
	l, _ = c.Next() // zString
	u, e1 := stringToNodeID(l)
	if e1 != nil || l.err {
		return e1
	}
	rr.Locator64 = u
	return slurpRemainder(c)
}

func (rr *UID) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 32)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad UID Uid", lex: l}
=======
		return &ParseError{"", "bad UID Uid", l}
>>>>>>> deathstrox/main
	}
	rr.Uid = uint32(i)
	return slurpRemainder(c)
}

func (rr *GID) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 32)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad GID Gid", lex: l}
=======
		return &ParseError{"", "bad GID Gid", l}
>>>>>>> deathstrox/main
	}
	rr.Gid = uint32(i)
	return slurpRemainder(c)
}

func (rr *UINFO) parse(c *zlexer, o string) *ParseError {
	s, e := endingToTxtSlice(c, "bad UINFO Uinfo")
	if e != nil {
		return e
	}
	if ln := len(s); ln == 0 {
		return nil
	}
	rr.Uinfo = s[0] // silently discard anything after the first character-string
	return nil
}

func (rr *PX) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 16)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad PX Preference", lex: l}
=======
		return &ParseError{"", "bad PX Preference", l}
>>>>>>> deathstrox/main
	}
	rr.Preference = uint16(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Map822 = l.token
	map822, map822Ok := toAbsoluteName(l.token, o)
	if l.err || !map822Ok {
<<<<<<< HEAD
		return &ParseError{err: "bad PX Map822", lex: l}
=======
		return &ParseError{"", "bad PX Map822", l}
>>>>>>> deathstrox/main
	}
	rr.Map822 = map822

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	rr.Mapx400 = l.token
	mapx400, mapx400Ok := toAbsoluteName(l.token, o)
	if l.err || !mapx400Ok {
<<<<<<< HEAD
		return &ParseError{err: "bad PX Mapx400", lex: l}
=======
		return &ParseError{"", "bad PX Mapx400", l}
>>>>>>> deathstrox/main
	}
	rr.Mapx400 = mapx400
	return slurpRemainder(c)
}

func (rr *CAA) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad CAA Flag", lex: l}
=======
		return &ParseError{"", "bad CAA Flag", l}
>>>>>>> deathstrox/main
	}
	rr.Flag = uint8(i)

	c.Next()        // zBlank
	l, _ = c.Next() // zString
	if l.value != zString {
<<<<<<< HEAD
		return &ParseError{err: "bad CAA Tag", lex: l}
=======
		return &ParseError{"", "bad CAA Tag", l}
>>>>>>> deathstrox/main
	}
	rr.Tag = l.token

	c.Next() // zBlank
	s, e1 := endingToTxtSlice(c, "bad CAA Value")
	if e1 != nil {
		return e1
	}
	if len(s) != 1 {
<<<<<<< HEAD
		return &ParseError{err: "bad CAA Value", lex: l}
=======
		return &ParseError{"", "bad CAA Value", l}
>>>>>>> deathstrox/main
	}
	rr.Value = s[0]
	return nil
}

func (rr *TKEY) parse(c *zlexer, o string) *ParseError {
	l, _ := c.Next()

	// Algorithm
	if l.value != zString {
<<<<<<< HEAD
		return &ParseError{err: "bad TKEY algorithm", lex: l}
=======
		return &ParseError{"", "bad TKEY algorithm", l}
>>>>>>> deathstrox/main
	}
	rr.Algorithm = l.token
	c.Next() // zBlank

	// Get the key length and key values
	l, _ = c.Next()
	i, e := strconv.ParseUint(l.token, 10, 8)
	if e != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad TKEY key length", lex: l}
=======
		return &ParseError{"", "bad TKEY key length", l}
>>>>>>> deathstrox/main
	}
	rr.KeySize = uint16(i)
	c.Next() // zBlank
	l, _ = c.Next()
	if l.value != zString {
<<<<<<< HEAD
		return &ParseError{err: "bad TKEY key", lex: l}
=======
		return &ParseError{"", "bad TKEY key", l}
>>>>>>> deathstrox/main
	}
	rr.Key = l.token
	c.Next() // zBlank

	// Get the otherdata length and string data
	l, _ = c.Next()
	i, e1 := strconv.ParseUint(l.token, 10, 8)
	if e1 != nil || l.err {
<<<<<<< HEAD
		return &ParseError{err: "bad TKEY otherdata length", lex: l}
=======
		return &ParseError{"", "bad TKEY otherdata length", l}
>>>>>>> deathstrox/main
	}
	rr.OtherLen = uint16(i)
	c.Next() // zBlank
	l, _ = c.Next()
	if l.value != zString {
<<<<<<< HEAD
		return &ParseError{err: "bad TKEY otherday", lex: l}
=======
		return &ParseError{"", "bad TKEY otherday", l}
>>>>>>> deathstrox/main
	}
	rr.OtherData = l.token
	return nil
}

func (rr *APL) parse(c *zlexer, o string) *ParseError {
	var prefixes []APLPrefix

	for {
		l, _ := c.Next()
		if l.value == zNewline || l.value == zEOF {
			break
		}
		if l.value == zBlank && prefixes != nil {
			continue
		}
		if l.value != zString {
<<<<<<< HEAD
			return &ParseError{err: "unexpected APL field", lex: l}
=======
			return &ParseError{"", "unexpected APL field", l}
>>>>>>> deathstrox/main
		}

		// Expected format: [!]afi:address/prefix

		colon := strings.IndexByte(l.token, ':')
		if colon == -1 {
<<<<<<< HEAD
			return &ParseError{err: "missing colon in APL field", lex: l}
=======
			return &ParseError{"", "missing colon in APL field", l}
>>>>>>> deathstrox/main
		}

		family, cidr := l.token[:colon], l.token[colon+1:]

		var negation bool
		if family != "" && family[0] == '!' {
			negation = true
			family = family[1:]
		}

		afi, e := strconv.ParseUint(family, 10, 16)
		if e != nil {
<<<<<<< HEAD
			return &ParseError{wrappedErr: fmt.Errorf("failed to parse APL family: %w", e), lex: l}
=======
			return &ParseError{"", "failed to parse APL family: " + e.Error(), l}
>>>>>>> deathstrox/main
		}
		var addrLen int
		switch afi {
		case 1:
			addrLen = net.IPv4len
		case 2:
			addrLen = net.IPv6len
		default:
<<<<<<< HEAD
			return &ParseError{err: "unrecognized APL family", lex: l}
=======
			return &ParseError{"", "unrecognized APL family", l}
>>>>>>> deathstrox/main
		}

		ip, subnet, e1 := net.ParseCIDR(cidr)
		if e1 != nil {
<<<<<<< HEAD
			return &ParseError{wrappedErr: fmt.Errorf("failed to parse APL address: %w", e1), lex: l}
		}
		if !ip.Equal(subnet.IP) {
			return &ParseError{err: "extra bits in APL address", lex: l}
		}

		if len(subnet.IP) != addrLen {
			return &ParseError{err: "address mismatch with the APL family", lex: l}
=======
			return &ParseError{"", "failed to parse APL address: " + e1.Error(), l}
		}
		if !ip.Equal(subnet.IP) {
			return &ParseError{"", "extra bits in APL address", l}
		}

		if len(subnet.IP) != addrLen {
			return &ParseError{"", "address mismatch with the APL family", l}
>>>>>>> deathstrox/main
		}

		prefixes = append(prefixes, APLPrefix{
			Negation: negation,
			Network:  *subnet,
		})
	}

	rr.Prefixes = prefixes
	return nil
}
