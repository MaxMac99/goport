package notifications

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/sideshow/apns2/payload"
)

type priority int

const (
	// PriorityLow will tell APNs to send the push message at a time that takes
	// into account power considerations for the device. Notifications with this
	// priority might be grouped and delivered in bursts. They are throttled,
	// and in some cases are not delivered.
	PriorityLow priority = 5

	// PriorityHigh will tell APNs to send the push message immediately.
	// Notifications with this priority must trigger an alert, sound, or badge
	// on the target device. It is an error to use this priority for a push
	// notification that contains only the content-available key.
	PriorityHigh priority = 10
)

type pushType string

const (
	// PushTypeAlert is used for notifications that trigger a user interaction —
	// for example, an alert, badge, or sound. If you set this push type, the
	// apns-topic header field must use your app’s bundle ID as the topic. The
	// alert push type is required on watchOS 6 and later. It is recommended on
	// macOS, iOS, tvOS, and iPadOS.
	PushTypeAlert pushType = "alert"

	// PushTypeBackground is used for notifications that deliver content in the
	// background, and don’t trigger any user interactions. If you set this push
	// type, the apns-topic header field must use your app’s bundle ID as the
	// topic. The background push type is required on watchOS 6 and later. It is
	// recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeBackground pushType = "background"

	// PushTypeVOIP is used for notifications that provide information about an
	// incoming Voice-over-IP (VoIP) call. If you set this push type, the
	// apns-topic header field must use your app’s bundle ID with .voip appended
	// to the end. If you’re using certificate-based authentication, you must
	// also register the certificate for VoIP services. The voip push type is
	// not available on watchOS. It is recommended on macOS, iOS, tvOS, and
	// iPadOS.
	PushTypeVOIP pushType = "voip"

	// PushTypeComplication is used for notifications that contain update
	// information for a watchOS app’s complications. If you set this push type,
	// the apns-topic header field must use your app’s bundle ID with
	// .complication appended to the end. If you’re using certificate-based
	// authentication, you must also register the certificate for WatchKit
	// services. The complication push type is recommended for watchOS and iOS.
	// It is not available on macOS, tvOS, and iPadOS.
	PushTypeComplication pushType = "complication"

	// PushTypeFileProvider is used to signal changes to a File Provider
	// extension. If you set this push type, the apns-topic header field must
	// use your app’s bundle ID with .pushkit.fileprovider appended to the end.
	// The fileprovider push type is not available on watchOS. It is recommended
	// on macOS, iOS, tvOS, and iPadOS.
	PushTypeFileProvider pushType = "fileprovider"

	// PushTypeMDM is used for notifications that tell managed devices to
	// contact the MDM server. If you set this push type, you must use the topic
	// from the UID attribute in the subject of your MDM push certificate.
	PushTypeMDM pushType = "mdm"
)

type NotificationRequest struct {
	AppName     string
	ApnsID      string
	Expiration  *time.Time
	Priority    priority
	CollapseID  string
	PushType    pushType
	Devices     []string
	Payload     *payload.Payload
	Development bool
}

func SendNotification(request NotificationRequest) {
	urlString := "http://apns-proxy_devcontainer-apnsproxy-1:8632/notifications/" + request.AppName
	if request.Development {
		urlString += "/development"
	}
	url, _ := url.Parse(urlString)

	bodyBytes, _ := json.Marshal(map[string]interface{}{
		"devices": request.Devices,
		"payload": request.Payload,
	})
	body := io.NopCloser(bytes.NewReader(bodyBytes))
	headers := map[string][]string{
		"Content-Type": {"application/json; charset=UTF-8"},
	}
	if request.ApnsID != "" {
		headers["apns-id"] = []string{request.ApnsID}
	}
	if request.Expiration != nil {
		headers["apns-expiration"] = []string{strconv.FormatInt(request.Expiration.Unix(), 10)}
	}
	if request.Priority != 0 {
		headers["apns-priority"] = []string{strconv.FormatInt(int64(request.Priority), 10)}
	}
	if request.CollapseID != "" {
		headers["apns-collapse-id"] = []string{request.CollapseID}
	}
	if request.PushType != "" {
		headers["apns-push-type"] = []string{string(request.PushType)}
	}
	req := &http.Request{
		Method: "POST",
		URL:    url,
		Header: headers,
		Body:   body,
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Notification failed:", err)
		return
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	log.Println("Finished with status \"" + res.Status + "\" body: \"" + string(data))
}
