/*
Copyright © 2023 Dex Wood
*/
package cmd

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509/pkix"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/fs"
	"log"
	"os"
	"ssltool/pkg/gen"
	"strings"
	"syscall"
)

var encryptKey = false

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Used to generate a certificate request.",
	Long: `You can set the required parameters with the following environmental variables
or you can enter them interactively: COUNTRY, ORG, OU, LOCALITY, and PROVINCE.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		country := promptForInfo(scanner, "COUNTRY", "COUNTRY: ")
		org := promptForInfo(scanner, "ORG", "ORG: ")
		ou := promptForInfo(scanner, "OU", "OU: ")
		locality := promptForInfo(scanner, "LOCALITY", "LOCALITY: ")
		province := promptForInfo(scanner, "PROVINCE", "PROVINCE: ")

		if len(commonName) == 0 {
			fmt.Println("The common name must not be blank.")
			os.Exit(1)
		}
		encryptKeyPass := ""
		if encryptKey {
			fmt.Print("Password: ")
			// int cast is required for windows
			password, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				fmt.Println("Couldn't get password.")
				os.Exit(1)
			}
			encryptKeyPass = string(password)
		}

		subj := getSubject(country, org, ou, locality, province, commonName)
		var csrOutput gen.CsrOutputInfo
		key, err := rsa.GenerateKey(rand.Reader, bits)
		if err != nil {
			fmt.Println("Couldn't generate private key")
			os.Exit(1)
		}
		csrInfo := gen.CsrInputInfo{
			CommonName: commonName,
			Sans:       trimStrings(sans),
			Name:       subj,
			PrivKey:    key,
		}

		csrOutput, err = gen.NewCsrSecure(csrInfo, encryptKey, encryptKeyPass)
		if err != nil {
			fmt.Println("Couldn't generate CSR")
			os.Exit(1)
		}
		if csrOut == "-" {
			fmt.Printf("%s\n", csrOutput.CsrPem)
		} else {
			err := os.WriteFile(csrOut, []byte(csrOutput.CsrPem), fs.FileMode(0600))
			if err != nil {
				log.Fatalln("Couldn't write out csr file.")
			}
		}

		if keyOut == "-" {
			fmt.Printf("%s\n", csrOutput.PrivateKeyPem)
		} else {
			err := os.WriteFile(keyOut, []byte(csrOutput.PrivateKeyPem), fs.FileMode(0600))
			if err != nil {
				log.Fatalln("Couldn't write out key file.")
			}
		}
	},
}

func promptForInfo(scan *bufio.Scanner, env, prompt string) string {
	data, exists := os.LookupEnv(env)
	if !exists {
		data = input(scan, prompt)
	}
	return data
}

func trimStrings(s []string) []string {
	for i := range s {
		s[i] = strings.Trim(s[i], " ")
	}
	return s
}
func input(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func getSubject(country, org, ou, locality, province, commonName string) pkix.Name {
	return pkix.Name{
		Country:            []string{country},
		Organization:       []string{org},
		OrganizationalUnit: []string{ou},
		Locality:           []string{locality},
		Province:           []string{province},
		CommonName:         commonName,
	}
}

var (
	commonName = ""
	sans       = make([]string, 0)
	keyOut     = ""
	csrOut     = ""
	bits       = 2048
)

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Example = "LOCALITY=\"Bowling Green\" PROVINCE=\"Kentucky\" COUNTRY=\"US\" ORG=\"Example ORG\" OU=\"Example OU\" ./ssltool gen -c www.example.com"
	genCmd.Flags().StringVarP(&commonName, "cn", "c", "", "Common name")
	genCmd.Flags().StringSliceVarP(&sans, "sans", "s", []string{}, "Sans list. In the form www.example.com,www-prod01.example.edu")
	genCmd.Flags().StringVarP(&csrOut, "csrout", "", "-", "Csr out filename. - for stdout")
	genCmd.Flags().StringVarP(&keyOut, "keyout", "", "-", "Key out filename. - for stdout")
	genCmd.Flags().BoolVarP(&encryptKey, "encryptkey", "e", false, "Encrypt private key")
	genCmd.Flags().IntVarP(&bits, "bits", "b", 2048, "RSA bits")
	err := genCmd.MarkFlagRequired("cn")
	if err != nil {
		log.Fatalln("Couldn't mark cn as required.")
	}
}
