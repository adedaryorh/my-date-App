import 'package:celebut/features/app_widget.dart';
import 'package:celebut/firebase_options_prod.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await SystemChrome.setPreferredOrientations([
    DeviceOrientation.portraitUp,
    DeviceOrientation.portraitDown,
  ]);
  await Firebase.initializeApp(
    name: 'celebut-prod',
    options: DefaultFirebaseOptions.currentPlatform,
  );

  FlutterError.onError = (errorDetails) {};

  PlatformDispatcher.instance.onError = (error, stackTrace) {
    return true;
  };
  runApp(const ProviderScope(child: AppWidget()));
}
