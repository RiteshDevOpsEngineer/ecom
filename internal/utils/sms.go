// internal/utils/sms.go
package utils

// func SendSMS(mobile string, message string) error {
// 	userID := "2000232563"
// 	password := "Rsplit%402023"
// 	apiURL := "https://enterpriseapi.smsgupshup.com/GatewayAPI/rest"

// 	encodedMessage := url.QueryEscape(message)
// 	requestURL := fmt.Sprintf("%s?method=SendMessage&send_to=%s&msg=%s&msg_type=TEXT&userid=%s&auth_scheme=plain&password=%s&v=1.1&format=text",
// 		apiURL, mobile, encodedMessage, userID, password)

// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", requestURL, nil)
// 	if err != nil {
// 		return err
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to send SMS: %s", body)
// 	}

// 	fmt.Println("SMS sent successfully!")
// 	return nil
// }

func SendSMS(mobile string, message string) error {

	return nil
}
