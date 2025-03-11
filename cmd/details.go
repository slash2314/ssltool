/*
Copyright © 2023 Dex Wood
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"ssltool/pkg/details"
	"time"

	"github.com/spf13/cobra"
)

// detailsCmd represents the details command
var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Retrieve certificates details.",
	Long:  `Retrieve details about certificates returned from a host.`,
	Run: func(cmd *cobra.Command, args []string) {
		retrieveDetails, err := details.RetrieveCertDetails(fmt.Sprintf("%s:%d", hostname, port), insecure)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, certDetails := range retrieveDetails {
			fmt.Printf("Issuer: %s\n  Expiration Date: %v\n  Issue Date: %v\n  Serial: %x\n", 
			certDetails.Issuer, 
			certDetails.NotAfter.Format(time.RFC3339), 
			certDetails.Cert.NotBefore.Format(time.RFC3339), 
			certDetails.Cert.SerialNumber)
			if len(certDetails.DNSNames) > 0 {
				fmt.Println("  DNS Names:")
				for _, name := range certDetails.DNSNames {
					fmt.Printf("  - %s\n", name)
				}
			}
			if displayCertPem {
				details.DisplayPemCertificate(certDetails)
			}
			fmt.Println()
		}
	},
}

var hostname = ""
var port = 443

var insecure = false

var displayCertPem = false
var detailsExample = `ssltool details --host www.example.com
ssltool details --host www.example.com --cert`

func init() {
	rootCmd.AddCommand(detailsCmd)
	detailsCmd.Example = detailsExample
	detailsCmd.Flags().StringVar(&hostname, "host", "", "hostname to check certificate.")
	detailsCmd.Flags().IntVar(&port, "port", 443, "port")
	detailsCmd.Flags().BoolVarP(&insecure, "insecure", "i", false, "Don't verify certificates.")
	detailsCmd.Flags().BoolVarP(&displayCertPem, "cert", "c", false, "Print certificate in pem format.")
	err := detailsCmd.MarkFlagRequired("host")
	if err != nil {
		log.Fatalln("Couldn't require the hostname argument.")
	}
}
