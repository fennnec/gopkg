// Copyright 2012 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Password generator
//
// Usage:
//	pwdgen [options]... [id]...
//
// Algorithm:
//	base58(sha512(md5hex(encrypt_key+encrypt_salt)+site_id+site_salt)[0:16]
//
// Example:
//	pwdgen id0
//	pwdgen id0 id1 id2
//
//	pwdgen --encrypt-key=111 id0
//	pwdgen --encrypt-key=111 id0 id1 id2
//
//	pwdgen --encrypt-key=111 --encrypt-salt=fuckcsdn --site-salt=site0 id0 id1
//	pwdgen --encrypt-key=111 --encrypt-salt=fuckcsdn --site-salt=site0 id0 id1
//	pwdgen --encrypt-key=111 --encrypt-salt=fuckcsdn --site-salt=site0 id0 id1
//
//	# KeePass: See config.ini
//	# output: *.ini -> *.keepass1x.csv
//	pwdgen --keepass-config=config.ini
//	pwdgen --keepass-config=config.ini --encrypt-key=111
//	pwdgen --keepass-config=config.ini --encrypt-key=111 --encrypt-salt=fuckcsdn
//
//	pwdgen --version
//	pwdgen --help
//	pwdgen -h
//
// Use pwdgen as a Go package:
//	package main
//
//	import (
//		"fmt"
//		pwdgen "bitbucket.org/chai2010/pwdgen"
//	)
//	func main() {
//		fmt.Println(pwdgen.PwdGen("id0", "site0", "111", "fuckcsdn"))
//		// Output: 2jNXfMGoXTSK9pFS
//	}
//
// Report bugs to <chaishushan{AT}gmail.com>.
package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/chai2010/encoding/base58"
	"github.com/chai2010/encoding/ini"
)

const (
	majorVersion = 2
	minorVersion = 4
)

var (
	encrypt_key  = flag.String("encrypt-key", "", "Set encrypt key.")
	encrypt_salt = flag.String("encrypt-salt", "", "Set encrypt salt.")
	site_salt    = flag.String("site-salt", "", "Set site salt.")

	keepass_config = flag.String("keepass-config", "", "Generate KeePass 1.x CSV.")

	version = flag.Bool("version", false, "Show version and exit.")
	help    = flag.Bool("help", false, "Show usage and exit.")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: pwdgen [options]... [id]...\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "%s\n", `
Algorithm:
  base58(sha512(md5hex(encrypt_key+encrypt_salt)+site_id+site_salt)[0:16]

Example:
  pwdgen id0
  pwdgen id0 id1 id2
  
  pwdgen --encrypt-key=111 id0
  pwdgen --encrypt-key=111 id0 id1 id2
  
  pwdgen --encrypt-key=111 --encrypt-salt=fuckcsdn --site-salt=site0 id0 id1
  pwdgen --encrypt-key=111 --encrypt-salt=fuckcsdn --site-salt=site0 id0 id1
  pwdgen --encrypt-key=111 --encrypt-salt=fuckcsdn --site-salt=site0 id0 id1
  
  # KeePass: See config.ini
  # output: *.ini -> *.keepass1x.csv
  pwdgen --keepass-config=config.ini
  pwdgen --keepass-config=config.ini --encrypt-key=111
  pwdgen --keepass-config=config.ini --encrypt-key=111 --encrypt-salt=fuckcsdn
  
  pwdgen --version
  pwdgen --help
  pwdgen -h

Use pwdgen as a Go package:
  package main
  
  import (
      "fmt"
      pwdgen "bitbucket.org/chai2010/pwdgen"
  )
  func main() {
      fmt.Println(pwdgen.PwdGen("id0", "site0", "111", "fuckcsdn"))
      // Output: 2jNXfMGoXTSK9pFS
  }

Report bugs to <chaishushan{AT}gmail.com>.`)
	}
}

func parseCmdLine() {
	flag.Parse()

	if *version {
		fmt.Printf("pwdgen-%d.%d\n", majorVersion, minorVersion)
		os.Exit(0)
	}
	if *help {
		flag.Usage()
		os.Exit(0)
	}
}

// base58(sha512(md5hex(encrypt_key+encrypt_salt)+site_id+site_salt)[0:16]
func PwdGen(site_id, site_salt, encrypt_key, encrypt_salt string) string {
	md5 := md5.New()
	md5.Write([]byte(encrypt_key + encrypt_salt))
	md5Hex := fmt.Sprintf("%x", md5.Sum(nil))

	sha := sha512.New()
	sha.Write([]byte(md5Hex + site_id + site_salt))
	shaSum := sha.Sum(nil)

	pwd := base58.EncodeBase58(shaSum)[0:16]
	return string(pwd)
}

func main() {
	parseCmdLine()
	if flag.NArg() <= 0 && len(*keepass_config) <= 0 {
		flag.Usage()
		os.Exit(0)
	}

	if len(*encrypt_key) <= 0 {
		fmt.Printf("Encryption key: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		(*encrypt_key) = strings.TrimSpace(input)
		if len(*encrypt_key) <= 0 {
			fmt.Fprintf(os.Stderr, "ERROR: Key must be at least 1 characters.\n")
			os.Exit(-1)
		}
	}

	if *keepass_config != "" {
		dict, err := ini.LoadFile(*keepass_config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Load <%s> failed.\n", *keepass_config)
			os.Exit(-1)
		}

		csvName := strings.Replace(*keepass_config, ".ini", ".keepass1x.csv", -1)
		file, err := os.Create(csvName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Create <%s> failed.\n", csvName)
			os.Exit(-1)
		}
		defer file.Close()

		// KeePass 1.x csv head line
		fmt.Fprintf(file, `"%s","%s","%s","%s","%s"`+"\n",
			"Account", "Login Name", "Password", "Web Site", "Comments",
		)

		keepass_site_list := dict.GetSections()
		for i := 0; i < len(keepass_site_list); i++ {
			keepass_site_name := keepass_site_list[i]
			keepass_site_id := dict[keepass_site_name]["LoginName"]
			keepass_site_salt := keepass_site_name + dict[keepass_site_name]["SiteSalt"]
			keepass_site_url := dict[keepass_site_name]["WebSite"]
			keepass_site_comments := dict[keepass_site_name]["Comments"]
			keepass_site_pwd := PwdGen(keepass_site_id, keepass_site_salt, *encrypt_key, *encrypt_salt)

			fmt.Fprintf(file, `"%s","%s","%s","%s","%s"`+"\n",
				keepass_site_name, keepass_site_id, keepass_site_pwd,
				keepass_site_url, keepass_site_comments,
			)
		}
	} else {
		for i := 0; i < flag.NArg(); i++ {
			password := PwdGen(flag.Arg(i), *site_salt, *encrypt_key, *encrypt_salt)
			fmt.Printf("%s\n", password)
		}
	}
}
