update:
	go get -u google.golang.org/appengine
	go get -u github.com/migore/paypal
	go get -u github.com/mjibson/goon

create-web-experience:
	curl -v POST https://api.sandbox.paypal.com/v1/payment-experience/web-profiles/ \
	-H 'Content-Type:application/json' \
	-H 'Authorization: Bearer A101.d_kV70FhGtfWMoJna_KeQA3Ok07ItTzuObC0rhC-owX4vpzWxNjmxHslyighD-Lr.POKhEfeyL6P7OUGIM9lgVMsWtMG' \
	-d '{ \
		"name": "Natal Renascer da Esperança 2015", \
		"presentation": { \
			"brand_name": "Natal Renascer da Esperança 2015", \
			"logo_image": "http://fb-canvas-dot-staging-api-getunseen.appspot.com/kids.jpg", \
			"locale_code": "BR" \
		}, \
		"input_fields": { \
			"allow_note": true, \
			"no_shipping": 1, \
			"address_override": 0 \
		}, \
		"flow_config": { \
			"landing_page_type": "billing", \
			"bank_txn_pending_url": "https://www.facebook.com/events/200508130284787/" \
		} \
	}'

update-web-experience:
	curl -v -X PUT https://api.sandbox.paypal.com/v1/payment-experience/web-profiles/XP-3L6B-V2T3-RGFL-2JBZ \
	-H 'Content-Type:application/json' \
	-H 'Authorization: Bearer A101.d_kV70FhGtfWMoJna_KeQA3Ok07ItTzuObC0rhC-owX4vpzWxNjmxHslyighD-Lr.POKhEfeyL6P7OUGIM9lgVMsWtMG' \
	-d '{ \
		"name": "Natal Renascer da Esperança 2015", \
		"presentation": { \
			"brand_name": "Natal Renascer da Esperança 2015", \
			"logo_image": "http://fb-canvas-dot-staging-api-getunseen.appspot.com/kids.jpg", \
			"locale_code": "BR" \
		}, \
		"input_fields": { \
			"allow_note": true, \
			"no_shipping": 1, \
			"address_override": 0 \
		}, \
		"flow_config": { \
			"landing_page_type": "login", \
			"bank_txn_pending_url": "https://www.facebook.com/events/200508130284787/" \
		} \
	}'