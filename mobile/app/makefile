run:
		flutter run \
		--flavor $(flavor) \
		--target lib/main_$(flavor).dart \
		--dart-define-from-file=lib/env/$(flavor)_creds.json


clean: 
		flutter clean
		cd ios && rm -rf Podfile.lock
		cd ios && rm -rf Pods
		flutter pub get
		cd ios && pod install

clean_gen:
		   dart run build_runner clean
		   dart run build_runner build --delete-conflicting-outputs

gen:
	 dart run build_runner build --delete-conflicting-outputs