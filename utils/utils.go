package utils

import(
	"tcm/types"
	"encoding/json"
	"os"
	"fmt"
)

func DeleteCertificateByDomain(data *types.CertificatesMap, domain string) {
    for key, certs := range *data {
        var filteredCerts []types.Certificate
        for _, cert := range certs.Certificates {
            if cert.Domain.Main != domain {
                filteredCerts = append(filteredCerts, cert)
            }
        }
        (*data)[key] = types.Certificates{
            Account:      certs.Account,
            Certificates: filteredCerts,
        }
    }
}

func DeleteSansByDomain(data *types.CertificatesMap, sans string) {
    for key, certs := range *data {
        for i, cert := range certs.Certificates {
            var filteredSans []string
            for _, s := range cert.Domain.Sans {
                if s != sans {
                    filteredSans = append(filteredSans, s)
                }
            }
            (*data)[key].Certificates[i].Domain.Sans = filteredSans
        }
    }
}

func LoadJSONFile(filepath string) (types.CertificatesMap, error) {
    // Open the JSON file
    file, err := os.Open(filepath)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    // Decode JSON
    var data types.CertificatesMap
    if err := json.NewDecoder(file).Decode(&data); err != nil {
        return nil, fmt.Errorf("error decoding JSON: %v", err)
    }

    return data, nil
}

// Write updated JSON to a file
func WriteJSONToFile(data types.CertificatesMap, filename string) error {
    // Marshal data into JSON format with indentation
    updatedJSON, err := json.MarshalIndent(data, "", "    ")
    if err != nil {
        return err
    }

    // Write JSON data to file
    err = os.WriteFile(filename, updatedJSON, 0644)
    if err != nil {
        return err
    }

    return nil
}

// Function to extract values based on selection
func ExtractValues(data types.CertificatesMap, isDomain bool) ([]string, bool) {
        var values []string
        for _, certs := range data {
                for _, cert := range certs.Certificates {
                        if isDomain {
                                values = append(values, cert.Domain.Main)
                        } else {
                                values = append(values, cert.Domain.Sans...)
                        }
                }
        }
        return values, isDomain
}
