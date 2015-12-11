update:
	go get -u google.golang.org/appengine
	go get -u github.com/migore/paypal
	go get -u github.com/mjibson/goon

create-web-experience:
	curl -v POST https://api.sandbox.paypal.com/v1/payment-experience/web-profiles/ \
	-H 'Content-Type:application/json' \
	-H 'Authorization: Bearer A101.d_kV70FhGtfWMoJna_KeQA3Ok07ItTzuObC0rhC-owX4vpzWxNjmxHslyighD-Lr.POKhEfeyL6P7OUGIM9lgVMsWtMG' \
	-d '{ \
		"name": "Faça Uma Criança Feliz", \
		"presentation": { \
			"brand_name": "Faça Uma Criança Feliz", \
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
		"name": "Faça Uma Criança Feliz", \
		"presentation": { \
			"brand_name": "Faça Uma Criança Feliz", \
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

bearer-token:
	curl https://api.paypal.com/v1/oauth2/token \
	-H "Accept: application/json" \
	-H "Accept-Language: en_US" \
	-u "ATRub8NK5m1iZV1EFPcs2Ad_lcKx6A7yasQaRSj6wdjKEDPBpzZ1UZBUr4qQtxg45fG-zO8OlZ85fJx4:EF4fNq7M9l_VztubdFCLsTsUnqGAoSj12WTnGWuguyQKisAC2aneCVNuXDAusmwE5EjDit67YYTMev3z" \
	-d "grant_type=client_credentials"