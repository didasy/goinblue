package goinblue

import (
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

const (
	WRONG_API_KEY = "wrongapikey"

	CODE_SUCCESS = "success"
	CODE_FAILURE = "failure"
	CODE_ERROR   = "error"

	METHOD = "POST"
)

var (
	validApiKey    string
	to             = map[string]string{}
	from           = []string{}
	toMobileNumber string
	fromSMSName    string
	id             int
)

func init() {
	validApiKey = os.Getenv("SENDINBLUE_API_KEY")

	var err error
	id, err = strconv.Atoi(os.Getenv("SENDINBLUE_EMAIL_TEMPLATE_ID"))
	if err != nil {
		panic(err)
	}

	toEmail := os.Getenv("SENDINBLUE_TO_EMAIL")
	toName := os.Getenv("SENDINBLUE_TO_NAME")

	to[toEmail] = toName

	fromEmail := os.Getenv("SENDINBLUE_FROM_EMAIL")
	fromName := os.Getenv("SENDINBLUE_FROM_NAME")

	from = append(from, fromEmail, fromName)

	toMobileNumber = os.Getenv("SENDINBLUE_TO_MOBILE_NUMBER")
	fromSMSName = os.Getenv("SENDINBLUE_FROM_SMS_NAME")
}

func TestSendMessageWrongAPIKey(t *testing.T) {
	Convey("Given a wrong API key", t, func() {
		apiKey := WRONG_API_KEY

		Convey("And a client", func() {
			client := NewClient(apiKey)

			Convey("And a POST method", func() {
				method := METHOD

				Convey("And a target url for email", func() {
					url := BASE_URL + EMAIL_URL

					Convey("And a proper body", func() {
						body := &bytes.Buffer{}

						Convey("Encoded as JSON", func() {
							encoder := json.NewEncoder(body)
							err := encoder.Encode(&Email{
								To:      to,
								Subject: "Test",
								From:    from,
								Text:    "This is just a test.",
							})

							So(err, ShouldBeNil)

							Convey("And correct body length", func() {
								length := body.Len()

								Convey("Should not return an error", func() {
									res, err := client.sendMessage(method, url, nil, ioutil.NopCloser(body), length)

									So(err, ShouldBeNil)
									Convey("But should return 401 StatusCode", func() {

										So(res.StatusCode, ShouldEqual, 401)

										// Drain and close the body to let the Transport reuse the connection
										io.Copy(ioutil.Discard, res.Body)
										res.Body.Close()
									})
								})
							})
						})
					})
				})
			})
		})
	})
}

func TestSendMessageValidAPIKey(t *testing.T) {
	Convey("Given a valid API key", t, func() {
		apiKey := validApiKey

		Convey("And a client", func() {
			client := NewClient(apiKey)

			Convey("And a POST method", func() {
				method := METHOD

				Convey("And a target url for email", func() {
					url := BASE_URL + EMAIL_URL

					Convey("And a proper body", func() {
						body := &bytes.Buffer{}

						Convey("Encoded as JSON", func() {
							encoder := json.NewEncoder(body)
							err := encoder.Encode(&Email{
								To:      to,
								Subject: "Test",
								From:    from,
								Text:    "This is just a test.",
							})

							So(err, ShouldBeNil)

							Convey("And correct body length", func() {
								length := body.Len()

								Convey("Should not return an error", func() {
									res, err := client.sendMessage(method, url, nil, ioutil.NopCloser(body), length)

									So(err, ShouldBeNil)
									Convey("And should not return 401 StatusCode", func() {

										So(res.StatusCode, ShouldNotEqual, 401)

										// Drain and close the body to let the Transport reuse the connection
										io.Copy(ioutil.Discard, res.Body)
										res.Body.Close()
									})
								})
							})
						})
					})
				})
			})
		})
	})
}

func TestSendEmailWrongAPIKey(t *testing.T) {
	Convey("Given a wrong API key", t, func() {
		apiKey := WRONG_API_KEY

		Convey("And an email", func() {
			email := &Email{
				To:      to,
				Subject: "Test",
				From:    from,
				Text:    "This is just a test.",
			}

			Convey("And create a Goinblue client using the key", func() {
				client := NewClient(apiKey)

				Convey("Should return an error", func() {
					_, err := client.SendEmail(email)

					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestSendEmailTemplateWrongAPIKey(t *testing.T) {
	Convey("Given a wrong API key", t, func() {
		apiKey := WRONG_API_KEY

		Convey("And an email template", func() {
			emailTemplate := &EmailTemplate{
				Id: id,
				To: to,
			}

			Convey("And create a Goinblue client using the key", func() {
				client := NewClient(apiKey)

				Convey("Should return an error", func() {
					_, err := client.SendEmailTemplate(emailTemplate)

					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestSendSMSWrongAPIKey(t *testing.T) {
	Convey("Given a wrong API key", t, func() {
		apiKey := WRONG_API_KEY

		Convey("And an sms", func() {
			sms := &SMS{
				To:   toMobileNumber,
				From: fromSMSName,
			}

			Convey("And create a Goinblue client using the key", func() {
				client := NewClient(apiKey)

				Convey("Should return an error", func() {
					_, err := client.SendSMS(sms)

					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestSendEmailValidAPIKey(t *testing.T) {
	Convey("Given a valid API key", t, func() {
		apiKey := validApiKey

		Convey("And an email", func() {
			email := &Email{
				To:      to,
				Subject: "Test",
				From:    from,
				Text:    "This is just a test.",
			}

			Convey("And create a Goinblue client using the key", func() {
				client := NewClient(apiKey)

				Convey("Should not return an error", func() {
					_, err := client.SendEmail(email)

					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestSendEmailTemplateValidAPIKey(t *testing.T) {
	Convey("Given a valid API key", t, func() {
		apiKey := validApiKey

		Convey("And an email template", func() {
			emailTemplate := &EmailTemplate{
				Id: id,
				To: to,
			}

			Convey("And create a Goinblue client using the key", func() {
				client := NewClient(apiKey)

				Convey("Should not return an error", func() {
					_, err := client.SendEmailTemplate(emailTemplate)

					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestSendSMSValidAPIKey(t *testing.T) {
	Convey("Given a valid API key", t, func() {
		apiKey := validApiKey

		Convey("And an sms", func() {
			sms := &SMS{
				To:   toMobileNumber,
				From: fromSMSName,
			}

			Convey("And create a Goinblue client using the key", func() {
				client := NewClient(apiKey)

				Convey("Should not return an error", func() {
					_, err := client.SendSMS(sms)

					So(err, ShouldBeNil)
				})
			})
		})
	})
}
