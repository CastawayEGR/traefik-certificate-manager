// utils_test.go
package utils_test

import (
    "testing"
    "os"
    "sort"
    "reflect"
    "encoding/json"

    "tcm/utils"
    "tcm/types"
)

func TestDeleteCertificateByDomain(t *testing.T) {
    // Create mock data
    certs := types.Certificates{
        Account: types.Account{},
        Certificates: []types.Certificate{
            {Domain: types.Domain{Main: "example.com"}},
            {Domain: types.Domain{Main: "example2.com"}},
        },
    }
    data := types.CertificatesMap{"key": certs}

    // Call the function
    utils.DeleteCertificateByDomain(&data, "example.com")

    // Assert that the certificate with domain "example.com" is deleted
    if len(data["key"].Certificates) != 1 || data["key"].Certificates[0].Domain.Main != "example2.com" {
        t.Errorf("DeleteCertificateByDomain() failed, expected certificate with domain 'example.com' to be deleted")
    }
}

func TestDeleteSansByDomain(t *testing.T) {
    // Create mock data
    certs := types.Certificates{
        Account: types.Account{},
        Certificates: []types.Certificate{
            {Domain: types.Domain{Main: "example.com", Sans: []string{"www.example.com", "test.example.com"}}},
            {Domain: types.Domain{Main: "example2.com", Sans: []string{"www.example2.com"}}},
        },
    }
    data := types.CertificatesMap{"key": certs}

    // Call the function
    utils.DeleteSansByDomain(&data, "www.example.com")

    // Assert that the SAN "www.example.com" is deleted
    if len(data["key"].Certificates[0].Domain.Sans) != 1 || data["key"].Certificates[0].Domain.Sans[0] != "test.example.com" {
        t.Errorf("DeleteSansByDomain() failed, expected SAN 'www.example.com' to be deleted")
    }
}


func TestWriteJSONToFile(t *testing.T) {
    // Create temporary file
    tmpfile, err := os.CreateTemp("", "test_json")
    if err != nil {
        t.Fatalf("Error creating temporary file: %v", err)
    }
    defer os.Remove(tmpfile.Name())

    // Create mock data with a single key
    data := types.CertificatesMap{
        "key1": types.Certificates{
            Account: types.Account{},
            Certificates: []types.Certificate{
                {Domain: types.Domain{Main: "example.com"}},
                {Domain: types.Domain{Main: "example2.com"}},
            },
        },
    }

    // Call the function
    err = utils.WriteJSONToFile(data, tmpfile.Name())
    if err != nil {
        t.Fatalf("WriteJSONToFile() failed: %v", err)
    }

    // Read the file
    content, err := os.ReadFile(tmpfile.Name())
    if err != nil {
        t.Fatalf("Error reading file: %v", err)
    }

    // Assert the content
    expectedJSON := `{
    "key1": {
        "Account": {
            "Email": "",
            "Registration": {
                "body": {
                    "status": "",
                    "contact": null
                },
                "uri": ""
            },
            "PrivateKey": "",
            "KeyType": ""
        },
        "Certificates": [
            {
                "domain": {
                    "main": "example.com"
                },
                "certificate": "",
                "key": "",
                "Store": ""
            },
            {
                "domain": {
                    "main": "example2.com"
                },
                "certificate": "",
                "key": "",
                "Store": ""
            }
        ]
    }
}`

    if string(content) != expectedJSON {
        t.Errorf("WriteJSONToFile() wrote incorrect JSON content:\nExpected:\n%s\nGot:\n%s", expectedJSON, string(content))
    }
}

func TestExtractValues(t *testing.T) {
    // Create mock data
    data := types.CertificatesMap{
        "key1": types.Certificates{
            Certificates: []types.Certificate{
                {Domain: types.Domain{Main: "example.com", Sans: []string{"www.example.com", "test.example.com"}}},
                {Domain: types.Domain{Main: "example2.com", Sans: []string{"www.example2.com"}}},
            },
        },
        "key2": types.Certificates{
            Certificates: []types.Certificate{
                {Domain: types.Domain{Main: "example3.com", Sans: []string{"www.example3.com"}}},
            },
        },
    }

    // Call the function with isDomain=true
    domains, isDomain := utils.ExtractValues(data, true)
    expectedDomains := []string{"example.com", "example2.com", "example3.com"}
    sort.Strings(domains)
    sort.Strings(expectedDomains)
    if !reflect.DeepEqual(domains, expectedDomains) || !isDomain {
        t.Errorf("ExtractValues(isDomain=true) returned unexpected result:\nExpected Domains: %v\nReturned Domains: %v\nisDomain: %t", expectedDomains, domains, isDomain)
    }

    // Call the function with isDomain=false
    sans, isDomain := utils.ExtractValues(data, false)
    expectedSans := []string{"www.example.com", "test.example.com", "www.example2.com", "www.example3.com"}
    sort.Strings(sans)
    sort.Strings(expectedSans)
    if !reflect.DeepEqual(sans, expectedSans) || isDomain {
        t.Errorf("ExtractValues(isDomain=false) returned unexpected result:\nExpected Sans: %v\nReturned Sans: %v\nisDomain: %t", expectedSans, sans, isDomain)
    }
}

func TestLoadJSONFile(t *testing.T) {
    // Define a temporary JSON file with mock data
    tempFile := "/tmp/test.json"
    mockData := types.CertificatesMap{
        "key1": types.Certificates{
            Certificates: []types.Certificate{
                {Domain: types.Domain{Main: "example.com"}},
                {Domain: types.Domain{Main: "example2.com"}},
            },
        },
    }
    if err := writeTestData(tempFile, mockData); err != nil {
        t.Fatalf("Error writing test data to file: %v", err)
    }
    defer deleteTestFile(tempFile)

    // Call the function
    data, err := utils.LoadJSONFile(tempFile)
    if err != nil {
        t.Fatalf("LoadJSONFile() returned error: %v", err)
    }

    // Compare the loaded data with the expected data
    if len(data) != len(mockData) {
        t.Fatalf("Length of loaded data does not match expected length")
    }
    for key, expectedCerts := range mockData {
        loadedCerts, ok := data[key]
        if !ok {
            t.Fatalf("Loaded data does not contain key: %s", key)
        }
        if len(loadedCerts.Certificates) != len(expectedCerts.Certificates) {
            t.Fatalf("Length of certificates for key %s does not match expected length", key)
        }
        for i, expectedCert := range expectedCerts.Certificates {
            loadedCert := loadedCerts.Certificates[i]
            if loadedCert.Domain.Main != expectedCert.Domain.Main {
                t.Fatalf("Loaded certificate domain for key %s at index %d does not match expected value", key, i)
            }
        }
    }
}

func writeTestData(filepath string, data types.CertificatesMap) error {
    file, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    if err := encoder.Encode(data); err != nil {
        return err
    }
    return nil
}

func deleteTestFile(filepath string) {
    os.Remove(filepath)
}
