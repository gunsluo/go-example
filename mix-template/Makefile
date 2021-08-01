
template:
	@gomplate --config web/gomplate/gomplate.yaml --input-dir web/fluxble --output-dir web/fluxbletmp
	@gomplate --config web/gomplate/gomplate.yaml --input-dir web/meera --output-dir web/meeratmp

clean-template:
	@rm -fr web/fluxbletmp
	@rm -fr web/meeratmp

.PHONY: template clean-template
